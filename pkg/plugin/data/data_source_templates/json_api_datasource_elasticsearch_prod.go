package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const jsonApiDatasourceElasticsearchProdContent string = `
{
  "orgId": 1,
  "name": "Elasticsearch-direct-prod--",
  "type": "marcusolsson-json-datasource",
  "typeName": "JSON API",
  "access": "proxy",
  "url": "",
  "basicAuth": false,
  "basicAuthUser": "",
  "jsonData": {
    "tlsSkipVerify": true
  },
  "secureJsonData": {
    "basicAuthPassword": ""
  },
  "readOnly": false
}
`

func init() {
	dataWarehouse.LoadGrafanaDataSources("json_api_datasource_elasticsearch_prod", jsonApiDatasourceElasticsearchProdContent)
}
