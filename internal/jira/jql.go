package jira

import (
	"fmt"
	"strings"
)

type Jql struct {
	jql string
}

func NewJql() *Jql {
	return &Jql{
		jql: "",
	}
}

func (j *Jql) String() string {
	return j.jql
}

func (j *Jql) Join(query string, operator string) *Jql {
	if j.jql == "" {
		j.jql = query
	} else {
		j.jql = strings.Join([]string{j.jql, query}, operator)
	}
	return j
}

func (j *Jql) And(query string) *Jql {
	return j.Join(query, " and ")
}

func (j *Jql) Or(query string) *Jql {
	return j.Join(query, " or ")
}

func (j *Jql) ByProjects(projects []string) *Jql {
	query := fmt.Sprintf("PROJECT in (%s)", strings.Join(projects, ", "))
	return j.And(query)
}

func (j *Jql) ByStatuses(statuses []string) *Jql {
	query := fmt.Sprintf("STATUS in (%s)", strings.Join(statuses, ", "))
	return j.And(query)
}

func (j *Jql) ByAssignee(users []string) *Jql {
	query := fmt.Sprintf("ASSIGNEE in (%s)", strings.Join(users, ", "))
	return j.And(query)
}

func (j *Jql) BySprints(sprints []string) *Jql {
	query := fmt.Sprintf("SPRINT in (%s)", strings.Join(sprints, ", "))
	return j.And(query)
}

func (j *Jql) ByReporter(users []string) *Jql {
	query := fmt.Sprintf("REPORTER in (%s)", strings.Join(users, ", "))
	return j.And(query)
}

func (j *Jql) OrderBy(columns []string) *Jql {
	query := fmt.Sprintf("ORDER BY %s", strings.Join(columns, ", "))
	return j.Join(query, " ")
}
