package plugin

import (
	"net/http"
)

// registerRoutes takes a *http.ServeMux and registers some HTTP handlers.
func (a *App) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/test_cluster", a.TestStatusHandler)
	mux.HandleFunc("/save", a.SaveHandler)
}
