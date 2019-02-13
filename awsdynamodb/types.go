package awsdynamodb

import "github.com/awslabs/aws-cdk-go/cdk"

type AttributeTypeIface interface {
}

type AttributeType struct{}

func AttributeType_String() AttributeTypeIface { return nil }

type AttributeIface interface {
	GetName() string
	GetType() AttributeTypeIface
}

type Attribute struct {
	Name string
	Type AttributeTypeIface
}

func (a Attribute) GetName() string             { return a.Name }
func (a Attribute) GetType() AttributeTypeIface { return a.Type }

type TablePropsIface interface {
	GetPartitionKey() AttributeIface
	GetTableName() *string
	//...
}

type TableProps struct {
	PartitionKey AttributeIface
	TableName    *string
}

func (t TableProps) GetPartitionKey() AttributeIface { return t.PartitionKey }
func (t TableProps) GetTableName() *string           { return t.TableName }

type TableIface interface {
	cdk.ConstructIface

	AddPartitionKey(key AttributeIface)
	TableName() string

	tablePrivate()
}

type Table struct {
	*cdk.Construct
}

func (Table) tablePrivate() {}

func NewTable(scope cdk.ConstructIface, id string, props TablePropsIface) *Table {
	return nil
}
func InternalNewTableAsBaseClass(jsiiID string) *Table {
	return nil
}
func ExtendTable(overrides TableIface, scope cdk.ConstructIface, id string, props TablePropsIface) *Table {
	return nil
}

func (t *Table) AddPartitionKey(key AttributeIface) {}
func (t *Table) TableName() string                  { return "" }
