package utils

import (
	"fmt"
	"strings"
)

func AskChoice() string {
	var action string
	fmt.Scanln(&action)
	return strings.ToLower(strings.TrimSpace(action))
}