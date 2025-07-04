package repository

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/cobanhub/madakaripura/internal/repository/model"
	"gopkg.in/yaml.v3"
)

func Save(config model.Integrations) error {
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
