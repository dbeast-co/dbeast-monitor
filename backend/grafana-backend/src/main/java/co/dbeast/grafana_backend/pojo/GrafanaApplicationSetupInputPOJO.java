package co.dbeast.grafana_backend.pojo;


import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSettingsPOJO;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class GrafanaApplicationSetupInputPOJO extends ClusterConnectionSettingsPOJO implements Cloneable {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("application_host")
    private String applicationHost;
    @JsonProperty("is_replace_keystore")
    private boolean isReplaceKeystore;

    @Override
    public GrafanaApplicationSetupInputPOJO clone() {
        return (GrafanaApplicationSetupInputPOJO) super.clone();
    }

    public ClusterConnectionSettingsPOJO generateClusterConnectionSettings() {
        ClusterConnectionSettingsPOJO clusterConnectionSettings = new ClusterConnectionSettingsPOJO();
        clusterConnectionSettings.setEs_host(getEs_host());
        clusterConnectionSettings.setAuthentication_enabled(isAuthentication_enabled());
        clusterConnectionSettings.setUsername(getUsername());
        clusterConnectionSettings.setPassword(getPassword());
        clusterConnectionSettings.setSsl_enabled(isAuthentication_enabled());
        clusterConnectionSettings.setSsl_file(getSsl_file());
        clusterConnectionSettings.setStatus(getStatus());
        return clusterConnectionSettings;
    }
    public ClusterConnectionSettingsPOJO generateApplicationConnectionSettings() {
        ClusterConnectionSettingsPOJO clusterConnectionSettings = new ClusterConnectionSettingsPOJO();
        clusterConnectionSettings.setEs_host(getApplicationHost());
        clusterConnectionSettings.setAuthentication_enabled(false);
        clusterConnectionSettings.setUsername(null);
        clusterConnectionSettings.setPassword(null);
        clusterConnectionSettings.setSsl_enabled(false);
        clusterConnectionSettings.setSsl_file(null);
        clusterConnectionSettings.setStatus(null);
        return clusterConnectionSettings;
    }

    public String getApplicationHost() {
        return applicationHost;
    }

    public void setApplicationHost(String applicationHost) {
        this.applicationHost = applicationHost;
    }

    @JsonProperty("is_replace_keystore")
    public boolean isReplaceKeystore() {
        return isReplaceKeystore;
    }

    public void setReplaceKeystore(boolean replaceKeystore) {
        isReplaceKeystore = replaceKeystore;
    }
}
