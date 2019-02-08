package main

import (
	"github.com/awslabs/aws-cdk-go/awsapigateway"
	"github.com/awslabs/aws-cdk-go/awsdynamodb"
	"github.com/awslabs/aws-cdk-go/awslambda"
	"github.com/awslabs/aws-cdk-go/cdk"
	"github.com/awslabs/aws-cdk-go/jsii"
)

func main() {
	app := cdk.NewApp()

	addHelloHandlerStack(app)

	app.Run()
}

func addHelloHandlerStack(app cdk.App) {
	stack := cdk.NewStack(app, "CDKWorkshopStack", nil)

	// Defines an AWS Lambda resource.
	helloFn := awslambda.NewFunction(stack, "HelloHandler", awslambda.FunctionProps_{
		Runtime_: awslambda.Runtime_NodeJS810(),  // Execution environment.
		Code_:    awslambda.Code_Asset("lambda"), // Code loaded from the "lambda" directory.
		Handler_: "hello.Handler",                // File is "hello", function is "handler.
	})

	hitCounterFn := WrapLambdaWithHitCounter(stack, "HelloHitCounter", helloFn)

	// Defines an API Gateway REST API resource backed by our "hello" function.
	awsapigateway.NewLambdaRestApi(stack, "Endpoint", awsapigateway.LambdaRestApiProps_{
		Handler_: hitCounterFn,
	})
}

func WrapLambdaWithHitCounter(stack cdk.Stack, id string, downstreamFn awslambda.Function) awslambda.Function {
	wrapper := cdk.NewConstruct(stack, id)

	table := awsdynamodb.NewTable(wrapper, "Hits")
	table.AddPartitionKey(awsdynamodb.PartitionKey_{
		Name_: "path",
		Type_: awsdynamodb.AttributeType_String(),
	})

	handler := awslambda.NewFunction(wrapper, "HitCounterHandler", awslambda.FunctionProps_{
		Runtime_: awslambda.Runtime_NodeJS810(),  // Execution environment.
		Code_:    awslambda.Code_Asset("lambda"), // Code loaded from the "lambda" directory.
		Handler_: "hitcounter.Handler",           // File is "hitcounter", function is "handler.
		Environment_: map[string]jsii.Any{
			"DOWNSTREAM_FUNCTION_NAME": downstreamFn.FunctionName(),
			"HITS_TABLE_NAME":          table.TableName(),
		},
	})

	return handler
}
