package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "dbeast-mon-index-corrupted-data-historical-template"
	const dbeastMinimalCompatibleVersion = "0.0.0"
	dataWarehouse.AppendIndexTemplate(componentName, dbeastMonIndexCorruptedDataHistoricalTemplateContent, dbeastMinimalCompatibleVersion)
}

const dbeastMonIndexCorruptedDataHistoricalTemplateContent string = `
{
  "index_patterns": [
    "dbeast-mon-index-corrupted-data-*"
  ],
  "template": {
    "settings": {
      "index": {
        "lifecycle": {
          "name": "dbeast-index-30d",
          "rollover_alias": "dbeast-mon-index-corrupted-data"
        }
      }
    },
    "aliases": {
      "dbeast-mon-corrupted-data": {}
    }
  },
  "composed_of": [
    "dbeast-mon-index-corrupted-data-mapping-component",
    "dbeast-mon-default-settings-component"
  ],
  "priority": 1,
  "_meta": {
    "description": "This template used for the corrupted data, that can't be sent to TSDS",
    "dbeast_version": "%DBEAST_VERSION%",
    "dbeast_version_int": %DBEAST_APP_FORMAT_VERSION%
  }
}
`
