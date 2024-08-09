package plugin

import (
	"encoding/json"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"io"
	"net/http"
	"strings"
)

var NewCluster Cluster

type Credentials struct {
	Host                  string `json:"host"`
	AuthenticationEnabled bool   `json:"authentication_enabled"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	Status                string `json:"status"`
}

type EnvironmentConfig struct {
	Prod struct {
		Elasticsearch Credentials `json:"elasticsearch"`
		Kibana        Credentials `json:"kibana"`
	} `json:"prod"`
	Mon struct {
		Elasticsearch Credentials `json:"elasticsearch"`
	} `json:"mon"`
}

var GrafanaDataSourcesMap = make(map[string]interface{})

var ESComponentTemplatesMap = make(map[string]string)

var ESIndexTemplatesMap = make(map[string]string)

var ESILMTemplatesMap = make(map[string]string)

func (a *App) NewClusterHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	newClusterForSend, err := json.MarshalIndent(NewCluster, "", "")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": "Failed to marshal cluster template"})
		return
	}
	response.Write(newClusterForSend)
}

func (a *App) SaveClusterHandler(response http.ResponseWriter, request *http.Request) {

	ctxLogger := log.DefaultLogger.FromContext(request.Context())
	ctxLogger.Info("Got request for the new cluster save")
	response.Header().Add("Content-Type", "application/json")

	var environmentConfig EnvironmentConfig
	if err := json.NewDecoder(request.Body).Decode(&environmentConfig); err != nil {
		log.DefaultLogger.Error("Failed to decode JSON data: " + err.Error())
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	defer request.Body.Close()

	clusterNameProd, uidProd, err := GetClusterInfo(environmentConfig.Prod.Elasticsearch)
	if err != nil {
		log.DefaultLogger.Error("Error while receiving cluster info: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	updatedTemplates := UpdateDataSourceTemplates(environmentConfig, clusterNameProd, uidProd)

	updatedTemplatesJSON, err := json.MarshalIndent(updatedTemplates, "", "")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(updatedTemplatesJSON)
}

func (a *App) AddClusterHandler(response http.ResponseWriter, request *http.Request) {

	ctxLogger := log.DefaultLogger.FromContext(request.Context())
	ctxLogger.Info("Got request for the new cluster save")
	response.Header().Add("Content-Type", "application/json")

	var cluster Cluster
	if err := json.NewDecoder(request.Body).Decode(&cluster); err != nil {
		log.DefaultLogger.Error("Failed to decode JSON data: " + err.Error())
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	defer request.Body.Close()

	clusterNameProd, uidProd, err := GetClusterInfo(cluster.ClusterConnectionSettings.Prod.Elasticsearch)
	if err != nil {
		log.DefaultLogger.Error("Error while receiving cluster info: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	//TODO Add validation for the templates already exists
	err = SendILMToMonitoringCluster(cluster.ClusterConnectionSettings.Mon.Elasticsearch)
	if err != nil {
		log.DefaultLogger.Error("Error while the ILM policy injection: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	//TODO Add validation for the templates already exists
	err = SendComponentTemplatesToMonitoringCluster(cluster.ClusterConnectionSettings.Mon.Elasticsearch)
	if err != nil {
		log.DefaultLogger.Error("Error while the Component template injection: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	//TODO Add validation for the templates already exists
	err = SendIndexTemplatesToMonitoringCluster(cluster.ClusterConnectionSettings.Mon.Elasticsearch)
	if err != nil {
		log.DefaultLogger.Error("Error while the Index template injection: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	err = SaveLogstashConfigurationFiles(cluster, ctxLogger)

	UpdatedTemplates := UpdateDataSourceTemplates(cluster.ClusterConnectionSettings, clusterNameProd, uidProd)

	updatedTemplatesJSON, err := json.MarshalIndent(UpdatedTemplates, "", "")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(updatedTemplatesJSON)
}

func (a *App) DeleteClusterHandler(response http.ResponseWriter, request *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(request.Context())
	ctxLogger.Info("Got request for the cluster delete")

	// Read the request body as a string
	clusterId, err := io.ReadAll(request.Body)
	if err != nil {
		ctxLogger.Error("Failed to read request body: " + err.Error())
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid request payload"))
		return
	}

	// Ensure request body is closed
	defer request.Body.Close()

	// Convert clusterId from byte slice to string
	clusterIdStr := string(clusterId)
	err = DeleteFolder(LogstashConfigurationsFolder+"/"+clusterIdStr+"*", ctxLogger)

	err = DeleteTextInFile(LogstashConfigurationsFolder+"/pipelines.yml", "### Configuration files for the cluster: "+clusterIdStr+", ",
		"### Configuration files for the cluster",
		ctxLogger)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	} else {
		response.WriteHeader(http.StatusOK)
	}
}

func UpdateDataSourceTemplates(environmentConfig EnvironmentConfig, clusterNameProd string, uidProd string) map[string]interface{} {
	var UpdatedTemplates = make(map[string]interface{})
	for name, template := range GrafanaDataSourcesMap {
		clonedTemplates := CloneObject(template)
		switch {
		case strings.HasPrefix(name, "json_api_datasource_elasticsearch_mon"):
			UpdateJsonTemplateValues(clonedTemplates, environmentConfig.Mon.Elasticsearch, clusterNameProd, uidProd)
			break
		case strings.HasPrefix(name, "json_api_datasource_elasticsearch_prod"):
			UpdateJsonTemplateValues(clonedTemplates, environmentConfig.Prod.Elasticsearch, clusterNameProd, uidProd)
			break
		case strings.HasPrefix(name, "json_api_datasource_kibana"):
			UpdateJsonTemplateValues(clonedTemplates, environmentConfig.Prod.Kibana, clusterNameProd, uidProd)
			break
		case strings.HasPrefix(name, "elasticsearch_datasource"):
			UpdateElasticsearchTemplateValues(clonedTemplates, environmentConfig.Mon.Elasticsearch, clusterNameProd, uidProd)
			break
		default:
		}
		UpdatedTemplates[name] = clonedTemplates
	}
	return UpdatedTemplates
}

/*
UpdateJsonTemplateValues updates the provided JSON template (json_api_datasource) with information from credentials.
It modifies the template name, UID, URL, basic authentication settings, and TLS skip verification based on the provided credentials.
*/
func UpdateJsonTemplateValues(clonedTemplates interface{}, credentials Credentials, clusterName string, uid string) {
	if OneClonedTemplate, ok := clonedTemplates.(map[string]interface{}); ok {

		OneClonedTemplate["name"] = OneClonedTemplate["name"].(string) + clusterName + "--" + uid
		OneClonedTemplate["uid"] = OneClonedTemplate["uid"].(string) + clusterName + "--" + uid

		OneClonedTemplate["url"] = credentials.Host
		OneClonedTemplate["basicAuth"] = credentials.AuthenticationEnabled

		if OneClonedTemplate["basicAuth"] == true {
			OneClonedTemplate["basicAuthUser"] = credentials.Username
			OneClonedTemplate["secureJsonData"].(map[string]interface{})["basicAuthPassword"] = credentials.Password
		}

		if url, ok := OneClonedTemplate["url"].(string); ok {
			if strings.Contains(url, "https") {
				OneClonedTemplate["jsonData"].(map[string]interface{})["tlsSkipVerify"] = true
			}
		}
	}
}

/*
UpdateElasticsearchTemplateValues updates the provided JSON template (elasticsearch_datasource) with information from  credentials.
It modifies the template name, UID, URL, basic authentication settings, and TLS skip verification based on the provided credentials.
*/
func UpdateElasticsearchTemplateValues(clonedTemplates interface{}, credentials Credentials, clusterName string, uid string) {
	if OneClonedTemplate, ok := clonedTemplates.(map[string]interface{}); ok {

		if database, ok := OneClonedTemplate["database"].(string); ok {
			database = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(database, "*", ""), "?", ""), ",", ""), ".", "")
			clusterName = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(clusterName, "*", ""), "?", ""), ",", ""), ".", "")

			OneClonedTemplate["name"] = OneClonedTemplate["name"].(string) + database + "--" + clusterName + "--" + uid
			OneClonedTemplate["uid"] = OneClonedTemplate["uid"].(string) + database + "--" + clusterName + "--" + uid

			OneClonedTemplate["url"] = credentials.Host
			OneClonedTemplate["basicAuth"] = credentials.AuthenticationEnabled

			if OneClonedTemplate["basicAuth"] == true {
				OneClonedTemplate["basicAuthUser"] = credentials.Username
				OneClonedTemplate["secureJsonData"].(map[string]interface{})["basicAuthPassword"] = credentials.Password
			}

			if url, ok := OneClonedTemplate["url"].(string); ok {
				if strings.Contains(url, "https") {
					OneClonedTemplate["jsonData"].(map[string]interface{})["tlsSkipVerify"] = true
				}
			}
		}
	}
}

func SendILMToMonitoringCluster(credentials Credentials) error {
	log.DefaultLogger.Info("ILM policies ingest")
	for templateName, templateContent := range ESILMTemplatesMap {
		log.DefaultLogger.Debug("Inject template: ", templateName, " To the cluster: ", credentials.Host)
		log.DefaultLogger.Debug("Template content: ", templateContent)
		_, err := SendILMToCluster(credentials, templateName, templateContent)
		if err != nil {
			return err
		}
	}
	return nil
}

func SendComponentTemplatesToMonitoringCluster(credentials Credentials) error {
	log.DefaultLogger.Info("Components templates ingest")
	for templateName, templateContent := range ESComponentTemplatesMap {
		log.DefaultLogger.Debug("Inject template: ", templateName, " To the cluster: ", credentials.Host)
		log.DefaultLogger.Debug("Template content: ", templateContent)
		_, err := SendComponentTemplateToCluster(credentials, templateName, templateContent)
		if err != nil {
			return err
		}
	}
	return nil
}

func SendIndexTemplatesToMonitoringCluster(credentials Credentials) error {
	log.DefaultLogger.Info("Index templates ingest")
	for templateName, templateContent := range ESIndexTemplatesMap {
		log.DefaultLogger.Debug("Inject template: ", templateName, " To the cluster: ", credentials.Host)
		log.DefaultLogger.Debug("Template content: ", templateContent)
		_, err := SendIndexTemplateToCluster(credentials, templateName, templateContent)
		if err != nil {
			return err
		}
	}
	return nil
}
