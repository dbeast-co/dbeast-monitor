# Changelog

## Version 0.7.1
- Bug fixes
  * Logstash Thread pools pipeline - parsing exception [#67](https://github.com/dbeast-co/dbeast-monitor/issues/67)
  * Application can't connect to the clusters without TLS [#64](https://github.com/dbeast-co/dbeast-monitor/issues/64)
  * Error in click to the "DBeast Monitor" In the Grafana[#61](https://github.com/dbeast-co/dbeast-monitor/issues/61)

- Features
  * Move Data source templates to the backend part  [#66](https://github.com/dbeast-co/dbeast-monitor/issues/66)
  * Logstash "elasticsearch-ingest-pipeline" Optimization  [#49](https://github.com/dbeast-co/dbeast-monitor/issues/49)

- Updated dashboards
  * Elasticsearch ingest pipeline overview
  * Elasticsearch ingest pipeline analytics
  * Elastic stack monitor
  
## Version 0.7.0
- Bug fixes
  * Incorrect Aggregation of Pipeline Data in Logstash Overview Dashboard [#48](https://github.com/dbeast-co/dbeast-monitor/issues/48)

- New data sources:
  * Elasticsearch Tasks API datasource

- New dashboards:
  * Elasticsearch Tasks analytics

- Features
  * The backend component moved to the Grafana plugin

- Updated dashboards
  * Elasticsearch host overview
  * Elasticsearch index monitor
  * Elastic stack monitor
  * Logstash instance monitor
  * Logstash overview

## Version 0.6.1
- Bug fixes
  * Incorrect host definition in the logstash-logs configuration file [#43](https://github.com/dbeast-co/dbeast-monitor/issues/43)

## Version 0.6.0
- New data sources:
  * Thread pools datasource

- Updated dashboards
  * Elasticsearch host overview
  * Elasticsearch ingest pipelines overview
  * Elasticsearch ingest pipeline analytics
  * Elasticsearch index monitor
  * Elasticsearch index patterns monitor
  * Elastic stack monitor
  * Logstash instance monitor
  * Logstash overview
  * Logstash pipeline analytics
  * Elasticsearch shard level monitor

## Version 0.5.1
- Fixed bug in the Logstash logs and metrics configuration files


## Version 0.5.0
- New data sources:
  * Metricbeat + Monitoring datasource
  * Logs datasource
  * Elasticsearch Ingest pipeline datasource

- New dashboards:
  * Elasticsearch host overview
  * Elasticsearch ingest pipelines overview
  * Elasticsearch ingest pipeline analytics
  * Elasticsearch index monitor
  * Elasticsearch index patterns monitor

- Updated dashboards
  * Elastic stack monitor
  * Logstash instance monitor
  * Logstash overview
  * Logstash pipeline analytics
  * Elasticsearch shard level monitor


- Name conventions update


## Version 0.4.0-RC

Initial release.