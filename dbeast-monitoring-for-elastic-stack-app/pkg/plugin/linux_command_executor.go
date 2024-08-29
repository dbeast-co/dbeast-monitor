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

		createOutput, createErr := createCmd.CombinedOutput()
		if createErr != nil {
			log.DefaultLogger.Error("Error creating Logstash keystore: %v\n", createErr)
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
	var createCmd *exec.Cmd
	createCmd = exec.Command("echo \"y\" | /usr/share/logstash/bin/logstash-keystore", "--path.settings", keystorePath, "create")

	createOutput, createErr := createCmd.CombinedOutput()
	if createErr != nil {
		log.DefaultLogger.Warn("Error creating Logstash keystore: %v\n", createErr)
		log.DefaultLogger.Warn(string(createOutput))
		return false
	}

	log.DefaultLogger.Info("Key added successfully")
	return true

}

func RemoveValueFromKeystore(keystorePath string, key string) bool {
	return true
}

func main() {
	// Define the path where the keystore should be created
	keystorePath := "/path/to/keystore" // Replace with your desired path or leave empty for default

	// Determine the full path to the keystore file
	var keystoreFile string
	if keystorePath != "" {
		keystoreFile = filepath.Join(keystorePath, "logstash.keystore")
	} else {
		// Use the default path settings if not specified
		keystoreFile = "logstash.keystore"
	}

	// Check if the keystore file already exists
	if _, err := os.Stat(keystoreFile); err == nil {
		fmt.Println("Keystore file already exists. Skipping creation.")
	} else if os.IsNotExist(err) {
		// Create the new keystore if it doesn't exist
		var createCmd *exec.Cmd
		if keystorePath != "" {
			// If a custom keystore path is provided
			createCmd = exec.Command("logstash-keystore", "create", "-f", "--path.settings", keystorePath)
		} else {
			// Default path
			createCmd = exec.Command("logstash-keystore", "create", "-f")
		}

		// Execute the command to create the keystore
		createOutput, createErr := createCmd.CombinedOutput()
		if createErr != nil {
			fmt.Printf("Error creating Logstash keystore: %v\n", createErr)
			fmt.Println(string(createOutput))
			return
		}

		fmt.Println("Keystore created successfully")
	} else {
		fmt.Printf("Error checking keystore file: %v\n", err)
		return
	}

	// Define the password and the keystore entry name
	password := "your_password"
	entryName := "keystore_entry_name"

	// Command to add a password to the keystore
	var addCmd *exec.Cmd
	if keystorePath != "" {
		// Add password to the custom path keystore
		addCmd = exec.Command("sh", "-c", fmt.Sprintf("echo \"%s\" | logstash-keystore add \"%s\" --stdin --path.settings %s", password, entryName, keystorePath))
	} else {
		// Add password to the default path keystore
		addCmd = exec.Command("sh", "-c", fmt.Sprintf("echo \"%s\" | logstash-keystore add \"%s\" --stdin", password, entryName))
	}

	// Execute the command to add the password to the keystore
	addOutput, addErr := addCmd.CombinedOutput()
	if addErr != nil {
		fmt.Printf("Error adding password to Logstash keystore: %v\n", addErr)
		fmt.Println(string(addOutput))
		return
	}

	fmt.Println("Password added to keystore successfully")
}
