package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	JiraURL        string `yaml:"jiraURL"`
	JiraAPI        string `yaml:"jiraAPI"`
	JiraMaxResults int    `yaml:"jiraMaxResults"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	filePath := "config/config.yaml"

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
