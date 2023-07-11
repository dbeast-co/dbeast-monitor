package co.dbeast.elk_clients.grafana;

import co.dbeast.elk_clients.elasticsearch.ElasticsearchDbProvider;
import co.dbeast.elk_clients.exceptions.ClusterConnectionException;
import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSettingsPOJO;
import org.apache.http.entity.ContentType;
import org.apache.http.nio.entity.NStringEntity;
import org.apache.http.util.EntityUtils;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.elasticsearch.client.Request;
import org.elasticsearch.client.Response;
import org.elasticsearch.client.RestClient;

import java.io.IOException;

public class GrafanaDbProvider {
    private static final Logger logger = LogManager.getLogger();
    ElasticsearchDbProvider elasticsearchDbProvider = new ElasticsearchDbProvider();

    public RestClient getLowLevelClient(final ClusterConnectionSettingsPOJO connectionSettings, String projectId) throws ClusterConnectionException {
        return elasticsearchDbProvider.getHighLevelClient(connectionSettings, projectId).getLowLevelClient();
    }

    public String processLowLevelPUTRequest(final RestClient client,
                                            final String api,
                                            final String body) {
        Request request = new Request("PUT", api);
        NStringEntity jsonBody = new NStringEntity(body,
                ContentType.APPLICATION_JSON);
        request.setEntity(jsonBody);
        return processLowLevelRequest(client, request);
    }

    public String processLowLevelPOSTRequest(final RestClient client,
                                             final String api,
                                             final String body) {
        Request request = new Request("POST", api);
        NStringEntity jsonBody = new NStringEntity(body,
                ContentType.APPLICATION_JSON);
        request.setEntity(jsonBody);
        return processLowLevelRequest(client, request);
    }

    public String processLowLevelGETRequest(final RestClient client,
                                            final String api) {
        Request request = new Request("GET", api);
        return processLowLevelRequest(client, request);
    }

    public String processLowLevelDELETERequest(final RestClient client,
                                            final String api) {
        Request request = new Request("DELETE", api);
        return processLowLevelRequest(client, request);
    }

    private String processLowLevelRequest(final RestClient client,
                                          final Request request) {
        logger.info("Request for API: " + request.getEndpoint());
        Response response;
        try {
            response = client.performRequest(request);
            if (response.getStatusLine().getStatusCode() == 200) {
                return EntityUtils.toString(response.getEntity());
            } else {
                logger.error("There is the error in the get index parameters of index: " + request.getEndpoint());
                return null;
            }
        } catch (IOException e) {
            logger.error(e);
            return String.valueOf(e);
        }
    }

}
