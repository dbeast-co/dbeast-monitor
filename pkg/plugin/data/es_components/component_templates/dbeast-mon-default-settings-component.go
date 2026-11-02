package component_templates

import dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"

func init() {
	const componentName = "dbeast-mon-default-settings-component"
	const dbeastMinimalCompatibleVersion = "2.0.0"
	dataWarehouse.AppendComponentTemplate(componentName, dbeastMonDefaultSettingsComponentContent, dbeastMinimalCompatibleVersion)
}

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
      "dynamic_templates": [
        {
          "strings_as_keyword": {
            "match_mapping_type": "string",
            "mapping": {
              "type": "keyword"
            }
          }
        }
      ],
      "properties": {
		"dbeast": {
		  "properties": {
			"version": {
			  "type":"keyword"
			}
		  }
		},
        "app_version": {
          "type":"keyword"
        }
      }
    }
  },
  "_meta": {
    "description": "The default settings for all monitoring indices",
    "version": "%DBEAST_VERSION%"
  }
}
`
