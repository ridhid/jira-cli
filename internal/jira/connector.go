package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Connector struct {
	Url    string
	Token  string
	Client *http.Client
}

type SearchByJQLBody struct {
	Jql string `json:"jql"`
}

type SearchByJQLResponse struct {
	Issues []Issue `json:"issues"`
}

func NewJiraConnector(url string, token string) *Connector {
	return &Connector{
		Url:   url,
		Token: token,
		Client: &http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
	}
}

func (c *Connector) Request(url string, method string, body []byte, output any) error {
	var requestBody *bytes.Buffer = nil

	if body != nil {
		requestBody = bytes.NewBuffer(body)
	}

	requestUrl := c.Url + url

	req, err := http.NewRequest(method, requestUrl, requestBody)
	if err != nil {
		return err
	}

	token := fmt.Sprintf("Bearer %s", c.Token)
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	rsp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != 200 {
		return fmt.Errorf("Status code is %s", rsp.StatusCode)
	}

	rawResponseBody, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil
	}
	if err := json.Unmarshal(rawResponseBody, output); err != nil {
		return err
	}

	return nil
}

func (c *Connector) SearchJQLRequest(payload SearchByJQLBody) (*SearchByJQLResponse, error) {
	requestBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	response := &SearchByJQLResponse{}
	if err := c.Request("/search", http.MethodPost, requestBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
