package co.dbeast.grafana_backend.rest;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import java.io.IOException;

import static spark.Spark.*;

public class MainRest {
    private static final Logger logger = LogManager.getLogger();

    public MainRest() throws IOException {
        initRestAPIs();
    }

    /**
     * Initialize all REST APIs
     */
    public void initRestAPIs() {
        logger.info("Initialize Grafana backend APIs");
        RESTElasticsearchDataSource restElasticsearchDataSource = new RESTElasticsearchDataSource();
        RESTGrafanaAppConfig restGrafanaAppConfig = new RESTGrafanaAppConfig();
        RESTTest restTest = new RESTTest();
        path("/grafana_backend", () -> {
            restElasticsearchDataSource.rest();
            restGrafanaAppConfig.rest();
            restTest.rest();
            get("/test", (request, response) -> {
                logger.info("Got test request for Grafana backend");
                return "{\"grafana_backend_test\": true}";
            });
        });

    }

}
