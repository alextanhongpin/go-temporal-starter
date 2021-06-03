package greet

import "go.temporal.io/sdk/workflow"

func SimpleWorkflow(ctx workflow.Context, msg string) (string, error) {
	// Do something.

	return "hello " + msg, nil
}
