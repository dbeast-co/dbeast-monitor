package index_templates

import (
	"github.com/dbeast/dbeastmonitor/pkg/plugin"
)

const dbeatsMonIndexEsIndexStatsStatusContent string = `
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

var DBEastMonIndexEsIndexStatsStatus = IndexTemplate{
	Name:    "dbeast-mon-index-es-index-stats-status-template",
	Content: dbeatsMonIndexEsIndexStatsStatusContent,
}

func init() {
	plugin.AppendIndexTemplate(DBEastMonIndexEsIndexStatsStatus.Name, DBEastMonIndexEsIndexStatsStatus.Content)
}
