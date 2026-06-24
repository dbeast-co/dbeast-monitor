package plugin

import (
	_ "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
	_ "github.com/dbeast/dbeastmonitor/pkg/plugin/data/data_source_templates"
	_ "github.com/dbeast/dbeastmonitor/pkg/plugin/data/es_components/component_templates"
	_ "github.com/dbeast/dbeastmonitor/pkg/plugin/data/es_components/first_indices"
	_ "github.com/dbeast/dbeastmonitor/pkg/plugin/data/es_components/ilm_policies"
	_ "github.com/dbeast/dbeastmonitor/pkg/plugin/data/es_components/index_templates"
	_ "github.com/dbeast/dbeastmonitor/pkg/plugin/data/logstash_config"
	_ "github.com/dbeast/dbeastmonitor/pkg/plugin/data/new_cluster"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func init() {
	log.DefaultLogger.Info("Data warehouse successfully loaded")
}
