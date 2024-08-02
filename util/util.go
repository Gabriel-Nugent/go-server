package util

import (
	"fmt"
	"os"
)

const (
	colorRED = "\033[0;31m"
	colorGREEN = "\033[0;32m"
	colorNONE = "\033[0m"
)

// Error prints a string to stderr with red text
func Error(message string) {
	fmt.Fprintf(os.Stderr, "%s[ERROR]: %s%s\n", colorRED, message, colorNONE)
}

// Log prints a string to stdout with no color
func Log(message string) {
	fmt.Fprintf(os.Stdout, "%s[LOG]: %s\n", colorNONE, message)
}

// Success prints a string to stdout with green text
func Success(message string)  {
	fmt.Fprintf(os.Stdout, "%s[SUCCESS]: %s%s\n", colorGREEN, message, colorNONE)
}