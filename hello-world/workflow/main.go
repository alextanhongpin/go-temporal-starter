package main

import (
	"context"
	"fmt"
	"log"

	app "github.com/alextanhongpin/go-temporal-starter/hello-world"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: app.GreetingTaskQueue,
	}

	name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, app.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete workflow", err)
	}

	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get workflow result", err)
	}

	fmt.Printf("\nWorkflowID: %s, RunID: %s\n", we.GetID(), we.GetRunID())
	fmt.Printf("\n%s\n", greeting)
}
