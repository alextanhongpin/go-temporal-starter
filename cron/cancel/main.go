package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{
		//HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("failed to start Temporal client", err)
	}
	defer c.Close()

	// Obtained from the dashboard.
	workflowID := "89f7ec3f-8969-4e28-8f1b-95b0352ca0ad"
	runID := "a03eb91e-693a-4d76-889a-be6e20776a54"

	// NOTE: You need to have the worker running in order to cancel it.
	// If the worker is already terminate,
	// then the cancellation message will
	// not be received.
	err = c.CancelWorkflow(context.Background(), workflowID, runID)
	if err != nil {
		log.Fatalln("unable to cancel workflow", err)
	}
	log.Println("WorkflowExecutionCancelled", "WorkflowID", workflowID, "RunID", runID)
}
