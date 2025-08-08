package commands

import(
	"fmt"
    "goSheild/utils"
	"net"
	"strconv"
	"time"
)

var startPort int = 1
var endPort int = 1024

func portScan() {
	fmt.Println("$ Type your host")
	host := utils.AskChoice()

	_, err := net.LookupHost(host)
	if err != nil {
		fmt.Printf("❌ Invalid host: %s\n", host)
		return
	}


	fmt.Printf("Scanning ports on %s from %d to %d...\n", host, startPort, endPort)

	for port := startPort; port <= endPort; port++ {
		address := net.JoinHostPort(host, strconv.Itoa(port))
		conn, err := net.DialTimeout("tcp", address, 300*time.Millisecond)
		if err != nil {
			fmt.Printf("Port %d is close\n", port)

			continue
		}
		conn.Close()
		fmt.Printf("Port %d is open\n", port)

		time.Sleep(100 * time.Millisecond)
	}

}