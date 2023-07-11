package co.dbeast.elk_clients.elasticsearch.pojo;

import com.fasterxml.jackson.databind.JsonNode;

public class EsShardStatsPOJO {
    private String indexName;
    private String nodeId;
    private String nodeName;
    private boolean isPrimary;
    private int shardNumber;
    private double sizeInGb;

    public EsShardStatsPOJO(JsonNode node) {
        this.isPrimary = node.get("routing").get("primary").asBoolean();
        this.sizeInGb =  Math.round(node.get("store").get("total_data_set_size_in_bytes").asDouble() / 1024 / 1024 / 1024 *1000)/1000.0;
    }

    public boolean isIndexNameStartWithDot(){
        return indexName.startsWith(".");
    }
    public String getIndexName() {
        return indexName;
    }

    public void setIndexName(String indexName) {
        this.indexName = indexName;
    }

    public String getNodeId() {
        return nodeId;
    }

    public void setNodeId(String nodeId) {
        this.nodeId = nodeId;
    }

    public String getNodeName() {
        return nodeName;
    }

    public void setNodeName(String nodeName) {
        this.nodeName = nodeName;
    }

    public boolean isPrimary() {
        return isPrimary;
    }

    public void setPrimary(boolean primary) {
        isPrimary = primary;
    }

    public int getShardNumber() {
        return shardNumber;
    }

    public void setShardNumber(int shardNumber) {
        this.shardNumber = shardNumber;
    }

    public double getSizeInGb() {
        return sizeInGb;
    }

    public void setSizeInGb(double sizeInGb) {
        this.sizeInGb = sizeInGb;
    }
}
