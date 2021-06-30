package util

import (
	"context"
	"time"
)

const (
	coordinatorTimeout = time.Second
)

func GetTimeoutContext() context.Context {
	ctx, _ := context.WithTimeout(context.TODO(), coordinatorTimeout)
	return ctx
}
