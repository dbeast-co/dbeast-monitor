package ilm_policies

import (
	data_warehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "dbeast-tsds-30d"
	const dbeastMinimalCompatibleVersion = 0
	data_warehouse.AppendILMPolicy(componentName, dbeatsTSDS30dPolicyContent, dbeastMinimalCompatibleVersion)
}

const dbeatsTSDS30dPolicyContent string = `
{
  "policy": {
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
