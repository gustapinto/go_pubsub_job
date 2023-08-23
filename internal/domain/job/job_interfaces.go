package job

type Executor interface {
	Execute(Job) (JobState, error)
}

type Publisher interface {
	Publish(Job) error

	PublishBatch([]Job) (uint, []error)
}

type JobStateHandler func(JobState)

type Consumer interface {
	Consume(JobStateHandler) error
}

type JobStateRepository interface {
	Save(JobState) error
}
