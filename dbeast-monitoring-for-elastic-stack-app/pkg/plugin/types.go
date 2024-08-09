package plugin

type Project struct {
	ClusterConnectionSettings  EnvironmentConfig                  `json:"cluster_connection_settings"`
	MonitoringClusterInjection MonitoringClusterInjectionSettings `json:"monitoring_cluster_injection"`
	LogstashConfigurations     LogstashConfigurations             `json:"logstash_configurations"`
}

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

type MonitoringClusterInjectionSettings struct {
	TemplatesInjection   bool `json:"templates_injection"`
	ILMPoliciesInjection bool `json:"ilm_policies_injection"`
	CreateFirsIndices    bool `json:"create_first_indices"`
}

type LogstashHost struct {
	ServerAddress      string `json:"server_address"`
	LogstashApiHost    string `json:"logstash_api_host"`
	LogstashLogsFolder string `json:"logstash_logs_folder"`
}

type LogstashConfigurations struct {
	EsMonitoringConfigurationFiles       []ConfigurationCheckbox              `json:"es_monitoring_configuration_files"`
	LogstashMonitoringConfigurationFiles LogstashMonitoringConfigurationFiles `json:"logstash_monitoring_configuration_files"`
}

type LogstashMonitoringConfigurationFiles struct {
	Configurations []ConfigurationCheckbox `json:"configurations"`
	Hosts          []LogstashHost          `json:"hosts"`
}

type ConfigurationCheckbox struct {
	Label     string `json:"label"`
	Id        string `json:"id"`
	IsChecked bool   `json:"is_checked"`
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

type Status struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
