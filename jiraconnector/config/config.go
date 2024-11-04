package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	JiraUrl           string `yaml:"jiraUrl"`
	JiraApi           string `yaml:"jiraApi"`
	IssueInOneRequest int    `yaml:"issueInOneRequest"`
	ThreadCount       int    `yaml:"threadCount"`
}

func NewConfig(path string) (*Config, error) {
	cfg := &Config{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
