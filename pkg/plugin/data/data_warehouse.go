package data

import (
	"encoding/json"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

var NewCluster Project
var ESFirstIndicesTemplatesMap = make(map[string]string)
var ESILMTemplatesMap = make(map[string]string)
var ESComponentTemplatesMap = make(map[string]string)
var ESIndexTemplatesMap = make(map[string]string)
var LSConfigsMap = make(map[string]string)
var GrafanaDataSourcesMap = make(map[string]interface{})

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

func LoadGrafanaDataSources(DataSourceName string, DataSourceContent string) {
	var templateData map[string]interface{}
	err := json.Unmarshal([]byte(DataSourceContent), &templateData)
	if err != nil {
		log.DefaultLogger.Error("Error parsing: " + DataSourceName + " " + err.Error())
		return
	}
	GrafanaDataSourcesMap[DataSourceName] = templateData
	log.DefaultLogger.Info("Grafana data source " + DataSourceName + " added to the map successfully")
}

func LoadNewCluster(NewClusterContent string) {
	err := json.Unmarshal([]byte(NewClusterContent), &NewCluster)
	if err != nil {
		log.DefaultLogger.Error("Failed to parse NewCluster: " + err.Error())
		return
	}
	log.DefaultLogger.Info("New cluster configuration loaded successfully")

}
