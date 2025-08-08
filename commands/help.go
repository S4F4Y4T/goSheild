package commands

import (
	"fmt"
	"strings"
)

func help() {
	fmt.Println("$ List of all available commands:")

	fmt.Printf("%-15s %-20s %s\n", "Command", "Name", "Description")
	fmt.Println(strings.Repeat("-", 60))

	for _, operation := range operations {
		fmt.Printf("%-15s %-20s %s\n", "["+operation.Action+"]", operation.Name, operation.Description)
	}
}