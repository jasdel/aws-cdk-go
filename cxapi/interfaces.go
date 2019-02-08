package cxapi

import "github.com/awslabs/aws-cdk-go/jsii"

type MissingContext interface {
	Props() map[string]jsii.Any
	Provider() string
}

type MissingContext_ struct {
	Props_    map[string]jsii.Any
	Provider_ string
}

func (m *MissingContext_) Props() map[string]jsii.Any { return m.Props_ }
func (m *MissingContext_) Provider() string           { return m.Provider_ }
