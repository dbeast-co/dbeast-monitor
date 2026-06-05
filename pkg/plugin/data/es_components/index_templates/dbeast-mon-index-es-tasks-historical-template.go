package index_templates

import (
	"github.com/dbeast/dbeastmonitor/pkg/plugin"
)

const dbeatsMonIndexEsTasksHistoricalContent string = `
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
    "version": "1.0.2"
  }
}
`

var DBEastMonIndexEsTasksHistorical = IndexTemplate{
	Name:    "dbeast-mon-index-es-tasks-historical-template",
	Content: dbeatsMonIndexEsTasksHistoricalContent,
}

func init() {
	plugin.AppendIndexTemplate(DBEastMonIndexEsTasksHistorical.Name, DBEastMonIndexEsTasksHistorical.Content)
}
