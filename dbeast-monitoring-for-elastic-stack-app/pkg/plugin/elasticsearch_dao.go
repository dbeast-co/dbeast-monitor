package plugin

import (
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"net/http"
)

var TemplatesMap map[string]interface{}
var LSConfigs map[string]string
var NewCluster Cluster

type Credentials struct {
	Host                  string `json:"host"`
	AuthenticationEnabled bool   `json:"authentication_enabled"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	Status                string `json:"status"`
}

type EnvironmentConfig struct {
	Prod struct {
		Elasticsearch Credentials `json:"elasticsearch"`
		Kibana        Credentials `json:"kibana"`
	} `json:"prod"`
	Mon struct {
		Elasticsearch Credentials `json:"elasticsearch"`
	} `json:"mon"`
}

type Status struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type StatusData struct {
	Prod struct {
		Elasticsearch Status `json:"elasticsearch"`
		Kibana        Status `json:"kibana"`
	} `json:"prod"`
	Mon struct {
		Elasticsearch Status `json:"elasticsearch"`
	} `json:"mon"`
}

type ConfigurationCheckbox struct {
	Label     string `json:"label"`
	Id        string `json:"id"`
	IsChecked bool   `json:"is_checked"`
}

type LogstashHost struct {
	ServerAddress        string `json:"server_address"`
	LogstashApiHost      string `json:"logstash_api_host"`
	LogstashConfigFolder string `json:"logstash_config_folder"`
	LogstashLogsFolder   string `json:"logstash_logs_folder"`
}

type LogstashConfigurations struct {
	EsMonitoringConfigurationFiles       []ConfigurationCheckbox `json:"es_monitoring_configuration_files"`
	LogstashMonitoringConfigurationFiles struct {
		Configurations []ConfigurationCheckbox `json:"configurations"`
		Hosts          []LogstashHost          `json:"hosts"`
	} `json:"logstash_monitoring_configuration_files"`
}
type Cluster struct {
	ClusterConnectionSettings EnvironmentConfig      `json:"cluster_connection_settings"`
	LogstashConfigurations    LogstashConfigurations `json:"logstash_configurations"`
}

/*
GetStatus makes an HTTP GET request to retrieve the cluster health status based on the provided credentials
and returns the HTTP response.
*/
func GetStatus(credentials Credentials) (*http.Response, error) {
	requestURL := credentials.Host + "/_cluster/health"
	log.DefaultLogger.Info("Request path: ", requestURL)
	log.DefaultLogger.Info("Request host: ", credentials.Host)
	response, err := ProcessGetRequest(credentials, requestURL)

	if err != nil {
		log.DefaultLogger.Warn("Error making HTTP request: " + err.Error())
		return nil, err
	}
	return response, err
}

/*
GetClusterNameAndUid makes an HTTP GET request to retrieve cluster name and UID information based on the provided credentials
and returns the HTTP response.
*/
func GetClusterNameAndUid(credentials Credentials) (*http.Response, error) {
	requestURL := credentials.Host + "/"
	response, err := ProcessGetRequest(credentials, requestURL)

	if err != nil {
		log.DefaultLogger.Warn("Error making HTTP request: " + err.Error())
		return nil, err
	}
	return response, err
}
