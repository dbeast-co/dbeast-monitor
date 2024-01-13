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
	FOLDER_PATH := "templates"
	lastIndex := strings.LastIndex(os.Args[0], string(os.PathSeparator))
	err := plugin.LoadTemplatesFromFolder(os.Args[0][:lastIndex] + string(os.PathSeparator) +
		"dbeast-dbeastmonitor-app" + string(os.PathSeparator) + FOLDER_PATH)
	if err != nil {
		return
	}

	if err := app.Manage("dbeast-dbeastmonitor-app", plugin.NewApp, app.ManageOpts{}); err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
