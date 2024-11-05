package dto

import "time"

type Author struct {
	ID   int
	Name string
}

type StatusChanges struct {
	IssueId    int
	AuthorId   int
	ChangeTime time.Time
	FromStatus string
	ToStatus   string
}

type Project struct {
	ID    int
	Title string
}

type Issue struct {
	ID          int
	ProjectId   int
	AuthorId    int
	AssigneeId  int
	Key         string
	CreatedTime time.Time
	UpdatedTime time.Time
	ClosedTime  time.Time
	TimeSpent   int
	Summary     string
	Description string
	Priority    string
	Status      string
	Type        string
}
