package connector

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/config"
	"github.com/internal/entities"
)

type JiraConnector struct {
	HttpClient *http.Client
}

func NewJiraConnector() *JiraConnector {
	jiraConnector := new(JiraConnector)
	jiraConnector.HttpClient = &http.Client{}
	return jiraConnector
}

func (jiraConnector *JiraConnector) GetProjectData(projectKey string) {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	project := jiraConnector.doRequest(fmt.Sprintf("%s/%s/search?jql=project=%s&maxResults=%d",
		cfg.JiraURL, cfg.JiraAPI, projectKey, cfg.JiraMaxResults))

	pageCount := project.IssuesCount / cfg.JiraMaxResults

	if project.IssuesCount%cfg.JiraMaxResults != 0 {
		pageCount++
	}

	for i := 1; i <= pageCount; i++ {
		startAt := i * cfg.JiraMaxResults
		data := jiraConnector.doRequest(fmt.Sprintf("%s/%s/search?jql=project=%s&startAt=%d&maxResults=%d",
			cfg.JiraURL, cfg.JiraAPI, projectKey, startAt, cfg.JiraMaxResults))
		project.Issues = append(project.Issues, data.Issues...)
	}
}

func (jiraConnector *JiraConnector) doRequest(link string) entities.Data {
	response, err := jiraConnector.HttpClient.Get(link)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	issuesData := entities.Data{}
	data, _ := io.ReadAll(response.Body)
	_ = json.Unmarshal(data, &issuesData)
	return issuesData
}

//TODO: разбить на потоки
