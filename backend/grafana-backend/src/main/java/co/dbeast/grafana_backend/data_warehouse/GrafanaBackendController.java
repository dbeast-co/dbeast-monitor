package co.dbeast.grafana_backend.data_warehouse;

import co.dbeast.elk_clients.constants.EClusterStatus;
import co.dbeast.elk_clients.elasticsearch.ElasticsearchController;
import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSettingsPOJO;
import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSourcePOJO;
import co.dbeast.elk_clients.exceptions.ClusterConnectionException;
import co.dbeast.general_utils.GeneralUtils;
import co.dbeast.general_utils.JSONUtils;
import co.dbeast.general_utils.keystore.KeystoreController;
import co.dbeast.grafana_backend.constants.EDataSourceTypes;
import co.dbeast.grafana_backend.constants.EGrafanaBackendSettings;
import co.dbeast.grafana_backend.pojo.GrafanaApplicationSetupInputPOJO;
import co.dbeast.grafana_backend.pojo.GrafanaClusterConnectionInputPOJO;
import co.dbeast.grafana_backend.pojo.data_sources.*;
import com.fasterxml.jackson.databind.JsonNode;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.elasticsearch.client.RestClient;

import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class GrafanaBackendController {
    private static final Logger logger = LogManager.getLogger();
    private final ElasticsearchController elasticsearchController = new ElasticsearchController();
    private DataSourceController dataSourceController;

    private static final GrafanaBackendController _instance = new GrafanaBackendController();
    private RestClient grafanaClient;
    private String keystorePassword;
    private GrafanaApplicationSetupInputPOJO appConfig = new GrafanaApplicationSetupInputPOJO();
    private final Map<String, DataSourceObjectPOJO> dataSourceObjectTemplates = new HashMap<>();

    public static GrafanaBackendController getInstance() {
        if (_instance == null) {
            return new GrafanaBackendController();
        }
        return _instance;
    }

    private GrafanaBackendController() {
    }

    public void init(final String keystorePassword) throws ClusterConnectionException {
        this.keystorePassword = keystorePassword;
        if (GeneralUtils.isFileExists(EGrafanaBackendSettings.GRAFANA_CLIENT_CONFIG.getSetting())) {
            ClusterConnectionSettingsPOJO grafanaConfig = GeneralUtils.readFromFileAndSerializeToObject(
                    Paths.get(EGrafanaBackendSettings.GRAFANA_CLIENT_CONFIG.getSetting()), ClusterConnectionSettingsPOJO.class);
            grafanaClient = elasticsearchController.getLowLevelClient(grafanaConfig);
            logger.info("Application connected to the Grafana in address: " + grafanaConfig.getEs_host());
        } else {
            logger.warn("Grafana config doesn't exists");
        }
        DataSourceController.getInstance().init(grafanaClient);
        this.dataSourceController = DataSourceController.getInstance();
        List<Path> filesList = GeneralUtils.readFilesFromFolderPathOneInnerFolder(EGrafanaBackendSettings.GRAFANA_DATASOURCE_FOLDER.getSetting());
        filesList.forEach(filePath -> {
            DataSourceObjectPOJO dataSource = readDatasource(filePath);
            dataSource.updateDataSourcePrefix();
            dataSourceObjectTemplates.put(
                    filePath.toFile().getName(),
                    dataSource
            );
        });
    }


    public boolean saveDataSource(final ClusterConnectionSourcePOJO clusterConnectionSource) {
        try {

            RestClient client;
            try {
                client = elasticsearchController.getLowLevelClient(clusterConnectionSource);
                JsonNode clusterInfo = elasticsearchController.getClusterStats(client);
                clusterConnectionSource.setSourceName(clusterInfo.get("cluster_name").asText() +
                        EGrafanaBackendSettings.CLUSTER_ID_DELIMITER.getSetting() + clusterInfo.get("cluster_uuid").asText());
                clusterConnectionSource.setSourceUid(clusterInfo.get("cluster_name").asText() +
                        EGrafanaBackendSettings.CLUSTER_ID_DELIMITER.getSetting() + clusterInfo.get("cluster_uuid").asText());

                dataSourceObjectTemplates.values().forEach(objectTemplate ->
                        ingestDataSource(clusterConnectionSource, objectTemplate.clone()));

                //TODO check that if the cluster not response there is no error
                return true;
            } catch (ClusterConnectionException e) {
                throw new RuntimeException(e);
            }
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    private void setupKeystore(final String grafanaPassword,
                               final boolean isReplaceExistingKeystore) throws Exception {
        boolean isKeystoreExists = KeystoreController.isKeystoreExists(EGrafanaBackendSettings.KEYSTORE_PATH.getSetting());
        if (isKeystoreExists && !isReplaceExistingKeystore) {
            logger.info("The keystore already exists and you decide does not replace it");
            return;
        } else {
            if (isKeystoreExists) {
                logger.info("The keystore already exists, and you decide to replace it");
                KeystoreController.deleteKeystore(EGrafanaBackendSettings.KEYSTORE_PATH.getSetting());
            }
            KeystoreController.createNewKeystoreEntry(
                    EGrafanaBackendSettings.GRAFANA_USERNAME_FIELD_FOR_KEYSTORE.getSetting(),
                    grafanaPassword,
                    EGrafanaBackendSettings.KEYSTORE_PATH.getSetting(),
                    EGrafanaBackendSettings.APPLICATION_KEYSTORE_PSW.getSetting());
        }
    }

    private boolean setupApplicationDatasources(final GrafanaApplicationSetupInputPOJO applicationConnectionSettings) {
        ApplicationDirectDataSourceObject applicationSourceObjectTemplate = GeneralUtils.readYamlFileFromFileAndSerializeToObject(
                Paths.get(EGrafanaBackendSettings.SETUP_APPLICATION_DATASOURCE.getSetting()), ApplicationDirectDataSourceObject.class);
        GrafanaAPIDataSourceObject grafanaAPISourceObjectTemplate = GeneralUtils.readYamlFileFromFileAndSerializeToObject(
                Paths.get(EGrafanaBackendSettings.SETUP_GRAFANA_API_DATASOURCE.getSetting()), GrafanaAPIDataSourceObject.class);
        if (applicationSourceObjectTemplate != null && grafanaAPISourceObjectTemplate != null) {
            ingestDataSource(
                    new ClusterConnectionSourcePOJO(applicationConnectionSettings.generateApplicationConnectionSettings(), null, null),
                    applicationSourceObjectTemplate);

            ingestDataSource(
                    new ClusterConnectionSourcePOJO(applicationConnectionSettings.generateClusterConnectionSettings(), null, null),
                    grafanaAPISourceObjectTemplate);
            return true;
        } else {
            return false;
        }
    }


    public boolean saveApplicationConfig(final GrafanaApplicationSetupInputPOJO connectionSettings) {
        try {
            grafanaClient = elasticsearchController.getLowLevelClient(connectionSettings.generateClusterConnectionSettings());
            dataSourceController.init(grafanaClient);
            JSONUtils.saveJSONToFile(EGrafanaBackendSettings.GRAFANA_CLIENT_CONFIG.getSetting(),
                    connectionSettings.generateClusterConnectionSettings());
            setupKeystore(connectionSettings.getPassword(), connectionSettings.isReplaceKeystore());
            setupApplicationDatasources(connectionSettings);
            return true;
        } catch (ClusterConnectionException e) {
            return false;
        } catch (Exception e) {
            logger.error(e);
            return false;
        }
    }

    public boolean isSetupCompleted() {
        boolean isSetupComplete = grafanaClient != null;
        if (!isSetupComplete) {
            return false;
        }
        try {
            isSetupComplete = KeystoreController.isKeystoreExists(EGrafanaBackendSettings.KEYSTORE_PATH.getSetting());
            isSetupComplete = isSetupComplete && KeystoreController.isKeyExists(EGrafanaBackendSettings.GRAFANA_USERNAME_FIELD_FOR_KEYSTORE.getSetting(),
                    EGrafanaBackendSettings.KEYSTORE_PATH.getSetting(),
                    EGrafanaBackendSettings.APPLICATION_KEYSTORE_PSW.getSetting()
            );
            isSetupComplete = isSetupComplete && dataSourceController.isDataSourceExists(EGrafanaBackendSettings.APPLICATION_DATA_SOURCE_ID.getSetting());
            isSetupComplete = isSetupComplete && dataSourceController.isDataSourceExists(EGrafanaBackendSettings.APPLICATION_GRAFANA_API_DATA_SOURCE_ID.getSetting());
            return isSetupComplete;
        } catch (Exception e) {
            return false;
        }
    }


    public String getAppConfig() {
        return JSONUtils.JSONToString(appConfig);
    }

    public boolean deleteDataSourceById(String dataSourceId) {
        boolean processingResult = true;
        dataSourceObjectTemplates.values().forEach(objectTemplate -> {
            String dataSourceUid = objectTemplate.getDataSourcePrefix() + dataSourceId;
            dataSourceController.deleteDataSource(dataSourceUid);
        });
        return processingResult;
    }


    public String getClusterStatus(final ClusterConnectionSettingsPOJO connectionSettings,
                                   final String projectId) {
        final ElasticsearchController elasticsearchController = new ElasticsearchController();
        try {
            return "{\"cluster_status\" : \"" + elasticsearchController.getClusterStatus(connectionSettings, projectId) + "\"}";
        } catch (Exception e) {
            return "{\"status\":{\"cluster_status\" : \"" +
                    EClusterStatus.ERROR +
                    "\"},\"error\":\"There is an error while command execution: " +
                    e.getMessage() +
                    "\"}";
        }
    }

    public String getGrafanaStatus(final GrafanaApplicationSetupInputPOJO connectionSettings) {
        try {
            RestClient grafanaTestClient = elasticsearchController.getLowLevelClient(connectionSettings.generateClusterConnectionSettings());
            String result = elasticsearchController.getHeath(grafanaTestClient, "/") == 200 ? "GREEN" : "ERROR";

            return "{\"cluster_status\" : \"" + result + "\"}";
        } catch (Exception e) {
            return "{\"status\":{\"cluster_status\" : \"" +
                    EClusterStatus.ERROR +
                    "\"},\"error\":\"There is an error while command execution: " +
                    e.getMessage() +
                    "\"}";
        }
    }


    private DataSourceObjectPOJO readDatasource(final Path filePath) {
        EDataSourceTypes sourceType = EDataSourceTypes.getByFileName(filePath.toFile().getName());
        switch (sourceType) {
            case APPLICATION_DS: {
                return GeneralUtils.readYamlFileFromFileAndSerializeToObject(filePath, ApplicationDirectDataSourceObject.class);
            }
            case ELASTICSEARCH_DIRECT_DS: {
                return GeneralUtils.readYamlFileFromFileAndSerializeToObject(filePath, ElasticsearchDirectDataSourceObject.class);
            }
            case KIBANA_DIRECT_DS: {
                return GeneralUtils.readYamlFileFromFileAndSerializeToObject(filePath, KibanaDirectDataSourceObject.class);
            }
            case ELASTICSEARCH_MONITORING_DS: {
                return GeneralUtils.readYamlFileFromFileAndSerializeToObject(filePath, ElasticsearchMonitoringDataSourceObject.class);
            }
            default: {
                return null;
            }
        }
    }


    private <T extends DataSourceObjectPOJO> boolean ingestDataSource(ClusterConnectionSourcePOJO clusterConnectionSource,
                                                                      T template) {
        T ingestObject = (T) template.clone();
        ingestObject.update(clusterConnectionSource);
        return dataSourceController.inject(JSONUtils.JSONToString(ingestObject),
                ingestObject.getName());
    }


}

