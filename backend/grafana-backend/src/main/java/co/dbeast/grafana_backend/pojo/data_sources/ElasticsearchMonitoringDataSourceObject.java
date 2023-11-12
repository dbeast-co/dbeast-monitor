package co.dbeast.grafana_backend.pojo.data_sources;

import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSourcePOJO;

import java.util.HashMap;

public class ElasticsearchMonitoringDataSourceObject extends DataSourceObjectPOJO {

    private String database;
    private String prefix;

    public ElasticsearchMonitoringDataSourceObject() {
        super();
    }

    public void update(final ClusterConnectionSourcePOJO clusterConnectionSource) {
        super.update(clusterConnectionSource);
        setUrl(clusterConnectionSource.getMonitoringHost());
        setName(prefix + database.replace("*", "").replace(",","").replace(".","") + "-" + clusterConnectionSource.getSourceName());
        setUid(prefix + database.replace("*", "").replace(",","").replace(".","") + "-" + clusterConnectionSource.getSourceUid());
        setBasicAuth(clusterConnectionSource.isMonitoringIsUseAuthentication());
        setBasicAuthUser(clusterConnectionSource.getMonitoringUsername());
        updatePassword(clusterConnectionSource.getMonitoringPassword());
    }

    @Override
    public ElasticsearchMonitoringDataSourceObject clone() {
        ElasticsearchMonitoringDataSourceObject result = (ElasticsearchMonitoringDataSourceObject) super.clone();
        result.setJsonData(new HashMap<>(getJsonData()));
        result.setSecureJsonData(new HashMap<>(getSecureJsonData()));
        return result;
    }

    public String getDatabase() {
        return this.database;
    }

    public void setDatabase(String database) {
        this.database = database;
    }

    @Override
    public String getDataSourcePrefix() {
        return prefix + database.replace("*", "") + "-";
    }

    @Override
    public void updateDataSourcePrefix() {
        this.prefix = new String(this.getName());
    }
}

