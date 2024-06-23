package plugin

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"io"
	"net/http"
	"strings"
)

type BasicAuthTransport struct {
	Username  string
	Password  string
	Transport http.RoundTripper
}

func (bat *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(bat.Username, bat.Password)
	return bat.Transport.RoundTrip(req)
}

// CreateHTTPClient creates an HTTP client based on the provided credentials.
// It takes a Credentials struct as input and returns an HTTP client.
func CreateHTTPClient(credentials Credentials) (*http.Client, error) {

	if credentials.Host == "" {
		log.DefaultLogger.Error("Host is empty")
		return nil, fmt.Errorf("host is empty")
	}

	var tr = &http.Transport{}
	var client *http.Client
	if credentials.AuthenticationEnabled == true {
		if strings.HasPrefix(credentials.Host, "https://") {
			tr = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
		}
		authTransport := &BasicAuthTransport{
			Username:  credentials.Username,
			Password:  credentials.Password,
			Transport: tr,
		}
		client = &http.Client{Transport: authTransport}
	} else {
		client = &http.Client{Transport: tr}
	}
	return client, nil
}

// ProcessGetRequest performs an HTTP GET request based on the provided credentials and request URL.
// It uses CreateHTTPClient to create an HTTP client, constructs a GET request, adds basic authentication if enabled,
// and returns the HTTP response.
func ProcessGetRequest(credentials Credentials, requestURL string) (*http.Response, error) {
	client, err := CreateHTTPClient(credentials)
	if err != nil {
		log.DefaultLogger.Warn("Failed to create HTTP client: " + err.Error())
		return nil, err
	}

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.DefaultLogger.Warn("Failed to create HTTP request: " + err.Error())
		return nil, err
	}
	if credentials.AuthenticationEnabled == true {
		req.SetBasicAuth(credentials.Username, credentials.Password)
	}

	response, err := client.Do(req)
	if err != nil {
		log.DefaultLogger.Warn("HTTP request failed: " + err.Error())
		return nil, err
	}
	if response.StatusCode == http.StatusUnauthorized {
		log.DefaultLogger.Warn("Unauthorized access. Check credentials")

		err := response.Body.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("unauthorized access. Check credentials")
	}

	return response, nil
}

/*
FetchClusterInfo retrieves cluster name and UID from Elasticsearch and returns them.
It uses the provided credentials to make a request to Elasticsearch and extracts cluster name and UID from the response.
*/
func FetchClusterInfo(credentials Credentials) (string, string) {

	var clusterName, uid string

	if credentials.Host != "" {
		response, err := GetClusterNameAndUid(credentials)

		if err != nil {
			log.DefaultLogger.Warn("Failed to get cluster name and UID: " + err.Error())
			return "ERROR", "ERROR"
		}

		if response.StatusCode == http.StatusOK {
			body, err := io.ReadAll(response.Body)
			err2 := response.Body.Close()
			if err2 != nil {
				return "ERROR", "ERROR"
			}

			if err != nil {
				log.DefaultLogger.Warn("Failed to read response body" + string(body) + err.Error())
				return "ERROR", "ERROR"
			} else if len(body) > 0 {
				result := map[string]interface{}{}
				err := json.Unmarshal([]byte(body), &result)
				if err != nil {
					log.DefaultLogger.Warn("Failed to unmarshal response body: " + string(body) + err.Error())
					return "ERROR", "ERROR"
				}
				if name, ok := result["cluster_name"].(string); ok {
					clusterName = name
				}
				if uidVal, ok := result["cluster_uuid"].(string); ok {
					uid = uidVal
				}
			}
		} else {
			log.DefaultLogger.Warn("Failed to get cluster name and UID. HTTP status: " + response.Status)
			return "ERROR", "ERROR"
		}
	}
	return clusterName, uid
}

/*
SendTemplateToServer takes a map of templates (Templates) as input, constructs an HTTP POST request for each template,
and sends it to a predefined HTTP endpoint.
*/
func SendTemplateToServer(Templates map[string]interface{}) {

	httpposturl := "http://localhost:3000"
	log.DefaultLogger.Warn("HTTP JSON POST URL:" + httpposturl)
	client := &http.Client{}

	for name, template := range Templates {
		requestBody, err := json.Marshal(template)

		if err != nil {
			log.DefaultLogger.Warn("Failed to marshal template: " + name + err.Error())
			return
		}

		request, err := http.NewRequest("POST", httpposturl+"/test", bytes.NewBuffer(requestBody))
		if err != nil {
			log.DefaultLogger.Warn("Failed to create request for template: " + name + err.Error())
			return
		}

		request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		response, err := client.Do(request)
		if err != nil {
			log.DefaultLogger.Warn("HTTP request failed for template " + name + ": " + err.Error())
			return
		}
		if err := response.Body.Close(); err != nil {
			log.DefaultLogger.Warn("Failed to close response body: " + err.Error())
			return
		}
	}
}
