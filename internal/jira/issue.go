package jira

import "fmt"

type Issue struct {
	Id     string      `json:"id"`
	Key    string      `json:"key"`
	Fields IssueFields `json:"fields"`
}

type IssueFields struct {
	Summary     string            `json:"summary,omitempty"`
	Description string            `json:"description,omitempty"`
	Project     IssueFieldProject `json:"project,omitempty"`
	Status      IssueFieldStatus  `json:"status,omitempty"`
}

type IssueFieldProject struct {
	Id   string `json:"id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

type IssueFieldStatus struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (s *IssueFieldStatus) String() string {
	return s.Name
}

func (i *Issue) String() string {
	return fmt.Sprintf("%s: [%s] %s", i.Key, i.Fields.Status.String(), i.Fields.Summary)
}
