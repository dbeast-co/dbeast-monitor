package component_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const dbeastMonDefaultSettingsComponentContent string = `
{
  "template": {
    "settings": {
      "index": {
        "codec": "best_compression",
        "refresh_interval": "30s",
        "number_of_shards": "1",
        "auto_expand_replicas": "0-1"
      }
    },
    "mappings": {
      "properties": {
        "app_version": {
          "type":"keyword"
        }
      }
    }
  },
  "_meta": {
    "description": "The default settings for all monitoring indices",
    "version": "1.0.2"
  }
}
`

func init() {
	dataWarehouse.AppendComponentTemplate("dbeast-mon-default-settings-component", dbeastMonDefaultSettingsComponentContent)
}
