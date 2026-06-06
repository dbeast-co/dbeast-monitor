package first_indices

import (
	"github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const dbeatsMonIndexCorruptedDataContent string = `
{
  "aliases": {
    "dbeast-mon-index-corrupted-data": {
      "is_write_index": true
    }
  }
}
`

func init() {
	data.AppendFirstIndex("dbeast-mon-index-corrupted-data", dbeatsMonIndexCorruptedDataContent)
}
