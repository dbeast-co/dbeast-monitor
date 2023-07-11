package co.dbeast.elk_clients.grafana;

import co.dbeast.elk_clients.exceptions.ClusterConnectionException;
import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSettingsPOJO;
import co.dbeast.general_utils.JSONUtils;
import com.fasterxml.jackson.databind.JsonNode;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.elasticsearch.client.RestClient;

import java.util.ArrayList;
import java.util.List;

public class GrafanaLibraryElementsDAO implements IGrafanaDAO {
    private static final Logger logger = LogManager.getLogger();
    private final String ROOT_API = "/api/library-elements";
    private final GrafanaDbProvider grafanaDbProvider = new GrafanaDbProvider();

    @Override
    public List<JsonNode> getAll(final ClusterConnectionSettingsPOJO connectionSettings) throws ClusterConnectionException {
        RestClient client = grafanaDbProvider.getLowLevelClient(connectionSettings, connectionSettings.getEs_host());
        return getAll(client);
    }

    @Override
    public List<JsonNode> getAll(final RestClient client) {
        String API = ROOT_API + "?kind=1&perPage=1000";
        String response = grafanaDbProvider.processLowLevelGETRequest(client, API);
        JsonNode parsedResponse = JSONUtils.stringToJSON(response);
        return JSONUtils.convertType(parsedResponse.get("result").get("elements"), new ArrayList<>());
    }

    @Override
    public List<JsonNode> getById(ClusterConnectionSettingsPOJO connectionSettings, String dataSourceId) throws ClusterConnectionException {
        return null;
    }

    @Override
    public List<JsonNode> getById(RestClient client, String dataSourceId) {
        return null;
    }

    @Override
    public JsonNode upsert(ClusterConnectionSettingsPOJO connectionSettings, String dataSourceId, String dataSource) throws ClusterConnectionException {
        return null;
    }

    @Override
    public JsonNode upsert(RestClient client, String dataSourceId, String dataSource) {
        return null;
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

//    public JsonNode updateDataSourceById(final ClusterConnectionSettingsPOJO connectionSettings,
//                                         final String dataSourceId,
//                                         final String dataSource) throws ClusterConnectionException {
//        RestClient client = grafanaDbProvider.getLowLevelClient(connectionSettings, connectionSettings.getEs_host());
//        return updateDataSourceById(client, dataSourceId, dataSource);
//    }
//
//    public JsonNode updateDataSourceById(final RestClient client,
//                                         final String dataSourceId,
//                                         final String dataSource) {
//        String API = ROOT_API + "/id/" + dataSourceId;
//        String response = grafanaDbProvider.processLowLevelPUTRequest(
//                client,
//                API,
//                dataSource);
//        return JSONUtils.stringToJSON(response);
//    }

    public JsonNode deletePanel(final ClusterConnectionSettingsPOJO connectionSettings,
                                final String dataSourceId) throws ClusterConnectionException {
        RestClient client = grafanaDbProvider.getLowLevelClient(connectionSettings, connectionSettings.getEs_host());
        return deletePanel(client, dataSourceId);
    }

    public JsonNode deletePanel(final RestClient client,
                                final String dataSourceUId) {
        String API = ROOT_API + "/" + dataSourceUId;
        String response = grafanaDbProvider.processLowLevelDELETERequest(
                client,
                API);
        return JSONUtils.stringToJSON(response);
    }


}
