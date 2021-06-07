package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
)

type GreetingPrefix struct {
	Prefix string
}

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

	{ // SIGNAL BLOCK
		var signalVal GreetingPrefix
		signalChan := workflow.GetSignalChannel(ctx, GreetingSignalName)
		s := workflow.NewSelector(ctx)
		s.AddReceive(signalChan, func(c workflow.ReceiveChannel, more bool) {
			c.Receive(ctx, &signalVal)
			workflow.GetLogger(ctx).Info("Received signal!", zap.String("signalName", GreetingSignalName), zap.Any("signalVal", signalVal))
		})
		s.Select(ctx)
	}

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err = workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)
	return result, err
}
