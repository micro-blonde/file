package grpc

import (
	"context"

	"github.com/micro-blonde/file/proto/file"
	"google.golang.org/grpc"
)

func (c *client) Store(ctx context.Context,
	in *file.StoreRequest, opts ...grpc.CallOption) (*file.StoreResponse, error) {
	if err := c.ensureService(); err != nil {
		return nil, err
	}
	return c.client.Store(ctx, in, opts...)
}
