package awsdynamodb

import "github.com/awslabs/aws-cdk-go/cdk"

type TableIface interface {
	cdk.ConstructIface

	AddPartitionKey(key PartitionKeyIface)
	TableName() string

	tablePrivate()
}

type Table struct {
	*cdk.Construct
}

func (Table) tablePrivate() {}

func NewTable(scope cdk.ConstructIface, id string) *Table {
	return nil
}
func InternalNewTableAsBaseClass(jsiiID string) *Table {
	return nil
}
func ExtendTable(overrides TableIface, scope cdk.ConstructIface, id string) *Table {
	return nil
}

func (t *Table) AddPartitionKey(key PartitionKeyIface) {}
func (t *Table) TableName() string                     { return "" }

type PartitionKeyIface interface {
	GetName() string
	GetType() AttributeTypeIface
}
type PartitionKey struct {
	Name string
	Type AttributeTypeIface
}

func (p PartitionKey) GetName() string             { return p.Name }
func (p PartitionKey) GetType() AttributeTypeIface { return p.Type }

type AttributeTypeIface interface{}
type AttributeType struct{}

func AttributeType_String() AttributeTypeIface { return nil }
