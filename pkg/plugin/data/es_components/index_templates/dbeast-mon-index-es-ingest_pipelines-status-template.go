package index_templates

import (
	"github.com/dbeast/dbeastmonitor/pkg/plugin"
)

const dbeatsMonIndexEsIngestPipelinesStatusContent string = `
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

var DBEastMonIndexEsIngestPipelinesStatus = IndexTemplate{
	Name:    "dbeast-mon-index-es-ingest_pipelines-status-template",
	Content: dbeatsMonIndexEsIngestPipelinesStatusContent,
}

func init() {
	plugin.AppendIndexTemplate(DBEastMonIndexEsIngestPipelinesStatus.Name, DBEastMonIndexEsIngestPipelinesStatus.Content)
}
