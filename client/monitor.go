package client

import (
	"context"

	"github.com/ginger-core/errors"
)

type Monitor interface {
	GetMonitorName() string
	CheckHealth(ctx context.Context) errors.Error
}

func (c *client[T]) GetMonitorName() string {
	return "file-client"
}

func (c *client[T]) CheckHealth(ctx context.Context) (err errors.Error) {
	if !c.grpcClient.IsConnected() {
		return errors.New().
			WithTrace("!grpcClient.IsConnected")
	}
	return nil
}
