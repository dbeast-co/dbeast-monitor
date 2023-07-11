package co.dbeast.elk_clients.elasticsearch;

import co.dbeast.elk_clients.elasticsearch.pojo.EsShardStatsPOJO;
import co.dbeast.elk_clients.exceptions.ClusterConnectionException;
import co.dbeast.elk_clients.elasticsearch.es_settings.ClusterConnectionSettingsPOJO;
import co.dbeast.elk_clients.elasticsearch.pojo.EsNodeStatsPOJO;
import co.dbeast.general_utils.JSONUtils;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ArrayNode;
import com.fasterxml.jackson.databind.node.ObjectNode;
import org.apache.http.entity.ContentType;
import org.apache.http.nio.entity.NStringEntity;
import org.apache.http.util.EntityUtils;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.elasticsearch.ElasticsearchStatusException;
import org.elasticsearch.action.admin.cluster.health.ClusterHealthRequest;
import org.elasticsearch.action.admin.cluster.health.ClusterHealthResponse;
import org.elasticsearch.action.search.ClearScrollRequest;
import org.elasticsearch.action.search.ClearScrollResponse;
import org.elasticsearch.client.*;
import org.elasticsearch.client.core.CountRequest;

import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.stream.Collectors;

public class ElasticsearchController {
    private static final Logger logger = LogManager.getLogger();
    private final ObjectMapper mapper = new ObjectMapper();

    private final ElasticsearchDbProvider elasticsearchClient = new ElasticsearchDbProvider();

    public RestClient getLowLevelClient(final ClusterConnectionSettingsPOJO connectionSettings) throws ClusterConnectionException {
        return elasticsearchClient.getLowLevelClient(connectionSettings, connectionSettings.getEs_host());
    }

    public JsonNode processPUTRequest(RestClient client) throws JsonProcessingException {
        String response = processLowLevelGETRequest(client, ".kibana/_search?q=type:index-pattern");
        JsonNode parsedResponse = mapper.readTree(response);
        return parsedResponse.get("hits").get("hits");
    }

    public JsonNode getIndexPatternsFromKibanaIndex(RestClient client) throws JsonProcessingException {
        String response = processLowLevelGETRequest(client, ".kibana/_search?q=type:index-pattern");
        JsonNode parsedResponse = mapper.readTree(response);
        return parsedResponse.get("hits").get("hits");
    }

    public JsonNode getPatternSearchDashVisFromKibanaIndex(RestClient client) throws JsonProcessingException {
        String response = processLowLevelGETRequest(client, ".kibana/_search?q=type:index-pattern%20or%20type:visualization%20o%20type:dashboard" +
                "%20or%20type:lens%20or%20type:search&_source_excludes=index-pattern.fieldFormatMap,index-pattern.fields&size=10000");
//        String response = processLowLevelGETRequest(client, ".kibana/_search?q=type:index-pattern");
        JsonNode parsedResponse = mapper.readTree(response);
        return parsedResponse.get("hits").get("hits");
    }

    public ArrayNode getCatShards(final ClusterConnectionSettingsPOJO connectionSettings, final String projectId) throws ClusterConnectionException {
        RestClient client = elasticsearchClient.getLowLevelClient(connectionSettings, projectId);
        Response response;
        try {
            response = client.performRequest(new Request("GET", "_cat/shards?h=index,shard,prirep,docs,store,node,state,prirep&bytes=m&format=json"));
            if (response.getStatusLine().getStatusCode() == 200) {
                // parse the JSON response
                return (ArrayNode) stringToJSON(EntityUtils.toString(response.getEntity()));
            } else {
                logger.error("There is an error in the cat shards request");
            }
        } catch (IOException e) {
            logger.error(e);
        }
        return null;
    }

