package awsstepfunctions

import "github.com/awslabs/aws-cdk-go/cdk"

var _ cdk.Construct = (*State_)(nil)
var _ cdk.Construct = (*Fail_)(nil)
var _ State = (*State_)(nil)
var _ State = (Fail)(nil)
var _ Fail = (*Fail_)(nil)
