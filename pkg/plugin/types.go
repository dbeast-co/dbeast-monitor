package plugin

// Plugin-specific types. Data-related types (Project, Credentials, etc.) are in pkg/plugin/data

type StatusData struct {
	Prod struct {
		Elasticsearch Status `json:"elasticsearch"`
		Kibana        Status `json:"kibana"`
	} `json:"prod"`
	Mon struct {
		Elasticsearch Status `json:"elasticsearch"`
		Grafana       Status `json:"grafana,omitempty"`
	} `json:"mon"`
}

type Status struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
