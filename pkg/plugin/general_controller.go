package plugin

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func (a *App) GetGrafanaAPIDatasourceHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		response.Header().Set("Allow", http.MethodGet)
		http.Error(response, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	template := dataWarehouse.GrafanaAPIDatasource
	if template == nil {
		HTTPErrorGenerator(response, errors.New("Grafana API datasource template is missing"), "Failed to get datasource template: ", http.StatusInternalServerError, log.DefaultLogger.FromContext(request.Context()))
		return
	}
	data, err := json.Marshal(template)
	if err != nil {
		HTTPErrorGenerator(response, err, "Failed to encode datasource template: ", http.StatusInternalServerError, log.DefaultLogger.FromContext(request.Context()))
		return
	}
	response.Header().Set("Content-Type", "application/json")
	if _, err := response.Write(data); err != nil {
		log.DefaultLogger.Error("Failed to write datasource template response: " + err.Error())
	}
}

func (a *App) GetVersionsHandler(response http.ResponseWriter, request *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(request.Context())
	ctxLogger.Info("Got request for the version backward compatibility map")

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	versionsForSend, err := json.MarshalIndent(dataWarehouse.BackCompatibilityVersionsMap, "", "")
	if err != nil {
		HTTPErrorGenerator(response, err, "Failed to normalize data, for versions request: ", http.StatusInternalServerError, ctxLogger)
		return
	}

	_, err = response.Write(versionsForSend)
	if err != nil {
		log.DefaultLogger.Error("Can't write to the response: " + err.Error())
		return
	}
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

func SanitizeEnvironmentConfig(config *dataWarehouse.EnvironmentConfig) {
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

func sanitizeHost(host *string) {
	if strings.HasSuffix(*host, "/") {
		*host = strings.TrimSuffix(*host, "/")
	}
}
