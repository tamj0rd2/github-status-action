package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetStagedFilesMatchingGlob(glob string) []string {
	output := RunCommand(fmt.Sprintf("list staged changes matching glob %s", glob), exec.Command("git", "diff", "--name-only", "--cached", "--diff-filter=d", "--", glob))
	if output == "" {
		return []string{}
	}
	return strings.Split(output, "\n")
}

func RunCommand(description string, cmd *exec.Cmd) string {
	// debug info
	PrintlnWithColour(description, ColourBlue)
	fmt.Println(cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		PrintlnWithColour(fmt.Sprintf("%s - Command failed - %s", description, err.Error()), ColourRed)
		fmt.Println("Command:", cmd)
		fmt.Println("Output:", string(output))
		PrintlnWithColour("Something went wrong. Do a git status and make sure all your changes are there. If not, have a look at the latest git stash", ColourRed)
		PrintlnWithColour("🏳  Skipping everything else until I have time to figure this out!", ColourYellow)
		os.Exit(0)
	}

	return strings.TrimSpace(string(output))
}

func PrintlnWithColour(text string, colour Colour) {
	fmt.Printf("%s%s%s\n", colour, text, ColourReset)
}

func PrintErr(err error) {
	PrintlnWithColour(err.Error(), ColourRed)
}

type Colour string

const (
	ColourRed    Colour = "\033[31m"
	ColourGreen  Colour = "\033[32m"
	ColourYellow Colour = "\033[33m"
	ColourBlue   Colour = "\033[34m"
	// ColourPurple Colour = "\033[35m"
	// ColourCyan   Colour = "\033[36m"
	// ColourWhite  Colour = "\033[37m".
	ColourReset Colour = "\033[0m"
)
