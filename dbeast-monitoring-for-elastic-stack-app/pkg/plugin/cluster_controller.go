package plugin

import (
	"encoding/json"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"net/http"
	"strings"
)

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
	sanitizeEnvironmentConfig(&environmentConfig)
	defer request.Body.Close()

	clusterNameProd, uidProd, err := GetClusterInfo(environmentConfig.Prod.Elasticsearch)
	if err != nil {
		log.DefaultLogger.Error("Error while receiving cluster info: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	updatedTemplates := UpdateGrafanaDataSourceTemplates(environmentConfig, clusterNameProd, uidProd)

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

	var cluster Project
	if err := json.NewDecoder(request.Body).Decode(&cluster); err != nil {
		log.DefaultLogger.Error("Failed to decode JSON data: " + err.Error())
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	sanitizeEnvironmentConfig(&cluster.ClusterConnectionSettings)
	defer request.Body.Close()

	clusterNameProd, clusterId, err := GetClusterInfo(cluster.ClusterConnectionSettings.Prod.Elasticsearch)
	if err != nil {
		log.DefaultLogger.Error("Error while receiving cluster info: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	for _, injectType := range cluster.MonitoringClusterInjection {
		if injectType.Id == "ilm_policies_injection" && injectType.IsChecked {
			err = SendILMToMonitoringCluster(cluster.ClusterConnectionSettings.Mon.Elasticsearch)
			if err != nil {
				log.DefaultLogger.Error("Error while the ILM policy injection: " + err.Error())
				response.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
				return
			}
		}

		if injectType.Id == "templates_injection" && injectType.IsChecked {
			err = SendComponentTemplatesToMonitoringCluster(cluster.ClusterConnectionSettings.Mon.Elasticsearch)
			if err != nil {
				log.DefaultLogger.Error("Error while the Component template injection: " + err.Error())
				response.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
				return
			}
			err = SendIndexTemplatesToMonitoringCluster(cluster.ClusterConnectionSettings.Mon.Elasticsearch)
			if err != nil {
				log.DefaultLogger.Error("Error while the Index template injection: " + err.Error())
				response.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
				return
			}
		}

		if injectType.Id == "create_first_indices" && injectType.IsChecked {
			err = SendFirstIndicesToMonitoringCluster(cluster.ClusterConnectionSettings.Mon.Elasticsearch)
			if err != nil {
				log.DefaultLogger.Error("Error while the First indices injection: " + err.Error())
				response.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
				return
			}
		}
	}

	err = DeleteTextBlockInFile(GrafanaLogstashConfigurationsFolder+"/pipelines.yml", "### Configuration files for the cluster Id: "+clusterId,
		"### Configuration files for the cluster Id: ",
		ctxLogger)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
	err = DeleteFolder(GrafanaLogstashConf_dConfigurationsFolder, clusterId, ctxLogger)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}

	err = SaveLogstashConfigurationFiles(cluster, ctxLogger)

	UpdatedTemplates := UpdateGrafanaDataSourceTemplates(cluster.ClusterConnectionSettings, clusterNameProd, clusterId)

	//Update Grafana datasource templates with actual values and return them to the Frontend for the future ingest into Grafana
	updatedTemplatesJSON, err := json.MarshalIndent(UpdatedTemplates, "", "")
	if err != nil {
		log.DefaultLogger.Error("Error while the updated templates parsing: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(updatedTemplatesJSON)
}

func (a *App) DeleteClusterHandler(response http.ResponseWriter, request *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(request.Context())

	clusterId := request.URL.Path[len("/delete_cluster/"):]
	ctxLogger.Info("Got request for the cluster delete. Cluster ID: " + clusterId)

	err := DeleteTextBlockInFile(GrafanaLogstashConfigurationsFolder+"/pipelines.yml", "### Configuration files for the cluster Id: "+clusterId,
		"### Configuration files for the cluster Id: ",
		ctxLogger)

	if err != nil {
		ctxLogger.Error("Error while clean pipeline.yml file: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		_, err := response.Write([]byte(err.Error()))
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			ctxLogger.Error("Can't write to the response: " + err.Error())
		}
	}
	err = DeleteFolder(GrafanaLogstashConf_dConfigurationsFolder, clusterId, ctxLogger)

	if err != nil {
		ctxLogger.Error("Error while delete folder: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		_, err := response.Write([]byte(err.Error()))
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			ctxLogger.Error("Can't write to the response: " + err.Error())
		}
	} else {
		ctxLogger.Info("We delete all files for the cluster")
		response.WriteHeader(http.StatusOK)
		_, err := response.Write([]byte(`{"status":"ok"}`))
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			ctxLogger.Error("Can't write to the response: " + err.Error())
		}
	}
}

func UpdateGrafanaDataSourceTemplates(environmentConfig EnvironmentConfig, clusterNameProd string, uidProd string) map[string]interface{} {
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

		clusterName = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(clusterName, "*", ""), "?", ""), ",", ""), ".", "")

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

func SendFirstIndicesToMonitoringCluster(credentials Credentials) error {
	log.DefaultLogger.Info("First indices ingest")
	for templateName, templateContent := range ESFirstIndicesTemplatesMap {
		log.DefaultLogger.Debug("Inject template: ", templateName, " To the cluster: ", credentials.Host)
		log.DefaultLogger.Debug("Template content: ", templateContent)
		exists, err := CheckIsIndexExists(credentials, templateName)
		if err != nil {
			return err
		}
		log.DefaultLogger.Info("Is index exists response: ", exists)
		if !exists {
			_, err = SendFirstIndexToCluster(credentials, templateName, templateContent)
			if err != nil {
				return err
			}
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
