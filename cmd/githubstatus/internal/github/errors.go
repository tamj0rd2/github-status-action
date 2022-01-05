package github

import (
	"errors"
	"fmt"
)

type ErrStepFailed struct {
	Step        step
	CommitURL   string
	PipelineURL string
}

func (e ErrStepFailed) Error() string {
	return fmt.Sprintf("%s failed", e.Step)
}

var ErrTimedOut = errors.New("timed out")
