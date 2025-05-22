package config

import "time"

type (
	MainConfig struct {
		Server ServerConfig `yaml:"Server"`
	}

	ServerConfig struct {
		Port            string        `yaml:"Port"`
		GracefulTimeout time.Duration `yaml:"GracefulTimeout"`
		ReadTimeout     time.Duration `yaml:"ReadTimeout"`
		WriteTimeout    time.Duration `yaml:"WriteTimeout"`
	}

	APIConfig struct {
		BasePath      string `yaml:"BasePath"`
		APITimeout    int    `yaml:"APITimeout"`
		EnableSwagger bool   `yaml:"EnableSwagger" default:"true"`
	}
)
