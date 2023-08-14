package result

import "time"

func NewRunningResult() Result {
	return Result{
		StartedAt: time.Now(),
		Status:    StatusRunning,
	}
}
