package job

import (
	"context"
	"encoding/json"
	"go_pubsub_job/internal/domain/job"
	"go_pubsub_job/internal/infrastructure/ctx"

	"cloud.google.com/go/pubsub"
)

type PubSubJobPublisher struct {
	Client pubsub.Client
	Topic  pubsub.Topic
}

func (p *PubSubJobPublisher) Publish(_job job.Job) error {
	jobJson, err := json.Marshal(_job)
	if err != nil {
		return err
	}

	res := p.Topic.Publish(context.Background(), &pubsub.Message{
		Data: jobJson,
	})

	_ctx, cancel := ctx.NewTimeoutContext()
	defer cancel()

	if _, err := res.Get(_ctx); err != nil {
		return err
	}

	return nil
}

func (p *PubSubJobPublisher) PublishBatch(jobs []job.Job) (publishCount uint, publishErrors []error) {
	for _, job := range jobs {
		if err := p.Publish(job); err != nil {
			publishErrors = append(publishErrors, err)
			continue
		}

		publishCount++
	}

	return publishCount, publishErrors
}
