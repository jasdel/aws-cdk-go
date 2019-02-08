package awsstepfunctions

import (
	"github.com/awslabs/aws-cdk-go/cdk"
	"github.com/awslabs/aws-cdk-go/jsii"
)

// StateProps are properties.
//
// TODO can this be a structure? Does the JSII for go have enough information
// to make this determination?
type StateProps struct {
	Comment   *string
	InputPath *string
	//...
}

// State provides the subtyping interface for JSII State_ class.
type State interface {
	// JSII for Go generator must generate Go interfaces for a JSII class with
	// extended classes in mind.
	cdk.Construct

	// State_ properties
	StartState() State
	EndStates() []INextable // Must be overridden
	ToStateJson() jsii.Any  // Must be overridden

	// State_ Methods
	//...

	state_private()
}

// State_Overrides provides the interface for overriding methods of a generated
// JSII class within the JSII kernel.
type State_Overrides interface {
	// Classes with abstract members and methods allow subclasses to extend
	// (aka replace) those members and methods. Need to provide a way for sub
	// classes to know which members can be extended.
	EndStates() []INextable
	ToStateJson() jsii.Any
}

// State_ is a JSII class.
type State_ struct {
	cdk.Construct

	base jsii.Base
}

func (*State_) state_private() {}

// NewState returns an initialized State_.
//
// State_ is supposed to be an abstract class, which means that its not
// supposed to be created via its constructor directly. What does this mean
// for Go? How should custom sub classes of State_ be created?
//
// TODO: should this be private, should there be a flag on the public,
// what if cross pkg constructs want to initialize a State_?
func NewState(scope cdk.Construct, id string, props StateProps) *State_ {
	return NewState_WithOverrides(scope, id, props, nil)
}

// NewState_AsBaseClass returns an initialized State_ initialized as a base type
// extended by another type. Used internally by generated JSII types.
func NewState_AsBaseClass(jsiiID string) *State_ {
	return &State_{
		base:      jsii.Base{ID: jsiiID},
		Construct: cdk.NewConstruct_AsBaseClass(jsiiID),
	}
}

// NewState_WithOverrides returns an initialized State_ initialized with
// overrides. Use when creating custom types extending JSII generated types.
func NewState_WithOverrides(scope cdk.Construct, id string, props StateProps, overrides State_Overrides) *State_ {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$awsstepfunctions$0.0.0.State",
		[]jsii.Any{},
		overrides,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &State_{
		base:      jsii.Base{ID: jsiiID},
		Construct: cdk.NewConstruct_AsBaseClass(jsiiID),
	}
}

func (c *State_) StartState() State {
	result, err := jsii.GlobalRuntime.Client().Get(c.base.ID, "startStates")
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	_ = result

	return nil
}

func (c *State_) EndStates() []INextable {
	panic("State_.EndStates implementation must be provided")
}

func (c *State_) ToStateJson() jsii.Any {
	panic("State_.ToStateJson implementation must be provided")
}
