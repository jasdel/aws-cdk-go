package main

import (
	"github.com/awslabs/aws-cdk-go/awsapigateway"
	"github.com/awslabs/aws-cdk-go/awslambda"
	"github.com/awslabs/aws-cdk-go/cdk"
)

func main() {
	app := cdk.NewApp()

	addHelloHandlerStack(app)

	app.Run()
}

func addHelloHandlerStack(app cdk.AppIface) {
	stack := cdk.NewStack(app, "CDKWorkshopStack", nil)

	// Defines an AWS Lambda resource.
	helloFn := awslambda.NewFunction(stack, "HelloHandler", awslambda.FunctionProps{
		Runtime: awslambda.Runtime_NodeJS810(),  // Execution environment.
		Code:    awslambda.Code_Asset("lambda"), // Code loaded from the "lambda" directory.
		Handler: "hello.Handler",                // File is "hello", function is "handler.
	})

	// Defines an API Gateway REST API resource backed by our "hello" function.
	awsapigateway.NewLambdaRestApi(stack, "Endpoint", awsapigateway.LambdaRestApiProps{
		Handler: helloFn,
	})
}
