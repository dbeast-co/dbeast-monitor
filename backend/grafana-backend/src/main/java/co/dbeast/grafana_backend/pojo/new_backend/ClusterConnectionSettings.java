package co.dbeast.grafana_backend.pojo.new_backend;

import org.elasticsearch.common.recycler.Recycler;

import java.util.HashMap;
import java.util.Map;
import java.util.TreeMap;

public class ClusterConnectionSettings {
    Map<String, ClusterInput> prod = new HashMap<>();
    Map<String, ClusterInput> mon = new HashMap<>();

    public ClusterConnectionSettings() {
        this.prod = new HashMap<String, ClusterInput>(){{
            put("elasticsearch", new ClusterInput());
            put("kibana", new ClusterInput());
        }};
        this.mon = new HashMap<String, ClusterInput>(){{
            put("elasticsearch", new ClusterInput());
        }};
    }

    public Map<String, ClusterInput> getProd() {
        return prod;
    }

    public void setProd(Map<String, ClusterInput> prod) {
        this.prod = prod;
    }

    public Map<String, ClusterInput> getMon() {
        return mon;
    }

    public void setMon(Map<String, ClusterInput> mon) {
        this.mon = mon;
    }
}
