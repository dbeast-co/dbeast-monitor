package co.dbeast.grafana_backend.constants;

import co.dbeast.grafana_backend.GrafanaBackend;

import java.util.Arrays;

public enum EGrafanaBackendSettings {
    GRAFANA_CLIENT_CONFIG(GrafanaBackend.HOME_FOLDER + "/config/grafana.json"),
    API_SERVER_CONFIG(GrafanaBackend.HOME_FOLDER + "/config/server.yml"),
    CLUSTERS_CONNECTION_SOURCE_FOLDER(GrafanaBackend.HOME_FOLDER + "/modules/connections/"),
    SETUP_FOLDER(GrafanaBackend.HOME_FOLDER + "/modules/grafana/setup/"),
    GRAFANA_OBJECTS_FOLDER(GrafanaBackend.HOME_FOLDER + "/modules/grafana/grafana_objects"),
    GRAFANA_POLICIES_FOLDER(GrafanaBackend.HOME_FOLDER + "/modules/grafana/policies/"),
    GRAFANA_DATASOURCE_FOLDER(GrafanaBackend.HOME_FOLDER + "/modules/grafana/grafana_objects/datasource/"),
    GRAFANA_CLUSTERS_FOLDER(GrafanaBackend.HOME_FOLDER + "/modules/grafana/clusters/"),
    GRAFANA_PANEL_FOLDER(GrafanaBackend.HOME_FOLDER + "/modules/grafana/grafana_objects/panel/"),
    SETUP_APPLICATION_DATASOURCE(EGrafanaBackendSettings.SETUP_FOLDER.getSetting() + "/datasources/json_api_application_datasource.json"),
    SETUP_GRAFANA_API_DATASOURCE(EGrafanaBackendSettings.SETUP_FOLDER.getSetting() + "/datasources/json_api_grafana_api_datasource.json"),
    APPLICATION_DATA_SOURCE_NAME("DBeast-toolkit"),
    APPLICATION_DATA_SOURCE_ID("DBeast-toolkit"),
    APPLICATION_GRAFANA_API_DATA_SOURCE_ID("Grafana-DS-API"),
    APPLICATION_KEYSTORE_PSW("qwe123"),
    GRAFANA_USERNAME_FIELD_FOR_KEYSTORE("grafana-user"),
    KEYSTORE_PATH(GrafanaBackend.HOME_FOLDER + "/config/dbeast-toolkit.jks"),
    CLUSTER_ID_DELIMITER("--"),
    ERROR("The error in the Settings object");

    private final String setting;

    EGrafanaBackendSettings(final String setting) {
        this.setting = setting;
    }

    public String getSetting() {
        return setting;
    }

    public static EGrafanaBackendSettings getByValue(String value) {
        return Arrays.stream(EGrafanaBackendSettings.values()).filter(enumValue -> enumValue.getSetting().equals(value)).findFirst().orElse(ERROR);
    }
}
