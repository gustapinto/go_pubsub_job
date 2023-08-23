package job

import (
	"errors"
	"go_pubsub_job/internal/domain/job"

	"github.com/gocolly/colly/v2"
)

var (
	ErrInvalidUrlParse = errors.New("failed to parse url to string")
)

func MakeJobExecutor(_job job.Job, repository job.JobStateRepository) (job.Executor, error) {
	switch _job.Executor {
	case "SubtitleScrapingExecutor":
		return &SubtitleScrapingExecutor{
			Repository: repository,
		}, nil
	}

	return nil, errors.New("invalid job executor")
}

type SubtitleScrapingExecutor struct {
	Repository job.JobStateRepository
}

func (j *SubtitleScrapingExecutor) Execute(_job job.Job) (job.JobState, error) {
	state := job.NewRunningJobStateFromJob(_job)

	if err := j.Repository.Save(state); err != nil {
		return state, err
	}

	collector := colly.NewCollector()
	subtitles := make([]string, 0)

	collector.OnHTML("h2", func(h *colly.HTMLElement) {
		subtitles = append(subtitles, h.Text)
	})

	url, ok := _job.Variables["url"].(string)
	if !ok {
		return state.FailedWithError(ErrInvalidUrlParse), ErrInvalidUrlParse
	}

	if err := collector.Visit(url); err != nil {
		return state.FailedWithError(err), nil
	}

	return state.SuccessWithData(map[string]any{
		"subtitles": subtitles,
	}), nil
}
