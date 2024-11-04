package entities

type Data struct {
	IssuesCount int     `json:"total"`
	Issues      []Issue `json:"issues"`
}
type Issue struct {
	Key    string `json:"key,omitempty"`
	Fields Fields `json:"fields,omitempty"`
}

type Project struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Key  string `json:"key,omitempty"`
}

type Fields struct {
	Project     Project   `json:"project,omitempty"`
	Creator     Creator   `json:"creator,omitempty"`
	Assignee    Creator   `json:"reporter,omitempty"`
	Summary     string    `json:"summary,omitempty"`
	Description string    `json:"description,omitempty"`
	Type        IssueType `json:"issuetype,omitempty"`
	Priority    string    `json:"priority,omitempty"`
	Status      Status    `json:"status,omitempty"`
	CreatedTime string    `json:"created"`
	UpdatedTime string    `json:"updated"`
	ClosedTime  string    `json:"resolutiondate"`
	TimeSpent   string    `json:"timespent"`
}

type IssueType struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	Description string `json:"description"`
}

type Creator struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type Status struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Id          string `json:"id"`
}
