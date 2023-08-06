package co.dbeast.elk_clients.elasticsearch.es_settings;


import co.dbeast.elk_clients.constants.EClusterStatus;

import java.util.Objects;

public class ClusterConnectionSettingsPOJO implements Cloneable {
    private String es_host = "http://localhost:9200";
    private boolean authentication_enabled = false;
    private String username;
    private String password;
    private boolean ssl_enabled = false;
    private String ssl_file;
    private EClusterStatus status = EClusterStatus.UNTESTED;

    public ClusterConnectionSettingsPOJO(ClusterConnectionSettingsPOJO clusterConnectionSettings) {
        setEs_host(clusterConnectionSettings.getEs_host());
        setAuthentication_enabled(clusterConnectionSettings.isAuthentication_enabled());
        setUsername(clusterConnectionSettings.getUsername());
        setPassword(clusterConnectionSettings.getPassword());
        setStatus(clusterConnectionSettings.getStatus());
    }

    public ClusterConnectionSettingsPOJO() {
    }
    public ClusterConnectionSettingsPOJO(final ClusterConnectionSourcePOJO clusterConnectionSettings,
                                         final String type) {
        switch (type) {
            case "kibana": {
                setEs_host(clusterConnectionSettings.getKibana_host());
                setAuthentication_enabled(clusterConnectionSettings.isAuthentication_enabled());
                setUsername(clusterConnectionSettings.getUsername());
                setPassword(clusterConnectionSettings.getPassword());
                break;
            }
            case "monitor":{
                setEs_host(clusterConnectionSettings.getMonitoringHost());
                setAuthentication_enabled(clusterConnectionSettings.isMonitoringIsUseAuthentication());
                setUsername(clusterConnectionSettings.getMonitoringUsername());
                setPassword(clusterConnectionSettings.getMonitoringPassword());
                break;
            }
            default:{
                setEs_host(clusterConnectionSettings.getEs_host());
                setAuthentication_enabled(clusterConnectionSettings.isAuthentication_enabled());
                setUsername(clusterConnectionSettings.getUsername());
                setPassword(clusterConnectionSettings.getPassword());
            }
        }

    }

    @Override
    public ClusterConnectionSettingsPOJO clone() {
        try {
            return (ClusterConnectionSettingsPOJO) super.clone();
        } catch (CloneNotSupportedException e) {
            e.printStackTrace();
            return this;
        }
    }

    public String getEs_host() {
        return es_host;
    }

    public void setEs_host(String es_host) {
        this.es_host = es_host;
        if (es_host.contains("https")) {
            this.ssl_enabled = true;
        }
    }

    public boolean isAuthentication_enabled() {
        return authentication_enabled;
    }

    public void setAuthentication_enabled(boolean authentication_enabled) {
        this.authentication_enabled = authentication_enabled;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public boolean isSsl_enabled() {
        return ssl_enabled;
    }

    public void setSsl_enabled(boolean ssl_enabled) {

    }

    public String getSsl_file() {
        return ssl_file;
    }

    public void setSsl_file(String ssl_file) {
        this.ssl_file = ssl_file;
    }

    public EClusterStatus getStatus() {
        return status;
    }

    public void setStatus(EClusterStatus status) {
        this.status = status;
    }

    public String host() {
        String[] splattedHost = es_host.split("//");
        return splattedHost[1].split(":")[0];
    }

    public String port() {
        String[] splattedHost = es_host.split("//");
        return splattedHost[1].split(":")[1];
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        ClusterConnectionSettingsPOJO that = (ClusterConnectionSettingsPOJO) o;
        return authentication_enabled == that.authentication_enabled && ssl_enabled == that.ssl_enabled && Objects.equals(es_host, that.es_host) && Objects.equals(username, that.username) && Objects.equals(password, that.password) && Objects.equals(ssl_file, that.ssl_file) && status == that.status;
    }

    @Override
    public int hashCode() {
        return Objects.hash(es_host, authentication_enabled, username, password, ssl_enabled, ssl_file, status);
    }

    @Override
    public String toString() {
        return "ClusterConnectionSettingsPOJO{" +
                "es_host='" + es_host + '\'' +
                ", authentication_enabled=" + authentication_enabled +
                ", username='" + username + '\'' +
                ", password='" + password + '\'' +
                ", ssl_enabled=" + ssl_enabled +
                ", ssl_file='" + ssl_file + '\'' +
                ", status=" + status +
                '}';
    }
}
