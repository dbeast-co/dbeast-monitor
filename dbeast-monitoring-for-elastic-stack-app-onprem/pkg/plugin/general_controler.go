package plugin

import (
	"encoding/json"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type Cluster struct {
	ClusterConnectionSettings EnvironmentConfig      `json:"cluster_connection_settings"`
	LogstashConfigurations    LogstashConfigurations `json:"logstash_configurations"`
}

/*
CloneObject creates a deep copy of the provided JSON template.
It marshals the template into JSON and then unmarshal it into a new interface{} to create a clone.
*/
func CloneObject(data interface{}) interface{} {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.DefaultLogger.Warn("Failed to marshal data: " + err.Error())
		return data
	}

	var clonedTemplate interface{}
	if err := json.Unmarshal(dataBytes, &clonedTemplate); err != nil {
		log.DefaultLogger.Warn("Failed to unmarshal cloned data: " + err.Error())
		return data
	}

	return clonedTemplate
}
