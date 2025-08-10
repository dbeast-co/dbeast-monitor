package plugin

import (
	"context"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/backend/resource/httpadapter"
	"net/http"
	"os"
	"strings"
)

var (
	_ backend.CallResourceHandler   = (*App)(nil)
	_ instancemgmt.InstanceDisposer = (*App)(nil)
	_ backend.CheckHealthHandler    = (*App)(nil)
)

type App struct {
	backend.CallResourceHandler
}

var applicationVersion string
var exists bool

func NewApp(_ context.Context, settings backend.AppInstanceSettings) (instancemgmt.Instance, error) {
	var app App
	applicationVersion, exists = os.LookupEnv("GF_PLUGIN_VERSION")
	if !exists || len(applicationVersion) == 0 {
		applicationVersion = "onprem"
	} else {
		applicationVersion = strings.ToLower(applicationVersion)
	}

	mux := http.NewServeMux()
	app.registerRoutes(mux)
	app.CallResourceHandler = httpadapter.New(mux)

	ctxLogger := log.DefaultLogger
	ctxLogger.Info(applicationVersion, " Version Installed")

	return &app, nil
}

func (a *App) Dispose() {
	// cleanup
}

func (a *App) IsContainerVersion() {

}

func (a *App) CheckHealth(_ context.Context, _ *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "ok",
	}, nil
}
