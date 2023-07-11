package co.dbeast.grafana_backend.rest;

import co.dbeast.general_utils.JSONUtils;
import co.dbeast.grafana_backend.data_warehouse.GrafanaBackendController;
import co.dbeast.grafana_backend.pojo.GrafanaApplicationSetupInputPOJO;
import co.dbeast.rest_server.ARest;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import static spark.Spark.*;

public class RESTGrafanaAppConfig extends ARest {
    private static final Logger logger = LogManager.getLogger();
    private final GrafanaBackendController grafanaBackendController = GrafanaBackendController.getInstance();
    public void rest() {
        path("/setup", () -> {
            get("/is_setup_completed", (request, response) -> {
                logger.info("Got requests for test is application ready to work");
                return grafanaBackendController.isSetupCompleted();
            });

            get("/app_config", (request, response) -> {
                logger.info("Got requests for the application config");
                return grafanaBackendController.getAppConfig();
            });

            post("/test_connection", (request, response) -> {
                logger.info("Got request for connection test");
                if (logger.isDebugEnabled()) {
                    logger.debug("Request body: " + request.body());
                }

                GrafanaApplicationSetupInputPOJO grafanaApplicationSetupInput = JSONUtils.jsonStringToObject(request.body(),  GrafanaApplicationSetupInputPOJO.class);
                String responseBody = grafanaBackendController.getGrafanaStatus(grafanaApplicationSetupInput);
                logger.info(responseBody);
                return responseBody;
            });

            post("/save", (request, response) -> {
                logger.info("Got request for save application settings");
                if (logger.isDebugEnabled()) {
                    logger.debug("Request body: " + request.body());
                }
                String req = request.body();
                GrafanaApplicationSetupInputPOJO grafanaApplicationSetupInput = JSONUtils.jsonStringToObject(req,  GrafanaApplicationSetupInputPOJO.class);
                return grafanaBackendController.saveApplicationConfig(grafanaApplicationSetupInput);
            });
        });
    }

}
