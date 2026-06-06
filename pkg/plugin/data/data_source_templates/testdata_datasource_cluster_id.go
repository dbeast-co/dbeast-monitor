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
  "jsonData": {},
  "readOnly": false
}
`

func init() {
	dataWarehouse.LoadGrafanaDataSources("testdata_datasource_cluster_id", testdataDatasourceClusterIdContent)
}
