package plugin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func GetClusterHealth(client *http.Client, host string) (*http.Response, error) {
	requestURL := host + "/_cluster/health"
	log.DefaultLogger.Debug("Request path: ", requestURL)
	response, err := ProcessGETRequest(client, requestURL)

	if err != nil {
		log.DefaultLogger.Error("Error making HTTP request: " + err.Error())
		return response, err
	}
	return response, err
}

func GetClusterInfo(client *http.Client, host string) (string, string, error) {
	var clusterName, uid string

	response, err := GetClusterInformation(client, host)
	if err != nil {
		log.DefaultLogger.Error("Failed to get cluster name and UID: " + err.Error())
		return "ERROR", "ERROR", err
	}

	defer DeferInternalHandler(response, log.DefaultLogger)

	if response.StatusCode == http.StatusOK {
		body, err := io.ReadAll(response.Body)
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
	return clusterName, uid, nil
}

func GetClusterInformation(client *http.Client, host string) (*http.Response, error) {
	requestURL := host + "/"

	response, err := ProcessGETRequest(client, requestURL)
	if err != nil {
		log.DefaultLogger.Error("Error making HTTP request: " + err.Error())
		return nil, err
	}

	return response, err
}

func SendILMToCluster(client *http.Client, host string, policyName string, templateContent string) (*http.Response, error) {
	requestURL := host + "/_ilm/policy/" + policyName
	exists, err := CheckIsILMExists(client, host, policyName)
	if err != nil {
		return nil, err
	} else if exists {
		log.DefaultLogger.Info("The ILM police: " + policyName + " already exists")
		return nil, nil
	} else {
		return ProcessPUTRequest(client, requestURL, templateContent)
	}
}

func SendComponentTemplateToCluster(client *http.Client, host string, templateName string, templateContent string) (*http.Response, error) {
	requestURL := host + "/_component_template/" + templateName
	return ProcessPUTRequest(client, requestURL, templateContent)
}

func SendIndexTemplateToCluster(client *http.Client, host string, templateName string, templateContent string) (*http.Response, error) {
	requestURL := host + "/_index_template/" + templateName
	return ProcessPUTRequest(client, requestURL, templateContent)
}

func SendFirstIndexToCluster(client *http.Client, host string, templateName string, templateContent string) (*http.Response, error) {
	requestURL := host + "/%3C" + templateName + "-%7Bnow%2Fd%7D-000001%3E"
	return ProcessPUTRequest(client, requestURL, templateContent)
}

func CheckIsIndexExists(client *http.Client, host string, indexName string) (bool, error) {
	requestUrl := host + "/_cat/indices/" + indexName + "?format=json&h=index"
	response, err := ProcessGETRequest(client, requestUrl)
	if err != nil {
		log.DefaultLogger.Error("Failed to check is index exists: " + err.Error())
		return false, err
	}

	defer DeferInternalHandler(response, log.DefaultLogger)

	if response.StatusCode == http.StatusOK {
		body, err := io.ReadAll(response.Body)

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

func CheckIsILMExists(client *http.Client, host string, policyName string) (bool, error) {
	requestUrl := host + "/_ilm/policy/" + policyName

	response, err := ProcessGETRequest(client, requestUrl)
	if err != nil {
		log.DefaultLogger.Error("Failed to check if ILM policy exists: " + err.Error())
		return false, err
	}

	defer DeferInternalHandler(response, log.DefaultLogger)

	if response.StatusCode == http.StatusOK {
		log.DefaultLogger.Debug("ILM policy: " + policyName + " exists")
		return true, nil
	} else if response.StatusCode == http.StatusNotFound {
		log.DefaultLogger.Debug("ILM policy: " + policyName + " does not exist")
		return false, nil
	} else {
		body, _ := io.ReadAll(response.Body)
		log.DefaultLogger.Error("Unexpected response while checking ILM policy. HTTP Status: " + response.Status + ", Body: " + string(body))
		return false, fmt.Errorf("unexpected response status: %d, response: %s", response.StatusCode, string(body))
	}
}

func SendRolloverCommand(client *http.Client, host string, rolloverAlias string) (*http.Response, error) {
	requestURL := host + "/" + rolloverAlias + "/_rollover"
	return ProcessPOSTRequest(client, requestURL, "")
}
