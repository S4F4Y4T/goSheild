package main

import (
	"fmt"
	"goSheild/commands"
	"goSheild/ui"
	"goSheild/utils"
)

func main() {
	ui.PrintHeader()

	for{
		fmt.Println("$ Select an action: ")
		choice := utils.AskChoice()

		if choice == "kill" {
			fmt.Println("See ya later ")
			break
		}

		err := commands.HandleAction(choice)
		if err != nil {
			fmt.Println(err)
		}
	}
	
}



