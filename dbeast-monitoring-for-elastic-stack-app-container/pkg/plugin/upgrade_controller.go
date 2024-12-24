package plugin

import (
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"net/http"
	"strings"
)

func (a *App) UpdateClusterHandler(response http.ResponseWriter, request *http.Request) {
	ctxLogger := log.DefaultLogger.FromContext(request.Context())
	ctxLogger.Info("Got request for the cluster update")
	response.Header().Add("Content-Type", "application/json")

	var monClusterConnectionSettings Credentials
	if err := json.NewDecoder(request.Body).Decode(&monClusterConnectionSettings); err != nil {
		log.DefaultLogger.Error("Failed to decode JSON data: " + err.Error())
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	sanitizeHost(&monClusterConnectionSettings.Host)
	defer request.Body.Close()

	err := SendComponentTemplatesToMonitoringCluster(monClusterConnectionSettings)
	if err != nil {
		log.DefaultLogger.Error("Error while the Component template injection: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	err = SendIndexTemplatesToMonitoringCluster(monClusterConnectionSettings)
	if err != nil {
		log.DefaultLogger.Error("Error while the Index template injection: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	err = SendIndexRolloverOrDeleteStatusIndexToMonitoringCluster(monClusterConnectionSettings)
	if err != nil {
		log.DefaultLogger.Error("Error in the index rollover: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	//Delete status indices

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode("True")

}

func SendIndexRolloverOrDeleteStatusIndexToMonitoringCluster(credentials Credentials) error {
	log.DefaultLogger.Info("Send Rollover Commands")
	for templateName, templateContent := range ESIndexTemplatesMap {
		if !strings.Contains(templateName, "status") {
			log.DefaultLogger.Debug("Template content: " + templateContent)
			rolloverAlias := GetRolloverAliasFromTemplate(templateContent)
			log.DefaultLogger.Debug("Rollover index: "+templateName+" Rollover alias: "+rolloverAlias+" Cluster: ", credentials.Host)
			isIndexExists, _ := CheckIsIndexExists(credentials, "*"+rolloverAlias+"*")
			if isIndexExists {
				_, err := SendIndexRolloverCommandToCluster(credentials, rolloverAlias)
				if err != nil {
					return err
				}
			} else {
				if _, isFirstIndex := ESFirstIndicesTemplatesMap[rolloverAlias]; isFirstIndex {
					log.DefaultLogger.Debug("Is index exists response: ", isFirstIndex)
					if !isFirstIndex {
						_, err := SendFirstIndexToCluster(credentials, templateName, templateContent)
						if err != nil {
							return err
						}
					}
				}
			}
		} else {
			for _, indexPattern := range GetIndexPatternsFromTemplate(templateContent) {
				log.DefaultLogger.Info("Index for delete:" + indexPattern)
				isIndexExists, _ := CheckIsIndexExists(credentials, indexPattern)
				if isIndexExists {
					indexDeleteResponse, err := DeleteIndex(credentials, indexPattern)
					if err != nil {
						return err
					}
					log.DefaultLogger.Info("Delete Index response: ", indexDeleteResponse)
				}

			}
		}
	}
	return nil
}

func GetRolloverAliasFromTemplate(stringTemplate string) string {
	var template IndexTemplate
	err := json.Unmarshal([]byte(stringTemplate), &template)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return ""
	}
	return template.Template.Settings.Index.Lifecycle.RolloverAlias
}
func GetIndexPatternsFromTemplate(stringTemplate string) []string {
	var template IndexTemplate
	err := json.Unmarshal([]byte(stringTemplate), &template)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	return template.IndexPatterns
}
