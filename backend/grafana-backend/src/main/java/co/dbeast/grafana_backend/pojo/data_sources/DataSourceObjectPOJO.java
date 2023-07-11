package co.dbeast.grafana_backend.pojo.data_sources;

import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSourcePOJO;

import java.util.HashMap;
import java.util.Map;

public abstract class DataSourceObjectPOJO implements Cloneable {
    private int orgId = 1;
    private String name = "";
    private String uid = "";
    private String type = "marcusolsson-json-datasource";
    private String typeName = "JSON";
    private String API = "";
    private String access = "proxy";
    private String url = "";
    private boolean basicAuth = false;
    private boolean withCredentials = false;
    private String basicAuthUser = "";
//    private String version = "";
    private boolean isDefault = false;
    private boolean readOnly = false;

    private Map<String, Object> jsonData = new HashMap<String, Object>() {{
        put("tlsSkipVerify", Boolean.FALSE);
    }};

    private Map<String, Object> secureJsonData = new HashMap<>();

    public DataSourceObjectPOJO() {
    }
    @Override
    public DataSourceObjectPOJO clone() {
        try {
            DataSourceObjectPOJO result = (DataSourceObjectPOJO) super.clone();
            result.setJsonData(new HashMap<>(getJsonData()));
            result.setSecureJsonData(new HashMap<>(getSecureJsonData()));
            return (DataSourceObjectPOJO) super.clone();
        } catch (CloneNotSupportedException e) {
            e.printStackTrace();
            return this;
        }
    }
    public void update(final ClusterConnectionSourcePOJO clusterConnectionSource) {
        setBasicAuth(clusterConnectionSource.getClusterConnectionSettings().isAuthentication_enabled());
        if (clusterConnectionSource.getClusterConnectionSettings().isAuthentication_enabled()) {
            setBasicAuthUser(clusterConnectionSource.getClusterConnectionSettings().getUsername());
            updatePassword(clusterConnectionSource.getClusterConnectionSettings().getPassword());
        }
    }
    public void update() {
    }

//    public String getVersion() {
//        return version;
//    }
//
//    public void setVersion(String version) {
//        this.version = version;
//    }

    public boolean isDefault() {
        return isDefault;
    }

    public void setDefault(boolean aDefault) {
        isDefault = aDefault;
    }

    public int getOrgId() {
        return orgId;
    }

    public void setOrgId(int orgId) {
        this.orgId = orgId;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getUid() {
        return uid;
    }

    public void setUid(String uid) {
        this.uid = uid;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public String getTypeName() {
        return typeName;
    }

    public void setTypeName(String typeName) {
        this.typeName = typeName;
    }

    public String getAPI() {
        return API;
    }

    public void setAPI(String API) {
        this.API = API;
    }

    public String getAccess() {
        return access;
    }

    public void setAccess(String access) {
        this.access = access;
    }

    public String getUrl() {
        return url;
    }

    public void setUrl(String url) {
        this.url = url;
    }

    public boolean isBasicAuth() {
        return basicAuth;
    }

    public void setBasicAuth(boolean basicAuth) {
        this.basicAuth = basicAuth;
    }

    public String getBasicAuthUser() {
        return basicAuthUser;
    }

    public void setBasicAuthUser(String basicAuthUser) {
        this.basicAuthUser = basicAuthUser;
    }

    public Map<String, Object> getJsonData() {
        return jsonData;
    }

    public void setJsonData(Map<String, Object> jsonData) {
        this.jsonData = jsonData;
    }

    public boolean isReadOnly() {
        return readOnly;
    }

    public void setReadOnly(boolean readOnly) {
        this.readOnly = readOnly;
    }

    public Map<String, Object> getSecureJsonData() {
        return secureJsonData;
    }

    public void setSecureJsonData(Map<String, Object> secureJsonData) {
        this.secureJsonData = secureJsonData;
    }

    public void updatePassword(final String password) {
        secureJsonData.put("basicAuthPassword", password);
    }

    public boolean isWithCredentials() {
        return withCredentials;
    }

    public void setWithCredentials(boolean withCredentials) {
        this.withCredentials = withCredentials;
    }

    public boolean getIsDefault() {
        return isDefault;
    }

    public void setIsDefault(boolean isDefault) {
        this.isDefault = isDefault;
    }

    public abstract String getDataSourcePrefix();
    public abstract void updateDataSourcePrefix();
}
