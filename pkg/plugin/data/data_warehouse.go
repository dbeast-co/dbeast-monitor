package data

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/build/buildinfo"
)

var NewCluster Project
var ESFirstIndicesTemplatesMap = make(map[string]string)
var ESILMTemplatesMap = make(map[string]string)
var ESComponentTemplatesMap = make(map[string]string)
var ESIndexTemplatesMap = make(map[string]string)
var LSConfigsMap = make(map[string]string)
var GrafanaDataSourcesMap = make(map[string]interface{})
var BackCompatibilityVersionsMap BackCompatibilityVersions
var dbeastVersion = loadDbeastVersion()

func AppendFirstIndex(IndexName string, IndexContent string) {
	ESFirstIndicesTemplatesMap[IndexName] = IndexContent
	log.DefaultLogger.Info("First index " + IndexName + " added to the map successfully")
}

func AppendILMPolicy(PolicyName string, PolicyContent string, minCompatibleVersion int) {
	PolicyContent = strings.ReplaceAll(PolicyContent, "%DBEAST_VERSION%", dbeastVersion)

	ESILMTemplatesMap[PolicyName] = PolicyContent

	version := VersionObject{Name: PolicyName, MinCompatibleVersion: minCompatibleVersion}
	BackCompatibilityVersionsMap.ElasticsearchILMPolicies = append(BackCompatibilityVersionsMap.ElasticsearchILMPolicies, version)

	log.DefaultLogger.Info("ILM policy " + PolicyName + " added to the map successfully")
}

func AppendComponentTemplate(TemplateName string, TemplateContent string, minCompatibleVersion int) {
	TemplateContent = strings.ReplaceAll(TemplateContent, "%DBEAST_VERSION%", dbeastVersion)
	ESComponentTemplatesMap[TemplateName] = TemplateContent

	version := VersionObject{Name: TemplateName, MinCompatibleVersion: minCompatibleVersion}
	BackCompatibilityVersionsMap.ElasticsearchComponentTemplates = append(BackCompatibilityVersionsMap.ElasticsearchComponentTemplates, version)

	log.DefaultLogger.Info("Component template " + TemplateName + " added to the map successfully")
}

func AppendIndexTemplate(TemplateName string, TemplateContent string, minCompatibleVersion int) {
	TemplateContent = strings.ReplaceAll(TemplateContent, "%DBEAST_VERSION%", dbeastVersion)

	ESIndexTemplatesMap[TemplateName] = TemplateContent

	version := VersionObject{Name: TemplateName, MinCompatibleVersion: minCompatibleVersion}
	BackCompatibilityVersionsMap.ElasticsearchIndexTemplates = append(BackCompatibilityVersionsMap.ElasticsearchIndexTemplates, version)

	log.DefaultLogger.Info("Index template " + TemplateName + " added to the map successfully")
}

func AppendLogstashConfig(ConfigName string, ConfigContent string, minCompatibleVersion int) {
	ConfigContent = strings.ReplaceAll(ConfigContent, "%DBEAST_VERSION%", dbeastVersion)

	LSConfigsMap[ConfigName] = ConfigContent

	version := VersionObject{Name: ConfigName, MinCompatibleVersion: minCompatibleVersion}
	BackCompatibilityVersionsMap.LogstashConfigFiles = append(BackCompatibilityVersionsMap.LogstashConfigFiles, version)

	log.DefaultLogger.Info("Logstash config " + ConfigName + " added to the map successfully")
}

func LoadGrafanaDataSources(DataSourceContent string, name string, minCompatibleVersion int) {
	DataSourceContent = UpdateGrafanaDataSources(DataSourceContent, name)
	var templateData map[string]interface{}
	err := json.Unmarshal([]byte(DataSourceContent), &templateData)
	if err != nil {
		log.DefaultLogger.Error("Error parsing grafana data source template: " + err.Error())
		return
	}
	GrafanaDataSourcesMap[templateData["name"].(string)] = templateData

	version := VersionObject{Name: name, MinCompatibleVersion: minCompatibleVersion}
	BackCompatibilityVersionsMap.GrafanaDataSourceTemplates = append(BackCompatibilityVersionsMap.GrafanaDataSourceTemplates, version)

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

func UpdateGrafanaDataSources(DataSourceContent string, name string) string {
	DataSourceContent = strings.ReplaceAll(DataSourceContent, "%COMPONENT_NAME%", name)
	DataSourceContent = strings.ReplaceAll(DataSourceContent, "%DBEAST_VERSION%", dbeastVersion)
	return DataSourceContent
}

func loadDbeastVersion() string {
	info, err := buildinfo.GetBuildInfo.GetInfo()
	if err != nil || info.Version == "" {
		return "dev"
	}

	parts := strings.Split(info.Version, ".")
	if len(parts) != 3 {
		return info.Version
	}

	major, majorErr := strconv.Atoi(parts[0])
	minor, minorErr := strconv.Atoi(parts[1])
	patch, patchErr := strconv.Atoi(parts[2])
	if majorErr != nil || minorErr != nil || patchErr != nil {
		return info.Version
	}

	return fmt.Sprintf("%d%02d%02d", major, minor, patch)
}
