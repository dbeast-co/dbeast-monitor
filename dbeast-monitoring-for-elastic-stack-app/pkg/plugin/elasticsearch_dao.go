package plugin

import (
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"net/http"
)

/*
GetStatus makes an HTTP GET request to retrieve the cluster health status based on the provided credentials
and returns the HTTP response.
*/
func GetStatus(credentials Credentials) (*http.Response, error) {
	requestURL := credentials.Host + "/_cluster/health"
	log.DefaultLogger.Debug("Request path: ", requestURL)
	log.DefaultLogger.Debug("Request host: ", credentials.Host)
	response, err := ProcessGetRequest(credentials, requestURL)

	if err != nil {
		log.DefaultLogger.Warn("Error making HTTP request: " + err.Error())
		return nil, err
	}
	return response, err
}

/*
GetClusterNameAndUid makes an HTTP GET request to retrieve cluster name and UID information based on the provided credentials
and returns the HTTP response.
*/
func GetClusterNameAndUid(credentials Credentials) (*http.Response, error) {
	requestURL := credentials.Host + "/"
	response, err := ProcessGetRequest(credentials, requestURL)

	if err != nil {
		log.DefaultLogger.Warn("Error making HTTP request: " + err.Error())
		return nil, err
	}
	return response, err
}
