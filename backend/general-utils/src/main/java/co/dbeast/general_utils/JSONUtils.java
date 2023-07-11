package co.dbeast.general_utils;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import java.io.File;
import java.io.IOException;
import java.nio.file.Path;
import java.nio.file.Paths;

public class JSONUtils {
    private static final Logger logger = LogManager.getLogger();
    private static final ObjectMapper mapper = new ObjectMapper();

    public static JsonNode getNewJsonNode(){
        return mapper.createObjectNode();
    }

    public static String JSONToString(Object object) {
        mapper.configure(JsonGenerator.Feature.ESCAPE_NON_ASCII, false);
        try {
            return mapper.writeValueAsString(object);
        } catch (JsonProcessingException e) {
            throw new RuntimeException(e);
        }
    }

    public static JsonNode stringToJSON(final String object) {
        try {
            return mapper.readTree(object);
        } catch (JsonProcessingException e) {
            throw new RuntimeException(e);
        }
    }

    public static <T> T convertType(final JsonNode source,
                                    final TypeReference<T> destinationClass) {
        return mapper.convertValue(source, destinationClass);
    }
    public static <T> T convertType(final JsonNode source,
                                    final T destinationClass) {
        TypeReference<T> typeReference = new TypeReference<T>() {
        };
        return mapper.convertValue(source, typeReference);
    }

    public static String jsonFileToString(final String filePath) throws IOException {
        logger.info("Read file: " + filePath);
        JsonNode fileContent = mapper.readTree(Paths.get(filePath).toFile());
        return JSONToString(fileContent);
    }
    public static String jsonFileToString(final Path filePath) throws IOException {
        logger.info("Read file: " + filePath.toString());
        JsonNode fileContent = mapper.readTree(filePath.toFile());
        return JSONToString(fileContent);
    }
    public static <T> T jsonFileToObject(final String filePath,
                                         final T destinationClass) throws IOException {
        logger.info("Read file: " + filePath);
        TypeReference<T> typeReference = new TypeReference<T>() {
        };
        return mapper.readValue(filePath, typeReference);
    }
    public static <T> T jsonStringToObject(final String context,
                                         final Class<T> destinationClass) throws IOException {
        TypeReference<T> typeReference = new TypeReference<T>() {
        };
        return mapper.readValue(context, destinationClass);
    }

    public static <T> T jsonStringToObject(final String context,
                                         final TypeReference<T> destinationClass) throws IOException {
        return mapper.readValue(context, destinationClass);
    }
    public static <T> T jsonStringToObject(final String context) throws IOException {
        TypeReference<T> typeReference = new TypeReference<T>() {
        };
        return mapper.readValue(context, typeReference);
    }

    public static boolean saveJSONToFile(final String fileAbsolutePath, final Object context) {
        File file = new File(fileAbsolutePath);
        ObjectMapper JSONMapper = new ObjectMapper();
        JSONMapper.configure(JsonGenerator.Feature.ESCAPE_NON_ASCII, false);
        try {
            JSONMapper.writeValue(file, context);
            if (logger.isDebugEnabled()) {
                logger.debug("File: " + fileAbsolutePath + " successfully created");
            }
            return true;
        } catch (IOException e) {
            logger.error("There is the error in creating file: " +
                    fileAbsolutePath + "\n" +
                    e);
            e.printStackTrace();
            return false;
        }
    }
}
