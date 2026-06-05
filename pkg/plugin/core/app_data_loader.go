package core

import (
	_ "github.com/dbeast/dbeastmonitor/pkg/plugin/data/es_components/ilm_policies"
	_ "github.com/dbeast/dbeastmonitor/pkg/plugin/data/new_cluster"
)

var ESILMTemplatesMap1 = make(map[string]string)

func GetNewCluster1() []string {
	return ProjectRegistry
}
