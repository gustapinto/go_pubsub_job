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
	projectName, topicName, err := flag.PublisherCliFlags()
	if err != nil {
		log.Fatalf("Err: %+v\n", err)
	}

	jobs := job.NewSubtitleScrapingJobFromUrls([]string{
		"https://g1.globo.com/",
		"https://www.bbc.com/portuguese",
	})

	_ctx, cancel := ctx.NewTimeoutContext()
	defer cancel()

	client, err := pubsub.NewClient(_ctx, projectName)
	if err != nil {
		log.Fatalf("Err: %+v", err)
	}

	topic := client.Topic(topicName)
	topic.PublishSettings.CountThreshold = 0

	publisher := jobapp.PubSubJobPublisher{
		Client: *client,
		Topic:  *topic,
	}
	published, errors := publisher.PublishBatch(jobs)

	log.Printf("Published: %d", published)
	log.Printf("Errors: %d", len(errors))
}
