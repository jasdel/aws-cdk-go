package awsstepfunctions

import (
	"github.com/awslabs/aws-cdk-go/cdk"
	"github.com/awslabs/aws-cdk-go/jsii"
)

type StatePropsIface interface {
	GetComment() *string
	GetInputPath() *string
}

type StateProps struct {
	Comment   *string
	InputPath *string
	//...
}

func (s StateProps) GetComment() *string   { return s.Comment }
func (s StateProps) GetInputPath() *string { return s.InputPath }

// State provides the subtyping interface for JSII State class.
type StateIface interface {
	// JSII for Go generator must generate Go interfaces for a JSII class with
	// extended classes in mind.
	cdk.ConstructIface

	// State properties
	StartState() StateIface
	EndStates() []INextable   // Must be overridden
	ToStateJson() interface{} // Must be overridden

	// State Methods
	//...

	statePrivate()
}

// State is a JSII class.
type State struct {
	*cdk.Construct

	base jsii.Base
}

func (*State) state_private() {}

// NewState returns an initialized State.
//
// State is supposed to be an abstract class, which means that its not
// supposed to be created via its constructor directly. What does this mean
// for Go? How should custom sub classes of State be created?
//
// TODO: should this be private, should there be a flag on the public,
// what if cross pkg constructs want to initialize a State?
func NewState(scope cdk.Construct, id string, props StatePropsIface) *State {
	return ExtendState(nil, scope, id, props)
}

// InternalNewStateAsBaseClass returns an initialized State initialized as a
// base type extended by another type. Used internally by generated JSII types.
func InternalNewStateAsBaseClass(jsiiID string) *State {
	return &State{
		base:      jsii.Base{ID: jsiiID},
		Construct: cdk.InternalNewConstructAsBaseClass(jsiiID),
	}
}

// ExtendState returns an initialized State initialized with overrides. Use
// when creating custom types extending JSII generated types.
func ExtendState(overrides StateIface, scope cdk.Construct, id string, props StatePropsIface) *State {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$awsstepfunctions$0.0.0.State",
		[]interface{}{},
		overrides,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &State{
		base:      jsii.Base{ID: jsiiID},
		Construct: cdk.InternalNewConstructAsBaseClass(jsiiID),
	}
}

func (c *State) StartState() StateIface {
	result, err := jsii.GlobalRuntime.Client().Get(c.base.ID, "startStates")
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	_ = result

	return nil
}

func (c *State) EndStates() []INextable {
	panic("State.EndStates implementation must be provided")
}

func (c *State) ToStateJson() interface{} {
	panic("State.ToStateJson implementation must be provided")
}
