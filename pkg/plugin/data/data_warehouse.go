package data

import (
	"encoding/json"
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
var dbeastAppFormatVersion = loadDbeastAppFormatVersion()

func AppendFirstIndex(IndexName string, IndexContent string) {
	ESFirstIndicesTemplatesMap[IndexName] = IndexContent
	log.DefaultLogger.Info("First index " + IndexName + " added to the map successfully")
}

func AppendILMPolicy(PolicyName string, PolicyContent string, minCompatibleVersion string) {
	PolicyContent = replaceVersions(PolicyContent)

	ESILMTemplatesMap[PolicyName] = PolicyContent

	version := GenerateVersionObject(PolicyName, minCompatibleVersion)
	BackCompatibilityVersionsMap.ElasticsearchILMPolicies = append(BackCompatibilityVersionsMap.ElasticsearchILMPolicies, version)

	log.DefaultLogger.Info("ILM policy " + PolicyName + " added to the map successfully")
}

func AppendComponentTemplate(TemplateName string, TemplateContent string, minCompatibleVersion string) {
	TemplateContent = replaceVersions(TemplateContent)
	ESComponentTemplatesMap[TemplateName] = TemplateContent

	version := GenerateVersionObject(TemplateName, minCompatibleVersion)
	BackCompatibilityVersionsMap.ElasticsearchComponentTemplates = append(BackCompatibilityVersionsMap.ElasticsearchComponentTemplates, version)

	log.DefaultLogger.Info("Component template " + TemplateName + " added to the map successfully")
}

func AppendIndexTemplate(TemplateName string, TemplateContent string, minCompatibleVersion string) {
	TemplateContent = replaceVersions(TemplateContent)

	ESIndexTemplatesMap[TemplateName] = TemplateContent

	version := GenerateVersionObject(TemplateName, minCompatibleVersion)
	BackCompatibilityVersionsMap.ElasticsearchIndexTemplates = append(BackCompatibilityVersionsMap.ElasticsearchIndexTemplates, version)

	log.DefaultLogger.Info("Index template " + TemplateName + " added to the map successfully")
}

func AppendLogstashConfig(ConfigName string, ConfigContent string, minCompatibleVersion string) {
	ConfigContent = replaceVersions(ConfigContent)

	LSConfigsMap[ConfigName] = ConfigContent

	version := GenerateVersionObject(ConfigName, minCompatibleVersion)
	BackCompatibilityVersionsMap.LogstashConfigFiles = append(BackCompatibilityVersionsMap.LogstashConfigFiles, version)

	log.DefaultLogger.Info("Logstash config " + ConfigName + " added to the map successfully")
}

func LoadGrafanaDataSources(DataSourceContent string, name string, minCompatibleVersion string) {
	DataSourceContent = UpdateGrafanaDataSources(DataSourceContent, name)
	var templateData map[string]interface{}
	err := json.Unmarshal([]byte(DataSourceContent), &templateData)
	if err != nil {
		log.DefaultLogger.Error("Error parsing grafana data source template: " + err.Error())
		return
	}
	GrafanaDataSourcesMap[templateData["name"].(string)] = templateData

	version := GenerateVersionObject(name, minCompatibleVersion)
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
	DataSourceContent = replaceVersions(DataSourceContent)
	return DataSourceContent
}

func loadDbeastVersion() string {
	info, err := buildinfo.GetBuildInfo.GetInfo()
	if err != nil {
		return "0.0.0"
	}
	return info.Version
}

func replaceVersions(content string) string {
	content = strings.ReplaceAll(content, "%DBEAST_VERSION%", dbeastVersion)
	return strings.ReplaceAll(content, "%DBEAST_APP_FORMAT_VERSION%", strconv.Itoa(dbeastAppFormatVersion))
}

func loadDbeastAppFormatVersion() int {
	info, err := buildinfo.GetBuildInfo.GetInfo()
	if err != nil {
		return 0
	}
	return formatVersion(info.Version)
}
