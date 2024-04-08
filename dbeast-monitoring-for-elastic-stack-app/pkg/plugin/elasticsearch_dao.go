package plugin

import (
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"net/http"
)

var TemplatesMap map[string]interface{}

type Credentials struct {
	Host                  string `json:"host"`
	AuthenticationEnabled bool   `json:"authentication_enabled"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	Status                string `json:"status"`
}
type EnvironmentConfig struct {
	Prod struct {
		Elasticsearch Credentials `json:"elasticsearch"`
		Kibana        Credentials `json:"kibana"`
	} `json:"prod"`
	Mon struct {
		Elasticsearch Credentials `json:"elasticsearch"`
	} `json:"mon"`
}

type Status struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type StatusData struct {
	Prod struct {
		Elasticsearch Status `json:"elasticsearch"`
		Kibana        Status `json:"kibana"`
	} `json:"prod"`
	Mon struct {
		Elasticsearch Status `json:"elasticsearch"`
	} `json:"mon"`
}

/*
GetStatus makes an HTTP GET request to retrieve the cluster health status based on the provided credentials
and returns the HTTP response.
*/
func GetStatus(credentials Credentials) (*http.Response, error) {

	requestURL := credentials.Host + "/_cluster/health"
	response, err := ProcessGetRequest(credentials, requestURL)

	if err != nil {
		log.DefaultLogger.Warn("Error making HTTP request: " + err.Error())
		return nil, err
	}
	return response, err
}

/*
GetClusterNameAndUID makes an HTTP GET request to retrieve cluster name and UID information based on the provided credentials
and returns the HTTP response.
*/
func GetClusterNameAndUid(dataToUpdate Credentials) (*http.Response, error) {

	requestURL := dataToUpdate.Host + "/"
	response, err := ProcessGetRequest(dataToUpdate, requestURL)

	if err != nil {
		log.DefaultLogger.Warn("Error making HTTP request: " + err.Error())
		return nil, err
	}
	return response, err
}
