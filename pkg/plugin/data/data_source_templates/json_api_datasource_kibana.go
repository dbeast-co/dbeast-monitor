package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const jsonApiDatasourceKibanaContent string = `
{
  "orgId": 1,
  "name": "Kibana-direct--",
  "type": "marcusolsson-json-datasource",
  "typeName": "JSON API",
  "access": "proxy",
  "url": "",
  "basicAuth": false,
  "basicAuthUser": "",
  "jsonData": {
    "tlsSkipVerify": true,
    "httpHeaderName1": "kbn-xsrf"
  },
  "secureJsonData": {
    "basicAuthPassword": "",
    "httpHeaderValue1": "true"
  },
  "readOnly": false
}
`

func init() {
	dataWarehouse.LoadGrafanaDataSources("json_api_datasource_kibana", jsonApiDatasourceKibanaContent)
}
