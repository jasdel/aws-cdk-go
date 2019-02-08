package cdk

import "github.com/awslabs/aws-cdk-go/jsii"

// Construct provides the subtyping interfaces for JSII Construct_ class.
type Construct interface {
	// members
	DependencyRoots() []IConstruct
	Node() IConstruct

	// methods
	IsConstruct() bool
	Prepare()         // protected method, must be exposed go has no protected
	ToString() string // override to Go's "String()" ?
	Validate() []string

	construct_private()
}

// Construct_Overrides provides the interface for overriding methods of a
// generated JSII class within the JSII kernel.
type Construct_Overrides interface{}

// Construct_ is a JSII class.
type Construct_ struct {
	base jsii.Base
}

func (*Construct_) construct_private() {}

// NewConstruct returns an initialized Construct_.
func NewConstruct(scope Construct, id string) *Construct_ {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$cdk$0.0.0.Construct_",
		[]jsii.Any{
			scope, id,
		},
		nil,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &Construct_{
		base: jsii.Base{ID: jsiiID},
	}
}

// NewConstruct_AsBaseClass returns an initialized Construct_ initialized as a
// base type extended by another type. Used internally by generated JSII types.
func NewConstruct_AsBaseClass(jsiiID string) *Construct_ {
	return &Construct_{
		base: jsii.Base{ID: jsiiID},
	}
}

// NewConstruct_WithOverrides returns an initialized Construct_ initialized with
// overrides. Use when creating custom types extending JSII generated types.
func NewConstruct_WithOverrides(scope Construct, id string, overrides Construct_Overrides) *Construct_ {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$cdk$0.0.0.Construct_",
		[]jsii.Any{
			scope, id,
		},
		overrides,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &Construct_{
		base: jsii.Base{ID: jsiiID},
	}
}

func (c *Construct_) DependencyRoots() []IConstruct { return nil }
func (c *Construct_) Node() IConstruct              { return nil }
func (c *Construct_) IsConstruct() bool             { return true }

// protected method, must be exposed go has no protected
func (c *Construct_) Prepare() {}

// override to Go's "String()" ?
func (c *Construct_) ToString() string   { return "" }
func (c *Construct_) Validate() []string { return nil }
