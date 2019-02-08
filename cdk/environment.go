package cdk

// Environment provides the interface for Environment_ datatype.
type Environment interface {
	Account() *string
	Region() *string
}

// Environment_ are properties.
type Environment_ struct {
	Account_ *string
	Region_  *string
}

func (d *Environment_) Account() *string { return d.Account_ }
func (d *Environment_) Region() *string  { return d.Region_ }
