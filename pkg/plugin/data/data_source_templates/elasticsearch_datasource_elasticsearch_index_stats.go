package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

// TODO Add field for dbeast version/min compatible version etc
// Build json object with the versions list based on the dataWarhouse
// Add "versions" API
// variable for the current version
// variable for the last updated version
// Application health includes data consistency and version verification dashboard

const dbeastMinimalCompatibleVersion = "0.0.0"
const elasticsearchDatasourceElasticsearchIndexStatsContent string = `
{
  "orgId": 1,
  "name": "Elasticsearch-mon-dbeast-mon-es-index-stats",
  "type": "elasticsearch",
  "typeName": "Elasticsearch",
  "access": "proxy", 
  "url": "",
  "database": "dbeast-mon-es-index-stats",
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
	"dbeastVersion": "%DBEAST_VERSION%",
	"dbeastMinimalCompatibleVersion": "2.0.0"
  },
  "secureJsonData": {
    "basicAuthPassword": ""
  },
  "readOnly": false
}
`

func init() {
	dataWarehouse.LoadGrafanaDataSources(elasticsearchDatasourceElasticsearchIndexStatsContent)
}
