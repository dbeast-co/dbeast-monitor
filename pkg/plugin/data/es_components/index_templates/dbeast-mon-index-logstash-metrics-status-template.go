package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const dbeastMonIndexLogstashMetricsStatusTemplateContent string = `
{
  "index_patterns": [
    "dbeast-mon-index-logstash-metrics-status"
  ],
  "template": {
    "aliases": {
      "dbeast-mon-logstash-metrics-status": {}
    }
  },
  "composed_of": [
    "dbeast-mon-tsds-logstash-metrics-mapping-component",
    "dbeast-mon-default-settings-component"
  ],
  "priority": 1,
  "_meta": {
    "description": "This template used for the Logstash metrics",
    "version": "1.0.2"
  }
}
`

func init() {
	dataWarehouse.AppendIndexTemplate("dbeast-mon-index-logstash-metrics-status-template", dbeastMonIndexLogstashMetricsStatusTemplateContent)
}
