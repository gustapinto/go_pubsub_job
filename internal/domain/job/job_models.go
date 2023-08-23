package job

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusRunning = "running"
	StatusSuccess = "success"
	StatusFailed  = "failed"
)

type Job struct {
	Id        uint           `json:"id"`
	Variables map[string]any `json:"variables"`
	Executor  string         `json:"executor"`
}

type JobState struct {
	Id         uuid.UUID      `json:"id"`
	StartedAt  time.Time      `json:"started_at"`
	FinishedAt time.Time      `json:"finished_at"`
	Data       map[string]any `json:"data"`
	Status     string         `json:"status"`
	Error      string         `json:"error"`
	Job        Job            `json:"job"`
}

func NewRunningJobStateFromJob(_job Job) JobState {
	return JobState{
		Id:        uuid.New(),
		StartedAt: time.Now(),
		Status:    StatusRunning,
		Job:       _job,
	}
}

func (r *JobState) SuccessWithData(data map[string]any) JobState {
	return JobState{
		Id:         r.Id,
		StartedAt:  r.StartedAt,
		FinishedAt: time.Now(),
		Data:       data,
		Status:     StatusSuccess,
		Error:      r.Error,
		Job:        r.Job,
	}
}

func (r *JobState) FailedWithError(err error) JobState {
	return JobState{

		Id:         r.Id,
		StartedAt:  r.StartedAt,
		FinishedAt: time.Now(),
		Data:       r.Data,
		Status:     StatusFailed,
		Error:      err.Error(),
		Job:        r.Job,
	}
}
