package ctx

import (
	"context"
	"time"
)

func NewTimeoutContext() (context.Context, context.CancelFunc) {
	parent := context.Background()
	timeout := 30 * time.Second

	return context.WithTimeout(parent, timeout)
}
