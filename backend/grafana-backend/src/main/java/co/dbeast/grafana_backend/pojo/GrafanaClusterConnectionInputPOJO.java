package co.dbeast.grafana_backend.pojo;

import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSettingsPOJO;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class GrafanaClusterConnectionInputPOJO extends ClusterConnectionSettingsPOJO {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("kibana_host")
    private String kibanaHost;

    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("monitoring_host")
    private String monitoringHost;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("monitoring_authentication_enabled")
    private boolean monitoringIsUseAuthentication;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("monitoring_username")
    private String monitoringUsername;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("monitoring_password")
    private String monitoringPassword;

    @JsonIgnore
    public ClusterConnectionSettingsPOJO getClusterConnectionSettings() {
        return new ClusterConnectionSettingsPOJO(this);
    }

    @JsonIgnore
    public ClusterConnectionSettingsPOJO getKibanaConnectionSettings() {
        return new ClusterConnectionSettingsPOJO(this);
    }

    public String getKibanaHost() {
        return kibanaHost;
    }

    public void setKibanaHost(String kibanaHost) {
        this.kibanaHost = kibanaHost;
    }

    public String getMonitoringHost() {
        return monitoringHost;
    }

    public void setMonitoringHost(String monitoringHost) {
        this.monitoringHost = monitoringHost;
    }

    public boolean isMonitoringIsUseAuthentication() {
        return monitoringIsUseAuthentication;
    }

    public void setMonitoringIsUseAuthentication(boolean monitoringIsUseAuthentication) {
        this.monitoringIsUseAuthentication = monitoringIsUseAuthentication;
    }

    public String getMonitoringUsername() {
        return monitoringUsername;
    }

    public void setMonitoringUsername(String monitoringUsername) {
        this.monitoringUsername = monitoringUsername;
    }

    public String getMonitoringPassword() {
        return monitoringPassword;
    }

    public void setMonitoringPassword(String monitoringPassword) {
        this.monitoringPassword = monitoringPassword;
    }
}
