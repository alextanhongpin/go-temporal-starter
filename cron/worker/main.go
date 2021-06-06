package main

import (
	"log"

	app "github.com/alextanhongpin/go-temporal-starter/cron"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("failed to start Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, app.CronTaskQueue, worker.Options{})
	w.RegisterWorkflow(app.SampleCronWorkflow)
	w.RegisterActivity(app.DoSomething)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}
