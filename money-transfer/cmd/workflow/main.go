package main

import (
	"context"
	"log"

	app "github.com/alextanhongpin/go-temporal-starter/money-transfer"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "transfer-money-workflow",
		TaskQueue: app.TransferMoneyTaskQueue,
	}
	transferDetails := app.TransferDetails{
		Amount:      54.99,
		FromAccount: "001-001",
		ToAccount:   "002-002",
		ReferenceID: uuid.New().String(),
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, app.TransferMoney, transferDetails)
	if err != nil {
		log.Fatalln("error starting TransferMoney workflow", err)
	}

	log.Printf(
		"\nTransfer of $%f from account %s t oaccount %s is processing. ReferenceId: %s",
		transferDetails.Amount,
		transferDetails.FromAccount,
		transferDetails.ToAccount,
		transferDetails.ReferenceID,
	)
	log.Printf(
		"\nWorkflowID: %s RunID: %s\n",
		we.GetID(),
		we.GetRunID(),
	)
}
