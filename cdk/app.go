package cdk

import (
	"github.com/awslabs/aws-cdk-go/jsii"
)

// AppIface provides the subtyping interfaces for JSII App class.
type AppIface interface {
	RootIface

	Run()

	appPrivate()
}

// App is a JSII class
type App struct {
	*Root

	base jsii.Base
}

func (*App) appPrivate() {}

// NewApp returns an initialized App.
func NewApp() *App {
	return ExtendApp(nil)
}

// InternalNewAppAsBaseClass returns an initialized App initialized as a base
// type extended by another type. Used internally by generated JSII types.
func InternalNewAppAsBaseClass(jsiiID string) *App {
	return &App{
		base: jsii.Base{ID: jsiiID},
		Root: InternalNewRootAsBaseClass(jsiiID),
	}
}

// ExtendApp returns an initialized App initialized with
// overrides. Use when creating custom types extending JSII generated types.
func ExtendApp(overrides AppIface) *App {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$cdk$0.0.0.App",
		[]interface{}{},
		nil,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &App{
		base: jsii.Base{ID: jsiiID},
		Root: InternalNewRootAsBaseClass(jsiiID),
	}
}

func (a *App) Run() {}
