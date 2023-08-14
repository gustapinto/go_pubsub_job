package job

import "go_pubsub_job/internal/domain/result"

type Executor interface {
	Execute(Job) (result.Result, error)
}

type Job struct {
	Id        uint
	Variables map[string]any
	Executor  string
}
