package awsstepfunctions

// IChainable is a JSII interface (not a datatype because of the leading `I` in
// name.
type IChainable interface {
	Id() string
	StartState() State
	EndStates() []INextable

	private()
}

// INextable is a JSII interface
type INextable interface {
	private()
}
