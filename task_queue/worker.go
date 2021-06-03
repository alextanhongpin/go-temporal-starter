package main

import (
	"log"

	"github.com/alextanhongpin/go-temporal-starter/worker/greet"
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

	w := worker.New(c, "your_task_queue", worker.Options{})
	w.RegisterWorkflow(greet.SimpleWorkflow)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start workflow", err)
	}
}
