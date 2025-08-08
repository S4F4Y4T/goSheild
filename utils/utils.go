package utils

import (
	"strings"
	"bufio"
    "io"
	"os"
)

// unexported helper for testability
func askChoiceFromReader(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return strings.ToLower(strings.TrimSpace(scanner.Text()))
}

// exported function for normal use
func AskChoice() string {
	return askChoiceFromReader(os.Stdin)
}