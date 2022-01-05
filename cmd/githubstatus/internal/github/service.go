package github

import (
	"fmt"
	"github.com/tamj0rd2/github-status-action/cmd/githubstatus/internal/helpers"
	"os"
	"time"
)

type Service struct {
	client *client
}

func NewService(githubToken string) *Service {
	return &Service{
		client: NewClient(githubToken),
	}
}

func (s *Service) WaitForChecksToSucceed(sha string, totalTimeout time.Duration) error {
	helpers.PrintlnWithColour(fmt.Sprintf("Waiting for statuses for commit %q to finish", sha), helpers.ColourBlue)
	progress := checksProgress{
		StepPipegen: false,
		StepCI:      false,
	}

	timeToWait := time.Second * 10
	attempts := int(totalTimeout / timeToWait)

	for i := 0; i < attempts; i++ {
		statuses, err := s.client.GetCommitStatuses(sha)
		if err != nil {
			return err
		}

		for step, wasComplete := range progress {
			if wasComplete {
				continue
			}

			isComplete, err := statuses.isStepComplete(step, sha)
			if err != nil {
				return err
			}
			progress.Update(step, isComplete)
		}

		if progress.AreAllComplete() {
			helpers.PrintlnWithColour("All checks passed :D", helpers.ColourGreen)
			os.Exit(0)
		}

		<-time.After(timeToWait)
	}

	return ErrTimedOut
}

type checksProgress map[step]bool

func (c checksProgress) Update(step step, isComplete bool) {
	c[step] = isComplete

	if isComplete {
		helpers.PrintlnWithColour(fmt.Sprintf("%s passed", step), helpers.ColourGreen)
	} else {
		helpers.PrintlnWithColour(fmt.Sprintf("%s pending", step), helpers.ColourYellow)
	}
}

func (c checksProgress) AreAllComplete() bool {
	for _, isComplete := range c {
		if !isComplete {
			return false
		}
	}
	return true
}
