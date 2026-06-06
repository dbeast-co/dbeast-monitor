package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const elasticsearchDatasourceElasticsearchIndexStatsStatusContent string = `
{
  "orgId": 1,
  "name": "Elasticsearch-mon-",
  "type": "elasticsearch",
  "typeName": "Elasticsearch",
  "access": "proxy",
  "url": "",
  "database": "dbeast-mon-es-index-stats-status",
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
    "tlsSkipVerify": true
  },
  "secureJsonData": {
    "basicAuthPassword": ""
  },
  "readOnly": false
}
`

func init() {
	dataWarehouse.LoadGrafanaDataSources("elasticsearch_datasource_elasticsearch_index_stats_status", elasticsearchDatasourceElasticsearchIndexStatsStatusContent)
}
