package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

// example tests. To run these tests, uncomment this file along with the
// example resource in test-cdk-sfn_test.go
func TestTestStateMachineStack(t *testing.T) {

	app := awscdk.NewApp(nil)

	stack := TestStateMachineStack(app, "MyStack", nil)

	template := assertions.Template_FromStack(stack, nil)

	template.HasResourceProperties(jsii.String("AWS::StepFunctions::StateMachine"), map[string]interface{}{
		"StateMachineName": "HelloWorld",
	})
}
