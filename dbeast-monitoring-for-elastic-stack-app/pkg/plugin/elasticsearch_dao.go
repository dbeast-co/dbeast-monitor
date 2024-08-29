package plugin

import (
	"encoding/json"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"io"
	"net/http"
)

/*
GetClusterHealth makes an HTTP GET request to retrieve the cluster health status based on the provided credentials
and returns the HTTP response.
*/
func GetClusterHealth(credentials Credentials) (*http.Response, error) {
	requestURL := credentials.Host + "/_cluster/health"
	log.DefaultLogger.Debug("Request path: ", requestURL)
	log.DefaultLogger.Debug("Request host: ", credentials.Host)
	response, err := ProcessGETRequest(credentials, requestURL)

	if err != nil {
		log.DefaultLogger.Error("Error making HTTP request: " + err.Error())
		return response, err
	}
	return response, err
}

/*
GetClusterInformation makes an HTTP GET request to retrieve cluster name and UID information based on the provided credentials
and returns the HTTP response.
*/
func GetClusterInformation(credentials Credentials) (*http.Response, error) {
	requestURL := credentials.Host + "/"
	response, err := ProcessGETRequest(credentials, requestURL)

	if err != nil {
		log.DefaultLogger.Error("Error making HTTP request: " + err.Error())
		return nil, err
	}
	return response, err
}

/*
GetClusterInfo retrieves cluster name and UID from Elasticsearch and returns them.
It uses the provided credentials to make a request to Elasticsearch and extracts cluster name and UID from the response.
*/
func GetClusterInfo(credentials Credentials) (string, string, error) {
	var clusterName, uid string

	if credentials.Host != "" {
		response, err := GetClusterInformation(credentials)

		if err != nil {
			log.DefaultLogger.Error("Failed to get cluster name and UID: " + err.Error())
			return "ERROR", "ERROR", err
		}

		if response.StatusCode == http.StatusOK {
			body, err := io.ReadAll(response.Body)
			err2 := response.Body.Close()
			if err2 != nil {
				log.DefaultLogger.Error("Failed to close response body" + string(body) + err.Error())
				return "ERROR", "ERROR", err2
			}

			if err != nil {
				log.DefaultLogger.Error("Failed to read response body" + string(body) + err.Error())
				return "ERROR", "ERROR", err
			} else if len(body) > 0 {
				result := map[string]interface{}{}
				err := json.Unmarshal([]byte(body), &result)
				if err != nil {
					log.DefaultLogger.Error("Failed to unmarshal response body: " + string(body) + err.Error())
					return "ERROR", "ERROR", err
				}
				if name, ok := result["cluster_name"].(string); ok {
					clusterName = name
				}
				if uidVal, ok := result["cluster_uuid"].(string); ok {
					uid = uidVal
				}
			}
		} else {
			log.DefaultLogger.Error("Failed to get cluster name and UID. HTTP status: " + response.Status)
			return "ERROR", "ERROR", err
		}
	}
	return clusterName, uid, nil
}

func SendILMToCluster(credentials Credentials, policyName string, policyContent string) (*http.Response, error) {
	requestURL := credentials.Host + "/_ilm/policy/" + policyName
	return SendDataToCluster(credentials, requestURL, policyContent)
}

func SendComponentTemplateToCluster(credentials Credentials, templateName string, templateContent string) (*http.Response, error) {
	requestURL := credentials.Host + "/_component_template/" + templateName
	return SendDataToCluster(credentials, requestURL, templateContent)
}

func SendIndexTemplateToCluster(credentials Credentials, templateName string, templateContent string) (*http.Response, error) {
	requestURL := credentials.Host + "/_index_template/" + templateName
	return SendDataToCluster(credentials, requestURL, templateContent)
}

func SendFirstIndicesToCluster(credentials Credentials, templateName string, templateContent string) (*http.Response, error) {
	requestURL := credentials.Host + "/%3C" + templateName + "-%7Bnow%2Fd%7D-000001%3E"
	return SendDataToCluster(credentials, requestURL, templateContent)
}
func SendDataToCluster(credentials Credentials, requestURL string, templateContent string) (*http.Response, error) {
	log.DefaultLogger.Info("Request path: " + requestURL)
	log.DefaultLogger.Debug("Request host: " + credentials.Host)
	log.DefaultLogger.Debug("Request body: " + templateContent)
	response, err := ProcessPUTRequest(credentials, requestURL, templateContent)

	if err != nil {
		log.DefaultLogger.Error("Error making HTTP request: " + err.Error())
		return response, err
	}
	return response, err
}
