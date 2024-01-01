package co.dbeast.grafana_backend.pojo.new_backend;

import org.elasticsearch.common.recycler.Recycler;

import java.util.HashMap;
import java.util.Map;
import java.util.TreeMap;

public class ClusterConnectionSettings {
    Map<String, ClusterInput> prod;
    Map<String, ClusterInput> mon;

    public ClusterConnectionSettings() {
        this.prod = new HashMap<String, ClusterInput>(){{
            put("elasticsearch", new ClusterInput());
            put("kibana", new ClusterInput());
        }};
        this.mon = new HashMap<String, ClusterInput>(){{
            put("elasticsearch", new ClusterInput());
        }};
    }
}
