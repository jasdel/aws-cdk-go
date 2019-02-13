package main

import (
	"github.com/awslabs/aws-cdk-go/awsec2"
	"github.com/awslabs/aws-cdk-go/awsecs"
	"github.com/awslabs/aws-cdk-go/cdk"
)

func main() {
	app := cdk.NewApp()

	addECSStack(app, "MyECSStack", nil)

	app.Run()
}

func addECSStack(app cdk.AppIface, id string, stackProps cdk.StackPropsIface) {
	stack := cdk.NewStack(app, id, stackProps)

	vpcnet := awsec2.NewVpcNetwork(stack, "VPC")
	cluster := awsecs.NewCluster(stack, "Cluster", awsecs.ClusterProps{
		Vpc: vpcnet,
	})

	awsecs.NewLoadBalancedFargateService(stack, "Fargate", awsecs.LoadBalancedFargateServiceProps{
		Cluster: cluster,
		Image:   awsecs.ContainerImage_FromDockerHub("amazon/amazon-ecs-sample"),
	})
}
