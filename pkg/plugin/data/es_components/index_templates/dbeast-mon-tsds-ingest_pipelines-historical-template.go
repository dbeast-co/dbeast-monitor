package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const dbeastMonTsdsIngestPipelinesHistoricalTemplateContent string = `
{
  "index_patterns": [
    "dbeast-mon-tsds-es-ingest_pipelines"
  ],
  "template": {
    "settings": {
      "index": {
        "lifecycle": {
          "name": "dbeast-tsds-30d",
          "rollover_alias": "dbeast-mon-tsds-es-ingest_pipelines"
        },
        "mode": "time_series",
        "routing_path": [
          "elasticsearch.cluster.id"
        ]
      }
    },
    "aliases": {
      "dbeast-mon-es-ingest_pipelines": {}
    }
  },
  "composed_of": [
    "dbeast-mon-tsds-es-ingest_pipelines-mapping-component",
    "dbeast-mon-default-settings-component"
  ],
  "data_stream": {
    "hidden": false
  },
  "priority": 1,
  "_meta": {
    "description": "This template used for the Elasticsearch pipelines metrics",
    "version": "1.0.2"
  }
}
`

func init() {
	dataWarehouse.AppendIndexTemplate("dbeast-mon-tsds-ingest_pipelines-historical-template", dbeastMonTsdsIngestPipelinesHistoricalTemplateContent)
}
