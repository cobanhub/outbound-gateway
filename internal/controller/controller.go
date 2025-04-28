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

type Config struct {
	IntegrationName string `yaml:"integration_name"`
	RequestMapping  string `yaml:"request_mapping"`
	ResponseMapping string `yaml:"response_mapping"`
}

func HandleOutbound(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	integrationName := vars["integration"]

	cfg, err := config.GetIntegrationConfig(integrationName)
	if err != nil {
		http.Error(w, "Integration not found", http.StatusBadRequest)
		return
	}

	var coreRequest map[string]interface{}
	json.NewDecoder(r.Body).Decode(&coreRequest)

	thirdPartyRequest, err := mapper.MapRequest(cfg, coreRequest)
	if err != nil {
		http.Error(w, "Mapping request failed", http.StatusInternalServerError)
		return
	}

	thirdPartyResp, err := client.Send(thirdPartyRequest, cfg)
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
	var config Config
	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		http.Error(w, "Invalid YAML format", http.StatusBadRequest)
		return
	}

	// Validate the config
	if config.IntegrationName == "" || config.RequestMapping == "" || config.ResponseMapping == "" {
		http.Error(w, "Missing required fields in YAML", http.StatusBadRequest)
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
func storeConfigFile(config Config) error {
	// Create the config directory if it doesn't exist
	configDir := "./config"
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.Mkdir(configDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create config directory: %v", err)
		}
	}

	// Create a unique filename for each config file (using integration name)
	configFileName := fmt.Sprintf("./config/%s.yaml", config.IntegrationName)

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
