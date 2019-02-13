package cdk

import (
	"github.com/awslabs/aws-cdk-go/jsii"
)

// RootIface provides the subtyping interfaces for JSII Root class.
type RootIface interface {
	ConstructIface

	rootPrivate()
}

// Root is a JSII class.
type Root struct {
	*Construct

	base jsii.Base
}

func (Root) rootPrivate() {}

// NewRoot returns an initialized Root.
func NewRoot() *Root {
	return ExtendRoot(nil)
}

// InternalNewRootAsBaseClass returns an initialized Root initialized as a base
// type extended by another type. Used internally by generated JSII types.
func InternalNewRootAsBaseClass(jsiiID string) *Root {
	return &Root{
		base:      jsii.Base{ID: jsiiID},
		Construct: InternalNewConstructAsBaseClass(jsiiID),
	}
}

// ExtendRoot returns an initialized Root initialized with overrides. Use when
// creating custom types extending JSII generated types.
func ExtendRoot(overrides RootIface) *Root {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$cdk$0.0.0.Root",
		[]interface{}{},
		nil,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &Root{
		base:      jsii.Base{ID: jsiiID},
		Construct: InternalNewConstructAsBaseClass(jsiiID),
	}
}
