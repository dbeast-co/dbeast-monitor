package plugin

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/grafana/grafana-plugin-sdk-go/backend/httpclient"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

var LSConfigs = make(map[string]string)

func (a *App) DownloadElasticsearchMonitoringConfigurationFilesHandler(response http.ResponseWriter, request *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(request.Context())
	ctxLogger.Info("Got request for the Elasticsearch configuration files download")

	GenerateLogstashConfigurationFiles(response, request, false, "ESConfigurationFiles.zip", a.httpClientOptions)
}

func (a *App) DownloadLogstashMonitoringConfigurationFilesHandler(response http.ResponseWriter, request *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(request.Context())
	ctxLogger.Info("Got request for the Logstash configuration files download")

	GenerateLogstashConfigurationFiles(response, request, true, "LogstashConfigurationFiles.zip", a.httpClientOptions)
}

func GenerateLogstashConfigurationFiles(response http.ResponseWriter, request *http.Request, isLogstash bool, resultZipFileName string, httpClientOptions httpclient.Options) {
	ctxLogger := log.DefaultLogger.FromContext(request.Context())

	response.Header().Add("Content-Disposition", "attachment; filename=\""+resultZipFileName+"\"")
	response.Header().Add("Content-Type", "application/zip")

	defer DeferHandler(request, ctxLogger)

	var project Project

	if err := json.NewDecoder(request.Body).Decode(&project); err != nil {
		HTTPErrorGenerator(response, err, "Failed to decode JSON data for generate Logstash configuration files request: ", http.StatusBadRequest, ctxLogger)
		return
	}

	client, err := CreateHTTPClient(project.ClusterConnectionSettings.Prod.Elasticsearch, httpClientOptions)
	if err != nil {
		HTTPErrorGenerator(response, err, "Error while creating HTTP client for Logstash configuration files request: ", http.StatusInternalServerError, ctxLogger)
		return
	}

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	clusterName, clusterId, err := GetClusterInfo(client, project.ClusterConnectionSettings.Prod.Elasticsearch.Host)
	if err != nil {
		HTTPErrorGenerator(response, err, "Error while receiving cluster info for Logstash configuration files request: ", http.StatusInternalServerError, ctxLogger)
		return
	}

	if isLogstash {
		GenerateLSLogstashConfigurationFiles(project, clusterId, zipWriter)
	} else {
		GenerateESLogstashConfigurationFiles(project, clusterId, clusterName, zipWriter)
	}
	err = zipWriter.Close()
	if err != nil {
		HTTPErrorGenerator(response, err, "Error while the ZIP file creation for Logstash configuration files request: ", http.StatusInternalServerError, ctxLogger)
		return
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write(buf.Bytes())
	if err != nil {
		log.DefaultLogger.Error("Can't write to the response: " + err.Error())
		return
	}
}

func GenerateESLogstashConfigurationFiles(project Project, clusterId string, clusterName string, zipWriter *zip.Writer) {
	pipelineFile := "### Configuration files for the cluster: " + clusterName + ", clusterId: " + clusterId + "\n"
	for _, configFile := range project.LogstashConfigurations.EsMonitoringConfigurationFiles {
		if configFile.IsChecked {
			configFileClone := strings.Clone(LSConfigs[configFile.Id])
			configFileClone = strings.ReplaceAll(configFileClone, "<CLUSTER_ID>", clusterId)
			configFileClone = UpdateMonConnectionSettings(configFileClone, project.ClusterConnectionSettings)
			configFileClone = UpdateProdConnectionSettings(configFileClone, project.ClusterConnectionSettings)

			fileInternalPath := clusterName + "-" + clusterId + "/" + configFile.Id
			//fileInternalPath := filepath.Join(clusterName+"-"+clusterId, configFile.Id)

			WriteFileToZip(zipWriter, fileInternalPath, configFileClone)

			pipelineId := strings.ReplaceAll(configFile.Id, ".conf", "") + "-" + clusterName + "-" + clusterId
			pipelineFile += fmt.Sprintf("- pipeline.id: %s\n", pipelineId)
			pipelineFile += fmt.Sprintf("  path.config: \"/etc/logstash/conf.d/%s\"\n\n", fileInternalPath)
		}
	}
	WriteFileToZip(zipWriter, "pipelines.yml", pipelineFile)
}

func GenerateLSLogstashConfigurationFiles(project Project, clusterId string, zipWriter *zip.Writer) {
	for _, logstashHost := range project.LogstashConfigurations.LogstashMonitoringConfigurationFiles.Hosts {
		pipelineFile := "### Configuration files for the Logstash monitoring\n"
		for _, configFile := range project.LogstashConfigurations.LogstashMonitoringConfigurationFiles.Configurations {
			if configFile.IsChecked {
				configFileClone := strings.Clone(LSConfigs[configFile.Id])
				configFileClone = strings.ReplaceAll(configFileClone, "<CLUSTER_ID>", clusterId)
				configFileClone = UpdateMonConnectionSettings(configFileClone, project.ClusterConnectionSettings)
				configFileClone = UpdateLogstashConnectionSettings(configFileClone, logstashHost)
				folderPath := filepath.Join(logstashHost.ServerAddress, "dbeast-mon", configFile.Id)
				WriteFileToZip(zipWriter, folderPath, configFileClone)
				pipelineId := strings.ReplaceAll(configFile.Id, ".conf", "")
				pipelineFile += fmt.Sprintf("- pipeline.id: %s\n", pipelineId)
				pipelineFile += fmt.Sprintf("  path.config: \"/etc/logstash/conf.d/dbeast-mon/%s\"\n\n", configFile.Id)
			}
		}
		WriteFileToZip(zipWriter, filepath.Join(logstashHost.ServerAddress, "pipelines.yml"), pipelineFile)
	}
}

func UpdateProdConnectionSettings(configFileContent string, environmentConfig EnvironmentConfig) string {
	return UpdateConnectionSettings(configFileContent, environmentConfig.Prod.Elasticsearch, "PROD")
}

func UpdateMonConnectionSettings(configFileContent string, environmentConfig EnvironmentConfig) string {
	return UpdateConnectionSettings(configFileContent, environmentConfig.Mon.Elasticsearch, "MON")
}

func UpdateConnectionSettings(configFileContent string, credentials Credentials, env string) string {
	configFileContent = strings.ReplaceAll(configFileContent, "<"+env+"_HOST>", credentials.Host)
	configFileContent = strings.ReplaceAll(configFileContent, "<"+env+"_USER>", credentials.Username)
	configFileContent = strings.ReplaceAll(configFileContent, "<"+env+"_PASSWORD>", credentials.Password)
	configFileContent = strings.ReplaceAll(configFileContent, "<"+env+"_SSL_ENABLED>", fmt.Sprintf("%t", strings.Contains(credentials.Host, "https")))
	return configFileContent
}

func UpdateLogstashConnectionSettings(configFileContent string, logstashHost LogstashHost) string {
	configFileContent = strings.ReplaceAll(configFileContent, "<PATH_TO_LOGS>", logstashHost.LogstashLogsFolder)
	configFileContent = strings.ReplaceAll(configFileContent, "<LOGSTASH-API>", logstashHost.LogstashApiHost)
	return configFileContent
}

func WriteFileToZip(zipWriter *zip.Writer, fileInternalPath string, configFile string) {
	fileWriter, err := zipWriter.Create(fileInternalPath)
	if err != nil {
		log.DefaultLogger.Error(err.Error())
	}

	_, err = fileWriter.Write([]byte(configFile))
	if err != nil {
		log.DefaultLogger.Error(err.Error())
	}
}
