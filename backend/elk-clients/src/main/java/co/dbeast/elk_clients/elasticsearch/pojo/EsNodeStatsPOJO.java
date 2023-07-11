package co.dbeast.elk_clients.elasticsearch.pojo;

import co.dbeast.general_utils.JSONUtils;
import com.fasterxml.jackson.databind.JsonNode;

import java.util.Comparator;
import java.util.LinkedList;
import java.util.List;
import java.util.stream.Collectors;

public class EsNodeStatsPOJO {
    private String nodeId;
    private String nodeName;
    private String host;
    private List<String> roles = new LinkedList<>();
    private List<EsShardStatsPOJO> shards = new LinkedList<>();
    private List<EsShardStatsPOJO> lowSizeShards = new LinkedList<>();
    private List<EsShardStatsPOJO> highSizeShards = new LinkedList<>();
    private long diskTotal;
    private long diskUsed;
    private long diskAvailable;
    private long diskFree;
    private int diskUsedPct;
    private long avgNodeSetAvailable;

    public EsNodeStatsPOJO(final String nodeId,
                           final String nodeName,
                           final String host,
                           final List<String> roles,
                           final long diskTotal,
                           final long diskUsed,
                           final long diskAvailable,
                           final long diskFree) {
        this.nodeId = nodeId;
        this.nodeName = nodeName;
        this.host = host;
        this.roles = roles;
        this.diskUsed = diskUsed;
        this.diskAvailable = diskAvailable;
        this.diskFree = diskFree;
        this.diskUsedPct = (int) (diskAvailable / diskTotal * 100);
    }

    public EsNodeStatsPOJO(JsonNode node) {
        this.nodeName = node.get("name").asText();
        this.host = node.get("host").asText();
        this.roles = JSONUtils.convertType(node.get("roles"), new LinkedList<>());
        this.diskTotal = node.get("fs").get("total").get("total_in_bytes").asLong();
        this.diskAvailable = node.get("fs").get("total").get("available_in_bytes").asLong();
        this.diskFree =  node.get("fs").get("total").get("free_in_bytes").asLong();
        this.diskUsedPct = (int) ((double)diskAvailable / (double)diskTotal * 100);
    }

    public void sortShardsBySize(){
        this.shards = shards.stream()
                .sorted(Comparator.comparingDouble(EsShardStatsPOJO::getSizeInGb))
                .collect(Collectors.toList());
    }
    public List<EsShardStatsPOJO> getShards() {
        return shards;
    }

    public void setShards(List<EsShardStatsPOJO> shards) {
        this.shards = shards.stream()
                .sorted(Comparator.comparingDouble(EsShardStatsPOJO::getSizeInGb))
                .collect(Collectors.toList());
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

    public String getHost() {
        return host;
    }

    public void setHost(String host) {
        this.host = host;
    }

    public List<String> getRoles() {
        return roles;
    }

    public void setRoles(List<String> roles) {
        this.roles = roles;
    }

    public long getDiskTotal() {
        return diskTotal;
    }

    public void setDiskTotal(long diskTotal) {
        this.diskTotal = diskTotal;
    }

    public long getDiskUsed() {
        return diskUsed;
    }

    public void setDiskUsed(long diskUsed) {
        this.diskUsed = diskUsed;
    }

    public long getDiskAvailable() {
        return diskAvailable;
    }

    public void setDiskAvailable(long diskAvailable) {
        this.diskAvailable = diskAvailable;
    }

    public long getDiskFree() {
        return diskFree;
    }

    public void setDiskFree(long diskFree) {
        this.diskFree = diskFree;
    }

    public int getDiskUsedPct() {
        return diskUsedPct;
    }

    public void setDiskUsedPct(int diskUsedPct) {
        this.diskUsedPct = diskUsedPct;
    }

    public List<EsShardStatsPOJO> getLowSizeShards() {
        return lowSizeShards;
    }

    public void setLowSizeShards(List<EsShardStatsPOJO> lowSizeShards) {
        this.lowSizeShards = lowSizeShards;
    }

    public List<EsShardStatsPOJO> getHighSizeShards() {
        return highSizeShards;
    }

    public void setHighSizeShards(List<EsShardStatsPOJO> highSizeShards) {
        this.highSizeShards = highSizeShards;
    }

    public long getAvgNodeSetAvailable() {
        return avgNodeSetAvailable;
    }

    public void setAvgNodeSetAvailable(long avgNodeSetAvailable) {
        this.avgNodeSetAvailable = avgNodeSetAvailable;
    }
}
