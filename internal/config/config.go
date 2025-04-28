package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type IntegrationConfig struct {
	Name            string
	MatchPath       string
	Endpoint        string
	Method          string
	RequestMapping  map[string]string
	ResponseMapping map[string]string
	Auth            *AuthConfig
	CodeMapping     map[string]string // third party codes => success/failed
	AdapterScript   string            // Lua adapter optional
}

type AuthConfig struct {
	Type  string `yaml:"type"`
	Token string `yaml:"token"`
}

func GetIntegrationConfig(name string) (*IntegrationConfig, error) {
	var cfg IntegrationConfig
	viper.SetConfigName("integration")
	viper.AddConfigPath("./configs")
	viper.ReadInConfig()

	sub := viper.Sub("integrations." + name)
	if sub == nil {
		return nil, fmt.Errorf("no config for integration: %s", name)
	}
	sub.Unmarshal(&cfg)
	return &cfg, nil
}
