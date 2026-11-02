package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "dbeast-mon-ds-logstash-logs-historical-template"
	const dbeastMinimalCompatibleVersion = "2.0.0"
	dataWarehouse.AppendIndexTemplate(componentName, dbeastMonDsLogstashLogsHistoricalTemplateContent, dbeastMinimalCompatibleVersion)
}

const dbeastMonDsLogstashLogsHistoricalTemplateContent string = `
{
  "index_patterns": [
    "dbeast-mon-ds-logstash-logs"
  ],
  "template": {
    "settings": {
      "index": {
        "lifecycle": {
          "name": "dbeast-ds-30d",
          "rollover_alias": "dbeast-mon-ds-logstash-logs"
        }
      }
    },
    "aliases": {
      "dbeast-mon-logstash-logs": {}
    }
  },
  "composed_of": [
    "dbeast-mon-ds-logstash-logs-mapping-component",
    "dbeast-mon-default-settings-component"
  ],
  "data_stream": {
    "hidden": false
  },
  "priority": 1,
  "_meta": {
    "description": "This template used for the Logstash logs",
    "version": "%DBEAST_VERSION%"
  }
}
`
