# DBeast monitoring toolkit for Elastic Stack
Welcome to the DBeast monitoring toolkit for Elastic Stack - 
The application, that provides powerful UI for the monitoring, analyzing and optimization your Elastic stack components( Elasticsearch, Logstash and Kibana_
The application support one or more clusters monitoring

Most of the data, that we use for the analytics, shipped by the Beats and Logstash and Stored in the Elasticsearch.
You can use your production Elasticsearch cluster for the data store, but we strongly recommend to use dedicated monitoring cluster

## Requirements
- Grafana 9+
- Java 8+
- Logstash 8+
- Elasticsearch 8+

## Installation guide
### Installation files
The toolkit installation contains several parts:
- Application plugin for the Grafana OSS application
- Elasticsearch cluster definition panel for the Grafana OSS application
- Backend application, using for the cluster adding or removing.
- Elasticsearch configurations - files with the index templates for the analytics data indexing
- Logstash configuration - files with the pipelines for the data shipment

### Elasticsearch configuration
1. We recommend to use dedicated monitoring cluster, but you can use your existing cluster
2. Copy the component templates from the all-components.json to the Kibana "Dev tools" and run them
3. Copy the index templates from the all-templates.json to the Kibana "Dev tools" and run them

### Logstash configuration
1. We strongly recommended to use the dedicated Logstash!!! (Not must, but strongly recommended)
2. Unzip Logstash configuration files from logstash_config.zip
3. Update the Logstash configuration files:
     - In all files update the Elasticsearch outputs. Add your MONITORING Elasticsearch cluster requisites (If you have dedicated monitoring cluster.. if no, use your main cluster)
     - In all files set up your Elasticsearch cluster id in the "[elasticsearch][cluster][id]". You can get it from your cluster via API: GET / in the cluster_uuid field
     - In the mon-logstash-logs.conf file, update the path to the Logstash logs in the "path" field (Linux default: /var/log/logstash)
     - in tge mon-es-index-stats.conf, mon-es-thread_pools.conf, mon-es-index-summary.conf files, update your SOURCE Elasticsearch cluster(the cluster, that you want to monitor) requisites
4. Copy all configuration files to your Logstash config folder
5. Add the pipelines definition into the pipelines.yml in the Logstash configuration folder, you can copy it from the pipeline.yml example file
6. Run Logstash

### Grafana configuration
#### Grafana folders:
- For the Linux package manager users (yum, apt):
  Configuration folder - /etc/grafana
  Plugins folder - /var/lib/grafana/plugins
- For the Linux users from zip file
  Configuration folder - <GRAFANA_DOR>/config
  Plugins folder - <GRAFANA_DOR>/data/plugins
- For Windows user
  Configuration folder - <GRAFANA_DOR>\config
  Plugins folder - <GRAFANA_DOR>\data\plugins

#### Backend setup
1. The backend required Java version 8+
2. Unzip backend application from the dbeast-monitoring-for-elastic-stack-VERSION.zip file
3. Edit application configuration file: config/server. You can change your host and port, but can use the default settings
4. Run the backup application:
    - For Linux users: 
      * chown +x dbeast-toolkit.sh 
      * ./dbeast-toolkit.sh (if you want that the application will be run permanently use nohup "./dneast-toolkit.sh &" command)
    - For Windows users:
      * ./dbeast-tookit.bat


#### Grafana setup
1. Install Grafana or use existing one. The application tested with the Grafana 9.x versions
2. The plugin haven't signature (in the beta version), so you have to change the Grafana to the development mode or defined it in to allow plugins list
   Both of them, placed in the default.ini file, that could be found in the Grafana configuration folder.
3. If you haven't "JSON datasource" (marcusolsson-json-datasource) installed in the Grafana you can install it from the Plugins
   or Copy and unzip the marcusolsson-json-datasource.zip into the plugin's folder.
4. If you haven't "Dynamic Text" (marcusolsson-dynamictext-panel) installed in the Grafana you can install it from the Plugins
   or Copy and unzip the marcusolsson-dynamictext-panel.zip into the plugin's folder.
5. If you haven't "Button Panel" (cloudspout-button-panel) installed in the Grafana you can install it from the Plugins or Copy and unzip the cloudspout-button-panel.zip into the plugin folder.
6. Copy and unzip the dbeast-monitoring-for-elastic-stack-app.zip into the plugin's folder:
7. Copy and unzip the dbeast-add_new_es_cluster-panel.zip into the plugin's folder.
8. Enable dbeast-monitoring_for_elastic_stack-app plugin in the plugins setup page
9. In the plugin Configuration page feel the plugin configuration:
    - Grafana host - Your current Grafana requisites
    - Application host - Backend host requisites, that you defined in the backend setup
    - If you already have installed backend, and want to update your Grafana or backend settings, check the "Is replace keystore" checkbox
    - Pres "Test" for check, is the Grafana defined correctly
    - Press "Save" for save settings

