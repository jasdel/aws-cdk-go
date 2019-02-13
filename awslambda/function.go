package awslambda

import (
	"github.com/awslabs/aws-cdk-go/cdk"
	"github.com/awslabs/aws-cdk-go/jsii"
)

type IFunction interface{}

type FunctionPropsIface interface {
	GetRuntime() string
	GetCode() AssetCodeIface
	GetHandler() string
	GetEnvironment() map[string]interface{}
}

type FunctionProps struct {
	Runtime     string
	Code        AssetCodeIface
	Handler     string
	Environment map[string]interface{}
}

func (f FunctionProps) GetRuntime() string                     { return f.Runtime }
func (f FunctionProps) GetCode() AssetCodeIface                { return f.Code }
func (f FunctionProps) GetHandler() string                     { return f.Handler }
func (f FunctionProps) GetEnvironment() map[string]interface{} { return f.Environment }

type FunctionBaseIface interface {
	cdk.ConstructIface

	FunctionName() string

	functionBasePrivate()
}
type FunctionBase struct {
	*cdk.Construct

	base jsii.Base
}

func (FunctionBase) functionBasePrivate() {}

func (*FunctionBase) FunctionName() string { return "" }

type FunctionIface interface {
	FunctionBaseIface

	functionPrivate()
}

type Function struct {
	*FunctionBase

	base jsii.Base
}

func (Function) functionPrivate() {}

func NewFunction(scope cdk.ConstructIface, id string, props FunctionPropsIface) *Function {
	return ExtendFunction(nil, scope, id, props)
}

func InternalNewFunctionAsBaseClass(jsiiID string) *Function {
	return &Function{
		base: jsii.Base{ID: jsiiID},
	}
}
func ExtendFunction(overrides FunctionIface, scope cdk.ConstructIface, id string, props FunctionPropsIface) *Function {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$awsapigateway$0.0.0.Function",
		[]interface{}{},
		overrides,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &Function{
		base: jsii.Base{ID: jsiiID},
	}
}
