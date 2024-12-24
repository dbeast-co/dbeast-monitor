package plugin

type Project struct {
	ClusterConnectionSettings  EnvironmentConfig       `json:"cluster_connection_settings"`
	MonitoringClusterInjection []ConfigurationCheckbox `json:"monitoring_cluster_injection"`
	LogstashConfigurations     LogstashConfigurations  `json:"logstash_configurations"`
}

type Credentials struct {
	Host                  string `json:"host"`
	AuthenticationEnabled bool   `json:"authentication_enabled"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	Status                string `json:"status"`
}

type EnvironmentConfig struct {
	Prod ProdEnvironmentConfig `json:"prod"`
	Mon  MonEnvironmentConfig  `json:"mon"`
}

type ProdEnvironmentConfig struct {
	Elasticsearch Credentials `json:"elasticsearch"`
	Kibana        Credentials `json:"kibana"`
}

type MonEnvironmentConfig struct {
	Elasticsearch Credentials `json:"elasticsearch"`
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

// Index template definition
// Root structure representing the entire JSON
type IndexTemplate struct {
	IndexPatterns []string `json:"index_patterns"`
	Template      Template `json:"template"`
	ComposedOf    []string `json:"composed_of"`
	Priority      int      `json:"priority"`
	Meta          Meta     `json:"_meta"`
}

// Represents the "template" field
type Template struct {
	Settings TemplateSettings `json:"settings"`
	Aliases  map[string]Alias `json:"aliases"`
}

// Represents the "settings" field under "template"
type TemplateSettings struct {
	Index IndexSettings `json:"index"`
}

// Represents the "index" field under "settings"
type IndexSettings struct {
	Lifecycle IndexLifecycle `json:"lifecycle"`
}

// Represents the "lifecycle" field under "index"
type IndexLifecycle struct {
	Name          string `json:"name"`
	RolloverAlias string `json:"rollover_alias"`
}

// Represents the "aliases" field under "template"
type Alias struct {
	// Empty object, no fields required currently
}

// Represents the "_meta" field
type Meta struct {
	Description string `json:"description"`
}
