package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

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

func init() {
	dataWarehouse.AppendIndexTemplate("dbeast-mon-index-es-shards-stats-status-template", dbeastMonIndexEsShardsStatsStatusTemplateContent)
}
