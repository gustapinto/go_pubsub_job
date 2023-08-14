package main

import (
	"flag"
	jobapp "go_pubsub_job/internal/app/job"
	"go_pubsub_job/internal/domain/job"
	"go_pubsub_job/internal/infrastructure/ctx"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	projectName := flag.String("project", "", "The Google Cloud Project Id")
	topicName := flag.String("topic", "", "The Pub/Sub Topic Name")
	flag.Parse()

	if *projectName == "" {
		log.Fatal("Please specify the Google Cloud Project Id!")
	}

	if *topicName == "" {
		log.Fatal("Please specify the Pub/Sub Topic Name")
	}

	jobs := job.NewSubtitleScrapingJobFromUrls([]string{
		"https://g1.globo.com/",
		"https://www.bbc.com/portuguese",
	})

	_ctx, cancel := ctx.NewTimeoutContext()
	defer cancel()

	client, err := pubsub.NewClient(_ctx, *projectName)
	if err != nil {
		log.Fatalf("Err: %+v", err)
	}

	publisher := jobapp.PubSubJobPublisher{
		Client: *client,
		Topic:  *client.Topic(*topicName),
	}
	published, errors := publisher.PublishBatch(jobs)

	log.Printf("Published: %d", published)
	log.Printf("Errors: %d", len(errors))
}
