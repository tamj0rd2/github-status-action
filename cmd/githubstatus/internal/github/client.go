package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type client struct {
	client *http.Client
	token  string
}

func NewClient(token string) *client {
	return &client{
		client: http.DefaultClient,
		token:  token,
	}
}

func (c *client) GetCommitStatuses(sha string) (CommitStatusesResponse, error) {
	var statusesRes CommitStatusesResponse

	url := fmt.Sprintf("https://api.github.com/repos/tamj0rd2/github-status-action/statuses/%s", sha)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", c.token))
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return statusesRes, fmt.Errorf("error getting statuses: %s", err)
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		err = json.NewDecoder(res.Body).Decode(&statusesRes)
		if err != nil {
			return statusesRes, err
		}
		return statusesRes, nil
	default:
		body, _ := io.ReadAll(res.Body)
		fmt.Println("res:", string(body))

		return statusesRes, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
}
