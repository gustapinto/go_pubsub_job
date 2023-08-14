package result

import (
	"time"
)

const (
	StatusRunning = "running"
	StatusSuccess = "success"
	StatusFailed  = "failed"
)

type Result struct {
	StartedAt  time.Time
	FinishedAt time.Time
	Data       map[string]any
	Status     string
	Error      string
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
