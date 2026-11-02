package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "dbeast-mon-index-es-thread_pools-status-template"
	const dbeastMinimalCompatibleVersion = 0
	dataWarehouse.AppendIndexTemplate(componentName, dbeastMonIndexEsThreadPoolsStatusTemplateContent, dbeastMinimalCompatibleVersion)
}

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
