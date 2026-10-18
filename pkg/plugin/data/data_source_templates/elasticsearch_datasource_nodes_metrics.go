package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "Elasticsearch-mon-metricbeat-monitoring-es"
	const dbeastMinimalCompatibleVersion = "2.0.0"
	dataWarehouse.LoadGrafanaDataSources(elasticsearchDatasourceNodesMetricsContent, componentName, dbeastMinimalCompatibleVersion)
}

const elasticsearchDatasourceNodesMetricsContent string = `
{
  "orgId": 1,
  "name": "%COMPONENT_NAME%",
  "type": "elasticsearch",
  "typeName": "Elasticsearch",
  "access": "proxy",
  "url": "",
  "database": "metricbeat-*,.monitoring-es-*",
  "basicAuth": true,
  "isDefault": false,
  "withCredentials": false,
  "basicAuthUser": "",
  "jsonData": {
    "esVersion": "8.0.0",
    "includeFrozen": false,
    "logLevelField": "",
    "logMessageField": "",
    "maxConcurrentShardRequests": 5,
    "timeField": "@timestamp",
    "tlsSkipVerify": true,
	"dbeastVersion": "%DBEAST_VERSION%"
  },
  "secureJsonData": {
    "basicAuthPassword": ""
  },
  "readOnly": false
}
`
