package plugin

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
	"github.com/grafana/grafana-plugin-sdk-go/backend/httpclient"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type BasicAuthTransport struct {
	Username  string
	Password  string
	Transport http.RoundTripper
}

func CreateHTTPClient(credentials dataWarehouse.Credentials, opts httpclient.Options) (*http.Client, error) {
	if credentials.Host == "" {
		log.DefaultLogger.Error("Host is empty")
		return nil, fmt.Errorf("host is empty")
	}

	if opts.Timeouts == nil {
		timeouts := httpclient.DefaultTimeoutOptions
		timeouts.Timeout = 10 * time.Second
		opts.Timeouts = &timeouts
	}

	if opts.TLS == nil && strings.HasPrefix(credentials.Host, "https://") && credentials.AuthenticationEnabled {
		log.DefaultLogger.Warn("Using InsecureSkipVerify for HTTPS with authentication - only use in development!")
		// #nosec G402 -- legacy support for clusters with self-signed certificates.
		opts.TLS = &httpclient.TLSOptions{InsecureSkipVerify: true}
	}

	configureTransport := opts.ConfigureTransport
	opts.ConfigureTransport = func(options httpclient.Options, transport *http.Transport) {
		transport.Proxy = http.ProxyFromEnvironment
		if configureTransport != nil {
			configureTransport(options, transport)
		}
	}

	client, err := httpclient.New(opts)
	if err != nil {
		return nil, fmt.Errorf("create HTTP client: %w", err)
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
