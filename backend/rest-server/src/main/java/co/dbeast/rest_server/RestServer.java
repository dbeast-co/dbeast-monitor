package co.dbeast.rest_server;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import java.io.File;
import java.io.IOException;

public class RestServer {
    private static final Logger logger = LogManager.getLogger();
    private static MainRest restApi;

    public static void main(String[] args) throws Exception {
        logger.info("Running REST server");
        if (args.length > 1) {
            System.exit(0);
        } else {
            runServer(args[0]);
        }
    }

    public static void runServer(final String configFileAbsolutePathPath) throws Exception {
        ObjectMapper mapper = new ObjectMapper(new YAMLFactory());
        ServerConfig serverConfig;
        try {
            File configFile = new File(configFileAbsolutePathPath);
            serverConfig = mapper.readValue(configFile, ServerConfig.class);
        } catch (IOException e) {
            throw new Exception("Configuration folder not present or Incorrect configuration file" + e);
        }
        restApi = new MainRest(serverConfig);
        restApi.runServer();
    }

    public static void initRestApi(Class apiClass ){
        restApi.initApi(apiClass);
    }

}
