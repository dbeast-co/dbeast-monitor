package plugin

import (
	"encoding/json"
	"net/http"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func HTTPErrorGenerator(response http.ResponseWriter, error error, errorMessage string, errorCode int, logger log.Logger) {
	logger.Error(errorMessage + error.Error())
	response.WriteHeader(errorCode)
	err := json.NewEncoder(response).Encode(map[string]interface{}{"error": error.Error()})
	if err != nil {
		logger.Error("Failed to encode error message: " + err.Error())
		return
	}
	return
}
