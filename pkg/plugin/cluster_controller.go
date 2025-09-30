package plugin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func (a *App) NewClusterHandler(response http.ResponseWriter, request *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(request.Context())
	ctxLogger.Info("Got request for the new cluster")

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	newClusterForSend, err := json.MarshalIndent(NewCluster, "", "")
	if err != nil {
		HTTPErrorGenerator(response, err, "Failed to normalize data, for new cluster request: ", http.StatusInternalServerError, ctxLogger)
		return
	}
	_, err = response.Write(newClusterForSend)
	if err != nil {
		log.DefaultLogger.Error("Can't write to the response: " + err.Error())
		return
	}
}

func (a *App) SaveClusterHandler(response http.ResponseWriter, request *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(request.Context())
	ctxLogger.Info("Got request for the new cluster save")
	response.Header().Add("Content-Type", "application/json")

	DeferHandler(response, request, ctxLogger)

	var environmentConfig EnvironmentConfig
	if err := json.NewDecoder(request.Body).Decode(&environmentConfig); err != nil {
		HTTPErrorGenerator(response, err, "Failed to decode JSON data for cluster save request: ", http.StatusInternalServerError, ctxLogger)
		return
	}
	sanitizeEnvironmentConfig(&environmentConfig)

	client, err := CreateHTTPClient(environmentConfig.Prod.Elasticsearch)
	if err != nil {
		HTTPErrorGenerator(response, err, "Error while creating HTTP client for cluster save request: ", http.StatusInternalServerError, ctxLogger)
		return
	}

	clusterNameProd, uidProd, err := GetClusterInfo(client, environmentConfig.Prod.Elasticsearch.Host)
	if err != nil {
		HTTPErrorGenerator(response, err, "Error while receiving cluster info for cluster save request: ", http.StatusInternalServerError, ctxLogger)
		return
	}

	updatedTemplates := UpdateGrafanaDataSourceTemplates(environmentConfig, clusterNameProd, uidProd)

	updatedTemplatesJSON, err := json.MarshalIndent(updatedTemplates, "", "")
	if err != nil {
		HTTPErrorGenerator(response, err, "Error while templates update, for cluster save request: ", http.StatusInternalServerError, ctxLogger)
		return
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write(updatedTemplatesJSON)
	if err != nil {
		log.DefaultLogger.Error("Can't write to the response: " + err.Error())
		return
	}
}

func (a *App) DeployElasticsearchConfigurations(response http.ResponseWriter, request *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(request.Context())
	ctxLogger.Info("Got request for the Elasticsearch components deployment")

	response.Header().Add("Content-Type", "application/json")

	DeferHandler(response, request, ctxLogger)

	var project Project
	if err := json.NewDecoder(request.Body).Decode(&project); err != nil {
		HTTPErrorGenerator(response, err, "Failed to decode JSON data for deploy ES configuration request: ", http.StatusBadRequest, ctxLogger)
		return
	}
	sanitizeEnvironmentConfig(&project.ClusterConnectionSettings)

	client, err := CreateHTTPClient(project.ClusterConnectionSettings.Mon.Elasticsearch)
	if err != nil {
		HTTPErrorGenerator(response, err, "Error while creating HTTP client for deploy ES configuration request: ", http.StatusInternalServerError, ctxLogger)
		return
	}

	for _, injectType := range project.MonitoringClusterInjection {
		if injectType.Id == "ilm_policies_injection" && injectType.IsChecked {
			err := SendILMToMonitoringCluster(client, project.ClusterConnectionSettings.Mon.Elasticsearch.Host)
			if err != nil {
				HTTPErrorGenerator(response, err, "Error while the ILM policy injection for deploy ES configuration request: ", http.StatusInternalServerError, ctxLogger)
				return
			}
		}

		if injectType.Id == "templates_injection" && injectType.IsChecked {
			err := SendComponentTemplatesToMonitoringCluster(client, project.ClusterConnectionSettings.Mon.Elasticsearch.Host)
			if err != nil {
				HTTPErrorGenerator(response, err, "Error while the Component template injection for deploy ES configuration request: ", http.StatusInternalServerError, ctxLogger)
				return
			}
			err = SendIndexTemplatesToMonitoringCluster(client, project.ClusterConnectionSettings.Mon.Elasticsearch.Host)
			if err != nil {
				HTTPErrorGenerator(response, err, "Error while the Index template injection for deploy ES configuration request: ", http.StatusInternalServerError, ctxLogger)
				return
			}
		}

		if injectType.Id == "create_first_indices" && injectType.IsChecked {
			err := SendFirstIndicesToMonitoringCluster(client, project.ClusterConnectionSettings.Mon.Elasticsearch.Host)
			if err != nil {
				HTTPErrorGenerator(response, err, "Error while the First indices injection for deploy ES configuration request: ", http.StatusInternalServerError, ctxLogger)
				return
			}
		}
	}
}

func (a *App) AddClusterHandlerToGrafana(response http.ResponseWriter, request *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(request.Context())
	ctxLogger.Info("Got request for the new cluster save")
	response.Header().Add("Content-Type", "application/json")

	DeferHandler(response, request, ctxLogger)

	var project Project
	if err := json.NewDecoder(request.Body).Decode(&project); err != nil {
		HTTPErrorGenerator(response, err, "Failed to decode JSON data for add new cluster to Grafana request: ", http.StatusBadRequest, ctxLogger)
		return
	}
	sanitizeEnvironmentConfig(&project.ClusterConnectionSettings)

	client, err := CreateHTTPClient(project.ClusterConnectionSettings.Prod.Elasticsearch)
	if err != nil {
		HTTPErrorGenerator(response, err, "Error while creating HTTP client for add new cluster to Grafana request: ", http.StatusInternalServerError, ctxLogger)
		return
	}

	clusterNameProd, clusterId, err := GetClusterInfo(client, project.ClusterConnectionSettings.Prod.Elasticsearch.Host)
	if err != nil {
		HTTPErrorGenerator(response, err, "Error while receiving cluster info for add new cluster to Grafana request: ", http.StatusInternalServerError, ctxLogger)
		return
	}

	//Update Grafana datasource templates with actual values and return them to the Frontend for the future ingesting into Grafana
	UpdatedTemplates := UpdateGrafanaDataSourceTemplates(project.ClusterConnectionSettings, clusterNameProd, clusterId)
	updatedTemplatesJSON, err := json.MarshalIndent(UpdatedTemplates, "", "")
	if err != nil {
		HTTPErrorGenerator(response, err, "Error while the updated templates parsing for add new cluster to Grafana request: ", http.StatusInternalServerError, ctxLogger)
		return
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write(updatedTemplatesJSON)
	if err != nil {
		log.DefaultLogger.Error("Can't write to the response for add new cluster to Grafana request: " + err.Error())
		return
	}
}

func (a *App) DeleteClusterHandler(response http.ResponseWriter, request *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(request.Context())
	clusterId := request.URL.Path[len("/delete_cluster/"):]
	ctxLogger.Info("Got request for the cluster delete. Cluster ID: " + clusterId)
	response.WriteHeader(http.StatusOK)
	_, err := response.Write([]byte(`{"status":"ok"}`))
	if err != nil {
		log.DefaultLogger.Error("Can't write to the response for delete cluster request: " + err.Error())
		return
	}
}

func SendILMToMonitoringCluster(client *http.Client, host string) error {
	log.DefaultLogger.Info("ILM policies ingest")
	for templateName, templateContent := range ESILMTemplatesMap {
		log.DefaultLogger.Debug("Inject template: ", templateName, " To the cluster: ", host)
		log.DefaultLogger.Debug("Template content: ", templateContent)
		_, err := SendILMToCluster(client, host, templateName, templateContent)
		if err != nil {
			return err
		}
	}
	return nil
}

func SendComponentTemplatesToMonitoringCluster(client *http.Client, host string) error {
	log.DefaultLogger.Info("Components templates ingest")
	for templateName, templateContent := range ESComponentTemplatesMap {
		log.DefaultLogger.Debug("Inject template: ", templateName, " To the cluster: ", host)
		log.DefaultLogger.Debug("Template content: ", templateContent)
		_, err := SendComponentTemplateToCluster(client, host, templateName, templateContent)
		if err != nil {
			return err
		}
	}
	return nil
}

func SendFirstIndicesToMonitoringCluster(client *http.Client, host string) error {
	log.DefaultLogger.Info("First indices ingest")
	for indexName, templateContent := range ESFirstIndicesTemplatesMap {
		log.DefaultLogger.Debug("Inject template: " + indexName + " To the cluster: " + host)
		log.DefaultLogger.Debug("Template content: " + templateContent)
		isIndexExists, err := CheckIsIndexExists(client, host, indexName)
		if err != nil {
			return err
		}
		if isIndexExists {
			log.DefaultLogger.Info("An index " + indexName + " already exists. Send rollover command")
			rolloverAlias, _ := ExtractRolloverAlias(templateContent)
			log.DefaultLogger.Info("Rollover alias: " + rolloverAlias)
			_, err = SendRolloverCommand(client, host, indexName)
			if err != nil {
				return err
			}
		} else {
			log.DefaultLogger.Info("An index " + indexName + " doesn't isIndexExists. Send create new index command")
			_, err = SendFirstIndexToCluster(client, host, indexName, templateContent)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SendIndexTemplatesToMonitoringCluster(client *http.Client, host string) error {
	log.DefaultLogger.Info("Index templates ingest")
	for templateName, templateContent := range ESIndexTemplatesMap {
		log.DefaultLogger.Debug("Inject template: ", templateName, " To the cluster: ", host)
		log.DefaultLogger.Debug("Template content: ", templateContent)
		_, err := SendIndexTemplateToCluster(client, host, templateName, templateContent)
		if err != nil {
			return err
		}
	}
	return nil
}

func ExtractRolloverAlias(jsonData string) (string, error) {
	var firstIndexTemplate map[string]map[string]map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &firstIndexTemplate)
	if err != nil {
		return "", err
	}
	aliasesMap, ok := firstIndexTemplate["aliases"]
	if !ok {
		return "", fmt.Errorf("key 'aliases' not found in JSON")
	}

	for aliasKey := range aliasesMap {
		return aliasKey, nil
	}
	return "", fmt.Errorf("key 'aliases' not found in JSON")
}

func UpdateTestDataTemplateValues(clonedTemplates interface{}, clusterName string, uid string) {
	if OneClonedTemplate, ok := clonedTemplates.(map[string]interface{}); ok {

		clusterName = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(clusterName, "*", ""), "?", ""), ",", ""), ".", "")

		OneClonedTemplate["name"] = OneClonedTemplate["name"].(string) + clusterName + "--" + uid
		OneClonedTemplate["uid"] = uid
	}
}

func UpdateJsonTemplateValues(clonedTemplates interface{}, credentials Credentials, clusterName string, uid string) {
	if OneClonedTemplate, ok := clonedTemplates.(map[string]interface{}); ok {

		clusterName = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(clusterName, "*", ""), "?", ""), ",", ""), ".", "")

		OneClonedTemplate["name"] = OneClonedTemplate["name"].(string) + clusterName + "--" + uid
		//OneClonedTemplate["uid"] = OneClonedTemplate["uid"].(string) + clusterName + "--" + uid

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

func UpdateElasticsearchTemplateValues(clonedTemplates interface{}, credentials Credentials, clusterName string, uid string) {
	if OneClonedTemplate, ok := clonedTemplates.(map[string]interface{}); ok {

		if database, ok := OneClonedTemplate["database"].(string); ok {
			database = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(database, "*", ""), "?", ""), ",", ""), ".", "")
			clusterName = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(clusterName, "*", ""), "?", ""), ",", ""), ".", "")

			OneClonedTemplate["name"] = OneClonedTemplate["name"].(string) + database + "--" + clusterName + "--" + uid

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
		case strings.HasPrefix(name, "testdata_datasource"):
			UpdateTestDataTemplateValues(clonedTemplates, clusterNameProd, uidProd)
			break
		default:
		}
		UpdatedTemplates[name] = clonedTemplates
	}
	return UpdatedTemplates
}
