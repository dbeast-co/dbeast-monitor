package co.dbeast.elk_clients.ds_warehouse;

import co.dbeast.elk_clients.constants.EElkClients;
import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSourcePOJO;
import co.dbeast.general_utils.GeneralUtils;
import co.dbeast.general_utils.keystore.KeystoreController;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import java.nio.file.Path;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class ESClustersConnectionWarehouse {
    private static final Logger logger = LogManager.getLogger();
    private final Map<String, ClusterConnectionSourcePOJO> clustersConnectionSourceMap = new HashMap<>();

    private static final ESClustersConnectionWarehouse _instance = new ESClustersConnectionWarehouse();

    public static synchronized ESClustersConnectionWarehouse getInstance() {
        if (_instance == null) {
            return new ESClustersConnectionWarehouse();
        }
        return _instance;
    }

    private ESClustersConnectionWarehouse() {

    }

    /**
     * Initialize all existing projects
     */
    public void init(final String keystoreLocation,
                     final String keystorePassword) {
        GeneralUtils.createFolderIfNotExists(EElkClients.CLUSTERS_CONNECTION_SOURCE_FOLDER.getSetting());
        List<Path> clustersFilesList = GeneralUtils.readFilesFromFolderPathOneInnerFolder(EElkClients.CLUSTERS_CONNECTION_SOURCE_FOLDER.getSetting());
        clustersFilesList.forEach(cluster -> readDataSource(
                cluster,
                keystoreLocation,
                keystorePassword));
    }
    /**
     * Initialize first run
     */
    public void init() {
        GeneralUtils.createFolderIfNotExists(EElkClients.CLUSTERS_CONNECTION_SOURCE_FOLDER.getSetting());
    }

    public Map<String, ClusterConnectionSourcePOJO> getClustersConnectionSourceMap() {
        return clustersConnectionSourceMap;
    }

    private void readDataSource(final Path file,
                                final String keystoreLocation,
                                final String keystorePassword) {
        ClusterConnectionSourcePOJO project = GeneralUtils.readFromFileAndSerializeToObject(file, ClusterConnectionSourcePOJO.class);
        if (project != null) {
            try {
                if (project.getClusterConnectionSettings().isAuthentication_enabled()) {
                    String password = null;
                    password = KeystoreController.getPasswordFromKeystore(project.getSourceUid(),
                            keystoreLocation,
                            keystorePassword
                    );
                    project.getClusterConnectionSettings().setPassword(password);
                    project.getKibanaConnectionSettings().setPassword(password);
                }
                clustersConnectionSourceMap.put(project.getSourceUid(), project);
            } catch (Exception e) {
                logger.error("There is an error in project loading");
                logger.error("Project name: " + project.getSourceName() + ". Project id: " + project.getSourceUid());
                logger.error(e);
            }
        }
    }
}
