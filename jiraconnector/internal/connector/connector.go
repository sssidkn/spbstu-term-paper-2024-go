package connector

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/jiraconnector/config"
	"github.com/jiraconnector/internal/entities"
	"gopkg.in/yaml.v3"
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

func (jiraConnector *JiraConnector) GetProjectData(projectKey string) (*entities.Data, error) {
	cfg := jiraConnector.Config

	firstRequestData := jiraConnector.doRequest(fmt.Sprintf("%s/%s/search?jql=project=%s&expand=changelog&maxResults=%d",
		cfg.JiraUrl, cfg.JiraApi, projectKey, cfg.IssueInOneRequest))

	projectData := new(entities.Data)
	_ = json.Unmarshal(firstRequestData, &projectData)

	pageCount := projectData.IssuesCount / cfg.IssueInOneRequest
	if projectData.IssuesCount%cfg.IssueInOneRequest != 0 {
		pageCount++
	}

	wg := sync.WaitGroup{}
	requests := make(chan int, pageCount)
	ch := make(chan []byte, pageCount)
	wg.Add(pageCount)

	for i := 0; i < cfg.ThreadCount; i++ {
		go func() {
			for startAt := range requests {
				data := jiraConnector.doRequest(fmt.Sprintf("%s/%s/search?jql=project=%s&expand=changelog&startAt=%d&maxResults=%d",
					cfg.JiraUrl, cfg.JiraApi, projectKey, startAt, cfg.IssueInOneRequest))
				ch <- data
				wg.Done()
			}
		}()
	}

	for i := 1; i <= pageCount; i++ {
		startAt := i * cfg.IssueInOneRequest
		requests <- startAt
	}
	close(requests)
	wg.Wait()
	close(ch)
	buffer := make([]byte, 0)
	var issues []entities.Issue
	for d := range ch {
		buffer = append(buffer, d...)
	}
	_ = yaml.Unmarshal(buffer, &issues)
	projectData.Issues = append(projectData.Issues, issues...)
	return projectData, nil
}

func (jiraConnector *JiraConnector) doRequest(link string) []byte {
	delay := jiraConnector.Config.Delay
	response, err := jiraConnector.HttpClient.Get(link)
	defer response.Body.Close()
	for {
		if response.StatusCode != 200 || err != nil {
			time.Sleep(time.Duration(delay))
			delay = min(delay*2, jiraConnector.Config.MaxDelay)
			response, err = jiraConnector.HttpClient.Get(link)
		} else {
			break
		}
	}
	data, _ := io.ReadAll(response.Body)
	return data
}
