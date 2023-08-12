package job

import (
	"errors"
	"go_pubsub_job/internal/domain/job"
	"log"
)

func MakeJobExecutor(_job job.Job) (job.Executor, error) {
	switch _job.Executor {
	case "SlaJobExecutor":
		return &SlaJobExecutor{}, nil
	}

	return nil, errors.New("invalid job executor")
}

type SlaJobExecutor struct{}

func (e *SlaJobExecutor) Execute(_job job.Job) error {
	message, exists := _job.Variables["message"]
	if !exists {
		return errors.New("message key must be present in the job variables")
	}

	log.Printf("Job %d: %s", _job.Id, message)

	return nil
}
