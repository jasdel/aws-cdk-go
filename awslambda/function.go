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
}

type FunctionProps_ struct {
	Runtime_ string
	Code_    AssetCode
	Handler_ string
}

func (f FunctionProps_) Runtime() string { return f.Runtime_ }
func (f FunctionProps_) Code() AssetCode { return f.Code_ }
func (f FunctionProps_) Handler() string { return f.Handler_ }

type Function_Overrides interface{}

type Function interface {
	function_private()
}
type Function_ struct {
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
