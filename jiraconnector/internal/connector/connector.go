package connector

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/jiraconnector/config"
	"github.com/jiraconnector/internal/entities"
)

type JiraConnector struct {
	HttpClient *http.Client
	Config     *config.Config
}

func NewJiraConnector(cfg *config.Config) *JiraConnector {
	jiraConnector := new(JiraConnector)
	jiraConnector.HttpClient = &http.Client{}
	jiraConnector.Config = cfg
	return jiraConnector
}

func (jiraConnector *JiraConnector) GetProjectData(projectKey string) {
	cfg := jiraConnector.Config
	project := jiraConnector.doRequest(fmt.Sprintf("%s/%s/search?jql=project=%s&maxResults=%d",
		cfg.JiraUrl, cfg.JiraApi, projectKey, cfg.IssueInOneRequest))

	pageCount := project.IssuesCount / cfg.IssueInOneRequest

	if project.IssuesCount%cfg.IssueInOneRequest != 0 {
		pageCount++
	}

	wg := sync.WaitGroup{}
	ch := make(chan entities.Data, pageCount)
	//TODO: ограничение количества горутин
	for i := 1; i <= pageCount; i++ {
		wg.Add(1)
		startAt := i * cfg.IssueInOneRequest
		go func() {
			defer wg.Done()
			data := jiraConnector.doRequest(fmt.Sprintf("%s/%s/search?jql=project=%s&startAt=%d&maxResults=%d",
				cfg.JiraUrl, cfg.JiraApi, projectKey, startAt, cfg.IssueInOneRequest))
			ch <- data
		}()
	}
	wg.Wait()
	close(ch)
	for data := range ch {
		project.Issues = append(project.Issues, data.Issues...)
	}
	fmt.Println(project.Issues)
}

func (jiraConnector *JiraConnector) doRequest(link string) entities.Data {
	response, err := jiraConnector.HttpClient.Get(link)
	//TODO: задержка между запросами для обхода блокировки JIRA API
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	issuesData := entities.Data{}
	data, _ := io.ReadAll(response.Body)
	_ = json.Unmarshal(data, &issuesData)
	return issuesData
}
