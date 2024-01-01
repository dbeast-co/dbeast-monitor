package co.dbeast.grafana_backend.rest;

import co.dbeast.general_utils.JSONUtils;
import co.dbeast.grafana_backend.data_warehouse.BackendMock;
import co.dbeast.grafana_backend.pojo.new_backend.ClusterConnectionSettings;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import java.io.IOException;

import static spark.Spark.path;
import static spark.Spark.post;

public class RESTTest extends MainRest{
    BackendMock backendMock = BackendMock.getInstance();
    private static final Logger logger = LogManager.getLogger();
    public RESTTest() throws IOException {
        path("/test", () -> {
            post("/test_cluster_new", (request, response) -> {
                logger.info("Got request for new Elasticsearch data source test (NEW)");
                if (logger.isDebugEnabled()) {
                    logger.debug("Request body: " + request.body());
                }
                ClusterConnectionSettings clusterConnectionSettings = JSONUtils.jsonStringToObject(request.body(), ClusterConnectionSettings.class);
                String responseBody = backendMock.test(clusterConnectionSettings);
                logger.info(responseBody);
                return responseBody;
            });
        });
    }
}
