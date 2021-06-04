package main

import (
	"log"

	app "github.com/alextanhongpin/go-temporal-starter/money-transfer"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// This worker hosts both Worker and Activity functions.
	w := worker.New(c, app.TransferMoneyTaskQueue, worker.Options{})
	w.RegisterWorkflow(app.TransferMoney)
	w.RegisterActivity(app.Withdraw)
	w.RegisterActivity(app.Deposit)

	// Start listening to Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
