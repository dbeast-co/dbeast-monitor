package co.dbeast.grafana_backend.data_warehouse;

import co.dbeast.general_utils.JSONUtils;
import co.dbeast.grafana_backend.pojo.new_backend.ClusterConnectionSettings;
import co.dbeast.grafana_backend.pojo.new_backend.ClusterResponse;

public class BackendMock {
    private static final BackendMock _instance = new BackendMock();

    public static BackendMock getInstance() {
        if (_instance == null) {
            return new BackendMock();
        }
        return _instance;
    }

    public String test(ClusterConnectionSettings connectionSettings) {
        ClusterResponse responseStatus = new ClusterResponse();
        return JSONUtils.JSONToString(responseStatus);
    }
}
