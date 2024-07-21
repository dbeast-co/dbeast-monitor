package plugin

import "path/filepath"

const DataFolder = "data"

var (
	NewClusterFile                      = filepath.Join(DataFolder, "new_cluster.json")
	DataSourceTemplatesFolder           = filepath.Join(DataFolder, "data_source_templates")
	LogstashTemplatesFolder             = filepath.Join(DataFolder, "logstash_config")
	EsComponentsTemplatesFolder         = filepath.Join(DataFolder, "es_components")
	EsILMTemplatesFolder                = filepath.Join(EsComponentsTemplatesFolder, "ilm_policies")
	EsIndexComponentsTemplatesFolder    = filepath.Join(EsComponentsTemplatesFolder, "component_templates")
	EsIndexTemplatesTemplatesFolder     = filepath.Join(EsComponentsTemplatesFolder, "index_templates")
	LogstashConfigurationsFolderLinux   = "/etc/logstash/conf.d"
	LogstashConfigurationsFolderWindows = "c:\\test2\\logstash-8.12.2-windows-x86_64\\logstash-8.12.2\\config"
)