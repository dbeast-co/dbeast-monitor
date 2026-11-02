package ilm_policies

import (
	data_warehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "dbeast-ds-30d"
	const dbeastMinimalCompatibleVersion = 0
	data_warehouse.AppendILMPolicy(componentName, dbeatsDS30dPolicyContent, dbeastMinimalCompatibleVersion)
}

const dbeatsDS30dPolicyContent string = `
{
  "policy": {
    "_meta": {
      "version": "%DBEAST_VERSION%"
    },
    "phases": {
      "hot": {
        "actions": {
          "rollover": {
            "max_age": "7d",
            "max_docs": 200000000,
            "max_primary_shard_size": "50gb"
          },
          "forcemerge": {
            "max_num_segments": 1
          }
        }
      },
      "delete": {
        "min_age": "30d",
        "actions": {
          "delete": {}
        }
      }
    }
  }
}
`
