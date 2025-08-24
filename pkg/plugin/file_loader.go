package plugin

import (
	"embed"
	"encoding/json"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed data/*
var dataFiles embed.FS

var NewCluster Project

var GrafanaDataSourcesMap = make(map[string]interface{})

var ESComponentTemplatesMap = make(map[string]string)

var ESIndexTemplatesMap = make(map[string]string)

var ESILMTemplatesMap = make(map[string]string)

var ESFirstIndicesTemplatesMap = make(map[string]string)

func LoadInitData() error {
	// This will hold: folderName → (fileName → content)
	folderMaps := make(map[string]map[string][]byte)

	err := fs.WalkDir(dataFiles, DataFolder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			dir := filepath.Base(filepath.Dir(path))
			name := strings.TrimSuffix(filepath.Base(path), ".json")

			if _, exists := folderMaps[dir]; !exists {
				folderMaps[dir] = make(map[string][]byte)
			}

			content, err := dataFiles.ReadFile(path)
			if err != nil {
				return err
			}
			folderMaps[dir][name] = content
		}
		return nil
	})

	if err != nil {
		return err
	}

	return parseData(folderMaps)
}

func parseData(folderMaps map[string]map[string][]byte) error {
	for folder, files := range folderMaps {
		for fileName, content := range files {
			log.DefaultLogger.Info("Folder for ingest: " + folder + ": " + fileName)
			switch folder {
			case NewClusterFolder:
				log.DefaultLogger.Info("Loading New cluster definition file")
				err := LoadNewClusterFile(fileName, content)
				if err != nil {
					log.DefaultLogger.Error("Error in the New cluster object loading")
					return err
				}
			case DataSourceTemplatesFolder:
				log.DefaultLogger.Info("Loading the Grafana data sources")
				err := LoadGrafanaDataSources(fileName, content)
				if err != nil {
					log.DefaultLogger.Error("Error in the Grafana data source loading")
					return err
				}
			case EsIndexComponentsTemplatesFolder:
				log.DefaultLogger.Info("Loading the Elasticsearch component templates")
				ESComponentTemplatesMap[fileName] = string(content)
			case EsIndexTemplatesTemplatesFolder:
				log.DefaultLogger.Info("Loading the Elasticsearch component templates")
				ESIndexTemplatesMap[fileName] = string(content)
			case EsILMTemplatesFolder:
				log.DefaultLogger.Info("Loading the Elasticsearch component templates")
				ESILMTemplatesMap[fileName] = string(content)
			case EsIndexFirstIndicesTemplatesFolder:
				log.DefaultLogger.Info("Loading the Elasticsearch component templates")
				ESFirstIndicesTemplatesMap[fileName] = string(content)
			case LogstashTemplatesFolder:
				LSConfigs[fileName] = string(content)
			}
		}
	}
	return nil
}

func LoadNewClusterFile(fileName string, content []byte) error {
	log.DefaultLogger.Debug("New cluster configuration file path: " + fileName)

	err := json.Unmarshal(content, &NewCluster)
	if err != nil {
		log.DefaultLogger.Error("Error parsing new cluster file: " + fileName + " " + err.Error())
		return err
	}
	return err
}

func LoadGrafanaDataSources(fileName string, fileContent []byte) error {
	log.DefaultLogger.Info("The templates folder path: " + fileName)
	var templateData map[string]interface{}
	err := json.Unmarshal(fileContent, &templateData)
	if err != nil {
		log.DefaultLogger.Error("Error parsing file: " + fileName + " " + err.Error())
		return err
	}
	GrafanaDataSourcesMap[fileName] = templateData
	return nil
}
