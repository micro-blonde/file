package client

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginger-core/errors"
	errorsgrpc "github.com/ginger-core/errors/grpc"
	"github.com/ginger-core/gateway"
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

func (c *client[T]) StoreFromRequest(request gateway.Request, key string,
	storeRequest *file.StoreRequest) (*file.StoreResponse[T], errors.Error) {
	ctx := request.GetContext()
	ginCtx := request.GetConn().(*gin.Context)
	f, fErr := ginCtx.FormFile(key)
	if fErr != nil {
		if fErr == http.ErrMissingFile {
			return nil, errors.Validation(fErr).
				WithId("FileRequiredError").
				WithMessage("file is missing to process your request")
		}
		return nil, errors.Validation(fErr)
	}
	uploadedFile, oErr := f.Open()
	if oErr != nil {
		return nil, errors.Validation(oErr).
			WithTrace("file.Open")
	}
	data := make([]byte, f.Size)
	_, rErr := uploadedFile.Read(data)
	if rErr != nil {
		return nil, errors.New(rErr).
			WithTrace("uploadedFile.Read")
	}
	storeRequest.Data = data
	storeRequest.Name = f.Filename
	return c.Store(ctx, storeRequest)
}
