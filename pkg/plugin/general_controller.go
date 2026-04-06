package plugin

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func (a *App) GetVersion(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("{ \"version\": \"" + applicationVersion + "\" }"))
	if err != nil {
		return
	}
	return
}

func CloneObject(data interface{}) interface{} {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.DefaultLogger.Warn("Failed to marshal data: " + err.Error())
		return data
	}

	var clonedTemplate interface{}
	if err := json.Unmarshal(dataBytes, &clonedTemplate); err != nil {
		log.DefaultLogger.Warn("Failed to unmarshal cloned data: " + err.Error())
		return data
	}

	return clonedTemplate
}

func sanitizeHost(host *string) {
	if strings.HasSuffix(*host, "/") {
		*host = strings.TrimSuffix(*host, "/")
	}
}

func sanitizeEnvironmentConfig(config *EnvironmentConfig) {
	sanitizeHost(&config.Prod.Elasticsearch.Host)
	sanitizeHost(&config.Prod.Kibana.Host)
	sanitizeHost(&config.Mon.Elasticsearch.Host)
	sanitizeHost(&config.Mon.Grafana.Host)
}

func DeferHandler(request *http.Request, logger log.Logger) {
	if request != nil && request.Body != nil {
		_, _ = io.Copy(io.Discard, request.Body)
		if err := request.Body.Close(); err != nil {
			logger.Error("Failed to close the body: " + err.Error())
		}
	}
}
func DeferInternalHandler(response *http.Response, logger log.Logger) {
	if response != nil && response.Body != nil {
		_, _ = io.Copy(io.Discard, response.Body)
		if err := response.Body.Close(); err != nil {
			logger.Error("Failed to close the body: " + err.Error())
		}
	}
}
