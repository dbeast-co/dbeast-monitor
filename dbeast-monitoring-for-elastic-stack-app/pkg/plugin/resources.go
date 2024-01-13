package plugin

import (
	"encoding/json"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"net/http"
	"strings"
)

// /api/plugins/app-with-backend/resources/ping

// handlePing is an example HTTP GET resource that returns a {"message": "ok"} JSON response.
//func (a *App) handlePing(w http.ResponseWriter, req *http.Request) {
//	ctxLogger := log.DefaultLogger.FromContext(req.Context())
//	w.Header().Add("Content-Type", "application/json")
//	if _, err := w.Write([]byte(`{"message": "Мяу"}`)); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	ctxLogger.Info("ping received")
//	w.WriteHeader(http.StatusOK)
//}

// handleEcho is an example HTTP POST resource that accepts a JSON with a "message" key and
// returns to the client whatever it is sent.
//func (a *App) handleEcho(w http.ResponseWriter, req *http.Request) {
//	if req.Method != http.MethodPost {
//		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
//		return
//	}
//	var body struct {
//		Message string `json:"message"`
//	}
//	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	w.Header().Add("Content-Type", "application/json")
//	if err := json.NewEncoder(w).Encode(body); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//}

/*
	TestStatusHandler handles HTTP requests to retrieve and update the status data based on the provided environment configuration.

It takes a http.ResponseWriter and http.Request as input, decodes the request body to extract environment configuration,
updates the status and sends the updated status data in JSON format as an HTTP response.
*/
func (a *App) TestStatusHandler(w http.ResponseWriter, req *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(req.Context())
	w.Header().Add("Content-Type", "application/json")

	var EnvironmentConfig EnvironmentConfig
	if err := json.NewDecoder(req.Body).Decode(&EnvironmentConfig); err != nil {
		log.DefaultLogger.Warn("Failed to decode JSON data: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid request payload"})
		return
	}
	defer req.Body.Close()

	var statusData StatusData

	statusData.Prod.Elasticsearch = UpdateStatusForType(&EnvironmentConfig.Prod.Elasticsearch)
	statusData.Mon.Elasticsearch = UpdateStatusForType(&EnvironmentConfig.Mon.Elasticsearch)

	statusDataJSON, err := json.MarshalIndent(statusData, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Failed to marshal status data"})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(statusDataJSON)
	ctxLogger.Info("Status data received")
}

func (a *App) SaveHandler(w http.ResponseWriter, req *http.Request) {

	ctxLogger := log.DefaultLogger.FromContext(req.Context())
	ctxLogger.Info("Got request for the new cluster save")
	w.Header().Add("Content-Type", "application/json")

	var environmentConfig EnvironmentConfig
	if err := json.NewDecoder(req.Body).Decode(&environmentConfig); err != nil {
		log.DefaultLogger.Warn("Failed to decode JSON data: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid request payload"})
		return
	}
	defer req.Body.Close()

	var UpdatedTemplates = make(map[string]interface{})
	clusterNameMon, uidMon := FetchClusterInfo(environmentConfig.Mon.Elasticsearch)
	clusterNameProd, uidProd := FetchClusterInfo(environmentConfig.Prod.Elasticsearch)
	//clusterNameKibana, uidKibana := UpdateNameAndUid(environmentConfig.Prod.Kibana)

	for name, template := range TemplatesMap {
		clonedTemplates := CloneTemplate(template)

		switch {
		case strings.HasPrefix(name, "json_api_datasource_elasticsearch_mon"):
			UpdateJsonTemplateValues(clonedTemplates, environmentConfig.Mon.Elasticsearch, clusterNameMon, uidMon)
			break

		case strings.HasPrefix(name, "json_api_datasource_elasticsearch_prod"):
			UpdateJsonTemplateValues(clonedTemplates, environmentConfig.Prod.Elasticsearch, clusterNameProd, uidProd)
			break
		case strings.HasPrefix(name, "json_api_datasource_kibana"):
			UpdateJsonTemplateValues(clonedTemplates, environmentConfig.Prod.Kibana, clusterNameMon, uidMon)
			break
		case strings.HasPrefix(name, "elasticsearch_datasource"):
			UpdateElasticsearchTemplateValues(clonedTemplates, environmentConfig.Mon.Elasticsearch, clusterNameMon, uidMon)
			break
		default:
		}
		UpdatedTemplates[name] = clonedTemplates

	}

	//SendTemplateToServer(UpdatedTemplates)

	updatedTemplatesJSON, err := json.MarshalIndent(UpdatedTemplates, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Failed to marshal updated templates"})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(updatedTemplatesJSON)
	ctxLogger.Info("Updated templates sent", string(updatedTemplatesJSON))
}

// registerRoutes takes a *http.ServeMux and registers some HTTP handlers.
func (a *App) registerRoutes(mux *http.ServeMux) {
	//mux.HandleFunc("/ping", a.handlePing)
	//mux.HandleFunc("/echo", a.handleEcho)
	mux.HandleFunc("/test_cluster", a.TestStatusHandler)
	mux.HandleFunc("/save", a.SaveHandler)
}
