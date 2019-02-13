package awsapigateway

import (
	"github.com/awslabs/aws-cdk-go/awslambda"
	"github.com/awslabs/aws-cdk-go/cdk"
	"github.com/awslabs/aws-cdk-go/jsii"
)

type LambdaRestApiPropsIface interface {
	GetHandler() awslambda.IFunction
}

type LambdaRestApiProps struct {
	Handler awslambda.IFunction
}

func (l LambdaRestApiProps) GetHandler() awslambda.IFunction { return l.Handler }

type LambdaRestApiIface interface {
	RestApiIface

	lambdaRestApiPrivate()
}

type LambdaRestApi struct {
	*RestApi

	base jsii.Base
}

func (*LambdaRestApi) lambdaRestApi_private() {}

func NewLambdaRestApi(scope cdk.ConstructIface, id string, props LambdaRestApiPropsIface) *LambdaRestApi {
	return ExtendLambdaRestApi(nil, scope, id, props)
}
func NewLambdaRestApi_AsBaseClass(jsiiID string) *LambdaRestApi {
	return &LambdaRestApi{
		base:    jsii.Base{ID: jsiiID},
		RestApi: InternalNewRestApiAsBaseClass(jsiiID),
	}
}
func ExtendLambdaRestApi(overrides LambdaRestApiIface, scope cdk.ConstructIface, id string, props LambdaRestApiPropsIface) *LambdaRestApi {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$awsapigateway$0.0.0.LambdaRestApi",
		[]interface{}{},
		overrides,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &LambdaRestApi{
		base:    jsii.Base{ID: jsiiID},
		RestApi: InternalNewRestApiAsBaseClass(jsiiID),
	}
}
