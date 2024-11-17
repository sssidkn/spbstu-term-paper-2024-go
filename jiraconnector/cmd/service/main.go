package main

import (
	"fmt"

	"github.com/jiraconnector/config"
	"github.com/jiraconnector/internal/connector"
)

func main() {
	cfg, err := config.NewConfig("jiraconnector/config/config.yaml")
	if err != nil {
		panic(err)
	}
	jiraConnector := connector.NewJiraConnector(cfg)
	data, err := jiraConnector.GetProjectData("AMQ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data.Issues[0])
}
