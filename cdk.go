package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type VpcStackProps struct {
	awscdk.StackProps
}

func NewVpcStack(scope constructs.Construct, id string, props *VpcStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	awsec2.NewVpc(stack, &id, &awsec2.VpcProps{
		VpcName:            jsii.String("vpc"),
		MaxAzs:             jsii.Number(3),
		Cidr:               jsii.String("10.0.0.0/16"),
		EnableDnsHostnames: jsii.Bool(true),
		EnableDnsSupport:   jsii.Bool(true),
		SubnetConfiguration: &[]*awsec2.SubnetConfiguration{
			{
				CidrMask:   jsii.Number(24),
				Name:       jsii.String("subnet-public-1"),
				SubnetType: awsec2.SubnetType_PUBLIC,
			},
			{
				CidrMask:   jsii.Number(24),
				Name:       jsii.String("subnet-public-2"),
				SubnetType: awsec2.SubnetType_PUBLIC,
			},
			{
				CidrMask:   jsii.Number(24),
				Name:       jsii.String("subnet-private-3"),
				SubnetType: awsec2.SubnetType_PUBLIC,
			},
		},
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewVpcStack(app, "vpc-stack", &VpcStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// // env determines the AWS environment (account+region) in which our stack is to
// // be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	// return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.a
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		// Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region: jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
