package co.dbeast.elk_clients.elasticsearch.pojo;


import java.util.*;
import java.util.stream.Collectors;

public class EsClusterStatsPOJO {

    private String clusterName;
    private String clusterUID;
    private List<EsNodeStatsPOJO> nodesStats = new LinkedList<>();
    private Map<String, ArrayList<EsNodeStatsPOJO>> nodesSetForBalance = new HashMap<>();


    public List<EsNodeStatsPOJO> getNodesStats() {
        return nodesStats;
    }

    public void setNodesStats(List<EsNodeStatsPOJO> nodesStats) {
        this.nodesStats = nodesStats;
    }

    public String getClusterName() {
        return clusterName;
    }

    public void setClusterName(String clusterName) {
        this.clusterName = clusterName;
    }

    public String getClusterUID() {
        return clusterUID;
    }

    public void setClusterUID(String clusterUID) {
        this.clusterUID = clusterUID;
    }

    public Map<String, ArrayList<EsNodeStatsPOJO>> getNodesSetForBalance() {
        return nodesSetForBalance;
    }

    public void setNodesSetForBalance(Map<String, ArrayList<EsNodeStatsPOJO>> nodesSetForBalance) {
        this.nodesSetForBalance = nodesSetForBalance;
    }

    public void addNodesSetForBalance(String setName,
                                      ArrayList<EsNodeStatsPOJO> setForBalance) {
        this.nodesSetForBalance.put(setName,setForBalance);
    }

}
