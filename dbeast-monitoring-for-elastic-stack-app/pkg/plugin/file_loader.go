package plugin

import (
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"os"
	"path/filepath"
	"strings"
)

func LoadInitData(applicationFolder string) error {
	err := LoadGrafanaDataSources(filepath.Join(applicationFolder, DataSourceTemplatesFolder))
	if err != nil {
		log.DefaultLogger.Error("Error in the Grafana data source loading")
		return err
	}

	err = LoadLogstashConfigFiles(filepath.Join(applicationFolder, LogstashTemplatesFolder))
	if err != nil {
		log.DefaultLogger.Error("Error in the Grafana data source loading")
		return err
	}
	err = LoadESComponentTemplates(filepath.Join(applicationFolder, EsIndexComponentsTemplatesFolder))
	if err != nil {
		log.DefaultLogger.Error("Error in the ES Component templates loading")
		return err
	}
	err = LoadESILMTemplates(filepath.Join(applicationFolder, EsILMTemplatesFolder))
	if err != nil {
		log.DefaultLogger.Error("Error in the ES ILM templates loading")
		return err
	}

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

/*
LoadGrafanaDataSources loads JSON templates from the specified folder and updates the global GrafanaDataSourcesMap.
It takes a folderPath string as input, reads the content of the folder, and parses each JSON file into the GrafanaDataSourcesMap.
If successful, it returns nil; otherwise, it returns an error.
The GrafanaDataSourcesMap is a global variable that represents a mapping of template names to their corresponding JSON data.
This map is expected to be used elsewhere in the application after the templates are loaded.
*/
func LoadGrafanaDataSources(folderPath string) error {
	log.DefaultLogger.Debug("The templates folder path: " + folderPath)

	stringSourceMap, err := ReadFilesFromFolderByteArrayType(folderPath, ".json", true)
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

func LoadESILMTemplates(folderPath string) error {
	var err error
	ESILMTemplatesMap, err = ReadFilesFromFolderStringType(folderPath, ".json", true)
	return err
}

func ReadFilesFromFolderStringType(folderPath string, filesExtension string, isRemoveExtension bool) (map[string]string, error) {
	tmpFilesContent := make(map[string]string)
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return tmpFilesContent, fmt.Errorf("failed to read files from folder: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && (len(filesExtension) == 0 || strings.HasSuffix(file.Name(), filesExtension)) {
			filePath := filepath.Join(folderPath, file.Name())
			fileContents, err := ReadFileToString(filePath)
			if err != nil {
				return tmpFilesContent, err
			}
			var templateName string
			if isRemoveExtension {
				templateName = file.Name()[:len(file.Name())-5]
			} else {
				templateName = file.Name()
			}
			tmpFilesContent[templateName] = fileContents
		}
	}
	return tmpFilesContent, nil
}

func ReadFilesFromFolderByteArrayType(folderPath string, filesExtension string, isRemoveExtension bool) (map[string][]byte, error) {
	tmpFilesContent := make(map[string][]byte)
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return tmpFilesContent, fmt.Errorf("failed to read files from folder: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && (len(filesExtension) == 0 || strings.HasSuffix(file.Name(), filesExtension)) {
			filePath := filepath.Join(folderPath, file.Name())
			fileContents, err := ReadFileToByteArray(filePath)
			if err != nil {
				return tmpFilesContent, err
			}
			var templateName string
			if isRemoveExtension {
				templateName = file.Name()[:len(file.Name())-5]
			} else {
				templateName = file.Name()
			}
			tmpFilesContent[templateName] = fileContents
		}
	}
	return tmpFilesContent, nil
}

func ReadFileToString(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.DefaultLogger.Error("Failed to read file %s: %v", filePath, err)
		return "", fmt.Errorf("failed to read files from folder: %v", err)
	}
	log.DefaultLogger.Info("Reading file: ", filePath)
	return string(data[:]), nil
}

func ReadFileToByteArray(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.DefaultLogger.Error("Failed to read file %s: %v", filePath, err)
		return nil, fmt.Errorf("failed to read files from folder: %v", err)
	}
	log.DefaultLogger.Debug("Reading file: ", filePath)
	return data, nil
}
