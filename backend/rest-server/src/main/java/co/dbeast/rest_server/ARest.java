package co.dbeast.rest_server;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public abstract class ARest {
    protected static Logger logger = LogManager.getLogger();

    public abstract void rest();

}
