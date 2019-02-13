package cdk

import "github.com/awslabs/aws-cdk-go/jsii"

// ConstructIface provides the subtyping interfaces for JSII Construct class.
type ConstructIface interface {
	// members
	DependencyRoots() []IConstruct
	Node() IConstruct

	// methods
	IsConstruct() bool
	Prepare()         // protected method, must be exposed go has no protected
	ToString() string // override to Go's "String()" ?
	Validate() []string

	constructPrivate()
}

// Construct is a JSII class.
type Construct struct {
	base jsii.Base
}

func (*Construct) constructPrivate() {}

// NewConstruct returns an initialized Construct.
func NewConstruct(scope ConstructIface, id string) *Construct {
	return ExtendConstruct(nil, scope, id)
}

// InternalNewConstructAsBaseClass returns an initialized Construct initialized
// as a base type extended by another type. Used internally by generated JSII
// for Go types.
func InternalNewConstructAsBaseClass(jsiiID string) *Construct {
	return &Construct{
		base: jsii.Base{ID: jsiiID},
	}
}

// ExtendConstruct returns an initialized Construct initialized with overrides.
// Use when creating custom types extending JSII generated types.
func ExtendConstruct(overrides, scope ConstructIface, id string) *Construct {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$cdk$0.0.0.Construct",
		[]interface{}{
			scope, id,
		},
		overrides,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &Construct{
		base: jsii.Base{ID: jsiiID},
	}
}

func (c *Construct) DependencyRoots() []IConstruct { return nil }
func (c *Construct) Node() IConstruct              { return nil }
func (c *Construct) IsConstruct() bool             { return true }
func (c *Construct) Prepare()                      {}
func (c *Construct) ToString() string              { return "" }
func (c *Construct) Validate() []string            { return nil }
