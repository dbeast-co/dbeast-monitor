package plugin

import "net/http"

func AddNewClusterToLogstash(credentials Credentials, clusterId string) (*http.Response, error) {
	DeployLogstashConfigFiles(credentials, []string{})
	AddClusterToPipelinesFile(credentials, clusterId)
	return nil, nil
}

func DeleteClusterFromLogstash(credentials Credentials, clusterId string) (*http.Response, error) {
	DeleteClusterFromPipelinesFile(credentials, clusterId)
	DeleteLogstashConfigurations(credentials, clusterId)
	return nil, nil
}

func AddClusterToPipelinesFile(credentials Credentials, pipelineFileContent string) (*http.Response, error) {
	requestURL := credentials.Host
	return ProcessPOSTRequest(credentials, requestURL, pipelineFileContent)
}

func DeployLogstashConfigFiles(credentials Credentials, logstashConfigFiles []string) (*http.Response, error) {
	requestURL := credentials.Host
	for _, file := range logstashConfigFiles {
		ProcessPOSTRequest(credentials, requestURL, file)
	}
	return nil, nil
}

func DeleteClusterFromPipelinesFile(credentials Credentials, clusterId string) (*http.Response, error) {
	return nil, nil
}

func DeleteLogstashConfigurations(credentials Credentials, clusterId string) (*http.Response, error) {
	return nil, nil
}
