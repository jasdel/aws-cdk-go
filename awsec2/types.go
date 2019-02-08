package awsec2

import "github.com/awslabs/aws-cdk-go/cdk"

type IVpcNetwork interface{}

type VpcNetwork interface {
	vpcNetwork_private()
}
type VpcNetwork_ struct {
}

func (VpcNetwork_) vpcNetwork_private() {}

func NewVpcNetwork(scope cdk.Construct, id string) *VpcNetwork_ {
	return nil
}
