package tmp

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func AddUserAndPasswordToKeystore(keystorePath string, clusterId string, environmentConfig EnvironmentConfig) error {
	AddKeyToKeystore(keystorePath, clusterId+"-PROD_USER", environmentConfig.Prod.Elasticsearch.Username)
	AddKeyToKeystore(keystorePath, clusterId+"-PROD_PSW", environmentConfig.Prod.Elasticsearch.Password)
	AddKeyToKeystore(keystorePath, clusterId+"-MON_USER", environmentConfig.Mon.Elasticsearch.Username)
	AddKeyToKeystore(keystorePath, clusterId+"-MON_PSW", environmentConfig.Mon.Elasticsearch.Password)
	return nil
}

func CheckAndCreateKeystore(keystorePath, keystoreFile string) bool {
	if _, err := os.Stat(filepath.Join(keystorePath, keystoreFile)); err == nil {
		log.DefaultLogger.Info("Keystore file already exists. Skipping creation.")
		return true
	} else if os.IsNotExist(err) {
		// Create the keystore creation command
		createCmd := exec.Command("/usr/share/logstash/bin/logstash-keystore", "create", "--path.settings", keystorePath)

		// Buffers to capture stdout and stderr
		var stdoutBuf, stderrBuf bytes.Buffer
		createCmd.Stdout = &stdoutBuf
		createCmd.Stderr = &stderrBuf

		// Create a pipe to send input to the command
		stdinPipe, err := createCmd.StdinPipe()
		if err != nil {
			log.DefaultLogger.Error("failed to create stdin pipe: " + err.Error())
			return false
		}

		// Start the keystore command
		if err := createCmd.Start(); err != nil {
			log.DefaultLogger.Error("failed to start keystore command: " + err.Error())
			return false
		}

		// Answer the "overwrite" prompt with "y" if it occurs
		go func() {
			scanner := bufio.NewScanner(&stdoutBuf)
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println("Logstash output:", line)

				// Look for the specific prompt in the output
				if strings.Contains(line, "Continue without password protection on the keystore? [y/N]") {
					log.DefaultLogger.Info("Responding to the password protection prompt.")
					stdinPipe.Write([]byte("y\n")) // Send "y" when the command asks for confirmation
				}
			}
		}()

		// Wait for the keystore command to finish
		if err := createCmd.Wait(); err != nil {
			log.DefaultLogger.Error("Error creating Logstash keystore: " + err.Error())
			log.DefaultLogger.Error("Stdout: " + stdoutBuf.String())
			log.DefaultLogger.Error("Stderr: " + stderrBuf.String())
			return false
		}

		// Log stdout and stderr output after successful execution
		log.DefaultLogger.Info("Keystore created successfully")
		log.DefaultLogger.Info("Stdout: " + stdoutBuf.String())
		log.DefaultLogger.Info("Stderr: " + stderrBuf.String())
		return true
	} else {
		log.DefaultLogger.Warn("Error checking keystore file: " + err.Error())
		return false
	}
}

func CreateKeystoreIfNotExists(keystorePath string, keystoreFile string) bool {
	if _, err := os.Stat(filepath.Join(keystorePath, keystoreFile)); err == nil {
		log.DefaultLogger.Info("Keystore file already exists. Skipping creation.")
		return true
	} else if os.IsNotExist(err) {
		var createCmd *exec.Cmd
		createCmd = exec.Command("/bin/sh", "-c", "echo y | /usr/share/logstash/bin/logstash-keystore create --path.settings "+keystorePath)
		log.DefaultLogger.Info("Execution string: " + createCmd.String())

		// Buffers to capture stdout and stderr
		var stdoutBuf, stderrBuf bytes.Buffer
		createCmd.Stdout = &stdoutBuf
		createCmd.Stderr = &stderrBuf

		if err := createCmd.Start(); err != nil {
			log.DefaultLogger.Error("Failed to start command: " + err.Error())
			return false
		}

		// Wait for the command to finish
		if err = createCmd.Wait(); err != nil {
			log.DefaultLogger.Error("Error creating Logstash keystore: " + err.Error())
			log.DefaultLogger.Error("Stdout: " + stdoutBuf.String())
			log.DefaultLogger.Error("Stderr: " + stderrBuf.String())
			return false
		}

		// Log stdout and stderr output after successful execution
		log.DefaultLogger.Info("Keystore created successfully")
		log.DefaultLogger.Info("Stdout: " + stdoutBuf.String())
		log.DefaultLogger.Info("Stderr: " + stderrBuf.String())
		return true

	} else {
		log.DefaultLogger.Warn("Error checking keystore file: " + err.Error())
		return false
	}
}

func AddKeyToKeystore(keystorePath, key, password string) error {
	log.DefaultLogger.Info("Ask to add key: " + key + " With password: " + password + " to the keystore: " + keystorePath)
	// Command to add key to Logstash keystore
	cmd := exec.Command(LogstashBinFolder+"/logstash-keystore", "add", key, "--stdin", "--path.settings", keystorePath)

	// Prepare to capture the output
	var outBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &outBuf

	// Get stdin pipe
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.DefaultLogger.Error("Error creating the stdin pipe for the keystore: %v\n", err)
		return fmt.Errorf("failed to get stdin pipe: %w", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		log.DefaultLogger.Error("failed to start command: %w", err)
		return fmt.Errorf("failed to start command: %w", err)
	}

	// Read the output and handle based on content
	scanner := bufio.NewScanner(&outBuf)
	for scanner.Scan() {
		line := scanner.Text()

		// Check if key already exists
		if strings.Contains(line, "already exists") {
			log.DefaultLogger.Info("Key already exists. Responding with 'y' and providing password.")
			stdin.Write([]byte("y\n"))
			break
		}

		// Check if the command is asking for the password for a new key
		if strings.Contains(line, "Enter value for") {
			log.DefaultLogger.Info("Key does not exist. Providing password.")
			break
		}
	}

	// After sending "y", or if it's a new key, send the password
	stdin.Write([]byte(password + "\n"))
	stdin.Close()

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("command failed: %w", err)
	}
	log.DefaultLogger.Info("Key added successfully")
	return nil
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
