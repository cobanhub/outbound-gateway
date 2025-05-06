package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/cobanhub/outbound-gateway/internal/client"
	"github.com/cobanhub/outbound-gateway/internal/config"
	"github.com/cobanhub/outbound-gateway/internal/mapper"
	"gopkg.in/yaml.v3"

	"github.com/gorilla/mux"
)

func HandleOutbound(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	integrationName := vars["integration"]
	var coreRequest map[string]interface{}

	cfg, err := config.GetIntegrationConfig(integrationName)
	if err != nil {
		http.Error(w, "Integration not found", http.StatusBadRequest)
		return
	}

	headers := r.Header

	thirdPartyRequest, thirdPartyHeadersReq, err := mapper.MapRequestWithHeaders(cfg, coreRequest, headers)
	if err != nil {
		http.Error(w, "Mapping request failed", http.StatusInternalServerError)
		return
	}

	thirdPartyResp, err := client.Send(thirdPartyRequest, thirdPartyHeadersReq, cfg)
	if err != nil {
		http.Error(w, "Failed to contact third party", http.StatusBadGateway)
		return
	}

	coreResponse, err := mapper.MapResponse(cfg, thirdPartyResp)
	if err != nil {
		http.Error(w, "Mapping response failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coreResponse)
	log.Println("Success outbound for integration:", integrationName)
}
func UploadConfigHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the YAML file from the request
	file, _, err := r.FormFile("config_file")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the file content
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file content", http.StatusInternalServerError)
		return
	}

	// Parse YAML
	var config config.Integrations
	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		http.Error(w, "Invalid YAML format", http.StatusBadRequest)
		return
	}

	// Validate the config
	if config.Integrations.Name == "" {
		http.Error(w, "Missing IntegrationName fields in YAML", http.StatusBadRequest)
		return
	}

	if config.Integrations.RequestMapping == nil {
		http.Error(w, "Missing RequestMapping fields in YAML", http.StatusBadRequest)
		return
	}

	if config.Integrations.ResponseMapping == nil {
		http.Error(w, "Missing ResponseMapping fields in YAML", http.StatusBadRequest)
		return
	}

	// Store the YAML file in the config folder
	err = storeConfigFile(config)
	if err != nil {
		http.Error(w, "Failed to store config file", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Config uploaded successfully"))
}

// storeConfigFile stores the YAML configuration in a file under the config folder
func storeConfigFile(config config.Integrations) error {
	// Create the config directory if it doesn't exist

	homedir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %v", err)
	}
	configDir := homedir + "/config"
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.Mkdir(configDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create config directory: %v", err)
		}
	}

	// Create a unique filename for each config file (using integration name)
	configFileName := fmt.Sprintf(homedir+"/config/%s.yaml", config.Integrations.Name)

	// Serialize the config back to YAML and save it
	configData, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("failed to marshal config to YAML: %v", err)
	}

	err = ioutil.WriteFile(configFileName, configData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}

	log.Printf("Config file saved: %s", configFileName)
	return nil
}
