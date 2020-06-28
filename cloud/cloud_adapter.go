package cloud

type Adapter interface {
	IsDomainExists(domainName string) bool
	CreateDomain(domainName string)
	IsFloatingIPExists() bool
	CreateFloatingIP()
}