    public ArrayNode getCatIndices(final ClusterConnectionSettingsPOJO connectionSettings, final String projectId) throws ClusterConnectionException {
        RestClient client = elasticsearchClient.getLowLevelClient(connectionSettings, projectId);
        Response response;
        try {
            response = client.performRequest(new Request("GET",
                    "_cat/indices?h=index,creation.date,docs.count,pri,pri.store.size&bytes=m&format=json"));
            if (response.getStatusLine().getStatusCode() == 200) {
                // parse the JSON response
                return (ArrayNode) stringToJSON(EntityUtils.toString(response.getEntity()));
            } else {
                logger.error("There is an error in the cat shards request");
            }
        } catch (IOException e) {
            logger.error(e);
        }
        return null;
    }

    public ArrayNode getCatAliases(final ClusterConnectionSettingsPOJO connectionSettings, final String projectId) throws ClusterConnectionException {
        RestClient client = elasticsearchClient.getLowLevelClient(connectionSettings, projectId);
        Response response;
        try {
            response = client.performRequest(new Request("GET",
                    "_cat/aliases?v&h=alias,index,is_write_index&format=json"));
            if (response.getStatusLine().getStatusCode() == 200) {
                // parse the JSON response
                return (ArrayNode) stringToJSON(EntityUtils.toString(response.getEntity()));
            } else {
                logger.error("There is an error in the cat aliases request");
            }
        } catch (IOException e) {
            logger.error(e);
        }
        return null;
    }

    public boolean closeScrollRequest(final String scrollId, RestHighLevelClient client) throws IOException {
        ClearScrollRequest clearScrollRequest = new ClearScrollRequest();
        clearScrollRequest.addScrollId(scrollId);
        ClearScrollResponse clearScrollResponse = client.clearScroll(clearScrollRequest, RequestOptions.DEFAULT);

        return clearScrollResponse.isSucceeded();
    }

    public long getIndexDocsCount(final ClusterConnectionSettingsPOJO connectionSettings, final String projectId, final String index) throws ClusterConnectionException {
        RestHighLevelClient client = elasticsearchClient.getHighLevelClient(connectionSettings, projectId);
        CountRequest countRequest = new CountRequest(index);
        try {
            return client
                    .count(countRequest, RequestOptions.DEFAULT).getCount();
        } catch (IOException | ElasticsearchStatusException e) {
            logger.error(e);
            return 0;
        }
    }

    public JsonNode getAllIndicesSettings(final RestHighLevelClient highLevelClient) throws JsonProcessingException {
        String response = processLowLevelGETRequest(highLevelClient, "*/_settings");
//        TypeReference<HashMap<String, HashMap<String, HashMap<String, HashMap<String, Object>>> typeRef =
//                new TypeReference<HashMap<String, HashMap<String, Object>>>() {
//                };
        return mapper.readTree(response);
    }

    public JsonNode getAllIndicesMappings(final RestHighLevelClient highLevelClient) throws JsonProcessingException {
        String response = processLowLevelGETRequest(highLevelClient, "*/_mapping");
        return mapper.readTree(response);
    }

    public List<HashMap<String, Object>> getLegacyTemplates(final RestHighLevelClient highLevelClient) throws JsonProcessingException {
        String response = processLowLevelGETRequest(highLevelClient, "_template/");
        TypeReference<HashMap<String, HashMap<String, Object>>> typeRef =
                new TypeReference<HashMap<String, HashMap<String, Object>>>() {
                };
        return new ArrayList<>(mapper.readValue(response, typeRef).values());
    }

    public List<HashMap<String, Object>> getLegacyTemplates(final RestClient lowLevelRestClient) throws JsonProcessingException {
        String response = processLowLevelGETRequest(lowLevelRestClient, "_template/");
        TypeReference<List<HashMap<String, Object>>> typeRef = new TypeReference<List<HashMap<String, Object>>>() {
        };
        return mapper.readValue(response, typeRef);
    }

