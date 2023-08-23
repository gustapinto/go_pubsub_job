package job

import (
	"fmt"
	"go_pubsub_job/internal/domain/job"
	"log"
)

type JobService struct {
	Publisher  job.Publisher
	Consumer   job.Consumer
	Repository job.JobStateRepository
}

func (js *JobService) errMissingField(field string) error {
	return fmt.Errorf("field %s must not be empty or nil", field)
}

func (js *JobService) InitDatabase() error {
	if js.Repository == nil {
		return js.errMissingField("Repository")
	}

	return js.Repository.Init()
}

func (js *JobService) RunJobs() error {
	if js.Consumer == nil {
		return js.errMissingField("Consumer")
	}

	return js.Consumer.Consume(func(r job.JobState) {
		log.Printf("Result: %+v\n", r)
	})
}
