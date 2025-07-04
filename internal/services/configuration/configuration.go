package configuration

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/cobanhub/madakaripura/internal/repository/model"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func (c *Configuration) UploadConfigHandler(file multipart.File) error {
	// Read the file content
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		// http.Error(w, "Failed to read file content", http.StatusInternalServerError)
		return errors.New("Failed to read file content")
	}

	// Parse YAML
	var config model.Integrations
	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		// http.Error(w, "Invalid YAML format", http.StatusBadRequest)
		return errors.New("Invalid YAML format")
	}

	// Validate the config
	if config.Integrations.Name == "" {
		// http.Error(w, "Missing IntegrationName fields in YAML", http.StatusBadRequest)
		return errors.New("Missing IntegrationName fields in YAML")
	}

	if config.Integrations.RequestMapping == nil {
		// http.Error(w, "Missing RequestMapping fields in YAML", http.StatusBadRequest)
		return errors.New("Missing RequestMapping fields in YAML")
	}

	if config.Integrations.ResponseMapping == nil {
		// http.Error(w, "Missing ResponseMapping fields in YAML", http.StatusBadRequest)
		return errors.New("Missing ResponseMapping fields in YAML")
	}

	// Store the YAML file in the config folder
	err = c.storeConfigFile(config)
	if err != nil {
		// http.Error(w, "Failed to store config file", http.StatusInternalServerError)
		return errors.New("Failed to store config file")
	}

	return nil
}

// storeConfigFile stores the YAML configuration in a file under the config folder
func (c *Configuration) storeConfigFile(config Integrations) error {
	// Create the config directory if it doesn't exist
	var err error
	switch c.Type {
	case Local:
		err = localConfigFilePath(config)
	case Cloud:
		err = nil // Implement remote storage logic here
	case Database:
		err = nil // Implement database storage logic here
	case SharedDrive:
		err = nil // Implement shared drive storage logic here
	default:
		err = errors.New("Invalid configuration type")
	}
	return err
}

func (c *Configuration) GetIntegrationConfig(name string) (*IntegrationConfig, error) {
	var cfg IntegrationConfig

	homedir, err := os.UserHomeDir()

	if err != nil {
		return nil, fmt.Errorf("failed to load home directory %v", err)
	}

	viper.SetConfigName(name)
	viper.AddConfigPath(homedir + "/config")
	err = viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	sub := viper.Sub("integrations")
	if sub == nil {
		return nil, fmt.Errorf("no config for integration: %s", name)
	}
	sub.Unmarshal(&cfg)
	return &cfg, nil
}
