package awsapigateway

import (
	"github.com/awslabs/aws-cdk-go/awslambda"
	"github.com/awslabs/aws-cdk-go/cdk"
	"github.com/awslabs/aws-cdk-go/jsii"
)

type LambdaRestApiProps interface {
	Handler() awslambda.IFunction
}

type LambdaRestApiProps_ struct {
	Handler_ awslambda.IFunction
}

func (l LambdaRestApiProps_) Handler() awslambda.IFunction { return l.Handler_ }

type LambdaRestApi interface {
	RestApi

	lambdaRestApi_private()
}

type LambdaRestApi_Overrides interface {
}

type LambdaRestApi_ struct {
	RestApi

	base jsii.Base
}

func (*LambdaRestApi_) lambdaRestApi_private() {}

func NewLambdaRestApi(scope cdk.Construct, id string, props LambdaRestApiProps) *LambdaRestApi_ {
	return NewLambdaRestApi_WithOverrides(scope, id, props, nil)
}
func NewLambdaRestApi_AsBaseClass(jsiiID string) *LambdaRestApi_ {
	return &LambdaRestApi_{
		base:    jsii.Base{ID: jsiiID},
		RestApi: NewRestApi_AsBaseClass(jsiiID),
	}
}
func NewLambdaRestApi_WithOverrides(scope cdk.Construct, id string, props LambdaRestApiProps, overrides LambdaRestApi_Overrides) *LambdaRestApi_ {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$awsapigateway$0.0.0.LambdaRestApi",
		[]jsii.Any{},
		overrides,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &LambdaRestApi_{
		base:    jsii.Base{ID: jsiiID},
		RestApi: NewRestApi_AsBaseClass(jsiiID),
	}
}
