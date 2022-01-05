package slack

import (
	"fmt"

	"github.com/tamj0rd2/github-status-action/cmd/githubstatus/internal/github"
)

func newDTO(err github.ErrStepFailed) slackDto {
	return slackDto{
		Blocks: []block{
			{
				Type: "section",
				Text: &text{
					Type:  "mrkdwn",
					Text:  fmt.Sprintf("Pipeline failed on *%s*", err.Step),
					Emoji: nil,
				},
			},
			{
				Type: "actions",
				Elements: []action{
					{
						Type:  "button",
						Text:  text{Type: "plain_text", Text: "See commit"},
						Value: "dunno",
						URL:   err.CommitURL,
					},
					{
						Type:  "button",
						Text:  text{Type: "plain_text", Text: "See pipeline"},
						Value: "dunno",
						URL:   err.PipelineURL,
					},
				},
			},
		},
	}
}

type slackDto struct {
	Blocks []block `json:"blocks"`
}

type block struct {
	Type     string   `json:"type"`
	Text     *text    `json:"text,omitempty"`
	Elements []action `json:"elements,omitempty"`
}

type action struct {
	Type  string `json:"type"`
	Text  text   `json:"text"`
	Value string `json:"value"`
	URL   string `json:"url"`
}

type text struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji *bool  `json:"emoji,omitempty"`
}
