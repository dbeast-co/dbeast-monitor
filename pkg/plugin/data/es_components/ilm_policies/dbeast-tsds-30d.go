package ilm_policies

import (
	"github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

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

func init() {
	data.AppendILMPolicy("dbeast-tsds-30d", dbeatsTSDS30dPolicyContent)
}
