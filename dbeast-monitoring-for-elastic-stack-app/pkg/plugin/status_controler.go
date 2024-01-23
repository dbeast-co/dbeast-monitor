package plugin

import (
	"encoding/json"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"io"
	"net/http"
)

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
