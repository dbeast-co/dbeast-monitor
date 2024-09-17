package plugin

import (
	"os"
)

const (
	DataFolder                          = "data" + string(os.PathSeparator)
	DataSourceTemplatesFolder           = DataFolder + "data_source_templates"
	LogstashTemplatesFolder             = DataFolder + "logstash_config"
	LogstashConfigurationsFileListFile  = DataFolder + "logstash_files.json"
	LogstashConfigurationsFolderLinux   = "/etc/logstash/conf.d"
	LogstashConfigurationsFolderWindows = "c:\\test2\\logstash-8.12.2-windows-x86_64\\logstash-8.12.2\\config"
)
