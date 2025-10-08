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
	requestURL := host + "/"
	log.DefaultLogger.Debug("Request path: ", requestURL)

	response, err := ProcessGETRequest(client, requestURL)
	if err != nil {
		log.DefaultLogger.Error("Request path: " + requestURL)
		log.DefaultLogger.Error("HTTP request failed: " + err.Error())
		return "ERROR", "ERROR", err
	}
	defer DeferInternalHandler(response, log.DefaultLogger)

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return "ERROR", "ERROR", fmt.Errorf("get cluster info: HTTP %s, body: %s",
			response.Status, string(body))
	}

	var resp struct {
		ClusterName string `json:"cluster_name"`
		ClusterUUID string `json:"cluster_uuid"`
	}
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return "ERROR", "ERROR", fmt.Errorf("decode cluster info: %w", err)
	}
	if resp.ClusterName == "" || resp.ClusterUUID == "" {
		return "ERROR", "ERROR", fmt.Errorf("cluster info missing fields")
	}
	return resp.ClusterName, resp.ClusterUUID, nil
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

func CheckIsIndexExists(client *http.Client, host, index string) (bool, error) {
	requestURL := host + "/" + index
	log.DefaultLogger.Debug("Request path: ", requestURL)

	response, err := ProcessGETRequest(client, requestURL)
	if err != nil {
		log.DefaultLogger.Error("Request path: " + requestURL)
		log.DefaultLogger.Error("HTTP request failed: " + err.Error())
		return false, err
	}
	defer DeferInternalHandler(response, log.DefaultLogger)

	switch response.StatusCode {
	case http.StatusOK:
		return true, nil
	case http.StatusNotFound:
		return false, nil
	default:
		body, readErr := io.ReadAll(response.Body)
		if readErr != nil {
			return false, fmt.Errorf("check index exists: HTTP %s, read body: %w",
				response.Status, readErr)
		}
		return false, fmt.Errorf("check index exists: HTTP %s, body: %s",
			response.Status, string(body))
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
