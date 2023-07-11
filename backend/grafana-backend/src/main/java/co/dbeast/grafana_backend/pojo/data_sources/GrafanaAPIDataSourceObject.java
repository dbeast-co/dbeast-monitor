package co.dbeast.grafana_backend.pojo.data_sources;

import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSourcePOJO;

import java.util.HashMap;

public class GrafanaAPIDataSourceObject extends DataSourceObjectPOJO implements Cloneable{

    public void update(final ClusterConnectionSourcePOJO clusterConnectionSource) {
        super.update(clusterConnectionSource);
        setUrl(clusterConnectionSource.getClusterConnectionSettings().getEs_host() +
                this.getUrl());
        setBasicAuth(clusterConnectionSource.getClusterConnectionSettings().isAuthentication_enabled());
        setBasicAuthUser(clusterConnectionSource.getClusterConnectionSettings().getUsername());
        updatePassword(clusterConnectionSource.getClusterConnectionSettings().getPassword());
        setName(getDataSourcePrefix());
        setUid(getDataSourcePrefix());
    }

    @Override
    public String getDataSourcePrefix() {
        return "Grafana-DS-API";
    }

    @Override
    public GrafanaAPIDataSourceObject clone() {
        GrafanaAPIDataSourceObject result = (GrafanaAPIDataSourceObject) super.clone();
        result.setJsonData(new HashMap<>(getJsonData()));
        result.setSecureJsonData(new HashMap<>(getSecureJsonData()));
        return result;
    }
    @Override
    public void updateDataSourcePrefix() {
    }
}
