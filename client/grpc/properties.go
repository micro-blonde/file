package grpc

import (
	"context"

	"github.com/micro-blonde/file/proto/file"
	"google.golang.org/grpc"
)

func (c *client) GetProperties(ctx context.Context,
	in *file.PropertiesRequest, opts ...grpc.CallOption) (*file.Properties, error) {
	if err := c.ensureService(); err != nil {
		return nil, err
	}
	return c.client.GetProperties(ctx, in, opts...)
}
