package connector

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/config"
	"github.com/internal/entities"
)

func GetProjectData(projectKey string) {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	httpClient := &http.Client{}
	response, err := httpClient.Get(fmt.Sprintf("%s/%s/search?jql=project=%s&maxResults=%d",
		cfg.JiraURL, cfg.JiraAPI, projectKey, cfg.JiraMaxResults))
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}
	project := entities.Project{}
	project.Key = projectKey
	data, _ := io.ReadAll(response.Body)
	_ = json.Unmarshal(data, &project)
	pageCount := project.IssuesCount / cfg.JiraMaxResults

	if project.IssuesCount%cfg.JiraMaxResults != 0 {
		pageCount++
	}

	for i := 1; i <= pageCount; i++ {
		startAt := i * cfg.JiraMaxResults
		response, err = httpClient.Get(fmt.Sprintf("%s/%s/search?jql=project=%s&startAt=%d&maxResults=%d",
			cfg.JiraURL, cfg.JiraAPI, projectKey, startAt, cfg.JiraMaxResults))
		tempProject := entities.Project{}
		data, _ = io.ReadAll(response.Body)
		_ = json.Unmarshal(data, &tempProject)
		project.Issues = append(project.Issues, tempProject.Issues...)
		defer response.Body.Close()
	}
}
