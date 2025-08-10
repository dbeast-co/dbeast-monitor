package plugin

import (
	"encoding/json"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"net/http"
	"strings"
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
}
