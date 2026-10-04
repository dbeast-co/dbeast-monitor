package data_source_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const testdataDatasourceClusterIdContent string = `
{
  "id": 57,
  "uid": "",
  "orgId": 1,
  "name": "Elasticsearch: ",
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
	"dbeastVersion": "%DBEAST_VERSION%",
	"dbeastMinimalCompatibleVersion": "2.0.0"
  },
  "readOnly": false
}
`

func init() {
	dataWarehouse.LoadGrafanaDataSources(testdataDatasourceClusterIdContent)
}
