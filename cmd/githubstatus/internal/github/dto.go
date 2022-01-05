package github

import (
	"fmt"
	"log"
)

type CommitStatusesResponse []status

type status struct {
	State       string `json:"state"`
	Description string `json:"description"`
	TargetURL   string `json:"target_url"`
	Context     string `json:"context"`
}

func (res CommitStatusesResponse) isStepComplete(step step, sha string) (isComplete bool, failure error) {
	for _, status := range res {
		if status.Context != string(step) {
			continue
		}

		switch status.State {
		case "success":
			return true, nil
		case "failure":
			return true, ErrStepFailed{
				Step:        step,
				CommitURL:   fmt.Sprintf("https://github.com/tamj0rd2/github-status-action/commit/%s", sha),
				PipelineURL: status.TargetURL,
			}
		case "pending":
			continue
		default:
			log.Fatalf("unknown state %s for %s:\n", status.State, step)
		}
	}

	return false, nil
}
