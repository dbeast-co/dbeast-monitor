package data_source_templates

import dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"

// TODO Add field for dbeast version/min compatible version etc
// Build json object with the versions list based on the dataWarhouse
// Add "versions" API
// variable for the current version
// variable for the last updated version
// Application health includes data consistency and version verification dashboard

func init() {
	const componentName = "Elasticsearch-mon-dbeast-mon-es-index-stats"
	const dbeastMinimalCompatibleVersion = "2.0.0"
	dataWarehouse.LoadGrafanaDataSources(elasticsearchDatasourceElasticsearchIndexStatsContent, componentName, dbeastMinimalCompatibleVersion)
	//dataWarehouse.LoadGrafanaDataSources(dataWarehouse.UpdateGrafanaDataSources(elasticsearchDatasourceElasticsearchIndexStatsContent, componentName))
	//dataWarehouse.AppendGrafanaDSMinCompatibleVersion(componentName, dbeastMinimalCompatibleVersion)
}

const elasticsearchDatasourceElasticsearchIndexStatsContent string = `
{
  "orgId": 1,
  "name": "%COMPONENT_NAME%",
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
	"dbeastVersion": "%DBEAST_VERSION%"
  },
  "secureJsonData": {
    "basicAuthPassword": ""
  },
  "readOnly": false
}
`
