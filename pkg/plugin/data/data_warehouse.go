package data

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/build/buildinfo"
)

const dbeastVersionPlaceholder = "%DBEAST_VERSION%"

var NewCluster Project
var ESFirstIndicesTemplatesMap = make(map[string]string)
var ESILMTemplatesMap = make(map[string]string)
var ESComponentTemplatesMap = make(map[string]string)
var ESIndexTemplatesMap = make(map[string]string)
var LSConfigsMap = make(map[string]string)
var GrafanaDataSourcesMap = make(map[string]interface{})
var BackCompatibilityVersionsMap BackCompatibilityVersions

func GenerateBackCompatibilityVersionsMap() {
	versions := make([]VersionObject, 0, len(GrafanaDataSourcesMap))
	for _, ds := range GrafanaDataSourcesMap {
		dsMap, ok := ds.(map[string]interface{})
		if !ok {
			continue
		}

		name, ok := dsMap["name"].(string)
		if !ok {
			continue
		}

		jsonData, ok := dsMap["jsonData"].(map[string]interface{})
		if !ok {
			continue
		}

		minVersion, ok := jsonData["dbeastMinimalCompatibleVersion"].(string)
		if !ok {
			continue
		}

		versions = append(versions, VersionObject{Name: name, MinCompatibleVersion: minVersion})
	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Name < versions[j].Name
	})

	BackCompatibilityVersionsMap.GrafanaDataSourceTemplates = versions
}

func AppendFirstIndex(IndexName string, IndexContent string) {
	ESFirstIndicesTemplatesMap[IndexName] = IndexContent
	log.DefaultLogger.Info("First index " + IndexName + " added to the map successfully")
}

func AppendILMPolicy(PolicyName string, PolicyContent string) {
	ESILMTemplatesMap[PolicyName] = PolicyContent
	log.DefaultLogger.Info("ILM policy " + PolicyName + " added to the map successfully")
}

func AppendComponentTemplate(TemplateName string, TemplateContent string) {
	ESComponentTemplatesMap[TemplateName] = TemplateContent
	log.DefaultLogger.Info("Component template " + TemplateName + " added to the map successfully")
}

func AppendIndexTemplate(TemplateName string, TemplateContent string) {
	ESIndexTemplatesMap[TemplateName] = TemplateContent
	log.DefaultLogger.Info("Index template " + TemplateName + " added to the map successfully")
}

func AppendLogstashConfig(ConfigName string, ConfigContent string) {
	LSConfigsMap[ConfigName] = ConfigContent
	log.DefaultLogger.Info("Logstash config " + ConfigName + " added to the map successfully")
}

func LoadGrafanaDataSources(DataSourceContent string) {
	DataSourceContent = strings.ReplaceAll(DataSourceContent, dbeastVersionPlaceholder, dbeastVersion())

	var templateData map[string]interface{}
	err := json.Unmarshal([]byte(DataSourceContent), &templateData)
	if err != nil {
		log.DefaultLogger.Error("Error parsing grafana data source template: " + err.Error())
		return
	}
	GrafanaDataSourcesMap[templateData["name"].(string)] = templateData
	log.DefaultLogger.Info("Grafana data source " + templateData["name"].(string) + " added to the map successfully")
}

func LoadNewCluster(NewClusterContent string) {
	err := json.Unmarshal([]byte(NewClusterContent), &NewCluster)
	if err != nil {
		log.DefaultLogger.Error("Failed to parse NewCluster: " + err.Error())
		return
	}
	log.DefaultLogger.Info("New cluster configuration loaded successfully")

}

func dbeastVersion() string {
	info, err := buildinfo.GetBuildInfo.GetInfo()
	if err != nil || info.Version == "" {
		return "dev"
	}
	return info.Version
}
