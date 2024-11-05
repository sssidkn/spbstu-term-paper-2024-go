package main

import (
	"github.com/jiraconnector/config"
	"github.com/jiraconnector/internal/connector"
)

func main() {
	cfg, err := config.NewConfig("jiraconnector/config/config.yaml")
	if err != nil {
		panic(err)
	}
	connector := connector.NewJiraConnector(cfg)
	connector.GetProjectData("AMQ")
}
