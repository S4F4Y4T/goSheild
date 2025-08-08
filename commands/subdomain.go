package commands

import (
	"fmt"
	"net"
	"goSheild/utils"
)

var commonSubdomains = []string{
	"www", "mail", "ftp", "cpanel", "webmail", "blog", "api", "dev", "test",
}

func subDomainScan() {
	fmt.Println("$ Enter domain (e.g. example.com):")
	domain := utils.AskChoice()

	fmt.Printf("Searching for subdomains of: %s\n", domain)

	for _, sub := range commonSubdomains {
		fullSubdomain := sub + "." + domain
		_, err := net.LookupHost(fullSubdomain)
		if err == nil {
			fmt.Printf("[FOUND] %s\n", fullSubdomain)
		}
	}

}