package main

import (
	jobapp "go_pubsub_job/internal/app/job"
	"go_pubsub_job/internal/infrastructure/ctx"
	"go_pubsub_job/internal/infrastructure/flag"
	"log"

	"cloud.google.com/go/firestore"
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

	_ctx, cancel = ctx.NewTimeoutContext()
	defer cancel()

	firestoreClient, err := firestore.NewClient(_ctx, projectName)
	if err != nil {
		log.Fatalf("Err: %+v", err)
	}

	repository := &jobapp.JobStateFirestoreRepository{
		Client: firestoreClient,
	}
	consumer := jobapp.PubSubJobConsumer{
		Client:       *client,
		Subscription: *client.Subscription(subscriptionName),
		Repository:   repository,
	}
	jobService := jobapp.JobService{
		Consumer:   &consumer,
		Repository: repository,
	}

	if err := jobService.RunJobs(); err != nil {
		log.Fatalf("Err: %+v", err)
	}
}
