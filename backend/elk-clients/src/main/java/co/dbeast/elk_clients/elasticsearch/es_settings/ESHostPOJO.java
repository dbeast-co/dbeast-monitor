package co.dbeast.elk_clients.elasticsearch.es_settings;


import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class ESHostPOJO {

    private String protocol;
    private String domain;
    private int port;

    public ESHostPOJO(final ClusterConnectionSettingsPOJO esSettings) {
        Pattern pattern = Pattern.compile("(https?)://([^:^/]*):(\\d*)?(.*)?");
        Matcher matcher = pattern.matcher(esSettings.getEs_host());
        try {
            matcher.find();
            this.protocol = matcher.group(1);
            this.domain = matcher.group(2);
            this.port = Integer.parseInt(matcher.group(3));
        } catch (IllegalStateException e) {
            e.printStackTrace();
        }
        this.protocol = this.protocol != null ? this.protocol : "http";
        this.domain = this.domain != null ? this.domain : "localhost";
        this.port = this.port != 0 ? this.port : 9200;
    }

    public String getProtocol() {
        return protocol;
    }

    public String getDomain() {
        return domain;
    }

    public int getPort() {
        return port;
    }
}
