package transformer

import (
	"github.com/jiraconnector/internal/dto"
	"github.com/jiraconnector/internal/entities"
	"strconv"
	"time"
)

func AuthorToDTO(creator *entities.Creator) dto.Author {
	return dto.Author{
		Name: creator.Name,
	}
}

func ProjectToDTO(project *entities.Project) dto.Project {
	return dto.Project{
		Title: project.Name,
	}
}

func IssueToDTO(issue *entities.Issue) dto.Issue {
	createdTime, _ := time.Parse("2006-01-02T15:04:05.999-0700", issue.Fields.CreatedTime)
	updatedTime, _ := time.Parse("2006-01-02T15:04:05.999-0700", issue.Fields.UpdatedTime)
	closedTime, _ := time.Parse("2006-01-02T15:04:05.999-0700", issue.Fields.ClosedTime)
	timeSpent, _ := strconv.Atoi(issue.Fields.TimeSpent)
	return dto.Issue{
		Key:         issue.Key,
		CreatedTime: createdTime,
		UpdatedTime: updatedTime,
		ClosedTime:  closedTime,
		TimeSpent:   timeSpent,
		Summary:     issue.Fields.Summary,
		Description: issue.Fields.Description,
		Priority:    issue.Fields.Priority,
		Status:      issue.Fields.Status.Name,
		Type:        issue.Fields.Type.Name,
	}
}
