package awsstepfunctions_test

import (
	"github.com/awslabs/aws-cdk-go/awsstepfunctions"
	"github.com/awslabs/aws-cdk-go/cdk"
	"github.com/awslabs/aws-cdk-go/jsii"
)

func ExampleFail() {
	app := cdk.NewApp()

	stack := cdk.NewStack(app, "ExampleFail", nil)

	awsstepfunctions.NewFail(stack, "aFail", nil)

	app.Run()
}

func ExampleCustomState() {
	app := cdk.NewApp()

	stack := cdk.NewStack(app, "ExampleCustomState", nil)

	NewCustomState(stack, "myState")

	app.Run()
}

// CustomState extends CDK AWS StepFunction's State construct.
type CustomState struct {
	*awsstepfunctions.State
}

// NewCustomState returns an initialized CustomState value with the underlying
// State CDK construct initialized in JSII kernel as well.
//
// All State methods will be invoked on the embedded State value, except the
// EndStates and ToStateJson methods which provide custom implementation.
func NewCustomState(scope cdk.Construct, id string) *CustomState {
	custState := &CustomState{}

	// Create the extended Base class with the method overrides that
	// CustomState will provide the implementation for.
	state := awsstepfunctions.ExtendState(custState, scope, id, nil)

	// Update custState's State value so that the underlying State construct value can be used for
	custState.State = state

	return custState
}

// EndStates custom implementation.
func (c *CustomState) EndStates() []awsstepfunctions.INextable {
	// Custom impl
	return nil
}

// ToStateJson custom implementation.
func (c *CustomState) ToStateJson() jsii.Any {
	// Custom impl
	return nil
}
