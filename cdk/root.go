package cdk

import (
	"github.com/awslabs/aws-cdk-go/jsii"
)

// Root provides the subtyping interfaces for JSII Root_ class.
type Root interface {
	Construct
}

// Root_Overrides provides the interface for overriding methods of a generated
// JSII class within the JSII kernel.
type Root_Overrides interface {
}

// Root_ is a JSII class.
type Root_ struct {
	Construct

	base jsii.Base
}

// NewRoot returns an initialized Root_.
func NewRoot() *Root_ {
	return NewRoot_WithOverrides(nil)
}

// NewRoot_AsBaseClass returns an initialized Root_ initialized as a base type
// extended by another type. Used internally by generated JSII types.
func NewRoot_AsBaseClass(jsiiID string) *Root_ {
	return &Root_{
		base:      jsii.Base{ID: jsiiID},
		Construct: NewConstruct_AsBaseClass(jsiiID),
	}
}

// NewRoot_WithOverrides returns an initialized Root_ initialized with
// overrides. Use when creating custom types extending JSII generated types.
func NewRoot_WithOverrides(overrides Root_Overrides) *Root_ {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$cdk$0.0.0.Root_",
		[]jsii.Any{},
		nil,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &Root_{
		base:      jsii.Base{ID: jsiiID},
		Construct: NewConstruct_AsBaseClass(jsiiID),
	}
}
