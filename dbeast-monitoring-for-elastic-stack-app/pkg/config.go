package main

import (
	"os"
)

const (
	DataFolder                         = "data" + string(os.PathSeparator)
	DataSourceTemplatesFolder          = DataFolder + "data_source_templates"
	LogstashTemplatesFolder            = DataFolder + "logstash_config"
	LogstashConfigurationsFileListFile = DataFolder + "logstash_files.json"
)
