package co.dbeast.grafana_backend.pojo.data_sources;

import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSourcePOJO;

import java.util.HashMap;

public class KibanaDirectDataSourceObject extends DataSourceObjectPOJO implements Cloneable{
    private String prefix;
    public KibanaDirectDataSourceObject() {
        super();
        getJsonData().put("httpHeaderName1", "kbn-xsrf");
        getSecureJsonData().put("httpHeaderValue1", "true");
    }

    public void update(final ClusterConnectionSourcePOJO clusterConnectionSource) {
        super.update(clusterConnectionSource);
        setUrl(clusterConnectionSource.getKibana_host());
        setName(this.prefix + clusterConnectionSource.getSourceName());
        setUid(this.prefix + clusterConnectionSource.getSourceUid());
        setBasicAuth(clusterConnectionSource.getClusterConnectionSettings().isAuthentication_enabled());
        setBasicAuthUser(clusterConnectionSource.getClusterConnectionSettings().getUsername());
        updatePassword(clusterConnectionSource.getClusterConnectionSettings().getPassword());
    }

    @Override
    public String getDataSourcePrefix() {
        return this.prefix;
    }

    @Override
    public KibanaDirectDataSourceObject clone() {
        KibanaDirectDataSourceObject result = (KibanaDirectDataSourceObject) super.clone();
        result.setJsonData(new HashMap<>(getJsonData()));
        result.setSecureJsonData(new HashMap<>(getSecureJsonData()));
        return result;
    }
    @Override
    public void updateDataSourcePrefix() {
        this.prefix = new String(this.getName());
    }
}
