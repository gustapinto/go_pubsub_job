package main

import (
	jobapp "go_pubsub_job/internal/app/job"
	"go_pubsub_job/internal/domain/job"
	"log"
)

func main() {
	jobs := []job.Job{
		{
			Id:       1,
			Executor: "SlaJobExecutor",
			Variables: map[string]any{
				"message": "Hello World!",
			},
		},
		{
			Id:        2,
			Executor:  "SlaJobExecutor",
			Variables: map[string]any{},
		},
		{
			Id:       3,
			Executor: "SlaJobExecutor",
			Variables: map[string]any{
				"message": "Hello World!",
			},
		},
		{
			Id:       4,
			Executor: "foobar",
			Variables: map[string]any{
				"message": "Hello World!",
			},
		},
	}

	for _, j := range jobs {
		executor, err := jobapp.MakeJobExecutor(j)
		if err != nil {
			log.Printf("Err: %+v\n", err)
			continue
		}

		if err := executor.Execute(j); err != nil {
			log.Printf("Err: %+v\n", err)
			continue
		}
	}
}
