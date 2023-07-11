package co.dbeast.general_utils.keystore;

import co.dbeast.general_utils.GeneralUtils;
import org.apache.commons.io.FileUtils;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import javax.crypto.SecretKey;
import javax.crypto.SecretKeyFactory;
import javax.crypto.spec.PBEKeySpec;
import java.io.*;
import java.security.KeyStore;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;
import java.security.cert.CertificateException;

public class KeystoreController {
    private static final Logger logger = LogManager.getLogger();

    public static String getPasswordFromKeystore(String key,
                                                 String keystoreLocation,
                                                 String keyStorePassword) throws Exception {

        KeyStore ks = KeyStore.getInstance("JCEKS");
        ks.load(null, keyStorePassword.toCharArray());
        KeyStore.PasswordProtection keyStorePP = new KeyStore.PasswordProtection(keyStorePassword.toCharArray());

        FileInputStream fIn = new FileInputStream(keystoreLocation);

        ks.load(fIn, keyStorePassword.toCharArray());

        SecretKeyFactory factory = SecretKeyFactory.getInstance("PBE");

        KeyStore.SecretKeyEntry ske =
                (KeyStore.SecretKeyEntry) ks.getEntry(key, keyStorePP);

        PBEKeySpec keySpec = (PBEKeySpec) factory.getKeySpec(
                ske.getSecretKey(),
                PBEKeySpec.class);

        char[] password = keySpec.getPassword();
        fIn.close();
        return new String(password);

    }


    public static void createNewKeystoreEntry(String entry,
                                              String entryPassword,
                                              String keyStoreLocation,
                                              String keyStorePassword)
            throws Exception {
        SecretKeyFactory factory = SecretKeyFactory.getInstance("PBE");
        SecretKey generatedSecret =
                factory.generateSecret(new PBEKeySpec(
                        entryPassword.toCharArray()));

        KeyStore ks = KeyStore.getInstance("JCEKS");
        if (GeneralUtils.isFileExists(keyStoreLocation)) {
            FileInputStream fIn = new FileInputStream(keyStoreLocation);
            ks.load(fIn, keyStorePassword.toCharArray());
            fIn.close();
        } else {
            ks.load(null, keyStorePassword.toCharArray());
        }
        KeyStore.PasswordProtection keyStorePP = new KeyStore.PasswordProtection(keyStorePassword.toCharArray());


        FileOutputStream fos = new java.io.FileOutputStream(keyStoreLocation);
        ks.setEntry(entry, new KeyStore.SecretKeyEntry(
                generatedSecret), keyStorePP);
        ks.store(fos, keyStorePassword.toCharArray());
        fos.close();
        logger.info("Key for project: " + entry + " successfully created");
    }

    public static boolean deleteKey(String key,
                                    String keystoreLocation,
                                    String adminPassword) throws Exception {
        KeyStore ks = KeyStore.getInstance("JCEKS");
        FileInputStream fIn = new FileInputStream(keystoreLocation);
        try {
            ks.load(fIn, adminPassword.toCharArray());
            fIn.close();
            ks.deleteEntry(key);
            FileOutputStream fos = new java.io.FileOutputStream(keystoreLocation);
            ks.store(fos, adminPassword.toCharArray());
        } catch (KeyStoreException e) {
            e.printStackTrace();
        }
        fIn.close();
        logger.info("Key for project: " + key + " successfully deleted");
        return true;
    }

    public static boolean isKeyExists(String key,
                                      String keystoreLocation,
                                      String adminPassword) throws Exception {
        KeyStore ks = KeyStore.getInstance("JCEKS");
        FileInputStream fIn = new FileInputStream(keystoreLocation);
        ks.load(fIn, adminPassword.toCharArray());
        boolean result = ks.containsAlias(key);
        fIn.close();
        return result;
    }

    public static boolean isKeystoreExists(String keystoreLocation) {
        try {
            FileInputStream fIn = new FileInputStream(keystoreLocation);
            fIn.close();
            return true;
        } catch (IOException e) {
            return false;
        }
    }

    public static void deleteKeystore(String keystoreLocation) throws IOException {
        File keystoreFile = new File(keystoreLocation);
        FileUtils.forceDelete(keystoreFile);
    }
}
