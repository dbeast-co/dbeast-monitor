package grafana_api_datasource

import dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"

const GrafanaAPIDatasourceName = "Grafana internal API"

func init() {
	const dbeastMinimalCompatibleVersion = "0.0.0"
	dataWarehouse.LoadGrafanaAPIDatasource(grafanaApiDatasource, GrafanaAPIDatasourceName, dbeastMinimalCompatibleVersion)
}

const grafanaApiDatasource string = `
{
	"uid": "dbeast-grafanaapi-datasource",
    "name": "%COMPONENT_NAME%",
	"type": "dbeast-grafanaapi-datasource",
	"typeLogoUrl": "public/plugins/dbeast-dbeastmonitor-app/grafana-api-ds/img/favicon-32x32.png",
	"access": "proxy",
	"url": "",
	"user": "",
	"database": "",
	"basicAuth": false,
	"basicAuthUser": "",
	"withCredentials": false,
	"isDefault": false,
	"jsonData": {
		"dbeast_version": "%DBEAST_VERSION%",
		"dbeast_version_int": %DBEAST_APP_FORMAT_VERSION%
	},
	"secureJsonFields": {},
	"version": 1,
	"readOnly": false,
	"apiVersion": ""
}
`
