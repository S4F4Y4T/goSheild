package commands

import (
	"fmt"
	"net"
	"errors"
	"time"
)

func checkInternet() error {
	_, err := net.DialTimeout("tcp", "google.com:80", 2*time.Second)


	if err != nil {
		return errors.New("🚫 No internet connection")
	}

	fmt.Println("✅ Internet is available.")

	return nil
}