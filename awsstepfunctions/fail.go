package awsstepfunctions

import (
	"github.com/awslabs/aws-cdk-go/cdk"
	"github.com/awslabs/aws-cdk-go/jsii"
)

// FailPropsIface provides the interface for FailProps datatype.
type FailPropsIface interface {
	GetComment() *string
	GetError() string
	GetReason() *string
}

// FailProps are properties.
type FailProps struct {
	Comment *string
	Error   string
	Reason  *string
}

func (d FailProps) GetComment() *string { return d.Comment }
func (d FailProps) GetError() string    { return d.Error }
func (d FailProps) GetReason() *string  { return d.Reason }

// Fail provides the subtyping interfaces for JSII Fail class.
type FailIface interface {
	StateIface

	faiPrivate()
}

// Fail is a JSII class.
type Fail struct {
	*State

	base jsii.Base
}

func (*Fail) fail_private() {}

// NewFail returns an initialized Fail.
func NewFail(scope cdk.ConstructIface, id string, props FailPropsIface) *Fail {
	return ExtendFail(nil, scope, id, props)
}

// InternalNewFailAsBaseClass returns an initialized Fail initialized as a base
// type extended by another type. Used internally by generated JSII types.
func InternalNewFailAsBaseClass(jsiiID string) *Fail {
	return &Fail{
		base:  jsii.Base{ID: jsiiID},
		State: InternalNewStateAsBaseClass(jsiiID),
	}
}

// ExtendFail returns an initialized Fail initialized with overrides. Use when
// creating custom types extending JSII generated types.
func ExtendFail(overrides FailIface, scope cdk.ConstructIface, id string, props FailPropsIface) *Fail {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$awsstepfunctions$0.0.0.Fail",
		[]interface{}{},
		overrides,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &Fail{
		base:  jsii.Base{ID: jsiiID},
		State: InternalNewStateAsBaseClass(jsiiID),
	}
}

func (c *Fail) EndStates() []INextable {
	result, err := jsii.GlobalRuntime.Client().Get(c.base.ID, "endStates")
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	_ = result

	return nil
}

func (c *Fail) ToStateJson() interface{} {
	result, err := jsii.GlobalRuntime.Client().Get(c.base.ID, "toStateJson")
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	_ = result

	return nil
}
