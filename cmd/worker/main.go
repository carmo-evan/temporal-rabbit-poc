package main

import (
	"log"

	"github.com/carmo-evan/temporal-poc/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

// ConvertImageTaskQueue is the name of the queue of convertImage jobs
const ConvertImageTaskQueue = "Convert Image Task Queue"

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// This worker hosts both Worker and Activity functions
	w := worker.New(c, ConvertImageTaskQueue, worker.Options{})

	w.RegisterWorkflow(workflow.ConvertImageWorkflow)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
