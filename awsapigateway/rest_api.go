package awsapigateway

import (
	"github.com/awslabs/aws-cdk-go/cdk"
	"github.com/awslabs/aws-cdk-go/jsii"
)

type RestApiProps interface {
	ResourceOptions

	restApiProps_Private()
}

// RestApiProps_ are properties.
type RestApiProps_ struct {
	ResourceOptions_

	ApiKeySourceType ApiKeySourceType
}

func (RestApiProps_) restApiProps_Private() {}

// RestApi provides the subtyping interfaces for JSII RestApi class.
type RestApi interface {
	cdk.Construct

	restApi_private()
}

// RestApi_Overrides provides the interface for overriding methods of a
// generated JSII class within the JSII kernel.
type RestApi_Overrides interface {
}

// RestApi_ is a JSII class.
type RestApi_ struct {
	cdk.Construct

	base jsii.Base
}

func (RestApi_) restApi_private() {}

// NewRestApi returns an initialized RestApi.
func NewRestApi(scope cdk.Construct, id string, props RestApiProps_) *RestApi_ {
	return NewRestApi_WithOverrides(scope, id, props, nil)
}

// NewRestApi_AsBaseClass returns an initialized RestApi initialized as a base type
// extended by another type. Used internally by generated JSII types.
func NewRestApi_AsBaseClass(jsiiID string) *RestApi_ {
	return &RestApi_{
		base:      jsii.Base{ID: jsiiID},
		Construct: cdk.NewConstruct_AsBaseClass(jsiiID),
	}
}

// NewRestApi_WithOverrides returns an initialized RestApi initialized with
// overrides. Use when creating custom types extending JSII generated types.
func NewRestApi_WithOverrides(scope cdk.Construct, id string, props RestApiProps_, overrides RestApi_Overrides) *RestApi_ {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$awsapigateway$0.0.0.RestApi",
		[]jsii.Any{},
		overrides,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &RestApi_{
		base:      jsii.Base{ID: jsiiID},
		Construct: cdk.NewConstruct_AsBaseClass(jsiiID),
	}
}
