#temporal-rabbit-poc

This is an example of how temporal can be used in tandem with a message queue such as Rabbit MQ.

In this example, a client sends a request to a server. The server passes that request forward to a temporal workflow, that will process it and wait for a rabbit mq message to come through before it completes. How can we achieve that?

We need three long-running processes: A temporal worker that processes temporal workflows, a server to listen for the client requests, and a consumer for the rabbit MQ messages.

## running
To see it in action, you will need four different terminal windows. In the first one, start your rabbitmq and temporal servees with:
`docker-compose up -d`

After that, start your temporal worker that will wait for workflow executions by doing:
`go run ./cmd/worker/main.go`

Then you will need to start your server:
`go run ./cmd/server/main.go`

After that, use the client command to send an http message into the server and start a workflow execution:
`go run ./cmd/client/main.go`

You should be able to see this newly created workflow execution on your temporal dashboard by going to `localhost:8088`. Finally, to complete the execution, you'd want to push a message into Rabbit by doing:
`go run ./cmd/producer/main.go`

And that's it! Of course this example is void of any actually useful business logic, but it illustrates a flow that could be used in real life by many developers.

