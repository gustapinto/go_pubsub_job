package job

import "time"

var (
	jobCount = 0
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
		Id:       uint(jobCount),
		Executor: "SubtitleScrapingExecutor",
		Variables: map[string]any{
			"url": url,
		},
	}
}

func NewRunningResult() Result {
	return Result{
		StartedAt: time.Now(),
		Status:    StatusRunning,
	}
}
