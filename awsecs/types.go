package awsecs

import (
	"github.com/awslabs/aws-cdk-go/awsec2"
	"github.com/awslabs/aws-cdk-go/cdk"
)

type LoadBalancedFargateServiceProps interface {
	Cluster() Cluster
	Image() ContainerImage
}
type LoadBalancedFargateServiceProps_ struct {
	Cluster_ Cluster
	Image_   ContainerImage
}

func (l LoadBalancedFargateServiceProps_) Cluster() Cluster      { return l.Cluster_ }
func (l LoadBalancedFargateServiceProps_) Image() ContainerImage { return l.Image_ }

type LoadBalancedFargateService interface {
	loadBalancedFargateService_private()
}

type LoadBalancedFargateService_ struct{}

func (LoadBalancedFargateService_) loadBalancedFargateService_private() {}

func NewLoadBalancedFargateService(scope cdk.Construct, id string, props LoadBalancedFargateServiceProps) *LoadBalancedFargateService_ {
	return nil
}

type ClusterProps interface {
	Vpc() awsec2.IVpcNetwork
}
type ClusterProps_ struct {
	Vpc_ awsec2.IVpcNetwork
}

func (c ClusterProps_) Vpc() awsec2.IVpcNetwork { return c.Vpc_ }

type Cluster interface {
	cluster_private()
}
type Cluster_ struct {
}

func (Cluster_) cluster_private() {}

func NewCluster(scope cdk.Construct, id string, prop ClusterProps) *Cluster_ {
	return nil
}

type ContainerImage interface{}

func ContainerImage_FromDockerHub(name string) ContainerImage { return nil }
