package awsec2

import "github.com/awslabs/aws-cdk-go/cdk"

type IVpcNetwork interface{}

type VpcNetworkIface interface {
	vpcNetwork_private()
}
type VpcNetwork struct {
}

func (VpcNetwork) vpcNetworkPrivate() {}

func NewVpcNetwork(scope cdk.ConstructIface, id string) *VpcNetwork {
	return nil
}
func InternalNewVpcNetworkAsBaseClass(jsiiID string) *VpcNetwork {
	return nil
}
func ExtendVpcNetwork(overrides VpcNetworkIface, scope cdk.ConstructIface, id string) *VpcNetwork {
	return nil
}
