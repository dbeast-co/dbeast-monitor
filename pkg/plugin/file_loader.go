package plugin

import (
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"os"
	"path/filepath"
)

var NewCluster Project

var GrafanaDataSourcesMap = make(map[string]interface{})

var ESComponentTemplatesMap = make(map[string]string)

var ESIndexTemplatesMap = make(map[string]string)

var ESILMTemplatesMap = make(map[string]string)

var ESFirstIndicesTemplatesMap = make(map[string]string)

func LoadInitData(applicationFolder string) error {
	log.DefaultLogger.Info("Loading the Grafana data sources")
	err := LoadGrafanaDataSources(filepath.Join(applicationFolder, DataSourceTemplatesFolder))
	if err != nil {
		log.DefaultLogger.Error("Error in the Grafana data source loading")
		return err
	}

	log.DefaultLogger.Info("Loading the Logstash config files")
	err = LoadLogstashConfigFiles(filepath.Join(applicationFolder, LogstashTemplatesFolder))
	if err != nil {
		log.DefaultLogger.Error("Error in the Grafana data source loading")
		return err
	}
	log.DefaultLogger.Info("Loading the Elasticsearch component templates")
	err = LoadESComponentTemplates(filepath.Join(applicationFolder, EsIndexComponentsTemplatesFolder))
	if err != nil {
		log.DefaultLogger.Error("Error in the ES Component templates loading")
		return err
	}
	log.DefaultLogger.Info("Loading the Elasticsearch index templates")
	err = LoadESIndexTemplates(filepath.Join(applicationFolder, EsIndexTemplatesTemplatesFolder))
	if err != nil {
		log.DefaultLogger.Error("Error in the ES ILM templates loading")
		return err
	}
	log.DefaultLogger.Info("Loading the ILM policies")
	err = LoadESILMTemplates(filepath.Join(applicationFolder, EsILMTemplatesFolder))
	if err != nil {
		log.DefaultLogger.Error("Error in the ES ILM templates loading")
		return err
	}
	log.DefaultLogger.Info("Loading the First indices")
	err = LoadESFirstIndices(filepath.Join(applicationFolder, EsIndexFirstIndicesTemplatesFolder))
	if err != nil {
		log.DefaultLogger.Error("Error in the ES first indices loading")
		return err
	}

	log.DefaultLogger.Info("Loading New cluster definition file")
	err = LoadNewClusterFile(filepath.Join(applicationFolder, NewClusterFile))
	if err != nil {
		log.DefaultLogger.Error("Error in the New cluster object loading")
		return err
	}

	return nil
}

func LoadNewClusterFile(filePath string) error {
	log.DefaultLogger.Debug("New cluster configuration file path: " + filePath)

	fileContent, err := ReadFileToByteArray(filePath)
	if err != nil {
		log.DefaultLogger.Error("Error loading new cluster file: " + filePath)
		return err
	}
	err = json.Unmarshal(fileContent, &NewCluster)
	if err != nil {
		log.DefaultLogger.Error("Error parsing new cluster file: " + filePath + " " + err.Error())
		return err
	}
	return err
}

func LoadGrafanaDataSources(folderPath string) error {
	log.DefaultLogger.Info("The templates folder path: " + folderPath)

	stringSourceMap, err := ReadFilesFromFolderByteArrayType(folderPath, ".json", true)
	if err != nil {
		log.DefaultLogger.Error("Error read Grafana templates folder: " + err.Error())
		return err
	}
	for templateName, stringDataSource := range stringSourceMap {
		var templateData map[string]interface{}
		err = json.Unmarshal(stringDataSource, &templateData)
		if err != nil {
			log.DefaultLogger.Error("Error parsing file: " + templateName + " " + err.Error())
			return err
		}
		GrafanaDataSourcesMap[templateName] = templateData
	}
	return nil
}

func LoadLogstashConfigFiles(folderPath string) error {
	var err error
	LSConfigs, err = ReadFilesFromFolderStringType(folderPath, ".conf", false)
	return err
}

func LoadESComponentTemplates(folderPath string) error {
	var err error
	ESComponentTemplatesMap, err = ReadFilesFromFolderStringType(folderPath, ".json", true)
	return err
}

func LoadESFirstIndices(folderPath string) error {
	var err error
	ESFirstIndicesTemplatesMap, err = ReadFilesFromFolderStringType(folderPath, ".json", true)
	return err
}

func LoadESIndexTemplates(folderPath string) error {
	var err error
	ESIndexTemplatesMap, err = ReadFilesFromFolderStringType(folderPath, ".json", true)
	return err
}

func LoadESILMTemplates(folderPath string) error {
	var err error
	ESILMTemplatesMap, err = ReadFilesFromFolderStringType(folderPath, ".json", true)
	return err
}

func ReadFilesFromFolderStringType(folderPath string, filesExtension string, isRemoveExtension bool) (map[string]string, error) {
	bytesFilesContent, err := ReadFilesFromFolderByteArrayType(folderPath, filesExtension, isRemoveExtension)
	tmpFilesContent := make(map[string]string)
	if err != nil {
		return nil, err
	}
	for templateName, templateContent := range bytesFilesContent {
		tmpFilesContent[templateName] = string(templateContent[:])
	}
	return tmpFilesContent, nil
}

func ReadFilesFromFolderByteArrayType(folderPath string, filesExtension string, isRemoveExtension bool) (map[string][]byte, error) {
	tmpFilesContent := make(map[string][]byte)
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.DefaultLogger.Error("failed to read files from folder: "+folderPath, err)
		return tmpFilesContent, fmt.Errorf("failed to read files from folder: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			//if !file.IsDir() && (len(filesExtension) == 0 || strings.HasSuffix(file.Name(), filesExtension)) {
			filePath := filepath.Join(folderPath, file.Name())
			fileContents, err := ReadFileToByteArray(filePath)
			if err != nil {
				return tmpFilesContent, err
			}
			var templateName string
			if isRemoveExtension {
				templateName = file.Name()[:len(file.Name())-len(filesExtension)]
			} else {
				templateName = file.Name()
			}
			tmpFilesContent[templateName] = fileContents
		}
	}
	return tmpFilesContent, nil
}

func ReadFileToString(filePath string) (string, error) {
	data, err := ReadFileToByteArray(filePath)
	return string(data[:]), err
}

func ReadFileToByteArray(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.DefaultLogger.Error("Failed to read file %s: %v", filePath, err)
		return nil, fmt.Errorf("failed to read files from folder: %v", err)
	}
	log.DefaultLogger.Info("Reading file: " + filePath)
	return data, nil
}
