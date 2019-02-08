package awsstepfunctions

import (
	"github.com/awslabs/aws-cdk-go/cdk"
	"github.com/awslabs/aws-cdk-go/jsii"
)

// FailProps provides the interface for FailProps_ datatype.
type FailProps interface {
	Comment() *string
	Error() string
	Reason() *string
}

// FailProps_ are properties.
type FailProps_ struct {
	Comment_ *string
	Error_   string
	Reason_  *string
}

func (d FailProps_) Comment() *string { return jsii.S("") }
func (d FailProps_) Error() string    { return "" }
func (d FailProps_) Reason() *string  { return jsii.S("") }

// Fail provides the subtyping interfaces for JSII Fail class.
type Fail interface {
	State

	fail_private()
}

// Fail_Overrides provides the interface for overriding methods of a generated
// JSII class within the JSII kernel.
type Fail_Overrides interface{}

// Fail_ is a JSII class.
type Fail_ struct {
	State

	base jsii.Base
}

func (*Fail_) fail_private() {}

// NewFail returns an initialized Fail.
func NewFail(scope cdk.Construct, id string, props FailProps) *Fail_ {
	return NewFail_WithOverrides(scope, id, props, nil)
}

// NewFail_AsBaseClass returns an initialized Fail initialized as a base type
// extended by another type. Used internally by generated JSII types.
func NewFail_AsBaseClass(jsiiID string) *Fail_ {
	return &Fail_{
		base:  jsii.Base{ID: jsiiID},
		State: NewState_AsBaseClass(jsiiID),
	}
}

// NewFail_WithOverrides returns an initialized Fail initialized with
// overrides. Use when creating custom types extending JSII generated types.
func NewFail_WithOverrides(scope cdk.Construct, id string, props FailProps, overrides Fail_Overrides) *Fail_ {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$awsstepfunctions$0.0.0.Fail",
		[]jsii.Any{},
		overrides,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &Fail_{
		base:  jsii.Base{ID: jsiiID},
		State: NewState_AsBaseClass(jsiiID),
	}
}

func (c *Fail_) EndStates() []INextable {
	result, err := jsii.GlobalRuntime.Client().Get(c.base.ID, "endStates")
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	_ = result

	return nil
}

func (c *Fail_) ToStateJson() jsii.Any {
	result, err := jsii.GlobalRuntime.Client().Get(c.base.ID, "toStateJson")
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	_ = result

	return nil
}
