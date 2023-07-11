package co.dbeast.grafana_backend.pojo.data_sources;

import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSourcePOJO;
import co.dbeast.grafana_backend.constants.EGrafanaBackendSettings;

import java.util.HashMap;

public class ApplicationDirectDataSourceObject extends DataSourceObjectPOJO implements Cloneable{

    public void update(final ClusterConnectionSourcePOJO clusterConnectionSource) {
        setUrl(clusterConnectionSource.getEs_host());
        setName(EGrafanaBackendSettings.APPLICATION_DATA_SOURCE_NAME.getSetting());
        setUid(EGrafanaBackendSettings.APPLICATION_DATA_SOURCE_ID.getSetting());
        setJsonData(new HashMap<>());
        setSecureJsonData(new HashMap<>());
    }

    @Override
    public String getDataSourcePrefix() {
        return "Dbeast-toolkit";
    }

    @Override
    public ApplicationDirectDataSourceObject clone() {
        ApplicationDirectDataSourceObject result = (ApplicationDirectDataSourceObject) super.clone();
        result.setJsonData(new HashMap<>(getJsonData()));
        result.setSecureJsonData(new HashMap<>(getSecureJsonData()));
        return result;
    }
    @Override
    public void updateDataSourcePrefix() {
    }
}
