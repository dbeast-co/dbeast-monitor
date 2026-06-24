package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

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
    "version": "1.0.2"
  }
}
`

func init() {
	dataWarehouse.AppendIndexTemplate("dbeast-mon-index-es-index-stats-status-template", dbeastMonIndexEsIndexStatsStatusTemplateContent)
}
