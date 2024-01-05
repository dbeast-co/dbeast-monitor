package co.dbeast.grafana_backend.rest;

import co.dbeast.general_utils.JSONUtils;
import co.dbeast.grafana_backend.data_warehouse.BackendMock;
import co.dbeast.grafana_backend.pojo.new_backend.ClusterConnectionSettings;
import co.dbeast.rest_server.ARest;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import java.io.IOException;

import static spark.Spark.path;
import static spark.Spark.post;

public class RESTTest  extends ARest {
    BackendMock backendMock = BackendMock.getInstance();
    private static final Logger logger = LogManager.getLogger();
    @Override
    public void rest() {
        path("/test", () -> {
            post("/test_cluster", (request, response) -> {
                logger.info("Got request for the new Elasticsearch data source test (NEW)");
                if (logger.isDebugEnabled()) {
                    logger.debug("Request body: " + request.body());
                }
                ClusterConnectionSettings clusterConnectionSettings = JSONUtils.jsonStringToObject(request.body(), ClusterConnectionSettings.class);
                String responseBody = backendMock.test(clusterConnectionSettings);
                logger.info(responseBody);
                return responseBody;
            });
            post("/save", (request, response) -> {
                logger.info("Got request for the new Elasticsearch data source save (NEW)");
                if (logger.isDebugEnabled()) {
                    logger.debug("Request body: " + request.body());
                }
                ClusterConnectionSettings clusterConnectionSettings = JSONUtils.jsonStringToObject(request.body(), ClusterConnectionSettings.class);
                boolean responseBody = backendMock.save(clusterConnectionSettings);
                logger.info(responseBody);
                return responseBody;
            });
        });
    }

}
