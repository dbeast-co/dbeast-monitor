package plugin

import (
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"os"
	"path/filepath"
	"strings"
)

/*
LoadTemplatesFromFolder loads JSON templates from the specified folder and updates the global TemplatesMap.
It takes a folderPath string as input, reads the content of the folder, and parses each JSON file into the TemplatesMap.
If successful, it returns nil; otherwise, it returns an error.
The TemplatesMap is a global variable that represents a mapping of template names to their corresponding JSON data.
This map is expected to be used elsewhere in the application after the templates are loaded.
*/
func LoadTemplatesFromFolder(folderPath string) error {
	log.DefaultLogger.Debug("The templates folder path: " + folderPath)
	TemplatesMap = make(map[string]interface{})

	files, err := os.ReadDir(folderPath)
	path, _ := filepath.Abs(folderPath)
	if err != nil {
		log.DefaultLogger.Error("The error in the folder read: " + err.Error())
		return fmt.Errorf("failed to read files from folder: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(path + string(os.PathSeparator) + file.Name())
			log.DefaultLogger.Info("Read file: " + filePath)
			data, err := os.ReadFile(filePath)
			if err != nil {
				log.DefaultLogger.Error("The error in the file: " + filePath + " " + err.Error())
				return fmt.Errorf("the error in the file: %v", err)
			}

			var tmp string
			err = json.Unmarshal(data, &tmp)

			var templateData map[string]interface{}
			err = json.Unmarshal(data, &templateData)
			if err != nil {
				log.DefaultLogger.Error("Error parsing file: " + filePath + " " + err.Error())
				return err
			}

			templateName := file.Name()[:len(file.Name())-5]

			TemplatesMap[templateName] = templateData

		}
	}
	TemplatesJSON, err := json.MarshalIndent(TemplatesMap, "", "")
	if err != nil {
		log.DefaultLogger.Error("Failed to marshal templates: " + err.Error())
		return err
	}

	log.DefaultLogger.Debug("Updated templates sent:" + string(TemplatesJSON))
	return nil

}

func LoadLogstashConfigFromFolder(folderPath string) error {
	LSConfigs = make(map[string]string)

	files, err := os.ReadDir(folderPath)
	if err != nil {
		return fmt.Errorf("failed to read files from folder: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".conf") {
			filePath := folderPath + "/" + file.Name()
			data, err := os.ReadFile(filePath)
			if err != nil {
				log.DefaultLogger.Error("Failed to read file %s: %v", file.Name(), err)
				continue
			}
			log.DefaultLogger.Info("Reading file: %s\n", filePath)

			//var tmp string
			//err = json.Unmarshal(data, &tmp)

			//strings.Replace(tmp, "<HOST>", "localhost", -1)
			//LSConfigs = append(LSConfigs, string(data[:]))

			//var templateData map[string]interface{}

			templateName := file.Name()[:len(file.Name())-5]

			LSConfigs[templateName] = string(data[:])
		}
	}

	return nil

}
