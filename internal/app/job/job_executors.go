package job

import (
	"errors"
	"go_pubsub_job/internal/domain/job"
	"go_pubsub_job/internal/domain/result"

	"github.com/gocolly/colly/v2"
)

var (
	ErrInvalidUrlParse = errors.New("failed to parse url to string")
)

func MakeJobExecutor(_job job.Job) (job.Executor, error) {
	switch _job.Executor {
	case "SubtitleScrapingExecutor":
		return &SubtitleScrapingExecutor{}, nil
	}

	return nil, errors.New("invalid job executor")
}

type SubtitleScrapingExecutor struct{}

func (j *SubtitleScrapingExecutor) Execute(_job job.Job) (result.Result, error) {
	jobResult := result.NewRunningResult()
	collector := colly.NewCollector()
	subtitles := make([]string, 0)

	collector.OnHTML("h2", func(h *colly.HTMLElement) {
		subtitles = append(subtitles, h.Text)
	})

	url, ok := _job.Variables["url"].(string)
	if !ok {
		return jobResult.FailedWithError(ErrInvalidUrlParse), ErrInvalidUrlParse
	}

	if err := collector.Visit(url); err != nil {
		return jobResult.FailedWithError(err), nil
	}

	return jobResult.SuccessWithData(map[string]any{
		"subtitles": subtitles,
	}), nil
}
