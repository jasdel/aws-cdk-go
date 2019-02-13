package awsstepfunctions

import "github.com/awslabs/aws-cdk-go/cdk"

var _ cdk.ConstructIface = (*State)(nil)
var _ cdk.ConstructIface = (*Fail)(nil)
var _ StateIface = (*State)(nil)
var _ StateIface = (FailIface)(nil)
var _ FailIface = (*Fail)(nil)
