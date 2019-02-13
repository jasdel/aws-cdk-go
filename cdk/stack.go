package cdk

import (
	"github.com/awslabs/aws-cdk-go/cxapi"
	"github.com/awslabs/aws-cdk-go/jsii"
)

// Stack_AnnotatePhysicalName is a stack method on the Stack JSII class.
func Stack_AnnotatePhysicalName(construct ConstructIface, phyisicalName string) {}

// Stack_Find is a stack method on the Stack JSII class.
func Stack_Find(scope IConstruct) StackIface { return nil }

// Stack_IsStack is a stack method on the Stack JSII class.
func Stack_IsStack(construct IConstruct) bool { return false }

// Stack_TryFind is a stack method on the Stack JSII class.
func Stack_TryFind(scope IConstruct) *bool { return jsii.B(false) }

// StackProps provides the interface for StackProps datatype.
type StackPropsIface interface {
	GetEnv() Environment
	GetNamingScheme() IAddressingScheme
}

// StackProps are properties.
type StackProps struct {
	Env          EnvironmentIface
	NamingScheme IAddressingScheme
}

func (s StackProps) GetEnv() EnvironmentIface           { return s.Env }
func (s StackProps) GetNamingScheme() IAddressingScheme { return s.NamingScheme }

// StackIface provides the subtyping interfaces for JSII Stack class.
type StackIface interface {
	ConstructIface

	AddDependency(stack StackIface)
	Dependencies() []StackIface
	FindResource(path string) ResourceIface
	FormatArn(components ArnComponentsIface) string
	ParentApp() AppIface
	ParseArn(arn string, sepIfToken *bool, hasName *bool) ArnComponentsIface
	RenameLogic(oldId string, newId string)
	ReportMissingConstext(key string, details cxapi.MissingContextIface)
	RequireAccountId(why *string) string
	RequireRegion(why *string) string
	ToCloudFormation() interface{}

	stackPrivate()
}

var _ StackIface = (*Stack)(nil)

// Stack is a JSII class.
type Stack struct {
	*Construct

	base jsii.Base
}

func (Stack) stackPrivate() {}

// NewStack returns an initialized Stack.
func NewStack(scope AppIface, id string, props StackPropsIface) *Stack {
	return ExtendStack(nil, scope, id, props)
}

// InternalNewStackAsBaseClass returns an initialized Stack initialized as a base type
// extended by another type. Used internally by generated JSII types.
func InternalNewStackAsBaseClass(jsiiID string) *Stack {
	return &Stack{
		base:      jsii.Base{ID: jsiiID},
		Construct: InternalNewConstructAsBaseClass(jsiiID),
	}
}

// ExtendStack returns an initialized Stack initialized with
// overrides. Use when creating custom types extending JSII generated types.
func ExtendStack(overrides StackIface, scope AppIface, id string, props StackPropsIface) *Stack {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$cdk$0.0.0.Stack",
		[]interface{}{
			scope, id,
		},
		nil,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &Stack{
		base:      jsii.Base{ID: jsiiID},
		Construct: InternalNewConstructAsBaseClass(jsiiID),
	}
}

func (*Stack) AddDependency(stack StackIface)                 {}
func (*Stack) Dependencies() []StackIface                     { return nil }
func (*Stack) FindResource(path string) ResourceIface         { return nil }
func (*Stack) FormatArn(components ArnComponentsIface) string { return "" }
func (*Stack) ParentApp() AppIface                            { return nil }
func (*Stack) ParseArn(arn string, sepIfToken *bool, hasName *bool) ArnComponentsIface {
	return nil
}
func (*Stack) Prepare()                                                            {}
func (*Stack) RenameLogic(oldId string, newId string)                              {}
func (*Stack) ReportMissingConstext(key string, details cxapi.MissingContextIface) {}
func (*Stack) RequireAccountId(why *string) string                                 { return "" }
func (*Stack) RequireRegion(why *string) string                                    { return "" }
func (*Stack) ToCloudFormation() interface{}                                       { return nil }
