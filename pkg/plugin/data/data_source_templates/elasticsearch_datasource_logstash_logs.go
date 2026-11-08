package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "Elasticsearch-mon-dbeast-mon-logstash-logs"
	const dbeastMinimalCompatibleVersion = "0.0.0"
	dataWarehouse.LoadGrafanaDataSources(elasticsearchDatasourceLogstashLogsContent, componentName, dbeastMinimalCompatibleVersion)
}

const elasticsearchDatasourceLogstashLogsContent string = `
{
  "orgId": 1,
  "name": "%COMPONENT_NAME%",
  "type": "elasticsearch",
  "typeName": "Elasticsearch",
  "access": "proxy",
  "url": "",
  "database": "dbeast-mon-logstash-logs",
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
	"dbeast_version": "%DBEAST_VERSION%",
	"dbeast_version_int": %DBEAST_APP_FORMAT_VERSION%
  },
  "secureJsonData": {
    "basicAuthPassword": ""
  },
  "readOnly": false
}
`
