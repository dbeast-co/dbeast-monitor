package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "dbeast-mon-index-es-shards-stats-status-template"
	const dbeastMinimalCompatibleVersion = 0
	dataWarehouse.AppendIndexTemplate(componentName, dbeastMonIndexEsShardsStatsStatusTemplateContent, dbeastMinimalCompatibleVersion)
}

const dbeastMonIndexEsShardsStatsStatusTemplateContent string = `
{
  "index_patterns": [
    "dbeast-mon-index-es-shards-stats-status"
  ],
  "template": {
    "aliases": {
      "dbeast-mon-es-shards-stats-status": {}
    }
  },
  "composed_of": [
    "dbeast-mon-tsds-es-shards-stats-mapping-component",
    "dbeast-mon-default-settings-component"
  ],
  "priority": 1,
  "_meta": {
    "description": "This template used for the shards monitoring",
    "version": "1.0.2"
  }
}
`
