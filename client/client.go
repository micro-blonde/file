package client

import (
	"context"
	"time"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/log/logger"
	"github.com/micro-blonde/file"
	"github.com/micro-blonde/file/client/grpc"
	pf "github.com/micro-blonde/file/proto/file"
)

type Client[T file.Model] interface {
	Monitor
	Initialize()
	GetGrpcClient() grpc.Client

	Store(ctx context.Context,
		request *file.StoreRequest) (*file.StoreResponse[T], errors.Error)
	StoreFromRequest(request gateway.Request, key string,
		storeRequest *file.StoreRequest) (*file.StoreResponse[T], errors.Error)

	GetDownloadUrlByKey(key string) (string, errors.Error)
}

type client[T file.Model] struct {
	logger log.Logger
	config config

	grpcClient grpc.Client

	properties *pf.Properties
}

func New[T file.Model](logger log.Logger, registry registry.Registry) Client[T] {
	var config config
	if err := registry.Unmarshal(&config); err != nil {
		panic(err)
	}
	config.Initialize()
	c := &client[T]{
		logger:     logger,
		config:     config,
		grpcClient: grpc.New(logger, registry),
	}
	return c
}

func (c *client[T]) GetGrpcClient() grpc.Client {
	return c.grpcClient
}

func (c *client[T]) Initialize() {
	if !c.config.enabled {
		return
	}
	c.grpcClient.Initialize()
	go c.loop()
}

func (c *client[T]) loop() {
	for {
		if !c.grpcClient.IsReady() &&
			!c.grpcClient.IsConnecting() {
			c.properties = nil
			err := c.grpcClient.Reconnect()
			if err != nil {
				c.logger.
					With(logger.Field{
						"error": err.Error(),
					}).
					WithTrace("grpcClient.Reconnect").
					Errorf("error on reconnect")
			}
		}
		if c.properties == nil && c.grpcClient.IsReady() {
			properties, err := c.grpcClient.GetProperties(
				context.Background(), &pf.PropertiesRequest{})
			if err != nil {
				c.logger.
					With(logger.Field{
						"error": err.Error(),
					}).
					WithTrace("grpcClient.GetProperties").
					Errorf("error on get properties")
			} else {
				c.properties = properties
			}
		}
		time.Sleep(time.Second)
	}
}
