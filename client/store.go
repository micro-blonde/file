package client

import (
	"context"

	"github.com/ginger-core/errors"
	errorsgrpc "github.com/ginger-core/errors/grpc"
	"github.com/micro-blonde/file"
	pf "github.com/micro-blonde/file/proto/file"
)

func (c *client[T]) Store(ctx context.Context,
	request *file.StoreRequest) (*file.StoreResponse[T], errors.Error) {
	if !c.config.enabled {
		if !c.config.enabled {
			return nil, notEnabledError
		}
	}
	req := &pf.StoreRequest{
		Category: request.Category,
		Type:     request.Type,
		Data:     request.Data,
		Name:     request.Name,
	}
	if request.BaseDir != nil {
		req.BaseDir = *request.BaseDir
	}
	if request.ExpiresIn != nil {
		req.ExpiresInSecs = int64(request.ExpiresIn.Seconds())
	}
	if request.AccountId != nil {
		req.AccountId = *request.AccountId
	}

	r, err := c.grpcClient.Store(ctx, req)
	if err != nil {
		return nil, errorsgrpc.Parse(err).
			WithTrace("Store.grpcClient.Store.Err")
	}

	resp := &file.StoreResponse[T]{
		File: getFile[T](r.File),
		Url:  r.Url,
	}
	return resp, nil
}
