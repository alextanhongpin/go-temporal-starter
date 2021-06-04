package main

import (
	"log"

	app "github.com/alextanhongpin/go-temporal-starter/task-queue"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the client object just once per process.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	w := worker.New(c, app.GreetTaskQueue, worker.Options{})
	w.RegisterWorkflow(app.SimpleWorkflow)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start workflow", err)
	}
}
