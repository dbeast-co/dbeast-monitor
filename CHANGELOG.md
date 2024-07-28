# Changelog

## Version 0.8.0
* All data sources have been updated to use aliases, directing them to Data Streams, TSDS, or Indexes.
* Note: There is no backward compatibility with the previous version.
* The minimum required version for the monitoring cluster is now 8.7. (Production clusters can still operate on older versions.)

- Features:
  * Application installer [#37](https://github.com/dbeast-co/dbeast-monitor/issues/37)
  Developed configurations downloader
  * [FEAT] Metric indices to the TSDS migration [#84](https://github.com/dbeast-co/dbeast-monitor/issues/84)
  * [FEAT] Add ILM policies for all Index templates [#32](https://github.com/dbeast-co/dbeast-monitor/issues/32)
  * [FIX] Fix API requests in the "Cluster list"[#23](https://github.com/dbeast-co/dbeast-monitor/issues/23)
  * [FEAT] In the "Tasks API" flow fix running time units [#92](https://github.com/dbeast-co/dbeast-monitor/issues/92)
  * [FEAT] Logstash "Elasticsearch tasks" pipeline - Add ESQL tasks parsing [#73](https://github.com/dbeast-co/dbeast-monitor/issues/73)
  
- Bug fixes
  * [BUG] Backend remove special chars from the cluster name [#100](https://github.com/dbeast-co/dbeast-monitor/issues/100)
  * [BUG] In the Add new cluster panel, the "Kibana host" field is supposed to be not required [#90](https://github.com/dbeast-co/dbeast-monitor/issues/90)

- Updated dashboards
  All dashboards updates to the aliases usage

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