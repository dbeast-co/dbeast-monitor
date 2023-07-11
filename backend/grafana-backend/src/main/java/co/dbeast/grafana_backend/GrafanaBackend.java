package co.dbeast.grafana_backend;

import co.dbeast.elk_clients.ElkClients;
import co.dbeast.elk_clients.ds_warehouse.ESClustersConnectionWarehouse;
import co.dbeast.elk_clients.exceptions.ClusterConnectionException;
import co.dbeast.general_utils.GeneralUtils;
import co.dbeast.grafana_backend.constants.EGrafanaBackendSettings;
import co.dbeast.grafana_backend.data_warehouse.GrafanaBackendController;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import java.io.IOException;

public class GrafanaBackend {
    public static String HOME_FOLDER;
    private static final Logger logger = LogManager.getLogger();

    public static void main(String[] args) {
        logger.info("Welcome to Grafana backend");
    }

    public static void init(final String homeFolder,
                            final String adminPassword) throws ClusterConnectionException {
        HOME_FOLDER = homeFolder;
        ElkClients.getInstance().init(homeFolder);
        if (!GeneralUtils.isFileExists(EGrafanaBackendSettings.KEYSTORE_PATH.getSetting())) {
            logger.warn("The keystore file: " + EGrafanaBackendSettings.KEYSTORE_PATH.getSetting() +
                    " doesn't exists! Please setup application from configuration page");
            ESClustersConnectionWarehouse.getInstance().init();
        } else {
            ESClustersConnectionWarehouse.getInstance().init(
                    EGrafanaBackendSettings.KEYSTORE_PATH.getSetting(),
                    adminPassword);
        }
        GrafanaBackendController.getInstance().init(adminPassword);
        logger.debug("Grafana backend initialization successfully completed");
    }
}