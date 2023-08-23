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
	Id        uuid.UUID      `json:"id"`
	Variables map[string]any `json:"variables"`
	Executor  string         `json:"executor"`
}

func (j *Job) Doc() string {
	return j.Id.String()
}

func (j *Job) ToMap() map[string]any {
	return map[string]any{
		"id":        j.Id.String(),
		"variables": j.Variables,
		"executor":  j.Executor,
	}
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

func (js *JobState) Doc() string {
	return js.Id.String()
}

func (js *JobState) ToMap() map[string]any {
	return map[string]any{
		"id":          js.Id.String(),
		"started_at":  js.StartedAt.Format("2006-01-02 15:04:05"),
		"finished_at": js.FinishedAt.Format("2006-01-02 15:04:05"),
		"data":        js.Data,
		"status":      js.Status,
		"error":       js.Error,
		"job":         js.Job.ToMap(),
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
