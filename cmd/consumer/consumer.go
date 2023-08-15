package main

import (
	jobapp "go_pubsub_job/internal/app/job"
	"go_pubsub_job/internal/domain/job"
	"go_pubsub_job/internal/infrastructure/ctx"
	"go_pubsub_job/internal/infrastructure/flag"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	projectName, subscriptionName, err := flag.ConsumerCliFlags()
	if err != nil {
		log.Fatalf("Err: %+v\n", err)
	}

	_ctx, cancel := ctx.NewTimeoutContext()
	defer cancel()

	client, err := pubsub.NewClient(_ctx, projectName)
	if err != nil {
		log.Fatalf("Err: %+v", err)
	}

	consumer := jobapp.PubSubJobConsumer{
		Client:       *client,
		Subscription: *client.Subscription(subscriptionName),
	}

	consumer.Consume(func(r job.Result) {
		log.Printf("Result: %+v", r)
	})
}
