package co.dbeast.elk_clients.grafana;

import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSettingsPOJO;
import co.dbeast.elk_clients.exceptions.ClusterConnectionException;
import co.dbeast.general_utils.JSONUtils;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.elasticsearch.client.RestClient;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

public class GrafanaDashboardsDAO implements IGrafanaDAO {
    private static final Logger logger = LogManager.getLogger();
    private final String ROOT_API = "/api/dashboards/db";
    private final GrafanaDbProvider grafanaDbProvider = new GrafanaDbProvider();

    @Override
    public List<JsonNode> getAll(final ClusterConnectionSettingsPOJO connectionSettings) throws ClusterConnectionException {
        RestClient client = grafanaDbProvider.getLowLevelClient(connectionSettings, connectionSettings.getEs_host());
        return getAll(client);
    }

    @Override
    public List<JsonNode> getAll(final RestClient client) {
        String response = grafanaDbProvider.processLowLevelGETRequest(client, ROOT_API);
        return JSONUtils.convertType(JSONUtils.stringToJSON(response), new ArrayList<>());
    }

    @Override
    public List<JsonNode> getById(final ClusterConnectionSettingsPOJO connectionSettings,
                                  final String dataSourceId) throws ClusterConnectionException {
        RestClient client = grafanaDbProvider.getLowLevelClient(connectionSettings, connectionSettings.getEs_host());
        return getById(client, dataSourceId);
    }

    @Override
    public List<JsonNode> getById(final RestClient client,
                                  final String dataSourceId) {
        String API = ROOT_API + "/" + dataSourceId;
        String response = grafanaDbProvider.processLowLevelGETRequest(client, API);
        JSONUtils.stringToJSON(response);
        return JSONUtils.convertType(JSONUtils.stringToJSON(response), new ArrayList<>());
    }

    @Override
    public JsonNode upsert(final ClusterConnectionSettingsPOJO connectionSettings,
                           final String dataSourceId,
                           final String dataSource) throws ClusterConnectionException {

        RestClient client = grafanaDbProvider.getLowLevelClient(connectionSettings, connectionSettings.getEs_host());
        return upsert(client, dataSourceId, dataSource);
    }

    @Override
    public JsonNode upsert(final RestClient client,
                           final String dataSourceName,
                           final String dataSource) {
        if (isExists(client, dataSourceName)) {
            return update(client, dataSourceName, dataSource);
        } else {
            return create(client, dataSource);
        }
    }

    public boolean isExists(final ClusterConnectionSettingsPOJO connectionSettings,
                            final String dataSourceId) throws ClusterConnectionException {
        RestClient client = grafanaDbProvider.getLowLevelClient(connectionSettings, connectionSettings.getEs_host());
        return isExists(client, dataSourceId);
    }
    public boolean isExists(final RestClient client,
                            final String dataSourceName) {
        ObjectMapper mapper = new ObjectMapper();
        List<JsonNode> dataSources = getAll(client);

        List<String> parsedList = new LinkedList<>();
        if (dataSources.size() > 0) {
            for (int i = 0; i < dataSources.size(); i++) {
                parsedList.add(mapper.convertValue(dataSources.get(i), new TypeReference<Map<String, Object>>() {
                }).get("name").toString());
            }
            return parsedList.stream().anyMatch(dataSourceName::equals);
        }
        return false;
    }

    @Override
    public JsonNode create(final ClusterConnectionSettingsPOJO connectionSettings,
                           final String dataSource) throws ClusterConnectionException {
        RestClient client = grafanaDbProvider.getLowLevelClient(connectionSettings, connectionSettings.getEs_host());
        return create(client, dataSource);
    }

    @Override
    public JsonNode create(final RestClient client,
                           final String dataSource) {
        String response = grafanaDbProvider.processLowLevelPOSTRequest(
                client,
                ROOT_API,
                dataSource);
        return JSONUtils.stringToJSON(response);
    }

    public JsonNode update(final ClusterConnectionSettingsPOJO connectionSettings,
                           final String dataSourceId,
                           final String dataSource) throws ClusterConnectionException {
        RestClient client = grafanaDbProvider.getLowLevelClient(connectionSettings, connectionSettings.getEs_host());
        return update(client, dataSourceId, dataSource);
    }

    public JsonNode update(final RestClient client,
                           final String dataSourceName,
                           final String dataSource) {
        String API = ROOT_API + "/uid/" + dataSourceName;
        String response = grafanaDbProvider.processLowLevelPUTRequest(
                client,
                API,
                dataSource);
        return JSONUtils.stringToJSON(response);
    }

    public JsonNode delete(final ClusterConnectionSettingsPOJO connectionSettings,
                           final String dataSourceId) throws ClusterConnectionException {
        RestClient client = grafanaDbProvider.getLowLevelClient(connectionSettings, connectionSettings.getEs_host());
        return delete(client, dataSourceId);
    }

    public JsonNode delete(final RestClient client,
                           final String dataSourceId) {
        String API = ROOT_API + "/uid/" + dataSourceId;
        String response = grafanaDbProvider.processLowLevelDELETERequest(
                client,
                API);
        return JSONUtils.stringToJSON(response);
    }
}
