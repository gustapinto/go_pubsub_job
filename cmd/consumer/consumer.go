package main

import (
	jobapp "go_pubsub_job/internal/app/job"
	"go_pubsub_job/internal/domain/job"
	"log"
)

func main() {
	results := make([]job.Result, 0)
	jobs := job.NewSubtitleScrapingJobFromUrls([]string{
		"https://g1.globo.com/",
		"https://www.bbc.com/portuguese",
	})

	for _, j := range jobs {
		executor, err := jobapp.MakeJobExecutor(j)
		if err != nil {
			log.Printf("Err: %+v\n", err)
			continue
		}

		result, err := executor.Execute(j)
		if err != nil {
			log.Printf("Err: %+v\n", err)
			continue
		}

		results = append(results, result)

		log.Printf("%+v\n", result)
	}
}
