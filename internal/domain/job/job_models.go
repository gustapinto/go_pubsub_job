package job

import (
	"time"
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
	StartedAt    time.Time      `json:"started_at"`
	FinishedAt   time.Time      `json:"finished_at"`
	Data         map[string]any `json:"data"`
	Status       string         `json:"status"`
	Error        string         `json:"error"`
	JobExecutor  string         `json:"job_executor"`
	JobVariables map[string]any `json:"job_variables"`
}

func NewRunningJobStateFromJob(_job Job) JobState {
	// TODO - Mover NewRunningJobState aqui

	return JobState{}
}

func (r *JobState) SuccessWithData(data map[string]any) JobState {
	return JobState{
		StartedAt:  r.StartedAt,
		FinishedAt: time.Now(),
		Data:       data,
		Status:     StatusSuccess,
		Error:      "",
	}
}

func (r *JobState) FailedWithError(err error) JobState {
	return JobState{
		StartedAt:  r.StartedAt,
		FinishedAt: time.Now(),
		Data:       nil,
		Status:     StatusFailed,
		Error:      err.Error(),
	}
}
