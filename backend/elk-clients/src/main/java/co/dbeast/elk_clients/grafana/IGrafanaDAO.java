package co.dbeast.elk_clients.grafana;

import co.dbeast.elk_clients.exceptions.ClusterConnectionException;
import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSettingsPOJO;
import com.fasterxml.jackson.databind.JsonNode;
import org.elasticsearch.client.RestClient;

import java.util.List;

public interface IGrafanaDAO {
    List<JsonNode> getAll(final ClusterConnectionSettingsPOJO connectionSettings) throws ClusterConnectionException;
    List<JsonNode> getAll(final RestClient client);
    List<JsonNode> getById(final ClusterConnectionSettingsPOJO connectionSettings,
                           final String dataSourceId) throws ClusterConnectionException;

    public List<JsonNode> getById(final RestClient client,
                                  final String dataSourceId);

    JsonNode upsert(ClusterConnectionSettingsPOJO connectionSettings,
                    String dataSourceId,
                    String dataSource) throws ClusterConnectionException;

    JsonNode upsert(RestClient client,
                    String dataSourceId,
                    String dataSource);

    public JsonNode create(final ClusterConnectionSettingsPOJO connectionSettings,
                           final String dataSource) throws ClusterConnectionException;

    public JsonNode create(final RestClient client,
                           final String dataSource);
}
