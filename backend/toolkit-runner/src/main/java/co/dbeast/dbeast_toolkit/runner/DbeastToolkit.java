package co.dbeast.dbeast_toolkit.runner;

import co.dbeast.elk_clients.ElkClients;
import co.dbeast.grafana_backend.GrafanaBackend;
import co.dbeast.rest_server.RestServer;
import org.apache.commons.cli.*;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class DbeastToolkit {
    public static String HOME_FOLDER;
    private static final Logger logger = LogManager.getLogger();

    public static void main(String[] args) throws Exception {
        System.out.println("Welcome to DBeast toolkit for Elasticsearch");
        try {
            initStartMenuAndRun(args);
        } catch (Exception exp) {
            System.err.println("Parsing failed.  Reason: " + exp.getMessage());
            System.exit(-1);
        }
    }

    private static void initStartMenuAndRun(String[] args) {
        Options options = initOptions();
        CommandLineParser parser = new DefaultParser();
        CommandLine cmd = null;
        HelpFormatter formatter = new HelpFormatter();
        try {
            cmd = parser.parse(options, args);
        } catch (ParseException e) {
            boolean isHelp = false;
            for (String arg : args) {
                if (arg.equals("-h") || arg.equals("--help")) {
                    isHelp = true;
                }
            }
            if (!isHelp) {
                System.err.println("Parsing failed.  Reason: " + e.getMessage());
                System.exit(-1);
            } else {
                formatter.printHelp("dbeast-toolkit.sh <ARGUMENTS>", options);
                System.exit(1);
            }
        }
        if (cmd.hasOption(options.getOption("help"))) {
            formatter.printHelp("dbeast-toolkit.sh <ARGUMENTS>", options);
        }
        HOME_FOLDER = cmd.getOptionValue(options.getOption("config"));
        try {
            run();
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    private static Options initOptions() {
        Options options = new Options();
        Option homeFolder = Option.builder("c")
                .longOpt("config")
                .desc("application config folder (REQUIRED FOR RUNNING WITHOUT SCRIPT!)")
                .required()
                .hasArg(true)
                .numberOfArgs(1)
                .build();
        Option help = Option.builder("h")
                .longOpt("help")
                .desc("print this message")
                .build();
        options.addOption(homeFolder);
        options.addOption(help);
        return options;
    }

    public static void run() throws Exception {
        ElkClients.HOME_FOLDER = HOME_FOLDER;
        GrafanaBackend.init(HOME_FOLDER,
                EDBeastToolkitRunnerSettings.APPLICATION_KEYSTORE_PSW.getSetting()
        );
        initPluginsRestApis();
    }

    public static void initPluginsRestApis() throws Exception {
        RestServer.runServer(EDBeastToolkitRunnerSettings.API_SERVER_CONFIG.getSetting());
        RestServer.initRestApi(co.dbeast.grafana_backend.rest.MainRest.class);
    }
}