package jira

import "fmt"

type Repo struct {
	connector *Connector
}

func NewRepo(connector *Connector) *Repo {
	return &Repo{
		connector: connector,
	}
}

func (r *Repo) SearchByJQL(jql *Jql) (*[]Issue, error) {
	response, err := r.connector.SearchJQLRequest(SearchByJQLBody{Jql: jql.String()})
	if err != nil {
		fmt.Printf("Error - %v", err)
		return nil, err
	}

	return &response.Issues, nil
}