    public List<HashMap<String, Object>> getIndexTemplates(final RestHighLevelClient highLevelClient) throws JsonProcessingException {
        String response = processLowLevelGETRequest(highLevelClient, "_index_template/");
        TypeReference<HashMap<String, List<HashMap<String, Object>>>> typeRef =
                new TypeReference<HashMap<String, List<HashMap<String, Object>>>>() {
                };
        List<HashMap<String, Object>> templates = mapper.readValue(response, typeRef).get("index_templates");
        return templates.stream()
                .map(tpl -> (HashMap<String, Object>) tpl.get("index_template"))
                .collect(Collectors.toList());
    }

    public List<HashMap<String, Object>> getIndexTemplates(final RestClient lowLevelRestClient) throws JsonProcessingException {
        String response = processLowLevelGETRequest(lowLevelRestClient, "_index_template/");
        TypeReference<List<HashMap<String, Object>>> typeRef = new TypeReference<List<HashMap<String, Object>>>() {
        };
        return mapper.readValue(response, typeRef);
    }

    public String getLegacyTemplate(final ClusterConnectionSettingsPOJO connectionSettings,
                                    final String template,
                                    final String projectId) throws ClusterConnectionException {
        RestClient client = elasticsearchClient.getLowLevelClient(connectionSettings, projectId);
        Response response;
        try {
            response = client.performRequest(new Request("GET", "_template/" + template));
            if (response.getStatusLine().getStatusCode() == 200) {
                // parse the JSON response
                return EntityUtils.toString(response.getEntity());
            } else {
                logger.error("There is the error in the get index parameters of index: " + template);
            }
        } catch (IOException e) {
            logger.error(e);
            return String.valueOf(e);
        }
        return null;
    }

    public String getIndexTemplate(final ClusterConnectionSettingsPOJO connectionSettings,
                                   final String template,
                                   final String projectId) throws ClusterConnectionException {
        RestClient client = elasticsearchClient.getLowLevelClient(connectionSettings, projectId);
        Response response;
        try {
            response = client.performRequest(new Request("GET", "_index_template/" + template));
            if (response.getStatusLine().getStatusCode() == 200) {
                // parse the JSON response
                return EntityUtils.toString(response.getEntity());
            } else {
                logger.error("There is the error in the get index parameters of index: " + template);
            }
        } catch (IOException e) {
            logger.error(e);
            return String.valueOf(e);
        }
        return null;
    }

    public String getIndexParameters(final ClusterConnectionSettingsPOJO connectionSettings,
                                     final String index,
                                     final String projectId) throws ClusterConnectionException {
        RestClient client = elasticsearchClient.getLowLevelClient(connectionSettings, projectId);
        Response response;
        try {
            response = client.performRequest(new Request("GET", index));
            if (response.getStatusLine().getStatusCode() == 200) {
                // parse the JSON response
                return EntityUtils.toString(response.getEntity());
            } else {
                logger.error("There is the error in the get index parameters of index: " + index);
            }
        } catch (IOException e) {
            logger.error(e);
            return String.valueOf(e);
        }
        return null;
    }

    public List<String> getClusterAllocation(RestClient client) {
        Response response;
        try {
            response = client.performRequest(new Request("GET", "_cat/allocation?format=json&bytes=gb"));
            if (response.getStatusLine().getStatusCode() == 200) {
                String rawBody = EntityUtils.toString(response.getEntity());
                TypeReference<List<String>> typeRef = new TypeReference<List<String>>() {
                };
                return mapper.readValue(rawBody, typeRef);

            } else {
                logger.error("There is the error in the get cat allocation." +
                        response.getStatusLine());
                return new LinkedList<>();
            }
        } catch (IOException e) {
            logger.error(e);
            return new LinkedList<>();
        }
    }

