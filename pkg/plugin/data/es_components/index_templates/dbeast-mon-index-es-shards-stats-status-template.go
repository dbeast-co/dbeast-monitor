package index_templates

import (
	"github.com/dbeast/dbeastmonitor/pkg/plugin"
)

const dbeatsMonIndexEsShardsStatsStatusContent string = `
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

var DBEastMonIndexEsShardsStatsStatus = IndexTemplate{
	Name:    "dbeast-mon-index-es-shards-stats-status-template",
	Content: dbeatsMonIndexEsShardsStatsStatusContent,
}

func init() {
	plugin.AppendIndexTemplate(DBEastMonIndexEsShardsStatsStatus.Name, DBEastMonIndexEsShardsStatsStatus.Content)
}
