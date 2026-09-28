package logstash_config

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

const dbeastMonEsXpackUsageContent string = `
input{
  http_poller {
    id => "Input-get-cat-indices-from-ES" 
    urls => {
      stats1 => "<PROD_HOST>_xpack/usage?filter_path=ccr,data_streams,data_tiers,frozen_indices,searchable_snapshots"
    }
	user => "<PROD_USER>"
	password => "<PROD_PASSWORD>"
	
    request_timeout => 60
    schedule => { "cron" => "0 * * * *"}
    codec => "json"
    ssl_verification_mode => "none"
    add_field => {
	  "[elasticsearch][cluster][id]" => "<CLUSTER_ID>"
	  "[event][module]" => "elasticsearch"
	  "[event][dataset]" => "xpack_usage"
	  "[dbeast][schema][version]" => "2.1.2"
	}	
  }
}
filter{
  mutate {
    id => "MUTATE-Rename-fields-to-ecs"
    rename => {
	  "ccr" => "[elasticsearch][xpack][ccr]"
	  "data_streams" => "[elasticsearch][xpack][data_streams]"
	  "data_tiers" => "[elasticsearch][xpack][data_tiers]"
	  "frozen_indices" => "[elasticsearch][xpack][frozen_indices]"
	  "searchable_snapshots" => "[elasticsearch][xpack][searchable_snapshots]"
	}
  }

}
output{
  if [elasticsearch][cluster][id] {
    elasticsearch {
      id => "es-output-send-to-xpack-usage-historical-index"
      hosts => ["<MON_HOST>"]
      user => "<MON_USER>"
      password => "<MON_PASSWORD>"
      ssl_enabled => <MON_SSL_ENABLED>
  	  ssl_verification_mode => none

      index => "dbeast-mon-tsds-es-xpack-usage"
      action => "create"
  	  manage_template => false
    }
  }
  else {
    elasticsearch {
      id => "es-output-send-to-corrupted-data"
  	  hosts => ["<MON_HOST>"]
  	  user => "<MON_USER>"
  	  password => "<MON_PASSWORD>"

  	  ssl_enabled => <MON_SSL_ENABLED>
  	  ssl_verification_mode => none
      index => "dbeast-mon-index-corrupted-data"
  	  manage_template => false
    }
  }
}
`

func init() {
	dataWarehouse.AppendLogstashConfig("dbeast-mon-es-xpack-usage", dbeastMonEsXpackUsageContent)
}
