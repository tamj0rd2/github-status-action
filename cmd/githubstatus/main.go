package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/tamj0rd2/github-status-action/cmd/githubstatus/internal/github"
	"github.com/tamj0rd2/github-status-action/cmd/githubstatus/internal/helpers"
	"github.com/tamj0rd2/github-status-action/cmd/githubstatus/internal/slack"
)

func main() {
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Fatal("GITHUB_TOKEN not set")
	}

	sha := os.Getenv("GITHUB_SHA")
	if sha == "" {
		log.Fatal("GITHUB_SHA not set")
	}

	slackURL := os.Getenv("SLACK_WEBHOOK")
	if slackURL == "" {
		log.Fatal("SLACK_WEBHOOK not set")
	}

	githubService := github.NewService(githubToken)
	slackService := slack.NewService(slackURL)

	if err := githubService.WaitForChecksToSucceed(sha, time.Minute*20); err != nil {
		helpers.PrintErr(err)

		var errStepFailed github.ErrStepFailed
		if errors.As(err, &errStepFailed) {
			if err := slackService.ReportPipelineFailure(errStepFailed); err != nil {
				log.Fatal(err)
			}
			return
		}

		os.Exit(1)
	}
}
