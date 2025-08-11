package commands

import (
	"fmt"
	"goSheild/utils"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/briandowns/spinner"
)

var (
	startPort = 1
	endPort   = 1024
	timeout   = 2 * time.Second
)

func scanPort(host string, port int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := net.JoinHostPort(host, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return // closed port
	}
	conn.Close()
	results <- port // open port
}

func portScan() {
	fmt.Println("$ Type your host")
	host := utils.AskChoice()

	if _, err := net.LookupHost(host); err != nil {
		fmt.Printf("❌ Invalid host: %s\n", host)
		return
	}

	fmt.Printf("🔍 Scanning ports on %s from %d to %d...\n", host, startPort, endPort)

	// Spinner for UX
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()

	results := make(chan int, endPort-startPort+1)
	var wg sync.WaitGroup

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go scanPort(host, port, results, &wg)
	}

	// Close channel when done
	go func() {
		wg.Wait()
		close(results)
	}()

	s.Stop()
	fmt.Println("✅ Scan Results:")

	for port := range results {
		fmt.Printf("Port %d is open\n", port)
	}
}
