package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "dbeast-mon-index-es-index-stats-status-template"
	const dbeastMinimalCompatibleVersion = "0.0.0"
	dataWarehouse.AppendIndexTemplate(componentName, dbeastMonIndexEsIndexStatsStatusTemplateContent, dbeastMinimalCompatibleVersion)
}

const dbeastMonIndexEsIndexStatsStatusTemplateContent string = `
{
  "index_patterns": [
    "dbeast-mon-index-es-index-stats-status"
  ],
  "template": {
    "aliases": {
      "dbeast-mon-es-index-stats-status": {}
    }
  },
  "composed_of": [
    "dbeast-mon-tsds-es-index-stats-mapping-component",
    "dbeast-mon-default-settings-component"
  ],
  "priority": 1,
  "_meta": {
    "description": "This template used for the index monitoring",
    "dbeast_version": "%DBEAST_VERSION%",
    "dbeast_version_int": %DBEAST_APP_FORMAT_VERSION%
  }
}
`
