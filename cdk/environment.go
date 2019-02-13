package cdk

// EnvironmentIface provides the interface for Environment_ datatype.
type EnvironmentIface interface {
	GetAccount() *string
	GetRegion() *string
}

// Environment are properties.
type Environment struct {
	Account *string
	Region  *string
}

func (d *Environment) GetAccount() *string { return d.Account }
func (d *Environment) GetRegion() *string  { return d.Region }
