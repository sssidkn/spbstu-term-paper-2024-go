package entities

type Issue struct {
	Id     string `json:"key,omitempty"`
	Fields Fields `json:"fields,omitempty"`
}

type Project struct {
	Name        string  `json:"name"`
	Id          string  `json:"id"`
	Key         string  `json:"key,omitempty"`
	IssuesCount int     `json:"total"`
	Issues      []Issue `json:"issues"`
}

type Fields struct {
	Creator     Creator   `json:"creator"`
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
