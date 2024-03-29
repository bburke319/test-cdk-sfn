package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type TestCdkSfnStackProps struct {
	awscdk.StackProps
}

func TestStateMachineStack(scope constructs.Construct, id string, props *TestCdkSfnStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	helloWorldLambda := lambda.NewFunction(stack, jsii.String("HelloWorld"), &lambda.FunctionProps{
		Runtime: lambda.Runtime_NODEJS_18_X(),
		Handler: jsii.String("index.handler"),
		Code: lambda.Code_FromInline(jsii.String(`
			exports.handler = (event, context, callback) => {
				callback(null, "Hello World!");
			};
		`)), // https://pkg.go.dev/github.com/aws/aws-cdk-go/awscdk/v2/awslambda#section-readme
	})

	invokeLambdaTask := tasks.NewLambdaInvoke(stack, jsii.String("InvokeLambda"), &tasks.LambdaInvokeProps{
		LambdaFunction: helloWorldLambda,
		OutputPath:     jsii.String("$.helloWorld"),
	})

	stateMachineDefinition := invokeLambdaTask.Next(sfn.NewSucceed(stack, jsii.String("Done"), &sfn.SucceedProps{}))

	sfn.NewStateMachine(stack, jsii.String("HelloWorldStateMachine"), &sfn.StateMachineProps{
		StateMachineName: jsii.String("HelloWorld"),
		DefinitionBody:   sfn.DefinitionBody_FromChainable(stateMachineDefinition),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	TestStateMachineStack(app, "TestStateMachineStack", &TestCdkSfnStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
