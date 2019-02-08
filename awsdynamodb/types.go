package awsdynamodb

import "github.com/awslabs/aws-cdk-go/cdk"

type Table interface {
	AddPartitionKey(key PartitionKey)

	table_private()
}

type Table_ struct {
}

func (Table_) table_private() {}

func NewTable(scope cdk.Construct, id string) *Table_ {
	return nil
}

func (t *Table_) AddPartitionKey(key PartitionKey) {}

type PartitionKey interface {
	Name() string
	Type() AttributeType
}
type PartitionKey_ struct {
	Name_ string
	Type_ AttributeType
}

func (p PartitionKey_) Name() string        { return p.Name_ }
func (p PartitionKey_) Type() AttributeType { return p.Type_ }

type AttributeType interface{}

func AttributeType_String() AttributeType { return nil }
