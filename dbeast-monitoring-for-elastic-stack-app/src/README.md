# DBeast Monitor
Welcome to the DBeast Monitor for Elastic Stack! This application provides a powerful UI for monitoring, analyzing, and optimizing your Elastic Stack components (Elasticsearch, Logstash, and Kibana).
The application supports monitoring one or more clusters.
Most of the data used for analytics is shipped by Beats and Logstash and stored in Elasticsearch. You can use your production Elasticsearch cluster as the data store, but we strongly recommend using a dedicated monitoring cluster.

## Requirements
- Grafana 9x
- Java 8+
- Logstash 8.1+
- Elasticsearch 8+

## Installation guide
### Installation files
The toolkit installation contains several parts:
- Application plugin for the Grafana OSS application
- Elasticsearch cluster definition panel for the Grafana OSS application
- Backend application used for cluster management (adding or removing)
- Elasticsearch configurations: files with index templates for indexing analytics data
- Logstash configuration: files with pipelines for data shipment

### Elasticsearch configuration
1. We recommend using a dedicated monitoring cluster, but you can use your existing cluster.
2. Copy the component templates from all-components.json to the Kibana "Dev Tools" and execute them.
3. Copy the index templates from all-templates.json to the Kibana "Dev Tools" and execute them.


### Logstash configuration
1. It is strongly recommended to use dedicated Logstash, but it is not mandatory.
2. Unzip the Logstash configuration files from logstash_config.zip.
3. Update the Logstash configuration files:
    - In all configuration files, update the Elasticsearch outputs. Add the requisites of your MONITORING Elasticsearch cluster (If you have a dedicated monitoring cluster, otherwise use your main cluster).
    - In all configuration files, set up your Elasticsearch cluster ID in the input in the "[elasticsearch][cluster][id]" field. You can obtain it from your cluster via the API: GET /, look for the cluster_uuid field.
    - In the mon-logstash-logs.conf file, update the path to the Logstash logs in the "path" field (default path on Linux: /var/log/logstash).
    - In the mon-es-index-stats.conf, mon-es-thread_pools.conf, and mon-es-index-summary.conf, mon-es-shards-stats.conf files, update the requisites of your SOURCE Elasticsearch cluster (the cluster you want to monitor).
4. Copy mon-logstash-logs.conf and mon-logstash-metrics.conf files to EACH monitored Logstash config folders
5. Add the pipeline definitions to the pipelines.yml file in the Logstash configuration folder. You can copy it from the provided example file pipeline_each_logstash.yml.
6. Restart Logstash if the config.reload.automatic: true not defined
7. Copy mon-es-index-stats.conf and mon-es-index-summary.conf configuration files to your MONITORING Logstash config folder.
8. Add the pipeline definitions to the pipelines.yml file in the Logstash configuration folder. You can copy it from the provided example file pipeline_monitoring.yml.
9. Run Logstash

### Grafana configuration
#### Grafana folders:
- For the Linux package manager users (yum, apt):
  Configuration folder - /etc/grafana
  Plugins folder - /var/lib/grafana/plugins
- For the Linux users from zip file
  Configuration folder - <GRAFANA_DIR>/config
  Plugins folder - <GRAFANA_DIR>/data/plugins
- For Windows user
  Configuration folder - <GRAFANA_DIR>\config
  Plugins folder - <GRAFANA_DIR>\data\plugins
  Note: Replace "<GRAFANA_DIR>" with the actual directory where Grafana is installed on your system.

#### Backend setup
1. The backend requires Java version 8 or higher.
2. Unzip the backend application from the dbeast-monitor-VERSION.zip file.
3. Edit the application configuration file: config/server. You can change the host and port if desired, or use the default settings.
4. To run the backend application:
    - For Linux users:
        * Make the dbeast-toolkit.sh script executable: chmod +x dbeast-toolkit.sh
        * Run the script: ./dbeast-toolkit.sh (If you want the application to run permanently, use the command nohup "./dbeast-toolkit.sh &")
    - For Windows users:
        * Run the dbeast-toolkit.bat script: ./dbeast-tookit.bat


#### Grafana setup
1. Install Grafana or use an existing one. The application has been tested with Grafana versions 9.x.
2. The plugin doesn't have a signature (in the beta version), so you have to change Grafana to the development mode or define it to allow plugins in the plugins list.
   Both options are located in the default.ini file, which can be found in the Grafana configuration folder.
3. If you don't have the "JSON datasource" (marcusolsson-json-datasource) installed in Grafana, you can install it from the Plugins section or copy and unzip the marcusolsson-json-datasource.zip file into the plugin's folder.
4. If you don't have the "Dynamic Text" (marcusolsson-dynamictext-panel) installed in Grafana, you can install it from the Plugins section or copy and unzip the marcusolsson-dynamictext-panel.zip file into the plugin's folder.
5. If you don't have the "Button Panel" (cloudspout-button-panel) installed in Grafana, you can install it from the Plugins section or copy and unzip the cloudspout-button-panel.zip file into the plugin's folder.
6. Copy and unzip the dbeast-monitoring-for-elastic-stack-app.zip file into the plugin's folder.
7. Copy and unzip the dbeast-add_new_es_cluster-panel.zip file into the plugin's folder.
8. Enable the dbeast-monitoring_for_elastic_stack-app plugin in the plugin's setup page.
9. In the plugin Configuration page, fill the following settings:
    - Grafana host: Your current Grafana requisites.
    - Application host: Backend host requisites that you defined in the backend setup.
    - If you already have the backend installed and want to update your Grafana or backend settings, check the "Is replace keystore" checkbox.
10. Press "Test" to check if the Grafana is defined correctly.
11. Press "Save" to save the settings.
12. For the Grafana version 9.5. If you want to see the application icon in the home page, add the following configurations in the config ini file:
```
[feature_toggles]
topnav = false
[navigation.app_sections]
dbeast-monitoringforelasticstack-app = root
```

#### Add new cluster
At the time of the new cluster setup, the backend was supposed to be started
1. Click on the "Add new cluster" in the application menu
2. Fill the all required fields
   ![img/new_cluster.png](img/new_cluster.png)
    - Elasticsearch host - The address of one of your PROD cluster nodes
    - Kibana host - The address of your Kibana (include port)
    - Use authentication, user, and password for your PROD cluster
    - Monitoring host - The address of one of your MONITORING cluster nodes (if you're using the PROD cluster as monitoring, you supposed to fill the PROD requisites)
    - Use authentication, user, and password for your MONITORING cluster
3. Press "Test" for the cluster health check
4. Press "Save" for adding a cluster to the application.