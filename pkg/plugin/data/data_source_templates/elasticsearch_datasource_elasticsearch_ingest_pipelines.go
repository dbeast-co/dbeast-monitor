package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const elasticsearchDatasourceElasticsearchIngestPipelinesContent string = `
{
  "orgId": 1,
  "name": "Elasticsearch-mon-",
  "type": "elasticsearch",
  "typeName": "Elasticsearch",
  "access": "proxy",
  "url": "",
  "database": "dbeast-mon-es-ingest_pipelines",
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
	dataWarehouse.LoadGrafanaDataSources("elasticsearch_datasource_elasticsearch_ingest_pipelines", elasticsearchDatasourceElasticsearchIngestPipelinesContent)
}
