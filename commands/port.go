package commands

import(
	"fmt"
    "goSheild/utils"
	"net"
	"strconv"
	"time"
	"github.com/briandowns/spinner"

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

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	for port := startPort; port <= endPort; port++ {
		address := net.JoinHostPort(host, strconv.Itoa(port))
		conn, err := net.DialTimeout("tcp", address, 300*time.Millisecond)

		if err != nil {
			s.Stop() // stop spinner before printing
			fmt.Printf("Port %d is close\n", port)
			s.Start() // restart spinner
			continue
		}

		conn.Close()
		s.Stop()
		fmt.Printf("Port %d is open\n", port)
		s.Start()

		time.Sleep(100 * time.Millisecond)
	}
}
