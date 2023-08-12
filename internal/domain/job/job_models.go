package job

type Executor interface {
	Execute(Job) error
}

type Job struct {
	Id        uint
	Variables map[string]any
	Executor  string
}
