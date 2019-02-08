package cdk

import "github.com/awslabs/aws-cdk-go/jsii"

// ConstructNode provides the subtyping interfaces for JSII ConstructNode_ class.
type ConstructNode interface {
	// Members
	Children() []IConstruct
	Host() Construct
	Id() string
	Locked() bool
	Metadata() []MetadataEntry
	Path() string
	Typename() string
	UniqueId() string
	Scope() IConstruct

	// Methods
	AddChild(child IConstruct, childName string)
	AddDependency(dependencies ...IDependable)
	AddError(message string)
	AddInfo(message string)
	AddMetadata(typeVal string, data jsii.Any, from *bool) // typeVal used instead of "type" param
	AddWarning(message string)
	Ancestors(upTo Construct) IConstruct // upTo is optional, optionals of class types are not pointers since Go interface zero value is nil.
	FindAll(order ConstructOrder) IConstruct
	FindChild(path string) IConstruct
	FindDependencies() []Dependency
	FindReferences() Token
	GetContext(key string) jsii.Any
	Lock()
	PrepareTree()
	RecordReference(refs ...Token)
	RequireContext(key string) jsii.Any
	Required(props jsii.Any, name string) jsii.Any
	Resolve(obj jsii.Any) jsii.Any
	SetContext(key string, value jsii.Any)
	StringifyJson(obj jsii.Any) string
	ToTreeString(depth *jsii.Number) string
	TryFindChild(path string) IConstruct
	Unlock()
	ValidateTree() []ValidationError

	cosntructNode_private()
}

// ConstructNode_Override provides the interface for overriding methods of a
// generated JSII class within the JSII kernel.
type ConstructNode_Override interface {
}

type ConstructNode_ struct {
	base jsii.Base
}

func (*ConstructNode_) construcNode_private() {}

// NewConstructNode returns an initialized ConstructNode_.
func NewConstructNode(host Construct, scope IConstruct, id string) *ConstructNode_ {
	return NewConstructNode_WithOverrides(host, scope, id, nil)
}

// NewConstructNode_AsBaseClass returns an initialized ConstructNode_
// initialized as a base type extended by another type. Used internally by
// generated JSII types.
func NewConstructNode_AsBaseClass(jsiiID string) *ConstructNode_ {
	return &ConstructNode_{
		base: jsii.Base{ID: jsiiID},
	}
}

// NewConstructNode_WithOverrides returns an initialized ConstructNode_
// initialized with overrides. Use when creating custom types extending JSII
// generated types.
func NewConstructNode_WithOverrides(host Construct, scope IConstruct, id string, overrides ConstructNode_Override) *ConstructNode_ {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$cdk$0.0.0.ConstructNode_",
		[]jsii.Any{
			host, scope, id,
		},
		overrides,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &ConstructNode_{
		base: jsii.Base{ID: jsiiID},
	}
}

func (c *ConstructNode_) Children() []IConstruct                                { return nil }
func (c *ConstructNode_) Host() Construct                                       { return nil }
func (c *ConstructNode_) Id() string                                            { return "" }
func (c *ConstructNode_) Locked() bool                                          { return false }
func (c *ConstructNode_) Metadata() []MetadataEntry                             { return nil }
func (c *ConstructNode_) Path() string                                          { return "" }
func (c *ConstructNode_) Typename() string                                      { return "" }
func (c *ConstructNode_) UniqueId() string                                      { return "" }
func (c *ConstructNode_) Scope() IConstruct                                     { return nil }
func (c *ConstructNode_) AddChild(child IConstruct, childName string)           {}
func (c *ConstructNode_) AddDependency(dependencies ...IDependable)             {}
func (c *ConstructNode_) AddError(message string)                               {}
func (c *ConstructNode_) AddInfo(message string)                                {}
func (c *ConstructNode_) AddMetadata(typeVal string, data jsii.Any, from *bool) {}
func (c *ConstructNode_) AddWarning(message string)                             {}
func (c *ConstructNode_) Ancestors(upTo Construct) IConstruct                   { return nil }
func (c *ConstructNode_) FindAll(order ConstructOrder) IConstruct               { return nil }
func (c *ConstructNode_) FindChild(path string) IConstruct                      { return nil }
func (c *ConstructNode_) FindDependencies() []Dependency                        { return nil }
func (c *ConstructNode_) FindReferences() Token                                 { return nil }
func (c *ConstructNode_) GetContext(key string) jsii.Any                        { return nil }
func (c *ConstructNode_) Lock()                                                 {}
func (c *ConstructNode_) PrepareTree()                                          {}
func (c *ConstructNode_) RecordReference(refs ...Token)                         {}
func (c *ConstructNode_) RequireContext(key string) jsii.Any                    { return nil }
func (c *ConstructNode_) Required(props jsii.Any, name string) jsii.Any         { return nil }
func (c *ConstructNode_) Resolve(obj jsii.Any) jsii.Any                         { return nil }
func (c *ConstructNode_) SetContext(key string, value jsii.Any)                 {}
func (c *ConstructNode_) StringifyJson(obj jsii.Any) string                     { return "" }
func (c *ConstructNode_) ToTreeString(depth *jsii.Number) string                { return "" }
func (c *ConstructNode_) TryFindChild(path string) IConstruct                   { return nil }
func (c *ConstructNode_) Unlock()                                               {}
func (c *ConstructNode_) ValidateTree() []ValidationError                       { return nil }
