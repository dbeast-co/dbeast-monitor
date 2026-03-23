package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

// GrafanaDataSource represents a Grafana datasource configuration
type GrafanaDataSource struct {
	Access          string                 `json:"access"`
	BasicAuth       bool                   `json:"basicAuth"`
	BasicAuthUser   string                 `json:"basicAuthUser,omitempty"`
	Database        string                 `json:"database"`
	ID              int                    `json:"id,omitempty"`
	IsDefault       bool                   `json:"isDefault"`
	JSONData        map[string]interface{} `json:"jsonData"`
	Name            string                 `json:"name"`
	OrgID           int                    `json:"orgId,omitempty"`
	ReadOnly        bool                   `json:"readOnly"`
	SecureJSONData  map[string]interface{} `json:"secureJsonData,omitempty"`
	Type            string                 `json:"type"`
	TypeLogoURL     string                 `json:"typeLogoUrl,omitempty"`
	TypeName        string                 `json:"typeName,omitempty"`
	UID             string                 `json:"uid,omitempty"`
	URL             string                 `json:"url"`
	User            string                 `json:"user,omitempty"`
	WithCredentials bool                   `json:"withCredentials,omitempty"`
}

func AddDataSource(client *http.Client, grafanaURL string, dataSource interface{}, dsName string) error {
	requestURL := grafanaURL + "/api/datasources"
	log.DefaultLogger.Debug("Request path: ", requestURL)

	dataSourceJSON, err := json.Marshal(dataSource)
	if err != nil {
		log.DefaultLogger.Error("Error marshaling datasource: " + err.Error())
		return fmt.Errorf("marshal datasource: %w", err)
	}

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(dataSourceJSON))
	if err != nil {
		log.DefaultLogger.Error("Error creating HTTP request: " + err.Error())
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		if response != nil && response.StatusCode == http.StatusConflict {
			return nil
		}
		log.DefaultLogger.Error("Error making HTTP request: " + err.Error())
		return fmt.Errorf("http request failed: %w", err)
	}
	defer DeferInternalHandler(response, log.DefaultLogger)

	return nil
}

// DeleteDataSourceByName deletes a datasource by its name
// Makes a DELETE request to /api/datasources/name/{name} endpoint
func DeleteDataSourceByName(client *http.Client, grafanaURL string, name string) error {
	requestURL := grafanaURL + "/api/datasources/name/" + name
	log.DefaultLogger.Debug("Request path: ", requestURL)
	log.DefaultLogger.Debug("Deleting datasource: ", name)

	req, err := http.NewRequest("DELETE", requestURL, nil)
	if err != nil {
		log.DefaultLogger.Error("Error creating HTTP request: " + err.Error())
		return fmt.Errorf("create request: %w", err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.DefaultLogger.Error("Error making HTTP request: " + err.Error())
		return fmt.Errorf("http request failed: %w", err)
	}
	defer DeferInternalHandler(response, log.DefaultLogger)

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return fmt.Errorf("delete datasource: HTTP %s, body: %s",
			response.Status, string(body))
	}

	log.DefaultLogger.Info("Datasource deleted successfully: " + name)
	return nil
}

// GetDataSources retrieves all datasources from Grafana
// Makes a GET request to /api/datasources endpoint
func GetDataSources(client *http.Client, grafanaURL string) ([]GrafanaDataSource, error) {
	requestURL := grafanaURL + "/api/datasources"
	log.DefaultLogger.Debug("Request path: ", requestURL)

	response, err := ProcessGETRequest(client, requestURL)
	if err != nil {
		log.DefaultLogger.Error("Error making HTTP request: " + err.Error())
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer DeferInternalHandler(response, log.DefaultLogger)

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("get datasources: HTTP %s, body: %s",
			response.Status, string(body))
	}

	var dataSources []GrafanaDataSource
	if err := json.NewDecoder(response.Body).Decode(&dataSources); err != nil {
		return nil, fmt.Errorf("decode datasources: %w", err)
	}

	log.DefaultLogger.Debug("Retrieved datasources count: " + fmt.Sprintf("%d", len(dataSources)))
	return dataSources, nil
}

// DeleteDataSourcesByClusterID deletes all datasources that end with a given cluster UUID
// This matches the pattern used in the UI (DataSourceItem component)
// It retrieves all datasources and filters those ending with the cluster UUID, then deletes them
func DeleteDataSourcesByClusterID(client *http.Client, grafanaURL string, clusterUUID string) error {
	log.DefaultLogger.Debug("Deleting datasources for cluster UUID: ", clusterUUID)

	dataSources, err := GetDataSources(client, grafanaURL)
	if err != nil {
		log.DefaultLogger.Error("Error retrieving datasources: " + err.Error())
		return fmt.Errorf("get datasources: %w", err)
	}

	deletedCount := 0
	for _, ds := range dataSources {
		// Check if datasource name ends with the cluster UUID
		if len(ds.Name) >= len(clusterUUID) && ds.Name[len(ds.Name)-len(clusterUUID):] == clusterUUID {
			if err := DeleteDataSourceByName(client, grafanaURL, ds.Name); err != nil {
				log.DefaultLogger.Error("Error deleting datasource: " + ds.Name + " " + err.Error())
				return fmt.Errorf("delete datasource %s: %w", ds.Name, err)
			}
			deletedCount++
			log.DefaultLogger.Info("Datasource deleted: " + ds.Name)
		}
	}

	log.DefaultLogger.Info("Deleted " + fmt.Sprintf("%d", deletedCount) + " datasources for cluster UUID: " + clusterUUID)
	return nil
}
