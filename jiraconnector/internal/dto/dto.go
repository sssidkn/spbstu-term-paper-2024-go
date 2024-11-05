package dto

import "time"

type Author struct {
	ID   uint
	Name string
}

type StatusChanges struct {
	IssueId    uint
	AuthorId   uint
	ChangeTime time.Time
	FromStatus string
	ToStatus   string
}

type Project struct {
	ID    uint
	Title string
}

type Issue struct {
	ID          uint
	ProjectId   uint
	AuthorId    uint
	AssigneeId  uint
	Key         string
	CreatedTime time.Time
	UpdatedTime time.Time
	ClosedTime  time.Time
	Summary     string
	Description string
	Priority    string
	Status      string
	Type        string
}
