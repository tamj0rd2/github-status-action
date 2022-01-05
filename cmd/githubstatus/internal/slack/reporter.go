package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tamj0rd2/github-status-action/cmd/githubstatus/internal/github"
)

type Service struct {
	webhookURL string
}

func NewService(webhookURL string) *Service {
	return &Service{
		webhookURL: webhookURL,
	}
}

func (s *Service) ReportPipelineFailure(stepFailure github.ErrStepFailed) error {
	dto := newDTO(stepFailure)
	requestJSON, err := json.Marshal(&dto)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, s.webhookURL, bytes.NewReader(requestJSON))
	if err != nil {
		return err
	}
	req.Header.Add("Content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		return nil
	default:
		body, _ := io.ReadAll(res.Body)
		fmt.Println(string(body))

		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
}
