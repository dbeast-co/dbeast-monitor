package index_templates

import (
	"github.com/dbeast/dbeastmonitor/pkg/plugin"
)

const dbeatsMonIndexEsThreadPoolsStatusContent string = `
{
  "index_patterns": [
    "dbeast-mon-index-es-thread_pools-status"
  ],
  "template": {
    "aliases": {
      "dbeast-mon-es-thread_pools-status": {}
    }
  },
  "composed_of": [
    "dbeast-mon-tsds-es-thread_pools-mapping-component",
    "dbeast-mon-default-settings-component"
  ],
  "priority": 1,
  "_meta": {
    "description": "This template used for the Elasticsearch thread pools metrics",
    "version": "1.0.2"
  }
}
`

var DBEastMonIndexEsThreadPoolsStatus = IndexTemplate{
	Name:    "dbeast-mon-index-es-thread_pools-status-template",
	Content: dbeatsMonIndexEsThreadPoolsStatusContent,
}

func init() {
	plugin.AppendIndexTemplate(DBEastMonIndexEsThreadPoolsStatus.Name, DBEastMonIndexEsThreadPoolsStatus.Content)
}
