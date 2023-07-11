package co.dbeast.elk_clients;

public class ElkClients {
    public static String HOME_FOLDER;

    private static final ElkClients _instance = new ElkClients();

    public static synchronized ElkClients getInstance() {
        if (_instance == null) {
            return new ElkClients();
        }
        return _instance;
    }

    private ElkClients() {
    }
    public void init(String homeFolder){
        HOME_FOLDER = homeFolder;
    }
}
