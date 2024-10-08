package plugin

import "path/filepath"

const DataFolder = "data"

var (
	NewClusterFile                           = filepath.Join(DataFolder, "new_cluster.json")
	DataSourceTemplatesFolder                = filepath.Join(DataFolder, "data_source_templates")
	LogstashTemplatesFolder                  = filepath.Join(DataFolder, "logstash_config")
	EsComponentsTemplatesFolder              = filepath.Join(DataFolder, "es_components")
	EsILMTemplatesFolder                     = filepath.Join(EsComponentsTemplatesFolder, "ilm_policies")
	EsIndexComponentsTemplatesFolder         = filepath.Join(EsComponentsTemplatesFolder, "component_templates")
	EsIndexTemplatesTemplatesFolder          = filepath.Join(EsComponentsTemplatesFolder, "index_templates")
	EsIndexFirstIndicesTemplatesFolder       = filepath.Join(EsComponentsTemplatesFolder, "first_indices")
	LogstashConfigurationsFolder             = "/usr/share/logstash/config"
	LogstashConfDConfigurationsFolder        = "/usr/share/logstash/config/conf.d"
	GrafanaLogstashConfigurationsFolder      = "/opt/dbeast/logstash/config"
	GrafanaLogstashConfDConfigurationsFolder = "/opt/dbeast/logstash/config/conf.d"
	//LogstashConfigurationsFolder       = "/etc/logstash"
	//LogstashConfDConfigurationsFolder  = "/etc/logstash/conf.d"
)
