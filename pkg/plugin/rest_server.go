package plugin

import (
	"net/http"
)

func (a *App) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/new_cluster", a.NewClusterHandler)
	mux.HandleFunc("/test_cluster", a.TestClusterHandler)
	mux.HandleFunc("/add_cluster", a.AddClusterHandler)
	mux.HandleFunc("/add_cluster_api", a.AddClusterViaAPIHandler)
	mux.HandleFunc("/delete_cluster_api", a.DeleteClusterViaAPIHandler)
	mux.HandleFunc("/delete_cluster/", a.DeleteClusterHandler)
	mux.HandleFunc("/deploy_elasticsearch_configuration", a.DeployElasticsearchConfigurations)
	mux.HandleFunc("/download_logstash_monitoring_configuration_files", a.DownloadLogstashMonitoringConfigurationFilesHandler)
	mux.HandleFunc("/download_es_monitoring_configuration_files", a.DownloadElasticsearchMonitoringConfigurationFilesHandler)
	mux.HandleFunc("/save", a.SaveClusterHandler)
	mux.HandleFunc("/versions", a.GetVersionsHandler)
	mux.HandleFunc("/grafana_api_datasource", a.GetGrafanaAPIDatasourceHandler)
}
