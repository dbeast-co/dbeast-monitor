package main

import (
	"os"
	"strings"

	"github.com/dbeast/dbeastmonitor/pkg/plugin"
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

	ctxLogger := log.DefaultLogger
	ctxLogger.Info("The app path: " + os.Args[0])
	lastIndex := strings.LastIndex(os.Args[0], "gpx_app-dbeast-dbeastmonitor-app")
	applicationFolder := os.Args[0][:lastIndex]

	err := plugin.LoadInitData(applicationFolder)
	if err != nil {
		ctxLogger.Info("Error in the init data loading. Please fix the init files ")
		os.Exit(1)
	}

	if err = app.Manage("dbeast-dbeastmonitor-app", plugin.NewApp, app.ManageOpts{}); err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
