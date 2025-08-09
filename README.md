# "DBeast monitor" for Elastic Stack

[![version: 2.0.0](https://img.shields.io/badge/version-2.0.0-green?style=flat-square)](https://github.com/dbeast-co/dbeast-monitor/releases/latest)
&nbsp;&nbsp;&nbsp;  [![docs](https://img.shields.io/badge/docs-latest-blue?style=flat-square)](https://github.com/dbeast-co/dbeast-monitor/wiki)

[![elastic stack support](https://img.shields.io/badge/contact%20us-support@dbeast.co-blue?style=plastic)](mailto:support@dbeast.co?subject=Elastic%20Stack%20Support%20Request)
&nbsp;&nbsp;&nbsp; [![sponsorship](https://img.shields.io/badge/sposorship-red?style=plastic)](https://github.com/sponsors/dbeast-co?frequency=recurring&sponsor=dbeast-co)

![Cluster Monitor](https://raw.githubusercontent.com/dbeast-co/dbeast-monitor/master/img/screenshots/ClusterMonitoring.jpg)

Welcome to the DBeast Monitor for Elastic Stack – an open-source plugin for Grafana!

Our plugin provides a powerful UI for monitoring, analyzing, and optimizing of your Elastic Stack components -
Elasticsearch, Logstash, and Kibana.

## Key features

1. Multi-cluster support – monitor and manage multiple Elasticsearch clusters from a single application.
2. Comprehensive dashboards providing monitoring and analytics for:
    - **Elastic Stack** overview
    - **Elasticsearch hosts** performance and health
    - **Elasticsearch Indices / index patterns** usage and metrics
    - **Elasticsearch shards** distribution and allocation
    - **Elasticsearch ingest pipelines** activity and throughput
    - **Elasticsearch tasks** progress and status
    - **Elasticsearch machine learning jobs** performance and status
    - **Logstash instances** performance and health
    - **Logstash pipelines** throughput and errors
3. Helm versions
    - The Helm charts can be found in our [Helm repository](https://github.com/dbeast-co/dbeast-monitor-helm)

## Documentation

The project documentation is available on our [wiki](https://github.com/dbeast-co/dbeast-monitor/wiki).

# Getting started

There are two plugin versions available (both are the same plugin but differ based on the environment variable
`DBEAST_MONITOR_VERSION`: either "Container" or "OnPrem" (default: OnPrem):

1. Regular (OnPrem) installation:  
   The OnPrem version includes only the application plugin. You will need to install and define Logstash separately for data
   shipment. The complete installation guid can be found in our [OnPrem installation instructions](https://github.com/dbeast-co/dbeast-monitor/wiki/Installation-Instructions).
2. Helm installation (Container version):  
   The Container version includes Grafana with pre-installed plugins and Logstash for data shipment.
   The complete installation guid can be found in our [Container installation instructions](https://github.com/dbeast-co/dbeast-monitor/wiki/Installation-Instructions-Container-Version).
3. The user guid for the adding new cluster can be found in the [OnPrem Add new cluster instructions](https://github.com/dbeast-co/dbeast-monitor/wiki/AddNewCluster) and [Container add new cluster instructions](https://github.com/dbeast-co/dbeast-monitor/wiki/AddNewClusterContainerVersion).

## Requirements

- Grafana: 10.1.0+
- Logstash: 8.1+
- Elasticsearch monitored cluster: 8+ (Elasticsearch 7X - partial support)
- Elasticsearch monitoring cluster: 8.7+

## Enterprise Support
Need expert help with Elastic Stack architecture, performance tuning, monitoring, or solving complex challenges?
Our team provides hands-on guidance and tailored solutions for your environment.

📩 Get in touch: support@dbeast.co

## For Contributors & Sponsors
We’re always looking for passionate contributors and sponsors to help us improve and grow the application.
Whether you want to contribute code, share ideas, or support the project financially — your help makes a difference.

📩 Reach out: support@dbeast.co

&copy; 2022 - 2025 DBeast. All rights reserved