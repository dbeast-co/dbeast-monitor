package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const dbeastMonIndexEsThreadPoolsStatusTemplateContent string = `
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

func init() {
	dataWarehouse.AppendIndexTemplate("dbeast-mon-index-es-thread_pools-status-template", dbeastMonIndexEsThreadPoolsStatusTemplateContent)
}
