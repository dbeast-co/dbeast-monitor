package plugin

import (
	"encoding/json"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"io"
	"net/http"
)

/*
	TestStatusHandler handles HTTP requests to retrieve and update the status data based on the provided environment configuration.

It takes a http.ResponseWriter and http.Request as input, decodes the request body to extract environment configuration,
updates the status and sends the updated status data in JSON format as an HTTP response.
*/
func (a *App) TestStatusHandler(w http.ResponseWriter, req *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(req.Context())
	w.Header().Add("Content-Type", "application/json")

	var EnvironmentConfig EnvironmentConfig
	if err := json.NewDecoder(req.Body).Decode(&EnvironmentConfig); err != nil {
		log.DefaultLogger.Warn("Failed to decode JSON data: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid request payload"})
		return
	}
	defer req.Body.Close()

	var statusData StatusData

	statusData.Prod.Elasticsearch = UpdateStatusForType(&EnvironmentConfig.Prod.Elasticsearch)
	statusData.Mon.Elasticsearch = UpdateStatusForType(&EnvironmentConfig.Mon.Elasticsearch)

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

/*
UpdateStatusForType retrieves the status information for a given type.
It uses the provided credentials to make a request to the corresponding status endpoint and extracts the status information from the response.
If successful, it returns a Status struct containing the status information; otherwise, it returns an error message in the Status struct.
*/
func UpdateStatusForType(credentials *Credentials) Status {
	var statusData = Status{}

	if credentials.Host != "" {
		response, err := GetStatus(*credentials)

		if err != nil {
			statusData.Error = err.Error()
			statusData.Status = "ERROR"
			log.DefaultLogger.Warn("Failed to get status information: " + statusData.Error)
			return statusData
		}

		if response.StatusCode == http.StatusOK {
			body, err := io.ReadAll(response.Body)
			//response.Body.Close()

			if err != nil {
				statusData.Error = err.Error()
				statusData.Status = "ERROR"
				log.DefaultLogger.Warn("Failed to read response body: " + statusData.Error)
			} else if len(body) > 0 {
				error := response.Body.Close()
				if error != nil {
					return Status{}
				}
				result := map[string]interface{}{}
				err := json.Unmarshal([]byte(body), &result)
				if err != nil {
					statusData.Error = err.Error()
					statusData.Status = "ERROR"
					log.DefaultLogger.Warn("Failed to unmarshal response body: " + statusData.Error)
				}
				if status, ok := result["status"].(string); ok {
					statusData.Status = status
				}
			}
		} else {
			statusData.Error = response.Status
			statusData.Status = "ERROR"
			log.DefaultLogger.Warn("HTTP request failed with status:" + statusData.Error)
		}
	}
	return statusData
}
