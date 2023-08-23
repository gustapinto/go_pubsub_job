package job

import (
	"time"

	"github.com/google/uuid"
)

func NewSubtitleScrapingJobFromUrls(urls []string) []Job {
	jobs := make([]Job, 0)

	for _, url := range urls {
		jobs = append(jobs, NewSubtitleScrapingJob(url))
	}

	return jobs
}

func NewSubtitleScrapingJob(url string) Job {
	return Job{
		Id:       uuid.New(),
		Executor: "SubtitleScrapingExecutor",
		Variables: map[string]any{
			"url": url,
		},
	}
}

func NewRunningJobState() JobState {
	return JobState{
		StartedAt: time.Now(),
		Status:    StatusRunning,
	}
}
