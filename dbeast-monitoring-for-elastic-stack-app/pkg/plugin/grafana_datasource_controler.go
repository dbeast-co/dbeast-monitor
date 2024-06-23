package plugin

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"io/ioutil"
	"net/http"
	"strings"
)

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
	//clusterNameMon, uidMon := FetchClusterInfo(environmentConfig.Mon.Elasticsearch)
	clusterNameProd, uidProd := FetchClusterInfo(environmentConfig.Prod.Elasticsearch)
	//clusterNameKibana, uidKibana := UpdateNameAndUid(environmentConfig.Prod.Kibana)

	for name, template := range TemplatesMap {
		clonedTemplates := CloneTemplate(template)

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

	//SendTemplateToServer(UpdatedTemplates)
	updatedTemplatesJSON, err := json.MarshalIndent(UpdatedTemplates, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Failed to marshal updated templates"})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(updatedTemplatesJSON)
}

func (a *App) NewClusterHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	newClusterForSend, err := json.MarshalIndent(NewCluster, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Failed to marshal cluster template"})
		return
	}
	w.Write(newClusterForSend)
}

func (a *App) GenerateLogstashMonitoringConfigurationFilesHandler(w http.ResponseWriter, req *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(req.Context())
	ctxLogger.Info("Got request for the Elasticsearch configuration files download")

	w.Header().Add("Content-Disposition", "attachment; filename=\"files.zip\"")
	w.Header().Add("Content-Type", "application/zip")

	var project Cluster

	if err := json.NewDecoder(req.Body).Decode(&project); err != nil {
		log.DefaultLogger.Warn("Failed to decode JSON data: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid request payload"})
		return
	}
	var strOut, _ = json.Marshal(project)
	ctxLogger.Info("The project: ", project)
	ctxLogger.Info("Request: " + string(strOut))
	var environmentConfig = project.ClusterConnectionSettings
	ctxLogger.Info("The Emv: ", environmentConfig)

	defer req.Body.Close()

	ctxLogger.Info("Connection settings: ", environmentConfig)
	buf := new(bytes.Buffer)

	// Create a new zip archive
	zipWriter := zip.NewWriter(buf)

	for fileName, configFile := range LSConfigs {
		if strings.HasPrefix(fileName, "dbeast-mon-es") {
			updatedConfigFileContent := UpdateLSConfigFile(configFile, environmentConfig)
			f1, err := zipWriter.Create(string(fileName) + ".conf")
			if err != nil {
				log.DefaultLogger.Error(err.Error())
			}
			// Create the first file
			_, err = f1.Write([]byte(updatedConfigFileContent))
			if err != nil {
				log.DefaultLogger.Error(err.Error())
			}
			//ctxLogger.Info("Updated file: ", fileName, " Config: ", updatedConfigFileContent)
			WriteConfigFileToDisk(ctxLogger, "c:\\test2\\grafana-9.5.10.windows-amd64\\grafana-9.5.10\\tmp\\"+fileName+".conf", updatedConfigFileContent)
		}
	}

	// Make sure to close the zip writer to flush the contents to the buffer
	err := zipWriter.Close()
	if err != nil {
		log.DefaultLogger.Error(err.Error())
	}
	// Set headers to instruct the browser to download the file as a zip

	// Write the buffer to the HTTP response
	_, err = w.Write(buf.Bytes())
	if err != nil {
		log.DefaultLogger.Error("Error writing response: ", err.Error())
	}

	//SendTemplateToServer(UpdatedTemplates)
}

func UpdateLSConfigFile(configFileContent string, environmentConfig EnvironmentConfig) string {
	_, uidProd := FetchClusterInfo(environmentConfig.Prod.Elasticsearch)
	configFileClone := strings.Clone(configFileContent)

	configFileClone = strings.ReplaceAll(configFileClone, "<PROD_HOST>", environmentConfig.Prod.Elasticsearch.Host)
	configFileClone = strings.ReplaceAll(configFileClone, "<CLUSTER_ID>", uidProd)
	if environmentConfig.Prod.Elasticsearch.AuthenticationEnabled {
		configFileClone = strings.ReplaceAll(configFileClone, "<PROD_USER>", environmentConfig.Prod.Elasticsearch.Username)
		configFileClone = strings.ReplaceAll(configFileClone, "<PROD_PASSWORD>", environmentConfig.Prod.Elasticsearch.Password)
		configFileClone = strings.ReplaceAll(configFileClone, "<PROD_SSL_ENABLED>", fmt.Sprintf("%t", strings.Contains(environmentConfig.Prod.Elasticsearch.Host, "https")))
	}

	configFileClone = strings.ReplaceAll(configFileClone, "<MON_HOST>", environmentConfig.Mon.Elasticsearch.Host)
	if environmentConfig.Mon.Elasticsearch.AuthenticationEnabled {
		configFileClone = strings.ReplaceAll(configFileClone, "<MON_USER>", environmentConfig.Mon.Elasticsearch.Username)
		configFileClone = strings.ReplaceAll(configFileClone, "<MON_PASSWORD>", environmentConfig.Mon.Elasticsearch.Password)
		if strings.Contains(environmentConfig.Mon.Elasticsearch.Host, "https") {
			configFileClone = strings.ReplaceAll(configFileClone, "<MON_SSL_ENABLED>", "true")
		} else {
			configFileClone = strings.ReplaceAll(configFileClone, "ssl_certificate_verification => false", "")
		}
	}
	return configFileClone
}

func WriteConfigFileToDisk(ctxLogger log.Logger, fileName string, content string) {
	data, err := json.Marshal(content)
	ctxLogger.Debug("Write file: ", fileName, " Content", data)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	// Save the JSON data to a file
	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	ctxLogger.Info("Object saved to file: ", fileName)

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

			OneClonedTemplate["name"] = OneClonedTemplate["name"].(string) + database + "-" + clusterName + "--" + uid
			OneClonedTemplate["uid"] = OneClonedTemplate["uid"].(string) + database + "-" + clusterName + "--" + uid

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

/*
CloneTemplate creates a deep copy of the provided JSON template.
It marshals the template into JSON and then unmarshals it into a new interface{} to create a clone.
*/
func CloneTemplate(data interface{}) interface{} {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.DefaultLogger.Warn("Failed to marshal data: " + err.Error())
		return data
	}

	var clonedTemplate interface{}
	if err := json.Unmarshal(dataBytes, &clonedTemplate); err != nil {
		log.DefaultLogger.Warn("Failed to unmarshal cloned data: " + err.Error())
		return data
	}

	return clonedTemplate
}
