package job

import (
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

	_ctx, cancel := ctx.NewTimeoutContext()
	defer cancel()

	p.Topic.Publish(_ctx, &pubsub.Message{
		Data: jobJson,
	})

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
