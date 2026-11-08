package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "Elasticsearch-mon-dbeast-mon-es-thread_pools-status"
	const dbeastMinimalCompatibleVersion = "0.0.0"
	dataWarehouse.LoadGrafanaDataSources(elasticsearchDatasourceElasticsearchThreadPoolsStatusContent, componentName, dbeastMinimalCompatibleVersion)
}

const elasticsearchDatasourceElasticsearchThreadPoolsStatusContent string = `
{
  "orgId": 1,
  "name": "%COMPONENT_NAME%",
  "type": "elasticsearch",
  "typeName": "Elasticsearch",
  "access": "proxy",
  "url": "",
  "database": "dbeast-mon-es-thread_pools-status",
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
