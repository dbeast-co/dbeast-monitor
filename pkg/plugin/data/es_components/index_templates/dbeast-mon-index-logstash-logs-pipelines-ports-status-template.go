package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const dbeastMonIndexLogstashLogsPipelinesPortsStatusTemplateContent string = `
{
  "index_patterns": [
    "dbeast-mon-index-logstash-logs-pipelines-ports-status"
  ],
  "template": {
    "aliases": {
      "dbeast-mon-logstash-logs-pipelines-ports-status": {}
    }
  },
  "composed_of": [
    "dbeast-mon-ds-logstash-logs-mapping-component",
    "dbeast-mon-default-settings-component"
  ],
  "priority": 1,
  "_meta": {
    "description": "This template used for the Logstash port usage status",
    "version": "1.0.2"
  }
}
`

func init() {
	dataWarehouse.AppendIndexTemplate("dbeast-mon-index-logstash-logs-pipelines-ports-status-template", dbeastMonIndexLogstashLogsPipelinesPortsStatusTemplateContent)
}
