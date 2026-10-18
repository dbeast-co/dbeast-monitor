package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "Elasticsearch-direct-prod"
	const dbeastMinimalCompatibleVersion = "2.0.0"
	dataWarehouse.LoadGrafanaDataSources(jsonApiDatasourceElasticsearchProdContent, componentName, dbeastMinimalCompatibleVersion)
}

const jsonApiDatasourceElasticsearchProdContent string = `
{
  "orgId": 1,
  "name": "%COMPONENT_NAME%",
  "type": "marcusolsson-json-datasource",
  "typeName": "JSON API",
  "access": "proxy",
  "url": "",
  "basicAuth": false,
  "basicAuthUser": "",
  "jsonData": {
    "tlsSkipVerify": true,
	"dbeastVersion": "%DBEAST_VERSION%"
  },
  "secureJsonData": {
    "basicAuthPassword": ""
  },
  "readOnly": false
}
`
