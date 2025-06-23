package plugin

import (
	"net/http"
)

// registerRoutes takes a *http.ServeMux and registers some HTTP handlers.
func (a *App) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/new_cluster", a.NewClusterHandler)
	mux.HandleFunc("/test_cluster", a.TestClusterHandler)
	mux.HandleFunc("/add_cluster", a.AddClusterHandler)
	mux.HandleFunc("/delete_cluster/", a.DeleteClusterHandler)
	mux.HandleFunc("/download_logstash_monitoring_configuration_files", a.DownloadLogstashMonitoringConfigurationFilesHandler)
	mux.HandleFunc("/download_es_monitoring_configuration_files", a.DownloadElasticsearchMonitoringConfigurationFilesHandler)
	mux.HandleFunc("/save", a.SaveClusterHandler)
}
