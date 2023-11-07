# DBeast monitoring toolkit for Elastic Stack
Welcome to the DBeast Monitor for Elastic Stack – a powerful open-source platform!

This platform provides a powerful UI for monitoring, analyzing, and optimizing of your Elastic Stack components - Elasticsearch, Logstash, and Kibana. This application supports the monitoring of one or more clusters. Most of the data used for the analytics is delivered by Beats and Logstash and stored in Elasticsearch. You can use your production Elasticsearch cluster as the data storage, but we strongly recommend using a dedicated monitoring cluster.

## Documentation
The project documentation is available on our [wiki](https://github.com/dbeast-co/dbeast-monitor/wiki).
- [Installation Instructions](https://github.com/dbeast-co/dbeast-monitor/wiki/InstallationInstructions)
- [Add new cluster](https://github.com/dbeast-co/dbeast-monitor/wiki/AddNewCluster)
- [Dashboards](https://github.com/dbeast-co/dbeast-monitor/wiki/Dashboards)

## Key features
DBeast monitor supports monitoring for multiple clusters. Right now, for each cluster, we have the following monitoring options:
1.	Elastic stack monitoring
2.	Logstash overview monitoring
3.	Logstash instance monitoring
4.	Logstash pipeline monitoring
5.	Index level monitoring
6.	Shards level monitoring
7.	Machine Learning Jobs monitoring

## Requirements
- Grafana 9x
- Java 8+
- Logstash 8.1+
- Elasticsearch 8+
  
## Architecture
![./img/overview/Architecture.jpg](https://github.com/dbeast-co/dbeast-monitor/wiki/img/overview/Architecture.jpg)

1.	Monitored Elastic Stack – Contains all the components of the cluster for monitoring. Right now, we are monitoring the Elasticsearch, Logstash, and Kibana. The Dbeast monitor is a multi-cluster system, therefore you can have more than one Monitored cluster connected to the system.
2.	Monitoring cluster – Dedicated cluster for storing the required data. You can use your monitored cluster also as a monitoring cluster; therefore, this component is not mandatory but strongly recommended.
3.	Grafana – Visualization UI application – This is the place for all analytics dashboards.





