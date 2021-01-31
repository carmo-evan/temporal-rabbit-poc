package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/carmo-evan/temporal-poc/workflow"
	"github.com/streadway/amqp"
	"go.temporal.io/sdk/client"
)

// this process emulates a server that might be serving data to a frontend client

func main() {
	ctx := context.Background()
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	go listenForRabbitMessage(ctx, c)

	http.HandleFunc("/", Handler(c))
	log.Println("Listening at port 1914")
	http.ListenAndServe(":1914", nil)
}

// Handler takes a payload and starts an activity
func Handler(c client.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var p Payload
		json.NewDecoder(r.Body).Decode(&p)
		we, err := c.ExecuteWorkflow(context.Background(), workflow.ConvertImageWorkflowOptions, workflow.ConvertImageWorkflow, p.Picture)
		if err != nil {
			log.Fatalln("unable to execute Workflow", err)
		}
		log.Println(we.GetRunID())
		log.Printf("%+v\n", p)
	}
}

// Payload represents a struct that holds a picture
type Payload struct {
	Picture string `json:"picture"`
}

func listenForRabbitMessage(ctx context.Context, c client.Client) {
	conn, err := amqp.Dial("amqp://rabbitmq:rabbitmq@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	log.Println("Listening to queue messages")
	if err != nil {
		panic(err)
	}

	for range msgs {
		log.Println("Got rabbit message. Pushing it to workflow.", workflow.ConvertImageWorkflowOptions.ID)
		c.SignalWorkflow(ctx, workflow.ConvertImageWorkflowOptions.ID, "", "message", "works!")
	}
}
