package cdk

import (
	"github.com/awslabs/aws-cdk-go/cxapi"
	"github.com/awslabs/aws-cdk-go/jsii"
)

// Stack_AnnotatePhysicalName is a stack method on the Stack JSII class.
func Stack_AnnotatePhysicalName(construct Construct, phyisicalName string) {}

// Stack_Find is a stack method on the Stack JSII class.
func Stack_Find(scope IConstruct) Stack { return nil }

// Stack_IsStack is a stack method on the Stack JSII class.
func Stack_IsStack(construct IConstruct) bool { return false }

// Stack_TryFind is a stack method on the Stack JSII class.
func Stack_TryFind(scope IConstruct) *bool { return jsii.B(false) }

// StackProps provides the interface for StackProps_ datatype.
type StackProps interface {
	Env() Environment
	NamingScheme() IAddressingScheme
}

// StackProps_ are properties.
type StackProps_ struct {
	Env_          Environment
	NamingScheme_ IAddressingScheme
}

func (s StackProps_) Env() Environment                { return s.Env_ }
func (s StackProps_) NamingScheme() IAddressingScheme { return s.NamingScheme_ }

// Stack provides the subtyping interfaces for JSII Stack class.
type Stack interface {
	Construct

	AddDependency(stack Stack)
	Dependencies() []Stack
	FindResource(path string) Resource
	FormatArn(components ArnComponents) string
	ParentApp() App
	ParseArn(arn string, sepIfToken *bool, hasName *bool) ArnComponents
	RenameLogic(oldId string, newId string)
	ReportMissingConstext(key string, details cxapi.MissingContext)
	RequireAccountId(why *string) string
	RequireRegion(why *string) string
	ToCloudFormation() jsii.Any

	stack_private()
}

// Stack_Overrides provides the interface for overriding methods of a
// generated JSII class within the JSII kernel.
type Stack_Overrides interface{}

// Stack_ is a JSII class.
type Stack_ struct {
	Construct

	base jsii.Base
}

func (Stack_) stack_private() {}

// NewStack returns an initialized Stack.
func NewStack(scope App, id string, props StackProps) *Stack_ {
	return NewStack_WithOverrides(scope, id, props, nil)
}

// NewStack_AsBaseClass returns an initialized Stack initialized as a base type
// extended by another type. Used internally by generated JSII types.
func NewStack_AsBaseClass(jsiiID string) *Stack_ {
	return &Stack_{
		base:      jsii.Base{ID: jsiiID},
		Construct: NewConstruct_AsBaseClass(jsiiID),
	}
}

// NewStack_WithOverrides returns an initialized Stack initialized with
// overrides. Use when creating custom types extending JSII generated types.
func NewStack_WithOverrides(scope App, id string, props StackProps, overrides Stack_Overrides) *Stack_ {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$cdk$0.0.0.Stack",
		[]jsii.Any{
			scope, id,
		},
		nil,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &Stack_{
		base:      jsii.Base{ID: jsiiID},
		Construct: NewConstruct_AsBaseClass(jsiiID),
	}
}

func (*Stack_) AddDependency(stack Stack)                 {}
func (*Stack_) Dependencies() []Stack                     { return nil }
func (*Stack_) FindResource(path string) Resource         { return nil }
func (*Stack_) FormatArn(components ArnComponents) string { return "" }
func (*Stack_) ParentApp() App                            { return nil }
func (*Stack_) ParseArn(arn string, sepIfToken *bool, hasName *bool) ArnComponents {
	return nil
}
func (*Stack_) Prepare()                                                       {}
func (*Stack_) RenameLogic(oldId string, newId string)                         {}
func (*Stack_) ReportMissingConstext(key string, details cxapi.MissingContext) {}
func (*Stack_) RequireAccountId(why *string) string                            { return "" }
func (*Stack_) RequireRegion(why *string) string                               { return "" }
func (*Stack_) ToCloudFormation() jsii.Any                                     { return nil }
