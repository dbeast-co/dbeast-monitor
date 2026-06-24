package first_indices

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const dbeatsMonIndexEsTasksContent string = `
{
  "aliases": {
    "dbeast-mon-index-es-tasks": {
      "is_write_index": true
    }
  }
}
`

func init() {
	dataWarehouse.AppendFirstIndex("dbeast-mon-index-es-tasks", dbeatsMonIndexEsTasksContent)
}
