package first_indices

import (
	"github.com/dbeast/dbeastmonitor/pkg/plugin"
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

var DBEastMonIndexEsTasks = FirstIndex{
	Name:    "dbeast-mon-index-es-tasks",
	Content: dbeatsMonIndexEsTasksContent,
}

func init() {
	plugin.AppendFirstIndex(DBEastMonIndexEsTasks.Name, DBEastMonIndexEsTasks.Content)
}
