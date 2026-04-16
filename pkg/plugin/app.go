package plugin

import (
	"context"
	"net/http"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/httpclient"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/backend/resource/httpadapter"
)

var (
	_ backend.CallResourceHandler   = (*App)(nil)
	_ instancemgmt.InstanceDisposer = (*App)(nil)
	_ backend.CheckHealthHandler    = (*App)(nil)
)

type App struct {
	backend.CallResourceHandler
	httpClientOptions httpclient.Options
}

const applicationVersion string = "OnPrem"

func NewApp(ctx context.Context, settings backend.AppInstanceSettings) (instancemgmt.Instance, error) {
	var app App

	opts, err := settings.HTTPClientOptions(ctx)
	if err != nil {
		log.DefaultLogger.Warn("Failed to get HTTP client options from Grafana settings, proxy will not be used: " + err.Error())
	} else {
		app.httpClientOptions = opts
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
