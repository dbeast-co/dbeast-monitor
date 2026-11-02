package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "dbeast-mon-index-es-ingest_pipelines-status-template"
	const dbeastMinimalCompatibleVersion = 0
	dataWarehouse.AppendIndexTemplate(componentName, dbeastMonIndexEsIngestPipelinesStatusTemplateContent, dbeastMinimalCompatibleVersion)
}

const dbeastMonIndexEsIngestPipelinesStatusTemplateContent string = `
{
  "index_patterns": [
    "dbeast-mon-index-es-ingest_pipelines-status"
  ],
  "template": {
    "aliases": {
      "dbeast-mon-es-ingest_pipelines-status": {}
    }
  },
  "composed_of": [
    "dbeast-mon-tsds-es-ingest_pipelines-mapping-component",
    "dbeast-mon-default-settings-component"
  ],
  "priority": 1,
  "_meta": {
    "description": "This template used for the Elasticsearch pipelines metrics",
    "version": "1.0.2"
  }
}
`
