package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	queryType := "current_state"
	currentState := "started"
	err := workflow.SetQueryHandler(ctx, queryType, func() (string, error) {
		return currentState, nil
	})

	currentState = "sleeping"
	if err := workflow.Sleep(ctx, 10*time.Second); err != nil {
		return "", err
	}
	currentState = "woke up"

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err = workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)
	return result, err
}
