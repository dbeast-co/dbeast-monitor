# DBeast monitoring toolkit for Elastic Stack

Welcome to the DBeast Monitor for Elastic Stack – a powerful open-source platform!

This platform provides a powerful UI for monitoring, analyzing, and optimizing of your Elastic Stack components -
Elasticsearch, Logstash, and Kibana. This application supports the monitoring of one or more clusters. Most of the data
used for the analytics is delivered by Beats and Logstash and stored in Elasticsearch. You can use your production
Elasticsearch cluster as the data storage, but we strongly recommend using a dedicated monitoring cluster.

## Helm versions

- [Helm repository](https://github.com/dbeast-co/dbeast-monitor-helm)

## Documentation

The project documentation is available on our [wiki](https://github.com/dbeast-co/dbeast-monitor/wiki).

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
8. Elasticsearch tasks analytics
9. Elasticsearch index level monitor
10. Elasticsearch index patterns monitor
11. Elasticsearch shard level monitor
12. Machine Learning Jobs monitoring

# Getting started

There are two plugin versions available (both are the same plugin but differ based on the environment variable
`DBEAST_MONITOR_VERSION`: either "Container" or "OnPrem"):

1. Regular (OnPrem) installation:  
   The OnPrem version includes only the application plugin. You will need to install and define Logstash separately for data
   shipment.
2. Helm installation (Container version):  
   The Container version includes Grafana with pre-installed plugins and Logstash for data shipment.


## Grafana Installation

### Helm Version
No additional configuration is required after deploying using Helm.

---

### On-Premise Version

#### Grafana Configuration

##### Grafana Folder Locations
- **Linux (Package Manager - yum/apt)**
   - Configuration folder: `/etc/grafana`
   - Plugins folder: `/var/lib/grafana/plugins`

- **Linux (ZIP File Installation)**
   - Configuration folder: `<GRAFANA_DIR>/config`
   - Plugins folder: `<GRAFANA_DIR>/data/plugins`

- **Windows**
   - Configuration folder: `<GRAFANA_DIR>\config`
   - Plugins folder: `<GRAFANA_DIR>\data\plugins`

> **Note**: Replace `<GRAFANA_DIR>` with the actual directory where Grafana is installed.

---

#### Grafana Setup

1. **Install Grafana or Use an Existing Installation**
   - Ensure Grafana version **10.1 or newer** is installed.
   - Follow the [Grafana installation guide](https://grafana.com/docs/grafana/latest/setup-grafana/installation/).

2. **Configure Grafana in Development Mode (Optional)**  
   For unsigned (dev version) plugins, enable development mode or explicitly allow the plugins in the `grafana.ini` file, located in the configuration folder.

   Update the `grafana.ini` file as needed:
   ![Change Grafana to Development Mode](https://raw.githubusercontent.com/dbeast-co/dbeast-monitor/master/img/readme/GrafanaChangeToDev.jpg)

3. **Restart Grafana**

4. **Install Required Plugins**
   - **JSON API Datasource Plugin**
      - Plugin: `marcusolsson-json-datasource`
      - Installation Guide: [JSON API Datasource](https://grafana.com/grafana/plugins/marcusolsson-json-datasource/?tab=installation)

   - **Dynamic Text Plugin**
      - Plugin: `marcusolsson-dynamictext-panel`
      - Installation Guide: [Dynamic Text Plugin](https://grafana.com/grafana/plugins/marcusolsson-dynamictext-panel/?tab=installation)

   - **Business Satellite (Grafana HTTP API) Plugin**
      - Plugin: `volkovlabs-grapi-datasource`
      - Installation Guide: [Business Satellite Plugin](https://grafana.com/grafana/plugins/volkovlabs-grapi-datasource/?tab=overview)

5. **Add DBeast Plugins**
   - Extract the `dbeast-dbeastmonitor-app-<VERSION>.zip` file (or `dbeast-dbeastmonitor-app-<VERSION>.tar.gz` for Linux users) into the Grafana plugins folder.

6. **Restart Grafana**

7. **Verify Plugin Installation**
   - Navigate to the Grafana Plugins page. 
   - Enable the "DBeast Monitor" plugin:
     ![Enable DBeast Plugins](https://raw.githubusercontent.com/dbeast-co/dbeast-monitor/master/img/readme/GrafanaEnablePlugin.jpg)


## Adding a New Cluster
Follow these steps to add and configure a cluster in the Grafana DBeast Monitor application.

---

### 1. Navigate to "Add New Cluster"

- In the Grafana application menu, click **"Add New Cluster"** to begin the configuration process.  
  ![Navigate to Add New Cluster](https://raw.githubusercontent.com/dbeast-co/dbeast-monitor/master/img/readme/AddNewCluster.jpg)

---

### 2. Fill in the Required Fields

- Enter all the required information in the provided fields.

  ![Add New Cluster Page](https://raw.githubusercontent.com/dbeast-co/dbeast-monitor/master/img/readme/AddNewClusterPage.jpg)

    - **Elasticsearch Host**: Provide the URL of one of your **PROD cluster** nodes.
    - **Kibana Host**: Enter the address of your Kibana instance, including the port (optional).
    - **Use Authentication**: Add the username and password for your **PROD cluster**.
    - **Monitoring Host**: Enter the URL of a monitoring cluster node. If monitoring uses the PROD cluster, provide the same PROD URL.
    - **Use Authentication**: Add the username and password for your **MONITORING cluster**.
    - **Test Connection**: Click the **Test** button to check the cluster health.
    - **Add Cluster**: If the connection's status is **GREEN** or **YELLOW**, click **Add** to finalize the addition. If the status is otherwise, troubleshoot the connection.

---

### 3. Monitor Cluster Injection Deployment

- For the application to function correctly, it needs to inject several index templates, ILM policies, and indices definitions into your **MONITORING cluster**.
- This process is performed automatically when clicking the **Deploy** button.
- **Note**: If the monitoring cluster is already used by other clusters, this step can be skipped.

---

### 4. Production Cluster Configuration

- **Logstash Configuration**: The application generates or injects the necessary Logstash configurations for data shipment. Details for the **Helm** or **On-Premise** installations are provided below.

---

### **Helm Version**

- Once the connection status is **GREEN** or **YELLOW**, the **Deploy** button becomes active.
- Press the **Deploy** button to save the configurations directly into the **Helm-deployed Logstash** instance.  
  ![Deploy with Helm Version](https://raw.githubusercontent.com/dbeast-co/dbeast-monitor/master/img/readme/AddNewClusterPageAfterTest.jpg)

---

### **On-Premise (OnPrem) Version**

The OnPrem version contains only the DBeast application plugin. You must manually install and configure Logstash.

1. **Download Configurations**  
   Once the connection status is **GREEN** or **YELLOW**, the **Download** button becomes active. Click it to download the Elasticsearch configurations.  
   ![Download OnPrem Configurations](https://raw.githubusercontent.com/dbeast-co/dbeast-monitor/master/img/readme/AddNewClusterPageAfterTestOnPrem.jpg)

2. **Unpack the Configuration Files**
    - Extract **_ESConfigurationFiles.zip_** and copy the files to your monitoring cluster.
    - For multi-cluster setups, retain the same folder structure from the archive.

3. **Update Pipelines Configuration**
    - Copy the contents of the **_pipelines.yml_** file into the `pipelines.yml` file in your Monitoring Logstash configuration folder.
    - Update the paths in the Logstash configuration to match your folder structure.  
      ![Update Pipelines Configuration](https://raw.githubusercontent.com/dbeast-co/dbeast-monitor/master/img/readme/ChangeConfigPathLS.png)

4. **Restart Logstash (if needed)**
    - Restart Logstash if the `auto.reload` setting is not defined or enabled.

---

## **Congratulations – Your Elasticsearch Monitor is Ready to Use!**

- You can add additional clusters for monitoring by following the same process.

### Setting Up "Logstash Inject Configuration"

To monitor external Logstash instances, you need to add monitoring pipelines to each Logstash instance. Follow the steps below:

---

#### 1. Add Logstash Configurations

Define one or more Logstash instances in your setup.

![Add New Logstash](https://raw.githubusercontent.com/dbeast-co/dbeast-monitor/master/img/readme/AddNewLogstash.png)

- **Server Address**: Provide the unique host address of your Logstash server. This will be used as the folder name in the configuration files.
- **Logstash API**: Specify the Logstash monitoring API address. The default is **localhost:9600**, unless changed in the Logstash configuration.
- **Logstash Logs Folder**: Enter the full path to the Logstash log folder. For Linux systems, the default path is **/var/log/logstash**.

---

#### 2. Set Up Configuration Files

Once you have added one or more Logstash instances:

1. Download the configuration files.
2. Add the downloaded configuration files to each Logstash instance.
3. Include the new configuration paths in the **pipelines.yml** file.

![Download Logstash Configuration](https://raw.githubusercontent.com/dbeast-co/dbeast-monitor/master/img/readme/DownloadLogstashConfig.png)

Next, update the folder path to the Logstash configuration files in the **pipelines.yml** file:

![Update Pipelines Configuration Path](https://raw.githubusercontent.com/dbeast-co/dbeast-monitor/master/img/readme/ChangeConfigPathLS.png)

---

#### 3. Restart Logstash

Restart Logstash if the `auto.reload` option is not defined in your setup.

---

## **Your Logstash Monitor is Now Ready to Use!**

## Requirements

- Grafana 10.1.0+
- Logstash 8.1+
- Elasticsearch 8+ (Elasticsearch 7X - partial support)
- Elasticsearch monitoring cluster 8.7+

## Enterprise Support

Looking for assistance with Elastic Stack architecture, tuning, monitoring, or any other Elastic Stack-related challenges?

📩 Contact us: [support@dbeast.co](mailto:support@dbeast.co)

&copy; 2022 - 2025 DBeast. All rights reserved