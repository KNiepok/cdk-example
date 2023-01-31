package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkExampleStackProps struct {
	awscdk.StackProps
}

func NewCdkExampleStack(scope constructs.Construct, id string, props *CdkExampleStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	awscdklambdagoalpha.NewGoFunction(stack, jsii.String("helloWorld"), &awscdklambdagoalpha.GoFunctionProps{
		Entry:   jsii.String("cmd/api/main.go"),
		Runtime: awslambda.Runtime_GO_1_X(),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkExampleStack(app, "CdkExampleStack", &CdkExampleStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String("319313959906"),
		Region:  jsii.String("eu-west-1"),
	}
}
