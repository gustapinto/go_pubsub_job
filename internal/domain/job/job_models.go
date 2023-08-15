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

type Result struct {
	StartedAt  time.Time      `json:"started_at"`
	FinishedAt time.Time      `json:"finished_at"`
	Data       map[string]any `json:"data"`
	Status     string         `json:"status"`
	Error      string         `json:"error"`
}

func (r *Result) SuccessWithData(data map[string]any) Result {
	return Result{
		StartedAt:  r.StartedAt,
		FinishedAt: time.Now(),
		Data:       data,
		Status:     StatusSuccess,
		Error:      "",
	}
}

func (r *Result) FailedWithError(err error) Result {
	return Result{
		StartedAt:  r.StartedAt,
		FinishedAt: time.Now(),
		Data:       nil,
		Status:     StatusFailed,
		Error:      err.Error(),
	}
}
