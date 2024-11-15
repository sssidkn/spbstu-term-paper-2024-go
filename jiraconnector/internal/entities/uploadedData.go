package entities

type Data struct {
	IssuesCount int     `json:"total"`
	Issues      []Issue `json:"issues"`
}
type Issue struct {
	Key           string        `json:"key,omitempty"`
	Fields        Fields        `json:"fields,omitempty"`
	StatusChanges StatusChanges `json:"changelog,omitempty"`
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
	Priority    Priority  `json:"priority,omitempty"`
	Status      Status    `json:"status,omitempty"`
	CreatedTime string    `json:"created,omitempty"`
	UpdatedTime string    `json:"updated,omitempty"`
	ClosedTime  string    `json:"resolutiondate,omitempty"`
	TimeSpent   string    `json:"timespent,omitempty"`
}

type IssueType struct {
	Name        string `json:"name,omitempty"`
	Id          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

type Creator struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
}

type Status struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Id          string `json:"id,omitempty"`
}

type StatusChanges struct {
	Histories []History `json:"histories,omitempty"`
}

type Priority struct {
	Name string `json:"name,omitempty"`
}

type History struct {
	Author     Creator `json:"author,omitempty"`
	ChangeTime string  `json:"created,omitempty"`
	Items      []Items `json:"items,omitempty"`
}
type Items struct {
	Field      string `json:"field,omitempty"`
	From       string `json:"from"`
	To         string `json:"to"`
	FromString string `json:"fromString"`
	ToString   string `json:"toString"`
}
