package co.dbeast.elk_clients.constants;

import co.dbeast.elk_clients.ElkClients;

import java.util.Arrays;

public enum EElkClients {
    API_SERVER_CONFIG(ElkClients.HOME_FOLDER + "/config/server.yml"),
    CLUSTERS_CONNECTION_SOURCE_FOLDER(ElkClients.HOME_FOLDER + "/modules/connections/"),
    ERROR("The error in the Settings object");

    private final String setting;

    EElkClients(final String setting) {
        this.setting = setting;
    }

    public String getSetting() {
        return setting;
    }

}
