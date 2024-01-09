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
        if (connectionSettings.getProd().get("elasticsearch").getHost().contains("9200")) {
            responseStatus.getProd().get("elasticsearch").setStatus("green");
        } else if (connectionSettings.getProd().get("elasticsearch").getHost().contains("9201")) {
            responseStatus.getProd().get("elasticsearch").setStatus("yellow");
        } else if (connectionSettings.getProd().get("elasticsearch").getHost().contains("9202")) {
            responseStatus.getProd().get("elasticsearch").setStatus("red");
        } else {
            responseStatus.getProd().get("elasticsearch").setStatus("error");
            responseStatus.getProd().get("elasticsearch").setError("This is the error");
        }

        if (connectionSettings.getProd().get("kibana").getHost().contains("9200")) {
            responseStatus.getProd().get("kibana").setStatus("green");
        } else if (connectionSettings.getProd().get("kibana").getHost().contains("9201")) {
            responseStatus.getProd().get("kibana").setStatus("yellow");
        } else if (connectionSettings.getProd().get("kibana").getHost().contains("9202")) {
            responseStatus.getProd().get("kibana").setStatus("red");
        } else {
            responseStatus.getProd().get("kibana").setStatus("error");
            responseStatus.getProd().get("kibana").setError("This is the error");
        }

        if (connectionSettings.getMon().get("elasticsearch").getHost().contains("9200")) {
            responseStatus.getMon().get("elasticsearch").setStatus("green");
        } else if (connectionSettings.getMon().get("elasticsearch").getHost().contains("9201")) {
            responseStatus.getMon().get("elasticsearch").setStatus("yellow");
        } else if (connectionSettings.getMon().get("elasticsearch").getHost().contains("9202")) {
            responseStatus.getMon().get("elasticsearch").setStatus("red");
        } else {
            responseStatus.getMon().get("elasticsearch").setStatus("error");
            responseStatus.getMon().get("elasticsearch").setError("This is the error");
        }

        return JSONUtils.JSONToString(responseStatus);
    }

    public boolean save(ClusterConnectionSettings connectionSettings) {
        return connectionSettings.getProd().get("elasticsearch").getHost().contains("9200");
    }
}
