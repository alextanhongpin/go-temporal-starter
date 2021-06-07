package main

import (
	"context"
	"log"

	app "github.com/alextanhongpin/go-temporal-starter/hello-world"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}

	workflowID := "greeting-workflow"
	runID := "d6f51f1c-5d3f-4619-9684-8c1a9b412d4f"
	signalName := app.GreetingSignalName
	signalVal := app.GreetingPrefix{
		Prefix: "Mr",
	}
	err = c.SignalWorkflow(context.Background(), workflowID, runID, signalName, signalVal)
	if err != nil {
		log.Fatalln("Error signalling client", err)
	}
}
