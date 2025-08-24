package plugin

import (
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"io"
	"net/http"
)

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
				log.DefaultLogger.Error("Failed to close response body" + string(body) + err2.Error())
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

func GetClusterInformation(credentials Credentials) (*http.Response, error) {
	requestURL := credentials.Host + "/"
	response, err := ProcessGETRequest(credentials, requestURL)

	if err != nil {
		log.DefaultLogger.Error("Error making HTTP request: " + err.Error())
		return nil, err
	}
	return response, err
}

func SendILMToCluster(credentials Credentials, policyName string, templateContent string) (*http.Response, error) {
	requestURL := credentials.Host + "/_ilm/policy/" + policyName
	exists, err := CheckIsILMExists(credentials, policyName)
	if err != nil {
		return nil, err
	} else if exists {
		log.DefaultLogger.Info("The ILM police: " + policyName + " already exists")
		return nil, nil
	} else {
		return ProcessPUTRequest(credentials, requestURL, templateContent)
	}
}

func SendComponentTemplateToCluster(credentials Credentials, templateName string, templateContent string) (*http.Response, error) {
	requestURL := credentials.Host + "/_component_template/" + templateName
	return ProcessPUTRequest(credentials, requestURL, templateContent)
}

func SendIndexTemplateToCluster(credentials Credentials, templateName string, templateContent string) (*http.Response, error) {
	requestURL := credentials.Host + "/_index_template/" + templateName
	return ProcessPUTRequest(credentials, requestURL, templateContent)
}

func SendFirstIndexToCluster(credentials Credentials, templateName string, templateContent string) (*http.Response, error) {
	requestURL := credentials.Host + "/%3C" + templateName + "-%7Bnow%2Fd%7D-000001%3E"
	return ProcessPUTRequest(credentials, requestURL, templateContent)
}

func CheckIsIndexExists(credentials Credentials, indexName string) (bool, error) {
	requestUrl := credentials.Host + "/_cat/indices/" + indexName + "?format=json&h=index"
	response, err := ProcessGETRequest(credentials, requestUrl)

	if err != nil {
		log.DefaultLogger.Error("Failed to check is index exists: " + err.Error())
		return false, err
	}

	if response.StatusCode == http.StatusOK {
		body, err := io.ReadAll(response.Body)
		err2 := response.Body.Close()
		if err2 != nil {
			log.DefaultLogger.Error("Failed to close response body" + string(body) + err2.Error())
			return false, err2
		}

		var result []map[string]interface{}
		err = json.Unmarshal([]byte(body), &result)
		if err != nil {
			log.DefaultLogger.Error("Failed to unmarshal response body: " + string(body) + err.Error())
			return false, err
		}

		if len(result) > 0 {
			log.DefaultLogger.Debug("Index: " + indexName + " exists")
			return true, nil
		} else {
			log.DefaultLogger.Debug("Index: " + indexName + " does not exist")
			return false, nil
		}

	} else {
		log.DefaultLogger.Error("Failed to check is index name exists. HTTP status: " + response.Status)
		return false, err
	}
}

func CheckIsILMExists(credentials Credentials, policyName string) (bool, error) {
	requestUrl := credentials.Host + "/_ilm/policy/" + policyName
	response, err := ProcessGETRequest(credentials, requestUrl)

	if err != nil {
		log.DefaultLogger.Error("Failed to check if ILM policy exists: " + err.Error())
		return false, err
	}

	if response.StatusCode == http.StatusOK {
		log.DefaultLogger.Debug("ILM policy: " + policyName + " exists")
		return true, nil
	} else if response.StatusCode == http.StatusNotFound { // Policy not found
		log.DefaultLogger.Debug("ILM policy: " + policyName + " does not exist")
		return false, nil
	} else {
		// Log any other unexpected HTTP status
		body, _ := io.ReadAll(response.Body) // Retain body if needed for debugging
		response.Body.Close()
		log.DefaultLogger.Error("Unexpected response while checking ILM policy. HTTP Status: " + response.Status + ", Body: " + string(body))
		return false, fmt.Errorf("unexpected response status: %d, response: %s", response.StatusCode, string(body))
	}
}

func SendRolloverCommand(credentials Credentials, rolloverAlias string) (*http.Response, error) {
	requestURL := credentials.Host + "/" + rolloverAlias + "/_rollover"
	return ProcessPOSTRequest(credentials, requestURL, "")
}
