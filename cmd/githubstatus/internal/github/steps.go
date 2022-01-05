package github

type step string

const (
	StepPipegen step = "pipegen"
	StepCI      step = "ci"
)
