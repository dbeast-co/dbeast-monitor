package plugin

import (
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"

	_ "github.com/dbeast/dbeastmonitor/pkg/plugin/data/es_components/first_indices"
	_ "github.com/dbeast/dbeastmonitor/pkg/plugin/data/es_components/ilm_policies"
	_ "github.com/dbeast/dbeastmonitor/pkg/plugin/data/new_cluster"
)

var NewCluster Project

// var GrafanaDataSourcesMap = make(map[string]interface{})
//
// var ESComponentTemplatesMap = make(map[string]string)
var ESIndexTemplatesMap = make(map[string]string)

var ESILMTemplatesMap = make(map[string]string)

var ESFirstIndicesMap = make(map[string]string)

//
//func LoadInitData() error {
//	// This will hold: folderName → (fileName → content)
//	folderMaps := make(map[string]map[string][]byte)
//
//	// grafana-plugin-validator: filesystem-access exception - this walks embedded files, not real filesystem
//	//err := fs.WalkDir(dataFiles, DataFolder, func(path string, d fs.DirEntry, err error) error {
//	//	if err != nil {
//	//		return err
//	//	}
//	//
//	//	if !d.IsDir() {
//	//		dir := filepath.Base(filepath.Dir(path))
//	//		name := strings.TrimSuffix(filepath.Base(path), ".json")
//	//
//	//		if _, exists := folderMaps[dir]; !exists {
//	//			folderMaps[dir] = make(map[string][]byte)
//	//		}
//	//
//	//		content, err := dataFiles.ReadFile(path)
//	//		if err != nil {
//	//			return err
//	//		}
//	//		folderMaps[dir][name] = content
//	//	}
//	//	return nil
//	//})
//	//
//	//if err != nil {
//	//	return err
//	//}
//
//	return parseData(folderMaps)
//}

//func parseData(folderMaps map[string]map[string][]byte) error {
//	//for folder, files := range folderMaps {
//	//	for fileName, content := range files {
//	//		log.DefaultLogger.Info("Folder for ingest: " + folder + ": " + fileName)
//	//		//switch folder {
//	//		//case DataSourceTemplatesFolder:
//	//		//	log.DefaultLogger.Info("Loading the Grafana data sources: " + fileName)
//	//		//	err := LoadGrafanaDataSources(fileName, content)
//	//		//	if err != nil {
//	//		//		log.DefaultLogger.Error("Error in the Grafana data source loading")
//	//		//		return err
//	//		//	}
//	//		//case EsIndexComponentsTemplatesFolder:
//	//		//	log.DefaultLogger.Info("Loading the Elasticsearch component templates: " + fileName)
//	//		//	ESComponentTemplatesMap[fileName] = string(content)
//	//		//case EsIndexTemplatesTemplatesFolder:
//	//		//	log.DefaultLogger.Info("Loading the Elasticsearch component templates: " + fileName)
//	//		//	ESIndexTemplatesMap[fileName] = string(content)
//	//		//case EsILMTemplatesFolder:
//	//		//	log.DefaultLogger.Info("Loading the Elasticsearch component templates: " + fileName)
//	//		//	ESILMTemplatesMap[fileName] = string(content)
//	//		//case EsIndexFirstIndicesTemplatesFolder:
//	//		//	log.DefaultLogger.Info("Loading the Elasticsearch component templates: " + fileName)
//	//		//	ESFirstIndicesTemplatesMap[fileName] = string(content)
//	//		//case LogstashTemplatesFolder:
//	//		//	LSConfigs[fileName] = string(content)
//	//		//}
//	//	}
//	//}
//	return nil
//}

func LoadNewCluster(newClusterData Project) {
	NewCluster = newClusterData
	log.DefaultLogger.Info("New cluster configuration loaded successfully")
}

func AppendILMPolicy(PolicyName string, PolicyContent string) {
	ESILMTemplatesMap[PolicyName] = PolicyContent
	log.DefaultLogger.Info("ILM policy " + PolicyName + " added to the map successfully")
}

func AppendFirstIndex(IndexName string, IndexContent string) {
	ESFirstIndicesMap[IndexName] = IndexContent
	log.DefaultLogger.Info("First index " + IndexName + " added to the map successfully")
}

//func LoadGrafanaDataSources(fileName string, fileContent []byte) error {
//	log.DefaultLogger.Info("The templates folder path: " + fileName)
//	var templateData map[string]interface{}
//	err := json.Unmarshal(fileContent, &templateData)
//	if err != nil {
//		log.DefaultLogger.Error("Error parsing file: " + fileName + " " + err.Error())
//		return err
//	}
//	GrafanaDataSourcesMap[fileName] = templateData
//	return nil
//}
