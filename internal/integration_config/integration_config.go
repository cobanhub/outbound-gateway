package integration_config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Integrations struct {
	Integrations IntegrationConfig `yaml:"integrations"`
}

type IntegrationConfig struct {
	Name               string            `yaml:"name"`
	MatchPath          string            `yaml:"matchPath"`
	Endpoint           string            `yaml:"endpoint"`
	Method             string            `yaml:"method"`
	RequestMapping     map[string]string `yaml:"requestMapping"`
	ResponseMapping    map[string]string `yaml:"responseMapping"`
	Auth               *AuthConfig       `yaml:"auth"`
	CodeMapping        map[string]string // third party codes => success/failed
	AdapterScript      string            // Lua adapter optional
	RetryCount         int               `yaml:"retryCount"`
	RetryInterval      int               `yaml:"retryInterval"`
	Timeout            int               `yaml:"timeout"`
	HeadersMapping     map[string]string `yaml:"headersMapping"`
	QueryParamsMapping map[string]string `yaml:"queryParamsMapping"`
}

type AuthConfig struct {
	Type  string `yaml:"type"`
	Token string `yaml:"token"`
}

func GetIntegrationConfig(name string) (*IntegrationConfig, error) {
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
