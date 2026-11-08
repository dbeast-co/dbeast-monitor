package index_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "dbeast-mon-index-logstash-metrics-status-template"
	const dbeastMinimalCompatibleVersion = "0.0.0"
	dataWarehouse.AppendIndexTemplate(componentName, dbeastMonIndexLogstashMetricsStatusTemplateContent, dbeastMinimalCompatibleVersion)
}

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
    "dbeast_version": "%DBEAST_VERSION%",
    "dbeast_version_int": %DBEAST_APP_FORMAT_VERSION%
  }
}
`
