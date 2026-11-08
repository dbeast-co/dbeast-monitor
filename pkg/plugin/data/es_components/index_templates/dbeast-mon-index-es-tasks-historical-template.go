package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "dbeast-mon-index-es-tasks-historical-template"
	const dbeastMinimalCompatibleVersion = "0.0.0"
	dataWarehouse.AppendIndexTemplate(componentName, dbeastMonIndexEsTasksHistoricalTemplateContent, dbeastMinimalCompatibleVersion)
}

const dbeastMonIndexEsTasksHistoricalTemplateContent string = `
{
  "index_patterns": [
    "dbeast-mon-index-es-tasks-*"
  ],
  "template": {
    "settings": {
      "index": {
        "lifecycle": {
          "name": "dbeast-index-30d",
          "rollover_alias": "dbeast-mon-index-es-tasks"
        }
      }
    },
    "aliases": {
      "dbeast-mon-es-tasks": {}
    }
  },
  "composed_of": [
    "dbeast-mon-index-es-tasks-mapping-component",
    "dbeast-mon-default-settings-component"
  ],
  "priority": 1,
  "_meta": {
    "description": "This template used for the Elasticsearch tasks",
    "dbeast_version": "%DBEAST_VERSION%",
    "dbeast_version_int": %DBEAST_APP_FORMAT_VERSION%
  }
}
`
