package cdk

// IDependable is a JSII interface.
type IDependable interface {
	DependencyRoots() []IConstruct
}

// IConstruct is a JSII interface.
type IConstruct interface {
	Node() ConstructNode
}
