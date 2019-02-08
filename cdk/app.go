package cdk

import (
	"github.com/awslabs/aws-cdk-go/jsii"
)

// App provides the subtyping interfaces for JSII App_[ class.
type App interface {
	Root

	Run()

	app_private()
}

// App_Overrides provides the interface for overriding methods of a generated
// JSII class within the JSII kernel.
type App_Overrides interface {
}

// App_ is a JSII class
type App_ struct {
	Root

	base jsii.Base
}

func (*App_) app_private() {}

// NewApp returns an initialized App_.
func NewApp() *App_ {
	return NewApp_WithOverrides(nil)
}

// NewApp_AsBaseClass returns an initialized App_ initialized as a base type
// extended by another type. Used internally by generated JSII types.
func NewApp_AsBaseClass(jsiiID string) *App_ {
	return &App_{
		base: jsii.Base{ID: jsiiID},
		Root: NewRoot_AsBaseClass(jsiiID),
	}
}

// NewApp_WithOverrides returns an initialized App_ initialized with
// overrides. Use when creating custom types extending JSII generated types.
func NewApp_WithOverrides(overrides App_Overrides) *App_ {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$cdk$0.0.0.App_",
		[]jsii.Any{},
		nil,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &App_{
		base: jsii.Base{ID: jsiiID},
		Root: NewRoot_AsBaseClass(jsiiID),
	}
}

func (a *App_) Run() {}
