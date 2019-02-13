package cdk

import "github.com/awslabs/aws-cdk-go/jsii"

// ConstructNodeIface provides the subtyping interfaces for JSII ConstructNode class.
type ConstructNodeIface interface {
	// Members
	Children() []IConstruct
	Host() ConstructIface
	Id() string
	Locked() bool
	Metadata() []MetadataEntryIface
	Path() string
	Typename() string
	UniqueId() string
	Scope() IConstruct

	// Methods
	AddChild(child IConstruct, childName string)
	AddDependency(dependencies ...IDependable)
	AddError(message string)
	AddInfo(message string)
	AddMetadata(typeVal string, data interface{}, from *bool) // typeVal used instead of "type" param
	AddWarning(message string)
	Ancestors(upTo ConstructIface) IConstruct // upTo is optional, optionals of class types are not pointers since Go interface zero value is nil.
	FindAll(order ConstructOrder) IConstruct
	FindChild(path string) IConstruct
	FindDependencies() []DependencyIface
	FindReferences() TokenIface
	GetContext(key string) interface{}
	Lock()
	PrepareTree()
	RecordReference(refs ...TokenIface)
	RequireContext(key string) interface{}
	Required(props interface{}, name string) interface{}
	Resolve(obj interface{}) interface{}
	SetContext(key string, value interface{})
	StringifyJson(obj interface{}) string
	ToTreeString(depth *jsii.Number) string
	TryFindChild(path string) IConstruct
	Unlock()
	ValidateTree() []ValidationErrorIface

	construcNodePrivate()
}

type ConstructNode struct {
	base jsii.Base
}

func (*ConstructNode) construcNodePrivate() {}

// NewConstructNode returns an initialized ConstructNode.
func NewConstructNode(host ConstructIface, scope IConstruct, id string) *ConstructNode {
	return ExtendConstructNode(nil, host, scope, id)
}

// InternalNewConstructNodeAsBaseClass returns an initialized ConstructNode
// initialized as a base type extended by another type. Used internally by
// generated JSII types.
func InternalNewConstructNodeAsBaseClass(jsiiID string) *ConstructNode {
	return &ConstructNode{
		base: jsii.Base{ID: jsiiID},
	}
}

// ExtendConstructNode returns an initialized ConstructNode initialized with
// overrides. Use when creating custom types extending JSII generated types.
func ExtendConstructNode(overrides ConstructNodeIface, host ConstructIface, scope IConstruct, id string) *ConstructNode {
	jsiiID, err := jsii.GlobalRuntime.Client().Create(
		"jsii$cdk$0.0.0.ConstructNode",
		[]interface{}{
			host, scope, id,
		},
		overrides,
	)
	if err != nil {
		panic("how are error handled?" + err.Error())
	}
	return &ConstructNode{
		base: jsii.Base{ID: jsiiID},
	}
}

func (c *ConstructNode) Children() []IConstruct                                   { return nil }
func (c *ConstructNode) Host() ConstructIface                                     { return nil }
func (c *ConstructNode) Id() string                                               { return "" }
func (c *ConstructNode) Locked() bool                                             { return false }
func (c *ConstructNode) Metadata() []MetadataEntryIface                           { return nil }
func (c *ConstructNode) Path() string                                             { return "" }
func (c *ConstructNode) Typename() string                                         { return "" }
func (c *ConstructNode) UniqueId() string                                         { return "" }
func (c *ConstructNode) Scope() IConstruct                                        { return nil }
func (c *ConstructNode) AddChild(child IConstruct, childName string)              {}
func (c *ConstructNode) AddDependency(dependencies ...IDependable)                {}
func (c *ConstructNode) AddError(message string)                                  {}
func (c *ConstructNode) AddInfo(message string)                                   {}
func (c *ConstructNode) AddMetadata(typeVal string, data interface{}, from *bool) {}
func (c *ConstructNode) AddWarning(message string)                                {}
func (c *ConstructNode) Ancestors(upTo ConstructIface) IConstruct                 { return nil }
func (c *ConstructNode) FindAll(order ConstructOrder) IConstruct                  { return nil }
func (c *ConstructNode) FindChild(path string) IConstruct                         { return nil }
func (c *ConstructNode) FindDependencies() []DependencyIface                      { return nil }
func (c *ConstructNode) FindReferences() TokenIface                               { return nil }
func (c *ConstructNode) GetContext(key string) interface{}                        { return nil }
func (c *ConstructNode) Lock()                                                    {}
func (c *ConstructNode) PrepareTree()                                             {}
func (c *ConstructNode) RecordReference(refs ...TokenIface)                       {}
func (c *ConstructNode) RequireContext(key string) interface{}                    { return nil }
func (c *ConstructNode) Required(props interface{}, name string) interface{}      { return nil }
func (c *ConstructNode) Resolve(obj interface{}) interface{}                      { return nil }
func (c *ConstructNode) SetContext(key string, value interface{})                 {}
func (c *ConstructNode) StringifyJson(obj interface{}) string                     { return "" }
func (c *ConstructNode) ToTreeString(depth *jsii.Number) string                   { return "" }
func (c *ConstructNode) TryFindChild(path string) IConstruct                      { return nil }
func (c *ConstructNode) Unlock()                                                  {}
func (c *ConstructNode) ValidateTree() []ValidationErrorIface                     { return nil }
