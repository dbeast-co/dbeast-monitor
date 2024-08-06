package plugin

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"net/http"
	"path/filepath"
	"strings"
)

type ConfigurationCheckbox struct {
	Label     string `json:"label"`
	Id        string `json:"id"`
	IsChecked bool   `json:"is_checked"`
}

type LogstashHost struct {
	ServerAddress   string `json:"server_address"`
	LogstashApiHost string `json:"logstash_api_host"`
	//LogstashConfigFolder string `json:"logstash_config_folder"`
	LogstashLogsFolder string `json:"logstash_logs_folder"`
}

type LogstashConfigurations struct {
	EsMonitoringConfigurationFiles       []ConfigurationCheckbox              `json:"es_monitoring_configuration_files"`
	LogstashMonitoringConfigurationFiles LogstashMonitoringConfigurationFiles `json:"logstash_monitoring_configuration_files"`
}

type LogstashMonitoringConfigurationFiles struct {
	Configurations []ConfigurationCheckbox `json:"configurations"`
	Hosts          []LogstashHost          `json:"hosts"`
}

var LSConfigs = make(map[string]string)

func (a *App) GenerateElasticsearchMonitoringConfigurationFilesHandler(w http.ResponseWriter, req *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(req.Context())
	ctxLogger.Info("Got request for the Elasticsearch configuration files generation")

	GenerateLogstashConfigurationFiles(w, req, false, "ESConfigurationFiles.zip")
}

func (a *App) GenerateLogstashMonitoringConfigurationFilesHandler(w http.ResponseWriter, req *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(req.Context())
	ctxLogger.Info("Got request for the Logstash configuration files generation")

	GenerateLogstashConfigurationFiles(w, req, true, "LogstashConfigurationFiles.zip")
}

func GenerateLogstashConfigurationFiles(w http.ResponseWriter, req *http.Request, isLogstash bool, resultZipFileName string) {
	ctxLogger := log.DefaultLogger.FromContext(req.Context())

	w.Header().Add("Content-Disposition", "attachment; filename=\""+resultZipFileName+"\"")
	w.Header().Add("Content-Type", "application/zip")

	var project Cluster

	if err := json.NewDecoder(req.Body).Decode(&project); err != nil {
		log.DefaultLogger.Error("Failed to decode JSON data: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid request payload"})
		return
	}
	ctxLogger.Debug("The project: ", project)
	defer req.Body.Close()

	buf := new(bytes.Buffer)

	zipWriter := zip.NewWriter(buf)

	clusterName, clusterId, err := GetClusterInfo(project.ClusterConnectionSettings.Prod.Elasticsearch)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	if isLogstash {
		GenerateLSLogstashConfigurationFiles(project, clusterId, zipWriter, ctxLogger)
	} else {
		GenerateESLogstashConfigurationFiles(project, clusterId, clusterName, zipWriter, ctxLogger)
	}
	err = zipWriter.Close()
	if err != nil {
		log.DefaultLogger.Error("Error closing ZIP: ", err.Error())
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		log.DefaultLogger.Error("Error writing response: ", err.Error())
	}
}

func GenerateESLogstashConfigurationFiles(project Cluster, clusterId string, clusterName string, zipWriter *zip.Writer, logger log.Logger) {
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
			pipelineFile += fmt.Sprintf("  path.config: \"%s\"\n\n", fileInternalPath)
		}
	}
	WriteFileToZip(zipWriter, "pipelines.yml", pipelineFile)
}

func GenerateLSLogstashConfigurationFiles(project Cluster, clusterId string, zipWriter *zip.Writer, logger log.Logger) {
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
				pipelineFile += fmt.Sprintf("  path.config: \"%s\"\n\n", "dbeast-mon"+"/"+configFile.Id)
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
