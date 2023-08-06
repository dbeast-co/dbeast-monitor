package co.dbeast.grafana_backend.rest;

import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSourcePOJO;
import co.dbeast.general_utils.JSONUtils;
import co.dbeast.grafana_backend.data_warehouse.GrafanaBackendController;
import co.dbeast.grafana_backend.pojo.GrafanaClusterConnectionInputPOJO;
import co.dbeast.rest_server.ARest;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import static spark.Spark.*;

public class RESTElasticsearchDataSource extends ARest {
    private static final Logger logger = LogManager.getLogger();
    private final GrafanaBackendController grafanaBackendController = GrafanaBackendController.getInstance();

    public void rest() {
        path("/data_sources", () -> {
            post("/test_cluster", (request, response) -> {
                logger.info("Got request for new Elasticsearch data source test");
                if (logger.isDebugEnabled()) {
                    logger.debug("Request body: " + request.body());
                }
                GrafanaClusterConnectionInputPOJO clusterConnectionSettings = JSONUtils.jsonStringToObject(request.body(), GrafanaClusterConnectionInputPOJO.class);
                clusterConnectionSettings.setSsl_file(null);
                String responseBody = grafanaBackendController.getClusterStatus(clusterConnectionSettings, "1234");
                logger.info(responseBody);
                return responseBody;
            });
            get("/delete/:id", (request, response) -> {
                logger.info("Got request for delete project with Id: " + request.params(":id"));
                return grafanaBackendController.deleteDataSourceById(request.params(":id"));
            });
            post("/save", (request, response) -> {
                logger.info("Got request for save project");
                if (logger.isDebugEnabled()) {
                    logger.debug("Request body: " + request.body());
                }
                ClusterConnectionSourcePOJO grafanaClusterConnectionSettings = JSONUtils.jsonStringToObject(request.body(), ClusterConnectionSourcePOJO.class);
                return grafanaBackendController.saveDataSource(grafanaClusterConnectionSettings);
            });
        });
    }


}
