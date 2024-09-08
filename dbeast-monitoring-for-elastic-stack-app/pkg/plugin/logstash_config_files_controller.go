package plugin

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var LSConfigs = make(map[string]string)

func (a *App) DownloadElasticsearchMonitoringConfigurationFilesHandler(w http.ResponseWriter, req *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(req.Context())
	ctxLogger.Info("Got request for the Elasticsearch configuration files generation")

	GenerateLogstashConfigurationFiles(w, req, false, "ESConfigurationFiles.zip")
}

func (a *App) DownloadLogstashMonitoringConfigurationFilesHandler(w http.ResponseWriter, req *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(req.Context())
	ctxLogger.Info("Got request for the Logstash configuration files generation")

	GenerateLogstashConfigurationFiles(w, req, true, "LogstashConfigurationFiles.zip")
}

func SaveLogstashConfigurationFiles(project Project, logger log.Logger) error {
	clusterName, clusterId, err := GetClusterInfo(project.ClusterConnectionSettings.Prod.Elasticsearch)
	if err != nil {
		return err
	}
	err = DeleteTextBlockInFile(GrafanaLogstashConfigurationsFolder+"/pipelines.yml", "### Configuration files for the cluster Id: "+clusterId,
		"### Configuration files for the cluster Id: ",
		logger)

	pipelineFile := "\n\n### Configuration files for the cluster Id: " + clusterId + ", cluster name: " + clusterName + "\n"
	for _, configFile := range project.LogstashConfigurations.EsMonitoringConfigurationFiles {
		if configFile.IsChecked {
			configFileClone := strings.Clone(LSConfigs[configFile.Id])
			configFileClone = strings.ReplaceAll(configFileClone, "<CLUSTER_ID>", clusterId)
			configFileClone = UpdateMonConnectionSettings(configFileClone, project.ClusterConnectionSettings)
			configFileClone = UpdateProdConnectionSettings(configFileClone, project.ClusterConnectionSettings)

			fileAbsolutePath := filepath.Join(GrafanaLogstashConf_dConfigurationsFolder, clusterName+"-"+clusterId, configFile.Id)
			WriteFilesToDisk(fileAbsolutePath, configFileClone, false, logger)

			pipelineId := strings.ReplaceAll(configFile.Id, ".conf", "") + "-" + clusterName + "-" + clusterId

			pipelinesPath := filepath.Join(LogstashOriginalConfigurationsFolder, clusterName+"-"+clusterId, configFile.Id)
			pipelineFile += fmt.Sprintf("- pipeline.id: %s\n", pipelineId)
			pipelineFile += fmt.Sprintf("  path.config: \"%s\"\n\n", pipelinesPath)
		}
	}
	fileAbsolutePath := filepath.Join(GrafanaLogstashConfigurationsFolder, "pipelines.yml")
	WriteFilesToDisk(fileAbsolutePath, pipelineFile, true, logger)
	return nil
}

func GenerateLogstashConfigurationFiles(w http.ResponseWriter, req *http.Request, isLogstash bool, resultZipFileName string) {
	ctxLogger := log.DefaultLogger.FromContext(req.Context())

	w.Header().Add("Content-Disposition", "attachment; filename=\""+resultZipFileName+"\"")
	w.Header().Add("Content-Type", "application/zip")

	var project Project

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

func GenerateESLogstashConfigurationFiles(project Project, clusterId string, clusterName string, zipWriter *zip.Writer, logger log.Logger) {
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

func GenerateLSLogstashConfigurationFiles(project Project, clusterId string, zipWriter *zip.Writer, logger log.Logger) {
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

func WriteFilesToDisk(fileInternalPath string, content string, isAppend bool, logger log.Logger) {
	//var fileAbsoluteInternalPath = filepath.Join(GrafanaLogstashConfigurationsFolder, fileInternalPath)

	logger.Debug("File content: ", content)
	dir := filepath.Dir(fileInternalPath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		logger.Error("Error creating directory:", err)
		return
	}

	// Save the JSON data to a file
	var file *os.File
	if isAppend {
		file, err = os.OpenFile(fileInternalPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		file, err = os.Create(fileInternalPath)

	}
	if err != nil {
		logger.Error("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content + "\n")
	if err != nil {
		logger.Error("Error writing to file:", err)
		return
	}

	// Ensure the buffer is flushed to disk
	err = writer.Flush()
	if err != nil {
		logger.Error("Error flushing writer to the file:", err)
		return
	}

	logger.Info("Object saved to the file: " + fileInternalPath)

}

func DeleteTextBlockInFile(filePath string, startMarker string, endMarkerPrefix string, logger log.Logger) error {
	// Read the entire file content
	logger.Info("File name: " + filePath)
	logger.Info("Start prefix: " + startMarker)
	logger.Info("End prefix: " + endMarkerPrefix)
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		logger.Error("Error opening file:", err)
		return err
	}
	defer file.Close()

	// Read the file line by line into a slice
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.Error("Error reading file:", err)
		return err
	}

	// Find the start marker and the next end marker prefix
	startIndex := -1
	endIndex := -1
	for i, line := range lines {
		if startIndex == -1 && strings.Contains(line, startMarker) {
			startIndex = i
		} else if startIndex != -1 && strings.Contains(line, endMarkerPrefix) {
			endIndex = i
			break
		}
	}

	// If a block was found, delete the lines between the markers
	if startIndex != -1 {
		if endIndex == -1 {
			// If no end marker was found, delete till the end of the file
			endIndex = len(lines)
		}
		// Remove lines between startIndex and endIndex
		lines = append(lines[:startIndex], lines[endIndex:]...)
	}

	// Rewrite the modified content back to the file
	err = os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		logger.Error("Error writing to file:", err)
		return err
	}

	logger.Info("Block between markers removed successfully.")
	return nil
}

func DeleteFolder(folderPath string, pattern string, logger log.Logger) error {
	entries, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return err
	}

	//logger.Info("The file list for delete: " + strings.Join(entries.Name(), ", "))
	logger.Info("The config folder path: " + folderPath)
	logger.Info("The file suffix for delete: " + pattern)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		logger.Info("Check folder for deletion: " + entry.Name())
		if entry.IsDir() && strings.HasSuffix(entry.Name(), pattern) { // Ensure it's the correct entry
			logger.Info("Deleting folder:" + entry.Name())
			fullPath := folderPath + "/" + entry.Name()
			logger.Info("Deleting folder full path:" + fullPath)

			// Delete the directory and its contents
			err := os.RemoveAll(fullPath)
			if err != nil {
				return err
			}

			fmt.Println("Deleted directory: ", fullPath)
		}
	}
	return nil
}
