package main

import (
	"github.com/awslabs/aws-cdk-go/awsapigateway"
	"github.com/awslabs/aws-cdk-go/awslambda"
	"github.com/awslabs/aws-cdk-go/cdk"
)

func main() {
	app := cdk.NewApp()

	addHandlerStack(app)

	app.Run()
}

func addHandlerStack(app cdk.App) {
	stack := cdk.NewStack(app, "MyHandlerStack", nil)

	// Create a validator function that will be used by ValidateLambdaFunction
	// when the lambda function CDK construct is generated into A
	// CloudFormation template
	validator := func(fn awslambda.Function) []string {
		if fn.FunctionName() != "HelloHandler" {
			return []string{"Wrong name for handler"}
		}
		return nil
	}

	// Custom Lambda Function construct with support for specifying a custom
	// validator that will be used when the CDK construct is synthesized into a
	// CloudFormation template.
	helloFn := NewValidateLambdaFunction(stack, "HelloHandler", validator, awslambda.FunctionProps_{
		Runtime_: awslambda.Runtime_NodeJS810(),  // Execution environment.
		Code_:    awslambda.Code_Asset("lambda"), // Code loaded from the "lambda" directory.
		Handler_: "hello.Handler",                // File is "hello", function is "handler.
	})

	// Defines an API Gateway REST API resource backed by our "hello" function.
	awsapigateway.NewLambdaRestApi(stack, "Endpoint", awsapigateway.LambdaRestApiProps_{
		Handler_: helloFn,
	})

}

type ValidateLambdaFunction struct {
	awslambda.Function

	validateFn LambdaValidator
}

type LambdaValidator func(fn awslambda.Function) []string

// NewValidateLambdaFunction returns a ValidateLambdaFunction extending the
// awslambda.Function CDK construct. The custom validator will be used when
// ValidateLambdaFunction's Validate method is invoked within the CDK runtime.
func NewValidateLambdaFunction(scope cdk.Construct, id string, validateFn LambdaValidator, props awslambda.FunctionProps) *ValidateLambdaFunction {
	v := &ValidateLambdaFunction{
		validateFn: validateFn,
	}

	lambda := awslambda.NewFunction_WithOverrides(scope, id, props, v)

	v.Function = lambda
	return v
}

// Validate overrides awslambda.Function's Validate method to perform
// validation on the lambda function during synthesis of the CloudFromation
// template.
func (l *ValidateLambdaFunction) Validate() []string {
	return l.validateFn(l)
}
