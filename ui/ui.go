package ui

import (
	"fmt"
)

func PrintHeader() {

	fmt.Println(`
		________        _________.__           .__.__       .___
		/  _____/  ____ /   _____/|  |__   ____ |__|  |    __| _/
		/   \  ___ /  _ \\_____  \ |  |  \_/ __ \|  |  |   / __ | 
		\    \_\  (  <_> )        \|   Y  \  ___/|  |  |__/ /_/ | 
		\______  /\____/_______  /|___|  /\___  >__|____/\____ | 
				\/              \/      \/     \/              \/ 
	`)

	fmt.Println("GoShield v0.1.0")
	fmt.Println("Author: Safayat Mahmud")
	fmt.Println("---------------------------")
	fmt.Println("use `h` to list available commands")
	fmt.Println()
}