package co.dbeast.grafana_backend.constants;

import co.dbeast.grafana_backend.GrafanaBackend;

import java.util.Arrays;

public enum EDataSourceTypes {
    APPLICATION_DS("json_api_application_datasource", "application_api"),
    ELASTICSEARCH_DIRECT_DS("json_api_datasource_elasticsearch", "elasticsearch_direct_api"),
    KIBANA_DIRECT_DS("json_api_datasource_kibana", "kibana_direct_api"),
    ELASTICSEARCH_MONITORING_DS("elasticsearch_datasource", "elasticsearch_api"),
    TESTDATA_DS("testdata_datasource", "testdata_api"),
    ERROR("error", "error");

    private final String fileName;
    private final String sourceType;

    EDataSourceTypes(final String fileName, String sourceType) {
        this.fileName = fileName;
        this.sourceType = sourceType;
    }

    public String getFileName() {
        return fileName;
    }
    public String getSourceType() {
        return sourceType;
    }
    public static EDataSourceTypes getByFileName(String fileName) {
        return Arrays.stream(EDataSourceTypes.values()).filter(enumValue -> fileName.contains(enumValue.getFileName())).findFirst().orElse(ERROR);
    }

    public static EDataSourceTypes getBySourceType(String sourceType) {
        return Arrays.stream(EDataSourceTypes.values()).filter(enumValue -> enumValue.getFileName().equals(sourceType)).findFirst().orElse(ERROR);
    }
}
