package plugin

import (
	"encoding/json"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"io"
	"net/http"
)

func (a *App) TestClusterHandler(w http.ResponseWriter, req *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(req.Context())
	w.Header().Add("Content-Type", "application/json")

	var environmentConfig EnvironmentConfig
	if err := json.NewDecoder(req.Body).Decode(&environmentConfig); err != nil {
		log.DefaultLogger.Warn("Failed to decode JSON data: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid request payload"})
		return
	}
	sanitizeEnvironmentConfig(&environmentConfig)
	defer req.Body.Close()

	var statusData StatusData

	statusData.Prod.Elasticsearch = UpdateStatus(&environmentConfig.Prod.Elasticsearch)
	statusData.Mon.Elasticsearch = UpdateStatus(&environmentConfig.Mon.Elasticsearch)

	statusDataJSON, err := json.MarshalIndent(statusData, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Failed to marshal status data"})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(statusDataJSON)
	ctxLogger.Info("Status data received")
}

func UpdateStatus(credentials *Credentials) Status {
	var statusData = Status{}

	if credentials.Host != "" {
		response, err := GetClusterHealth(*credentials)

		if err != nil {
			GenerateStatusError(&statusData, err.Error(), "Failed to get status information: ")
			return statusData
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			GenerateStatusError(&statusData, err.Error(), "Failed to read response body: ")
			return statusData
		}

		if response.StatusCode == http.StatusOK {
			if len(body) > 0 {
				err := response.Body.Close()
				if err != nil {
					return Status{}
				}
				result := map[string]interface{}{}
				err = json.Unmarshal([]byte(body), &result)
				if err != nil {
					GenerateStatusError(&statusData, err.Error(), "Failed to unmarshal response body: ")
				} else {
					if status, ok := result["status"].(string); ok {
						statusData.Status = status
					}
				}
			}
		} else {
			if len(body) > 0 {
				err := response.Body.Close()
				if err != nil {
					return Status{}
				}
				GenerateStatusError(&statusData, string(body), "Status fetch error: ")
			} else {
				GenerateStatusError(&statusData, response.Status, "HTTP request failed with status: ")
			}
		}
	}
	return statusData
}

func GenerateStatusError(statusData *Status, error string, errorMessage string) {
	statusData.Error = error
	statusData.Status = "ERROR"
	log.DefaultLogger.Error(errorMessage + statusData.Error)
}
