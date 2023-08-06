package co.dbeast.grafana_backend.data_warehouse;

import co.dbeast.elk_clients.grafana.GrafanaDataSourcesDAO;
import com.fasterxml.jackson.databind.JsonNode;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.elasticsearch.client.RestClient;

public class DataSourceController {
    private static final Logger logger = LogManager.getLogger();
    GrafanaDataSourcesDAO grafanaDataSourcesDAO = new GrafanaDataSourcesDAO();
    RestClient grafanaClient;
    private static final DataSourceController _instance = new DataSourceController();

    public DataSourceController() {
    }

    public static DataSourceController getInstance() {
        if (_instance == null) {
            return new DataSourceController();
        }
        return _instance;
    }

    public void init(final RestClient grafanaClient) {
        this.grafanaClient = grafanaClient;
    }


    public boolean isDataSourceExists(final String dataSourceName) {
        return grafanaDataSourcesDAO.isExists(grafanaClient, dataSourceName);
    }

    public boolean inject(final String payload,
                          final String dataSourceName) {
        if (logger.isDebugEnabled()) {
            logger.debug(payload);
        }
        grafanaDataSourcesDAO.upsert(grafanaClient, dataSourceName, payload);
        return true;
    }

    public boolean deleteDataSource(String dataSourceId) {
        JsonNode res;
        try {
            res = grafanaDataSourcesDAO.delete(grafanaClient, dataSourceId);
        } catch (RuntimeException e) {
            return false;
        }
        return "Data source deleted".contains(res.get("message").asText());
    }

}
