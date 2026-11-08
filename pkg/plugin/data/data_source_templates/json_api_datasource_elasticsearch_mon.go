package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "Elasticsearch-direct-mon"
	const dbeastMinimalCompatibleVersion = "0.0.0"
	dataWarehouse.LoadGrafanaDataSources(jsonApiDatasourceElasticsearchMonContent, componentName, dbeastMinimalCompatibleVersion)
}

const jsonApiDatasourceElasticsearchMonContent string = `
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
	"dbeast_version": "%DBEAST_VERSION%",
	"dbeast_version_int": %DBEAST_APP_FORMAT_VERSION%
  },
  "secureJsonData": {
    "basicAuthPassword": ""
  },
  "readOnly": false
}
`
