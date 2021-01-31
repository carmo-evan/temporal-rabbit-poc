package domain

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
)

// ConvertImageTaskQueue is the name of the queue of convertImage jobs
const ConvertImageTaskQueue = "Convert Image Task Queue"

// ConvertImageWorkflowOptions are the shared options used to start this workflow
var ConvertImageWorkflowOptions = client.StartWorkflowOptions{
	ID:        "convert-image-workflow",
	TaskQueue: ConvertImageTaskQueue,
}

// ConvertImageWorkflow takes an image and puts it through a conversion process that depends on a rabbit MQ message
func ConvertImageWorkflow(ctx workflow.Context, image string) (string, error) {
	var result string
	msg := workflow.GetSignalChannel(ctx, "message")
	msg.Receive(ctx, &result)
	log.Println("Received rabbit message", result)
	return result, nil
}
