package co.dbeast.elk_clients.elasticsearch.es_settings;

import co.dbeast.general_utils.GeneralUtils;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class ClusterConnectionSourcePOJO extends ClusterConnectionSettingsPOJO implements Cloneable{
    @JsonProperty("source_id")
    private String sourceUid = GeneralUtils.generateNewUID();
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("source_name")
    private String SourceName;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("kibana_host")
    private String kibana_host;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("monitoring_host")
    private String monitoringHost;

    public ClusterConnectionSourcePOJO(final ClusterConnectionSettingsPOJO clusterConnectionSettings,
                                       final String kibanaHost,
                                       final String monitoringHost) {
        super(clusterConnectionSettings);
        this.kibana_host = kibanaHost;
        if (monitoringHost != null) {
            this.monitoringHost = monitoringHost;
        } else {
            this.monitoringHost = getEs_host();
        }
    }

    public ClusterConnectionSourcePOJO() {
        super(new ClusterConnectionSettingsPOJO());
        this.monitoringHost = getEs_host();
        String[] splattedHost = getEs_host().split("//");
        kibana_host = "http://" + splattedHost[1].split(":")[0] + ":5601";
    }

    @Override
    public ClusterConnectionSourcePOJO clone() {
        return (ClusterConnectionSourcePOJO) super.clone();
    }

    @JsonIgnore
    public ClusterConnectionSettingsPOJO getClusterConnectionSettings() {
        return new ClusterConnectionSettingsPOJO(this);
    }
    @JsonIgnore
    public ClusterConnectionSettingsPOJO getKibanaConnectionSettings() {
        return new ClusterConnectionSettingsPOJO(this, "kibana");
    }
    @JsonIgnore
    public ClusterConnectionSettingsPOJO getMonitoringClusterConnectionSettings() {
        return new ClusterConnectionSettingsPOJO(this, "monitor");
    }
    public String getSourceUid() {
        return sourceUid;
    }

    public void setSourceUid(String sourceUid) {
        this.sourceUid = sourceUid;
    }

    public String getSourceName() {
        return SourceName;
    }

    public void setSourceName(String sourceName) {
        SourceName = sourceName;
    }

    public String getKibana_host() {
        return kibana_host;
    }

    public void setKibana_host(String kibana_host) {
        this.kibana_host = kibana_host;
    }

    public String getMonitoringHost() {
        return monitoringHost;
    }

    public void setMonitoringHost(String monitoringHost) {
        this.monitoringHost = monitoringHost;
    }
}
