# DBeast monitoring toolkit for Elastic Stack

Welcome to the DBeast Monitor for Elastic Stack – a powerful open-source platform!

This platform provides a powerful UI for monitoring, analyzing, and optimizing of your Elastic Stack components -
Elasticsearch, Logstash, and Kibana. This application supports the monitoring of one or more clusters. Most of the data
used for the analytics is delivered by Beats and Logstash and stored in Elasticsearch. You can use your production
Elasticsearch cluster as the data storage, but we strongly recommend using a dedicated monitoring cluster.

## Documentation

The project documentation is available on our [wiki](https://github.com/dbeast-co/dbeast-monitor/wiki).

- [Installation Instructions](https://github.com/dbeast-co/dbeast-monitor/wiki/InstallationInstructions)
- [Add new cluster](https://github.com/dbeast-co/dbeast-monitor/wiki/AddNewCluster)
- [Dashboards](https://github.com/dbeast-co/dbeast-monitor/wiki/Dashboards)

## Key features

DBeast monitor supports monitoring for multiple clusters. Right now, for each cluster, we have the following monitoring
options:

1. Elastic stack monitor
2. Elasticsearch host overview
3. Logstash overview
4. Logstash instance monitor
5. Logstash pipeline analytics
6. Elasticsearch ingest pipelines overview
7. Elasticsearch ingest pipeline analytics
8. Elasticsearch index level monitor
9. Elasticsearch index patterns monitor
10. Elasticsearch shard level monitor
11. Machine Learning Jobs monitoring

## Requirements

- Grafana 9x
- Java 8+
- Logstash 8.1+
- Elasticsearch 8+
