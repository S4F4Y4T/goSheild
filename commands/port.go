package commands

import (
	"fmt"
	"net"
	"time"
	"strconv"
	"sync"

	"github.com/briandowns/spinner"
)

var startPort int = 1
var endPort int = 1000
var counter int = 0

func validateHost(host string) bool {
	_, err := net.DialTimeout("tcp", net.JoinHostPort(host, "80"), time.Second*3)

	if err != nil {
		fmt.Println("$ Invalid host")
		return false
	}
	
	return true
}

func scan(host string, port string, ch chan<- string, wg *sync.WaitGroup) { 

	defer wg.Done()

	_, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Second*2)

	if err == nil {
		ch <- port
	}
	
}

func portScan() {

	fmt.Println("$ Please type your host here")
	var host string
	fmt.Scanln(&host)

	validate := validateHost(host)
	if !validate {
		return
	}

	var wg sync.WaitGroup
	var ch = make(chan string)
	
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Start() 

	for i := startPort; i <= endPort; i++ {
		wg.Add(1)
		go scan(host, strconv.Itoa(i), ch, &wg)
	}

	go func() {
		wg.Wait()   
		close(ch) 
		s.Stop() 
	}()


	for {
		select {
			case p, ok := <-ch:

				if(!ok){
					return
				}

				fmt.Println(p)
		}
	}
	
	fmt.Println("\nScan completed.")

}

