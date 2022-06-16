package client

import (
	"context"
	"time"

	"github.com/mjlee1983/go-micro/go-micro/v2/util/backoff"
)

type BackoffFunc func(ctx context.Context, req Request, attempts int) (time.Duration, error)

func exponentialBackoff(ctx context.Context, req Request, attempts int) (time.Duration, error) {
	return backoff.Do(attempts), nil
}
