package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

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

func init() {
	dataWarehouse.AppendIndexTemplate("dbeast-mon-index-es-ingest_pipelines-status-template", dbeastMonIndexEsIngestPipelinesStatusTemplateContent)
}