    public List<EsNodeStatsPOJO> getShardsAllocation(RestClient client) {
        Response response;
        try {
            response = client.performRequest(new Request("GET", "_nodes/data:true/stats?level=shards&" +
                    "filter_path=cluster_name,nodes.*.roles,nodes.*.name,nodes.*.host,nodes.*.indices.store.size_in_bytes," +
                    "nodes.*.fs.total.total_in_bytes,nodes.*.fs.total.available_in_bytes,nodes.*.fs.total.free_in_bytes," +
                    "nodes.*.indices.shards.*.*.store.total_data_set_size_in_bytes,nodes.*.indices.shards.*.*.routing.primary"));
            if (response.getStatusLine().getStatusCode() == 200) {
                List<EsNodeStatsPOJO> nodesAllocationList = new LinkedList<>();
                String rawBody = EntityUtils.toString(response.getEntity());
                ObjectNode jsonResult = (ObjectNode) JSONUtils.stringToJSON(rawBody);
                String clusterName = jsonResult.get("cluster_name").asText();
                ObjectNode nodes = (ObjectNode) jsonResult.get("nodes");
                nodes.fields().forEachRemaining(node -> {
                    EsNodeStatsPOJO nodeAllocation = new EsNodeStatsPOJO(node.getValue());
                    nodeAllocation.setNodeId(node.getKey());
                    node.getValue().get("indices").get("shards").fields()
                            .forEachRemaining(index ->
                                    index.getValue().forEach(indexValues ->
                                            indexValues.fields().forEachRemaining(shard -> {
                                                EsShardStatsPOJO newShard = new EsShardStatsPOJO(shard.getValue());
                                                newShard.setIndexName(index.getKey());
                                                newShard.setShardNumber(Integer.parseInt(shard.getKey()));
                                                newShard.setNodeId(node.getKey());
                                                newShard.setNodeName(node.getValue().get("name").asText());
                                                nodeAllocation.getShards().add(newShard);
                                            })
                                    )
                );
                nodesAllocationList.add(nodeAllocation);
            });
            return nodesAllocationList;
        } else{
            logger.error("There is the error in the get cat allocation." +
                    response.getStatusLine());
            return new LinkedList<>();
        }
    } catch(
    IOException e)

    {
        logger.error(e);
        return new LinkedList<>();
    }

}

    public List<HashMap<String, String>> getIndexList(final ClusterConnectionSettingsPOJO connectionSettings,
                                                      final String projectId) {
        return getTemplateOrIndexList(connectionSettings, "/_cat/indices?h=index&format=json", projectId);
    }

    public List<HashMap<String, String>> getTemplateList(final ClusterConnectionSettingsPOJO connectionSettings,
                                                         final String projectId) {
        return getTemplateOrIndexList(connectionSettings, "/_cat/templates?h=name&format=json", projectId);
    }


    //TODO add catch exception
    private List<HashMap<String, String>> getTemplateOrIndexList(final ClusterConnectionSettingsPOJO connectionSettings,
                                                                 final String endPoint,
                                                                 final String projectId) {
        Response response;
        try {
            RestClient client = elasticsearchClient.getLowLevelClient(connectionSettings, projectId);
            response = client.performRequest(new Request("GET", endPoint));
            if (response.getStatusLine().getStatusCode() == 200) {
                // parse the JSON response
                String rawBody = EntityUtils.toString(response.getEntity());
                TypeReference<List<HashMap<String, String>>> typeRef = new TypeReference<List<HashMap<String, String>>>() {
                };
                return mapper.readValue(rawBody, typeRef);
            }
        } catch (IOException | ClusterConnectionException e) {
            logger.error(e.getMessage());
        }
        return new LinkedList<>();
    }

    public JsonNode getClusterInfo(ClusterConnectionSettingsPOJO clusterConnectionSettings, String projectId) {
        try {
            Response response;
            try {
                RestClient client = elasticsearchClient.getLowLevelClient(clusterConnectionSettings, projectId);
                response = client.performRequest(new Request("GET", "/"));
                if (response.getStatusLine().getStatusCode() == 200) {
                    // parse the JSON response
                    String rawBody = EntityUtils.toString(response.getEntity());
                    TypeReference<List<HashMap<String, String>>> typeRef = new TypeReference<List<HashMap<String, String>>>() {
                    };
                    return mapper.readTree(rawBody);
                }
            } catch (IOException | ClusterConnectionException e) {
                logger.error(e.getMessage());
                return null;
            }
        } catch (Exception e) {
            return null;
        }
        return null;
    }

