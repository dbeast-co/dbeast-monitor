package plugin

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func (a *App) TestClusterHandler(response http.ResponseWriter, request *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(request.Context())

	defer DeferHandler(request, ctxLogger)

	response.Header().Add("Content-Type", "application/json")

	var environmentConfig EnvironmentConfig
	if err := json.NewDecoder(request.Body).Decode(&environmentConfig); err != nil {
		HTTPErrorGenerator(response, err, "Failed to decode JSON data for test cluster request: ", http.StatusInternalServerError, ctxLogger)
		return
	}
	sanitizeEnvironmentConfig(&environmentConfig)

	clientProd, err := CreateHTTPClient(environmentConfig.Prod.Elasticsearch)
	if err != nil {
		HTTPErrorGenerator(response, err, "Error while creating HTTP client for test cluster request: ", http.StatusInternalServerError, ctxLogger)
		return
	}
	clientMon, err := CreateHTTPClient(environmentConfig.Mon.Elasticsearch)
	if err != nil {
		HTTPErrorGenerator(response, err, "Error while creating HTTP client for test cluster request: ", http.StatusInternalServerError, ctxLogger)
		return
	}

	var statusData StatusData

	statusData.Prod.Elasticsearch = UpdateStatus(clientProd, environmentConfig.Prod.Elasticsearch.Host)
	statusData.Mon.Elasticsearch = UpdateStatus(clientMon, environmentConfig.Mon.Elasticsearch.Host)

	statusDataJSON, err := json.MarshalIndent(statusData, "", "")
	if err != nil {
		HTTPErrorGenerator(response, err, "Error while templates update, for cluster save request: ", http.StatusInternalServerError, ctxLogger)
		return
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write(statusDataJSON)
	if err != nil {
		log.DefaultLogger.Error("Can't write to the response: " + err.Error())
		return
	}
}

func UpdateStatus(client *http.Client, host string) Status {
	var statusData = Status{}

	if host != "" {
		response, err := GetClusterHealth(client, host)

		defer DeferInternalHandler(response, log.DefaultLogger)

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
