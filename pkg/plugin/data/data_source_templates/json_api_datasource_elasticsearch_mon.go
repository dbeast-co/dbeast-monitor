package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const jsonApiDatasourceElasticsearchMonContent string = `
{
  "orgId": 1,
  "name": "Elasticsearch-direct-mon--",
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

func init() {
	dataWarehouse.LoadGrafanaDataSources(jsonApiDatasourceElasticsearchMonContent)
}
