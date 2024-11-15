package transformer

import (
	"time"

	"github.com/jiraconnector/internal/dto"
	"github.com/jiraconnector/internal/entities"
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
	return dto.Issue{
		//ProjectId: ,
		//AuthorId:,
		//AssigneeId: ,
		Key:         issue.Key,
		CreatedTime: createdTime,
		UpdatedTime: updatedTime,
		ClosedTime:  closedTime,
		Summary:     issue.Fields.Summary,
		Description: issue.Fields.Description,
		Priority:    issue.Fields.Priority.Name,
		Status:      issue.Fields.Status.Name,
		Type:        issue.Fields.Type.Name,
	}
}
