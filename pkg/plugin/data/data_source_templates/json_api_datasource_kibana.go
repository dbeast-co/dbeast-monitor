package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "Kibana-direct"
	const dbeastMinimalCompatibleVersion = "2.0.0"
	dataWarehouse.LoadGrafanaDataSources(jsonApiDatasourceKibanaContent, componentName, dbeastMinimalCompatibleVersion)
}

const jsonApiDatasourceKibanaContent string = `
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
    "httpHeaderName1": "kbn-xsrf",
    "dbeastVersion": "%DBEAST_VERSION%"
  },
  "secureJsonData": {
    "basicAuthPassword": "",
    "httpHeaderValue1": "true"
  },
  "readOnly": false
}
`
