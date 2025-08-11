package main

import (
	"fmt"
	"goSheild/commands"
	"goSheild/ui"
	"goSheild/utils"

	"time"
)

func main() {
	concurrency()
	return

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

func concurrency() {
	fmt.Println("concurrency starts here")

	ch := make(chan string)

	go func () {
		defer close(ch)
		ch <- "a"
		fmt.Println("go routine 1")
	}()

	go func () {
		fmt.Println("go routine 2")
	}()

	go func () {
		fmt.Println("go routine 3")
	}()

	for data := range ch {
		fmt.Println(data)
	}

	time.Sleep(1*time.Second)
}



