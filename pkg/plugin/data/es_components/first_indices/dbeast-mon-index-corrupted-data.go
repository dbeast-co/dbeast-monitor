package first_indices

import (
	"github.com/dbeast/dbeastmonitor/pkg/plugin"
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

var DBEastMonIndexCorruptedData = FirstIndex{
	Name:    "dbeast-mon-index-corrupted-data",
	Content: dbeatsMonIndexCorruptedDataContent,
}

func init() {
	plugin.AppendFirstIndex(DBEastMonIndexCorruptedData.Name, DBEastMonIndexCorruptedData.Content)
}
