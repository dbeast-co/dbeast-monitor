package co.dbeast.grafana_backend.pojo.new_backend;

import java.util.HashMap;
import java.util.Map;

public class ClusterResponse {
    Map<String, ClusterStatus> prod;
    Map<String, ClusterStatus> mon;

    public ClusterResponse() {
        this.prod = new HashMap<String, ClusterStatus>(){{
            put("elasticsearch", new ClusterStatus());
            put("kibana", new ClusterStatus());
        }};
        this.mon = new HashMap<String, ClusterStatus>(){{
            put("elasticsearch", new ClusterStatus());
        }};
    }

    public Map<String, ClusterStatus> getProd() {
        return prod;
    }

    public void setProd(Map<String, ClusterStatus> prod) {
        this.prod = prod;
    }

    public Map<String, ClusterStatus> getMon() {
        return mon;
    }

    public void setMon(Map<String, ClusterStatus> mon) {
        this.mon = mon;
    }
}
