package cloud

type DigitalOceanAdapter struct {
}

func (adapter DigitalOceanAdapter) CreateFloatingIP() {

}

func (adapter DigitalOceanAdapter) IsFloatingIPExists() bool {
	return true
}

func (adapter DigitalOceanAdapter) CreateDomain(domainName string) {

}

func (adapter DigitalOceanAdapter) IsDomainExists(domainName string) bool {
	return true
}
