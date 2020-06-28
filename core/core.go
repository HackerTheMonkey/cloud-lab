package core

import (
	"fmt"
	"github.com/HackerTheMonkey/codematters/cloud"
)

var url = "https://httpbin.org/get"

func Init(domainName string, cloudAdapter cloud.Adapter) string {

	if !cloudAdapter.IsDomainExists(domainName) {
		cloudAdapter.CreateDomain(domainName)
	}

	if !cloudAdapter.IsFloatingIPExists() {
		cloudAdapter.CreateFloatingIP()
	}

	fmt.Printf("initializing codematters root instance for domain: %s\n", domainName)
	fmt.Println("initialization completed")
	return "OK"
}
