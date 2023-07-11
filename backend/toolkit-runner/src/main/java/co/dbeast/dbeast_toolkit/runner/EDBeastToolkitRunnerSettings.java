package co.dbeast.dbeast_toolkit.runner;

import java.util.Arrays;

public enum EDBeastToolkitRunnerSettings {
    API_SERVER_CONFIG(DbeastToolkit.HOME_FOLDER + "/config/server.yml"),
    KEYSTORE_PATH(DbeastToolkit.HOME_FOLDER + "/config/dbeast-toolkit.jks"),
    APPLICATION_KEYSTORE_PSW("qwe123"),
    ERROR("The error in the Settings object");

    private final String setting;

    EDBeastToolkitRunnerSettings(final String setting) {
        this.setting = setting;
    }

    public String getSetting() {
        return setting;
    }

    public static EDBeastToolkitRunnerSettings getByValue(String value) {
        return Arrays.stream(EDBeastToolkitRunnerSettings.values()).filter(enumValue -> enumValue.getSetting().equals(value)).findFirst().orElse(ERROR);
    }
}