    public JsonNode getClusterStats(RestClient client) {
        try {
            Response response;
            try {
                response = client.performRequest(new Request("GET", "/_cluster/stats"));
                if (response.getStatusLine().getStatusCode() == 200) {
                    // parse the JSON response
                    String rawBody = EntityUtils.toString(response.getEntity());
//                    TypeReference<List<HashMap<String, String>>> typeRef = new TypeReference<List<HashMap<String, String>>>() {
//                    };
                    return mapper.readTree(rawBody);
                }
            } catch (IOException e) {
                client.close();
                logger.error(e.getMessage());
                return null;
            }
        } catch (Exception e) {
            return null;
        }
        return null;
    }

    public JsonNode getClusterStats(ClusterConnectionSettingsPOJO clusterConnectionSettings) {
        try {
            Response response;
            RestClient client = null;
            try {
                client = elasticsearchClient.getLowLevelClient(clusterConnectionSettings, clusterConnectionSettings.getEs_host());
                return getClusterStats(client);
            } catch (ClusterConnectionException e) {
                logger.error(e.getMessage());
                throw e;
            }
        } catch (Exception e) {
            return null;
        }
    }

    public String getClusterStatus(final ClusterConnectionSettingsPOJO connectionSettings,
                                   final String projectId) throws Exception {
        RestHighLevelClient client = elasticsearchClient.getHighLevelClient(connectionSettings, projectId);
        try {
            ClusterHealthRequest request = new ClusterHealthRequest();
            request.timeout("30s");
            ClusterHealthResponse response = client.cluster().health(request, RequestOptions.DEFAULT);
            return response.getStatus().toString();
        } catch (Exception e) {
            throw e;
        } finally {
            try {
                if (client != null) {
                    client.close();
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }

    public int getHeath(final RestClient client,
                        final String api) {
        Response response;
        logger.info("Request API: " + api);
        try {
            response = client.performRequest(new Request("GET", api));
            return response.getStatusLine().getStatusCode();
        } catch (IOException e) {
            logger.error(e);
            return 400;
        }
    }
    public String processLowLevelGETRequest(final RestClient client,
                                             final String api) {
        Response response;
        logger.info("Request API: " + api);
        try {
            response = client.performRequest(new Request("GET", api));
            if (response.getStatusLine().getStatusCode() == 200) {
                // parse the JSON response
                return EntityUtils.toString(response.getEntity());
            } else {
                logger.error("There is the error in the get index parameters of index: " + api);
            }
        } catch (IOException e) {
            logger.error(e);
            return String.valueOf(e);
        }
        return null;
    }
    public String processLowLevelPUTRequest(final RestClient client,
                                             final String api,
                                             final String body) {
        Response response;
        logger.info("Request API: " + api);
        try {
            Request request = new Request("PUT", api);
            request.setEntity(new NStringEntity(
                    body,
                    ContentType.APPLICATION_JSON));
            response = client.performRequest(request);
            if (response.getStatusLine().getStatusCode() == 200) {
                // parse the JSON response
                return EntityUtils.toString(response.getEntity());
            } else {
                logger.error("There is the error in the get PUT request for: " + api
                + " and body: " + body);
            }
        } catch (IOException e) {
            logger.error(e);
            return String.valueOf(e);
        }
        return null;
    }

    private String processLowLevelGETRequest(final RestHighLevelClient highLevelClient,
                                             final String api) {
        RestClient client = highLevelClient.getLowLevelClient();
        return processLowLevelGETRequest(client, api);
    }

    private JsonNode stringToJSON(final String string) {
        try {
            JsonNode jsonNode = mapper.readTree(string);
            return jsonNode;
        } catch (JsonProcessingException e) {
            e.printStackTrace();
        }
        return null;
    }


}
