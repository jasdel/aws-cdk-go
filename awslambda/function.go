package awslambda

import (
	"github.com/awslabs/aws-cdk-go/cdk"
	"github.com/awslabs/aws-cdk-go/jsii"
)

type IFunction interface{}

type FunctionProps interface {
	Runtime() string
	Code() AssetCode
	Handler() string
	Environment() map[string]jsii.Any
}

type FunctionProps_ struct {
	Runtime_     string
	Code_        AssetCode
	Handler_     string
	Environment_ map[string]jsii.Any
}

func (f FunctionProps_) Runtime() string                  { return f.Runtime_ }
func (f FunctionProps_) Code() AssetCode                  { return f.Code_ }
func (f FunctionProps_) Handler() string                  { return f.Handler_ }
func (f FunctionProps_) Environment() map[string]jsii.Any { return f.Environment_ }

type Function_Overrides interface{}

type FunctionBase interface {
	cdk.Construct
	FunctionName() string
}
type FunctionBase_ struct {
	cdk.Construct
	base jsii.Base
}

func (FunctionBase_) functionBase_private() {}

type Function interface {
	FunctionBase

	function_private()
}
type Function_ struct {
	FunctionBase

	base jsii.Base
}

func (Function_) function_private() {}

func NewFunction(scope cdk.Construct, id string, props FunctionProps) *Function_ {
	return NewFunction_WithOverrides(scope, id, props, nil)
}
func NewFunction_AsBaseClass(jsiiID string) *Function_ {
	return &Function_{
		base: jsii.Base{ID: jsiiID},
	}
}
func NewFunction_WithOverrides(scope cdk.Construct, id string, props FunctionProps, overrides Function_Overrides) *Function_ {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$awsapigateway$0.0.0.Function",
		[]jsii.Any{},
		overrides,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &Function_{
		base: jsii.Base{ID: jsiiID},
	}
}
