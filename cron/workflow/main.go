package main

import (
	"context"
	"log"

	app "github.com/alextanhongpin/go-temporal-starter/cron"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("failed to start Temporal client", err)
	}
	defer c.Close()

	opts := client.StartWorkflowOptions{
		TaskQueue: app.CronTaskQueue,
		// Run every minute.
		CronSchedule: "* * * * *",
	}
	we, err := c.ExecuteWorkflow(context.Background(), opts, app.SampleCronWorkflow)
	if err != nil {
		log.Fatalln("unable to complete workflow", err)
	}

	// Use the WorkflowExecution to get the result.
	var workflowResult app.CronResult
	err = we.Get(context.Background(), &workflowResult)
	if err != nil {
		log.Fatalln("failed to get workflow result", err)
	}
	log.Printf("%+v", workflowResult)
}
