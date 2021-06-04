package main

import (
	"context"
	"log"

	app "github.com/alextanhongpin/go-temporal-starter/task-queue"
	"go.temporal.io/sdk/client"
)

func main() {
	// Create the client object just once per process.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		TaskQueue: app.GreetTaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, app.SimpleWorkflow, "world")
	if err != nil {
		log.Fatalf("failed to execute workflow: %s", err)
	}

	// Use the WorkflowExecution to get the result.
	// Get is blocking call and will wait for the Workflow to complete.
	var workflowResult string
	err = we.Get(context.Background(), &workflowResult)
	if err != nil {
		log.Fatalf("failed to get: %s", err)
	}

	// Do something with the result.
	log.Println(workflowResult)
	log.Println(we.GetID())
	log.Println(we.GetRunID())
}
