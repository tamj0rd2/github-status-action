package helpers

import (
	"fmt"
)

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
