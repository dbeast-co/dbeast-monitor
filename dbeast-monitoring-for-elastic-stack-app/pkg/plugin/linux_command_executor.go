package plugin

import (
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"os"
	"os/exec"
	"path/filepath"
)

func CreateKeystoreIfNotExists(keystorePath string, keystoreFile string) bool {
	if _, err := os.Stat(filepath.Join(keystorePath, keystoreFile)); err == nil {
		log.DefaultLogger.Warn("Keystore file already exists. Skipping creation.")
		return true
	} else if os.IsNotExist(err) {
		var createCmd *exec.Cmd
		createCmd = exec.Command("echo \"y\" | /usr/share/logstash/bin/logstash-keystore", "--path.settings", keystorePath, "create")

		createOutput, executionErr := createCmd.CombinedOutput()
		if executionErr != nil {
			log.DefaultLogger.Error("Error creating Logstash keystore: %v\n", executionErr)
			log.DefaultLogger.Error(string(createOutput))
			return false
		}

		log.DefaultLogger.Info("Keystore created successfully")
		return true
	} else {
		log.DefaultLogger.Warn("Error checking keystore file: %v\n", err)
		return false
	}
}

func AddValueToKeystore(keystorePath string, key string, value string) bool {
	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", fmt.Sprintf("echo \"%s\" | logstash-keystore add \"%s\" --stdin --path.settings %s", value, key, keystorePath))
	createOutput, executionErr := cmd.CombinedOutput()
	if executionErr != nil {
		log.DefaultLogger.Warn("Error adding the variable into the keystore: %v\n", executionErr)
		log.DefaultLogger.Warn(string(createOutput))
		return false
	}

	log.DefaultLogger.Info("Key added successfully")
	return true

}

func RemoveValueFromKeystore(keystorePath string, key string, value string) bool {
	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", fmt.Sprintf("echo \"%s\" | logstash-keystore remove  \"%s\" --stdin --path.settings %s", value, key, keystorePath))
	createOutput, executionErr := cmd.CombinedOutput()
	if executionErr != nil {
		log.DefaultLogger.Warn("Error deleting the variable from the keystore: %v\n", executionErr)
		log.DefaultLogger.Warn(string(createOutput))
		return false
	}

	log.DefaultLogger.Info("Key deleted successfully")
	return true
}
