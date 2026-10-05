package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const jsonApiDatasourceKibanaContent string = `
{
  "orgId": 1,
  "name": "Kibana-direct",
  "type": "marcusolsson-json-datasource",
  "typeName": "JSON API",
  "access": "proxy",
  "url": "",
  "basicAuth": false,
  "basicAuthUser": "",
  "jsonData": {
    "tlsSkipVerify": true,
    "httpHeaderName1": "kbn-xsrf",
    "dbeastVersion": "%DBEAST_VERSION%",
	"dbeastMinimalCompatibleVersion": "2.0.0"
  },
  "secureJsonData": {
    "basicAuthPassword": "",
    "httpHeaderValue1": "true"
  },
  "readOnly": false
}
`

func init() {
	dataWarehouse.LoadGrafanaDataSources(jsonApiDatasourceKibanaContent)
}
