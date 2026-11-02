package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "dbeast-mon-index-logstash-logs-pipelines-ports-status-template"
	const dbeastMinimalCompatibleVersion = 0
	dataWarehouse.AppendIndexTemplate(componentName, dbeastMonIndexLogstashLogsPipelinesPortsStatusTemplateContent, dbeastMinimalCompatibleVersion)
}

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
