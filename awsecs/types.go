package awsecs

import (
	"github.com/awslabs/aws-cdk-go/awsec2"
	"github.com/awslabs/aws-cdk-go/cdk"
)

type LoadBalancedFargateServicePropsIface interface {
	GetCluster() ClusterIface
	GetImage() ContainerImageIface
}
type LoadBalancedFargateServiceProps struct {
	Cluster ClusterIface
	Image   ContainerImageIface
}

func (l LoadBalancedFargateServiceProps) GetCluster() ClusterIface      { return l.Cluster }
func (l LoadBalancedFargateServiceProps) GetImage() ContainerImageIface { return l.Image }

type LoadBalancedFargateServiceIface interface {
	loadBalancedFargateServicePrivate()
}

type LoadBalancedFargateService struct{}

func (LoadBalancedFargateService) loadBalancedFargateServicePrivate() {}

func NewLoadBalancedFargateService(scope cdk.ConstructIface, id string, props LoadBalancedFargateServicePropsIface) *LoadBalancedFargateService {
	return nil
}
func InternalNewLoadBalancedFargateServiceAsBaseClass(jsiiID string) *LoadBalancedFargateService {
	return nil
}
func ExtendLoadBalancedFargateService(overrides LoadBalancedFargateServiceIface, scope cdk.ConstructIface, id string, props LoadBalancedFargateServicePropsIface) *LoadBalancedFargateService {
	return nil
}

type ClusterPropsIface interface {
	GetVpc() awsec2.IVpcNetwork
}
type ClusterProps struct {
	Vpc awsec2.IVpcNetwork
}

func (c ClusterProps) GetVpc() awsec2.IVpcNetwork { return c.Vpc }

type ClusterIface interface {
	clusterPrivate()
}
type Cluster struct {
}

func (Cluster) clusterPrivate() {}

func NewCluster(scope cdk.ConstructIface, id string, prop ClusterPropsIface) *Cluster {
	return nil
}
func InternalNewClusterAsBaseClass(jsiiID string) *Cluster {
	return nil
}
func ExtendCluster(overrides ClusterIface, scope cdk.ConstructIface, id string, prop ClusterPropsIface) *Cluster {
	return nil
}

type ContainerImageIface interface {
	containerImagePrivate()
}
type ContainerImage struct{}

func (ContainerImage) containerImagePrivate() {}

func ContainerImage_FromDockerHub(name string) ContainerImageIface { return nil }
