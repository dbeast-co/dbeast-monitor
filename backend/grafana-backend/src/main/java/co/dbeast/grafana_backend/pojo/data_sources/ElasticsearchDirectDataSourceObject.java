package co.dbeast.grafana_backend.pojo.data_sources;

import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSourcePOJO;

import java.util.HashMap;

public class ElasticsearchDirectDataSourceObject extends DataSourceObjectPOJO implements Cloneable {
    private String prefix;

    public ElasticsearchDirectDataSourceObject() {
        super();
    }

    public void update(final ClusterConnectionSourcePOJO clusterConnectionSource) {
        super.update(clusterConnectionSource);
        if (prefix.contains("prod")) {
            setUrl(clusterConnectionSource.getClusterConnectionSettings().getEs_host());
            setBasicAuth(clusterConnectionSource.getClusterConnectionSettings().isAuthentication_enabled());
            setBasicAuthUser(clusterConnectionSource.getClusterConnectionSettings().getUsername());
            updatePassword(clusterConnectionSource.getClusterConnectionSettings().getPassword());
        } else {
            setUrl(clusterConnectionSource.getMonitoringHost());
            setBasicAuth(clusterConnectionSource.isMonitoringIsUseAuthentication());
            setBasicAuthUser(clusterConnectionSource.getMonitoringUsername());
            updatePassword(clusterConnectionSource.getMonitoringPassword());
        }
        setName(this.prefix + clusterConnectionSource.getSourceName());
        setUid(this.prefix + clusterConnectionSource.getSourceUid());
    }

    @Override
    public String getDataSourcePrefix() {
        return this.prefix;
    }

    @Override
    public ElasticsearchDirectDataSourceObject clone() {
        ElasticsearchDirectDataSourceObject result = (ElasticsearchDirectDataSourceObject) super.clone();
        result.setJsonData(new HashMap<>(getJsonData()));
        result.setSecureJsonData(new HashMap<>(getSecureJsonData()));
        return result;
    }
    @Override
    public void updateDataSourcePrefix() {
        this.prefix = new String(this.getName());
    }
}

