package plugin

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend/httpclient"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type BasicAuthTransport struct {
	Username  string
	Password  string
	Transport http.RoundTripper
}

func CreateHTTPClient(credentials Credentials, opts httpclient.Options) (*http.Client, error) {
	_ = opts // API compatibility - opts may be used in future enhancements

	if credentials.Host == "" {
		log.DefaultLogger.Error("Host is empty")
		return nil, fmt.Errorf("host is empty")
	}

	// Note: opts parameter is accepted for API compatibility with the caller
	// The HTTP client is built manually to maintain control over TLS and proxy settings
	// Future enhancement: Can integrate opts.ConfigureClient() or opts.NewClient() if available

	// Create transport with proxy support from environment (respects HTTP_PROXY, HTTPS_PROXY, NO_PROXY)
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}

	// Apply TLS config if needed (HTTPS)
	if strings.HasPrefix(credentials.Host, "https://") {
		if tr.TLSClientConfig == nil {
			tr.TLSClientConfig = &tls.Config{}
		}
		// Apply InsecureSkipVerify if needed (development scenarios)
		// Note: In production, this should be configurable based on credentials
		if credentials.AuthenticationEnabled {
			log.DefaultLogger.Warn("Using InsecureSkipVerify for HTTPS with authentication - only use in development!")
			tr.TLSClientConfig.InsecureSkipVerify = true
		}
	}

	// Create HTTP client
	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	// Wrap with basic auth if needed
	if credentials.AuthenticationEnabled == true {
		authTransport := &BasicAuthTransport{
			Username:  credentials.Username,
			Password:  credentials.Password,
			Transport: client.Transport,
		}
		client.Transport = authTransport
	}

	return client, nil
}

func ProcessGETRequest(client *http.Client, requestURL string) (*http.Response, error) {
	log.DefaultLogger.Debug("Request path: " + requestURL)
	log.DefaultLogger.Debug("Request method: GET")

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.DefaultLogger.Warn("Request path: " + requestURL)
		log.DefaultLogger.Warn("Failed to create HTTP request: " + err.Error())
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		log.DefaultLogger.Error("Request path: " + requestURL)
		log.DefaultLogger.Error("HTTP request failed: " + err.Error())
		return response, err
	}

	return response, nil
}

func ProcessPOSTRequest(client *http.Client, requestURL string, body string) (*http.Response, error) {
	return ProcessRequestWithBody(client, requestURL, body, "POST")
}

func ProcessPUTRequest(client *http.Client, requestURL string, body string) (*http.Response, error) {
	return ProcessRequestWithBody(client, requestURL, body, "PUT")
}

func ProcessRequestWithBody(client *http.Client, requestURL string, body string, method string) (*http.Response, error) {
	log.DefaultLogger.Debug("Request path: " + requestURL)
	log.DefaultLogger.Debug("Request body: " + body)
	log.DefaultLogger.Debug("Request method: " + method)

	req, err := http.NewRequest(method, requestURL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.DefaultLogger.Warn("Request path: " + requestURL)
		log.DefaultLogger.Warn("Failed to create HTTP request: " + err.Error())
		return nil, err
	}
	Header := map[string][]string{
		"Content-Type": {"application/json"},
	}
	req.Header = Header

	response, err := client.Do(req)
	if err != nil {
		log.DefaultLogger.Error("Request path: " + requestURL)
		log.DefaultLogger.Error("HTTP request failed: " + err.Error())
		return response, err
	}

	return response, nil
}

func (bat *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(bat.Username, bat.Password)
	return bat.Transport.RoundTrip(req)
}
