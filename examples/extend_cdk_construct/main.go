package main

import (
	"strings"

	"github.com/awslabs/aws-cdk-go/awsapigateway"
	"github.com/awslabs/aws-cdk-go/awslambda"
	"github.com/awslabs/aws-cdk-go/cdk"
)

func main() {
	app := cdk.NewApp()

	addHandlerStack(app, cdk.StackProps_{
		NamingScheme_: NamingSchemeFunc(func(parts []string) string {
			return strings.Join(parts, ":")
		}),
	})

	app.Run()
}

type NamingSchemeFunc func([]string) string

func (fn NamingSchemeFunc) AlocateAddress(parts []string) string {
	return fn(parts)
}

func addHandlerStack(app cdk.App, stackProps cdk.StackProps) {
	stack := cdk.NewStack(app, "MyHandlerStack", stackProps)

	validator := func(fn awslambda.Function) []string {
		if fn.FunctionName() != "HelloHandler" {
			return []string{"Wrong name for handler"}
		}
		return nil
	}

	// Custom Lambda Funtion construct with support for sepcifing a custom
	// validator that will be used when the function construct is syntisized.
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

	ValidateFn LambdaValidator
}

type LambdaValidator func(fn awslambda.Function) []string

func NewValidateLambdaFunction(scope cdk.Construct, id string, validateFn LambdaValidator, props awslambda.FunctionProps) *ValidateLambdaFunction {
	v := &ValidateLambdaFunction{
		ValidateFn: validateFn,
	}

	lambda := awslambda.NewFunction_WithOverrides(scope, id, props, v)

	v.Function = lambda
	return v
}

func (l *ValidateLambdaFunction) Validate() []string {
	return l.ValidateFn(l)
}
