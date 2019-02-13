package awsapigateway

import (
	"github.com/awslabs/aws-cdk-go/cdk"
	"github.com/awslabs/aws-cdk-go/jsii"
)

type RestApiPropsIface interface {
	ResourceOptionsIface

	GetApiKeySourceType() ApiKeySourceTypeIface
}

// RestApiProps are properties.
type RestApiProps struct {
	ResourceOptions

	ApiKeySourceType ApiKeySourceTypeIface
}

func (r RestApiProps) GetApiKeySourceType() ApiKeySourceTypeIface { return r.ApiKeySourceType }

// RestApiIface provides the subtyping interfaces for JSII RestApi class.
type RestApiIface interface {
	cdk.ConstructIface

	restApiPrivate()
}

// RestApi_Overrides provides the interface for overriding methods of a
// generated JSII class within the JSII kernel.
type RestApi_Overrides interface {
}

// RestApi is a JSII class.
type RestApi struct {
	*cdk.Construct

	base jsii.Base
}

func (RestApi) restApiPrivate() {}

// NewRestApi returns an initialized RestApi.
func NewRestApi(scope cdk.ConstructIface, id string, props RestApiPropsIface) *RestApi {
	return ExtendRestApi(nil, scope, id, props)
}

// InternalNewRestApiAsBaseClass returns an initialized RestApi initialized as
// a base type extended by another type. Used internally by generated JSII
// types.
func InternalNewRestApiAsBaseClass(jsiiID string) *RestApi {
	return &RestApi{
		base:      jsii.Base{ID: jsiiID},
		Construct: cdk.InternalNewConstructAsBaseClass(jsiiID),
	}
}

// NewRestApi_WithOverrides returns an initialized RestApi initialized with
// overrides. Use when creating custom types extending JSII generated types.
func ExtendRestApi(overrides RestApiIface, scope cdk.ConstructIface, id string, props RestApiPropsIface) *RestApi {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$awsapigateway$0.0.0.RestApi",
		[]interface{}{},
		overrides,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &RestApi{
		base:      jsii.Base{ID: jsiiID},
		Construct: cdk.InternalNewConstructAsBaseClass(jsiiID),
	}
}
