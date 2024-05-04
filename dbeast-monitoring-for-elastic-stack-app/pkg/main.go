package main

import (
	"os"
	"strings"

	"github.com/grafana/app-with-backend/pkg/plugin"
	"github.com/grafana/grafana-plugin-sdk-go/backend/app"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func main() {
	// Start listening to requests sent from Grafana. This call is blocking so
	// it won't finish until Grafana shuts down the process or the plugin choose
	// to exit by itself using os.Exit. Manage automatically manages life cycle
	// of app instances. It accepts app instance factory as first
	// argument. This factory will be automatically called on incoming request
	// from Grafana to create different instances of `App` (per plugin
	// ID).
	//DATA_FOLDER := "resources" + string(os.PathSeparator)
	//DATA_SOURCE_TEMPLATES_FOLDER := DATA_FOLDER + "data_source_templates"
	//LOGSTASH_TEMPLATES_FOLDER := DATA_FOLDER + "logstash_config"
	//LOGSTASH_FILES_LIST_FILE := DATA_FOLDER + "logstash_files"
	ctxLogger := log.DefaultLogger
	ctxLogger.Info("The app path: " + os.Args[0])
	lastIndex := strings.LastIndex(os.Args[0], "gpx_app-dbeast-dbeastmonitor-app")
	//lastIndex := strings.LastIndex(os.Args[0][:lastIndex1], "/")
	//string(os.PathSeparator) +
	err := plugin.LoadTemplatesFromFolder(os.Args[0][:lastIndex] + DataSourceTemplatesFolder)
	if err != nil {
		return
	}
	err = plugin.LoadLogstashConfigFromFolder(os.Args[0][:lastIndex] + LogstashTemplatesFolder)
	if err != nil {
		return
	}
	err = plugin.LoadLogstashConfigurationFileList(os.Args[0][:lastIndex] + LogstashConfigurationsFileListFile)
	if err != nil {
		return
	}

	if err := app.Manage("dbeast-dbeastmonitor-app", plugin.NewApp, app.ManageOpts{}); err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
