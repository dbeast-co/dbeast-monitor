package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "Elasticsearch: "
	const dbeastMinimalCompatibleVersion = "2.0.0"
	dataWarehouse.LoadGrafanaDataSources(testdataDatasourceClusterIdContent, componentName, dbeastMinimalCompatibleVersion)
}

const testdataDatasourceClusterIdContent string = `
{
  "id": 57,
  "uid": "",
  "orgId": 1,
  "name": "%COMPONENT_NAME%",
  "type": "testdata",
  "typeName": "TestData",
  "typeLogoUrl": "public/app/plugins/datasource/testdata/img/testdata.svg",
  "access": "proxy",
  "url": "",
  "user": "",
  "database": "",
  "basicAuth": false,
  "isDefault": false,
  "jsonData": {
	"dbeastVersion": "%DBEAST_VERSION%"
  },
  "readOnly": false
}
`
