package co.dbeast.grafana_backend.pojo.new_backend;

public class ClusterStatus {
    String status = "green";
    String error = "";

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public String getError() {
        return error;
    }

    public void setError(String error) {
        this.error = error;
    }
}
