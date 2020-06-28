package core_test

import (
	"github.com/HackerTheMonkey/codematters/core"
	"testing"
)

var isDomainExists func(domainName string) bool
var createDomain func(domainName string)
var isFloatingIPExists func() bool
var createFloatingIP func()

func Test_given_domainExists_then_noNewDomainIsCreated(t *testing.T) {
	// Given
	domainNameExists()
	createDomainShouldNotBeInteractedWith(t)
	floatingIPExists()

	// When
	core.Init("example.com", mockCloudAdapter{})
}

func Test_given_domainDoesNotExist_then_newDomainIsCreated(t *testing.T) {
	// Given
	isDomainCreated := false
	actualDomainName := ""
	var expectedDomainName string = "example.com"

	domainNameDoesNotExists()
	createDomain = func(domainName string) {
		isDomainCreated = true
		actualDomainName = domainName
	}

	// When
	core.Init(expectedDomainName, mockCloudAdapter{})

	// Then
	if !isDomainCreated {
		t.Errorf("createDomain() was not interacted with as we expect!\n")
	}

	if actualDomainName != expectedDomainName {
		t.Errorf("expecting: 'example.com' got: '%s' \n", actualDomainName)
	}
}

func Test_given_floatingIPDoesNotExist_then_newFloatingIPCreated(t *testing.T) {
	// Given
	domainNameExists()
	floatingIPExists()

	// When
	core.Init("example.com", mockCloudAdapter{})

	// Then
	createFloatingIPShouldNotBeInteractedWith(t)
}

func Test_given_floatingIPExists_then_noFloatingIPCreated(t *testing.T) {
	// Given
	isFloatingIPCreated := false
	domainNameExists()
	floatingIPDoesNotExist()
	createFloatingIP = func() {
		isFloatingIPCreated = true
	}

	// When
	core.Init("example.com", mockCloudAdapter{})

	// Then
	if !isFloatingIPCreated {
		t.Errorf("createFloatingIP() was not interacted with as we expect!\n")
	}
}

func floatingIPDoesNotExist() {
	isFloatingIPExists = func() bool {
		return false
	}
}

func createDomainShouldNotBeInteractedWith(t *testing.T) {
	createDomain = func(domainName string) {
		t.Errorf("createDomain() should not be interacted with! But it did!")
	}
}

func createFloatingIPShouldNotBeInteractedWith(t *testing.T) {
	createDomain = func(domainName string) {
		t.Errorf("createFloatingIP() should not be interacted with! But it did!")
	}
}

func domainNameExists() {
	isDomainExists = func(domainName string) bool {
		return true
	}
}

func floatingIPExists() {
	isFloatingIPExists = func() bool {
		return true
	}
}

func domainNameDoesNotExists() {
	isDomainExists = func(domainName string) bool {
		return false
	}
}

type mockCloudAdapter struct {
}

func (m mockCloudAdapter) CreateFloatingIP() {
	createFloatingIP()
}

func (m mockCloudAdapter) IsFloatingIPExists() bool {
	return isFloatingIPExists()
}

func (m mockCloudAdapter) IsDomainExists(domainName string) bool {
	return isDomainExists(domainName)
}

func (m mockCloudAdapter) CreateDomain(domainName string) {
	createDomain(domainName)
}
