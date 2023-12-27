package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type DatabaseConf struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}
type ServerConf struct {
	Port int `yaml:"port"`
}

type ServerConfig struct {
	Database DatabaseConf `yaml:"database"`
	Server   ServerConf   `yaml:"server"`
}

func LoadServerConfig(configPath string) (*ServerConfig, error) {
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config ServerConfig
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return &config, nil
}
