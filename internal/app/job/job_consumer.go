package job

import (
	"context"
	"encoding/json"
	"errors"
	"go_pubsub_job/internal/domain/job"
	"log"

	"cloud.google.com/go/pubsub"
)

var (
	ErrNoJob = errors.New("no new job to process")
)

type PubSubJobConsumer struct {
	Client       pubsub.Client
	Subscription pubsub.Subscription
}

func (c *PubSubJobConsumer) Consume(handler job.JobStateHandler) error {
	if err := c.Subscription.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		var _job job.Job

		if err := json.Unmarshal(m.Data, &_job); err != nil {
			log.Printf("Err: %+v\n", err)
			return
		}

		executor, err := MakeJobExecutor(_job)
		if err != nil {
			log.Printf("Err: %+v\n", err)
			return
		}

		result, err := executor.Execute(_job)
		if err != nil {
			log.Printf("Err: %+v\n", err)
			return
		}

		m.Ack()
		handler(result)
	}); err != nil {
		return err
	}

	return nil
}
