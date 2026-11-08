package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "dbeast-mon-tsds-es-index-stats-historical-template"
	const dbeastMinimalCompatibleVersion = "0.0.0"
	dataWarehouse.AppendIndexTemplate(componentName, dbeastMonTsdsEsIndexStatsHistoricalTemplateContent, dbeastMinimalCompatibleVersion)
}

const dbeastMonTsdsEsIndexStatsHistoricalTemplateContent string = `
{
  "index_patterns": [
    "dbeast-mon-tsds-es-index-stats"
  ],
  "template": {
    "settings": {
      "index": {
        "lifecycle": {
          "name": "dbeast-tsds-30d",
          "rollover_alias": "dbeast-mon-tsds-es-index-stats"
        },
        "mode": "time_series",
        "routing_path": [
          "elasticsearch.cluster.id"
        ]
      }
    },
    "aliases": {
      "dbeast-mon-es-index-stats": {}
    }
  },
  "composed_of": [
    "dbeast-mon-tsds-es-index-stats-mapping-component",
    "dbeast-mon-default-settings-component"
  ],
  "data_stream": {
    "hidden": false
  },
  "priority": 1,
  "_meta": {
    "description": "This template used for the index monitoring",
    "dbeast_version": "%DBEAST_VERSION%",
    "dbeast_version_int": %DBEAST_APP_FORMAT_VERSION%
  }
}
`
