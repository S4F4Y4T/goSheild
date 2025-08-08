package commands

import (
	"fmt"
	"errors"
)

type Operation struct {
	Name string
	Description string
	Action string
}

var operations = []Operation{
	{"Help", "List all available commands", "h"},
	{"Internet Checker", "Check internet availability", "ic"},
	{"Port Scanner", "Scan every port of given host", "ps"},
	{"Sub-domain Finder", "Find subdomains of a given domain", "sf"},
	{"Password Generator", "Generate random password", "pgen"},
	{"Kill", "Exit the program", "kill"},
}

func HandleAction(action string) error {

	switch action {
		case "h":
			help()
		case "ic":
			return checkInternet()
		case "ps":
			portScan()
		case "sf":
			subDomainScan()
		case "pgen":
			passwordGenerate()
		default:
			return errors.New("Invalid command")
	}

	return nil
}


func process() {
	fmt.Println("Under Process")
 }