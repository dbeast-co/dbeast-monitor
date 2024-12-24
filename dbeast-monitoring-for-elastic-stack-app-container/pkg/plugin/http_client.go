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
	"time"
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
		client = &http.Client{
			Transport: authTransport,
			Timeout:   10 * time.Second,
		}
	} else {
		client = &http.Client{
			Transport: tr,
			Timeout:   10 * time.Second,
		}
	}
	return client, nil
}

// ProcessGETRequest performs an HTTP GET request based on the provided credentials and request URL.
// It uses CreateHTTPClient to create an HTTP client, constructs a GET request, adds basic authentication if enabled,
// and returns the HTTP response.
func ProcessGETRequest(credentials Credentials, requestURL string) (*http.Response, error) {
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

	response, err := client.Do(req)
	if err != nil {
		log.DefaultLogger.Error("HTTP request failed: " + err.Error())
		return response, err
	}

	return response, nil
}

// ProcessPUTRequest performs an HTTP POST request based on the provided credentials and request URL.
// It uses CreateHTTPClient to create an HTTP client, constructs a GET request, adds basic authentication if enabled,
// and returns the HTTP response.
func ProcessPOSTRequest(credentials Credentials, requestURL string, body string) (*http.Response, error) {
	client, err := CreateHTTPClient(credentials)
	if err != nil {
		log.DefaultLogger.Warn("Failed to create HTTP client: " + err.Error())
		return nil, err
	}

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.DefaultLogger.Warn("Failed to create HTTP request: " + err.Error())
		return nil, err
	}
	Header := map[string][]string{
		"Content-Type": {"application/json"},
	}
	req.Header = Header

	response, err := client.Do(req)
	if err != nil {
		log.DefaultLogger.Error("HTTP request failed: " + err.Error())
		return response, err
	} else {
		body, err := io.ReadAll(response.Body)
		err2 := response.Body.Close()
		if err2 != nil {
			log.DefaultLogger.Error("Failed to close response body" + string(body) + err.Error())
		}

		result := map[string]interface{}{}
		if err != nil {
			log.DefaultLogger.Error("Failed to read response body" + string(body) + err.Error())
		} else if len(body) > 0 {
			err := json.Unmarshal(body, &result)
			if err != nil {
				log.DefaultLogger.Error("Failed to unmarshal response body: " + string(body) + err.Error())
			}
		}
		log.DefaultLogger.Info("Response from the Post operation: " + string(body))
	}

	return response, nil
}

// ProcessPUTRequest performs an HTTP PUT request based on the provided credentials and request URL.
// It uses CreateHTTPClient to create an HTTP client, constructs a GET request, adds basic authentication if enabled,
// and returns the HTTP response.
func ProcessPUTRequest(credentials Credentials, requestURL string, body string) (*http.Response, error) {
	client, err := CreateHTTPClient(credentials)
	if err != nil {
		log.DefaultLogger.Warn("Failed to create HTTP client: " + err.Error())
		return nil, err
	}

	req, err := http.NewRequest("PUT", requestURL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.DefaultLogger.Warn("Failed to create HTTP request: " + err.Error())
		return nil, err
	}
	Header := map[string][]string{
		"Content-Type": {"application/json"},
	}
	req.Header = Header

	response, err := client.Do(req)
	if err != nil {
		log.DefaultLogger.Error("HTTP request failed: " + err.Error())
		return response, err
	} else {
		body, err := io.ReadAll(response.Body)
		err2 := response.Body.Close()
		if err2 != nil {
			log.DefaultLogger.Error("Failed to close response body" + string(body) + err.Error())
		}

		result := map[string]interface{}{}
		if err != nil {
			log.DefaultLogger.Error("Failed to read response body" + string(body) + err.Error())
		} else if len(body) > 0 {
			err := json.Unmarshal(body, &result)
			if err != nil {
				log.DefaultLogger.Error("Failed to unmarshal response body: " + string(body) + err.Error())
			}
		}
		log.DefaultLogger.Info("Response from the Put operation: " + string(body))
	}

	return response, nil
}

func ProcessHEADRequest(credentials Credentials, requestURL string) (*http.Response, error) {
	client, err := CreateHTTPClient(credentials)
	if err != nil {
		log.DefaultLogger.Warn("Failed to create HTTP client: " + err.Error())
		return nil, err
	}

	req, err := http.NewRequest("HEAD", requestURL, nil)
	if err != nil {
		log.DefaultLogger.Warn("Failed to create HTTP request: " + err.Error())
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		log.DefaultLogger.Error("HTTP request failed: " + err.Error())
		return response, err
	}

	return response, nil
}

func ProcessDELETERequest(credentials Credentials, requestURL string) (*http.Response, error) {
	client, err := CreateHTTPClient(credentials)
	if err != nil {
		log.DefaultLogger.Warn("Failed to create HTTP client: " + err.Error())
		return nil, err
	}

	req, err := http.NewRequest("DELETE", requestURL, nil)
	if err != nil {
		log.DefaultLogger.Warn("Failed to create HTTP request: " + err.Error())
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		log.DefaultLogger.Error("HTTP request failed: " + err.Error())
		return response, err
	} else {
		body, err := io.ReadAll(response.Body)
		err2 := response.Body.Close()
		if err2 != nil {
			log.DefaultLogger.Error("Failed to close response body" + string(body) + err.Error())
		}

		result := map[string]interface{}{}
		if err != nil {
			log.DefaultLogger.Error("Failed to read response body" + string(body) + err.Error())
		} else if len(body) > 0 {
			err := json.Unmarshal(body, &result)
			if err != nil {
				log.DefaultLogger.Error("Failed to unmarshal response body: " + string(body) + err.Error())
			}
		}
		log.DefaultLogger.Info("Response from the DELETE operation: " + string(body))
	}
	return response, nil
}
