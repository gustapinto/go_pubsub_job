package job

type Executor interface {
	Execute(Job) (Result, error)
}

type Publisher interface {
	Publish(Job) error

	PublishBatch([]Job) (uint, []error)
}

type Consumer interface {
	Consume(func(Result)) error
}
