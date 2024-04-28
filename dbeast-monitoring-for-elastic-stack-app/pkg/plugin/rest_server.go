package plugin

import (
	"net/http"
)

// registerRoutes takes a *http.ServeMux and registers some HTTP handlers.
func (a *App) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/test_cluster", a.TestStatusHandler)
	mux.HandleFunc("/new_cluster", a.NewClusterHandler)
	mux.HandleFunc("/download_logstash_monitoring_configuration_files", a.GenerateLogstashMonitoringConfigurationFilesHandler)
	mux.HandleFunc("/download_es_monitoring_configuration_files", a.GenerateLogstashMonitoringConfigurationFilesHandler)
	mux.HandleFunc("/save", a.SaveHandler)
}
