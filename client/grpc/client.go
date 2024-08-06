package grpc

import (
	"time"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/file/proto/file"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	file.ServiceClient
	Initialize()
	IsReady() bool
	IsConnecting() bool
	Reconnect() errors.Error
}

type client struct {
	logger log.Logger
	config config

	conn        *grpc.ClientConn
	dialOptions []grpc.DialOption
	client      file.ServiceClient
}

func New(logger log.Logger, registry registry.Registry) Client {
	c := &client{
		logger: logger,
		dialOptions: []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(32 * 10e6)),
		},
	}
	if err := registry.Unmarshal(&c.config); err != nil {
		panic(err)
	}
	c.config.Initialize()
	return c
}

func (c *client) Initialize() {
	//var state = connectivity.Shutdown
	//for c.conn == nil || c.conn.GetState() != connectivity.Ready {
	//	if c.conn != nil {
	//		state = c.conn.GetState()
	//	}
	//	c.logger.
	//		With(log.Field{
	//			"state": state,
	//		}).
	//		Infof("connection is not ready yet, waiting...")
	//	time.Sleep(time.Second)
	//}
}

func (c *client) ensureService() errors.Error {
	if c.conn == nil {
		return errors.Internal().
			WithId("ServiceUnavailable").
			WithMessage("Service is unavailable right now")
	}
	return nil
}

func (c *client) getState() connectivity.State {
	if c.conn == nil {
		return connectivity.Idle
	}
	return c.conn.GetState()
}

func (c *client) isState(state connectivity.State) bool {
	if c.conn == nil {
		return false
	}
	return state == c.conn.GetState()
}

func (c *client) IsReady() bool {
	return c.isState(connectivity.Ready)
}

func (c *client) IsConnecting() bool {
	if c.conn == nil {
		return false
	}
	state := c.conn.GetState()
	return state == connectivity.Connecting
}

func (c *client) Reconnect() errors.Error {
	if c.conn != nil {
		c.conn.Close()
	}
	c.logger.Infof("connecting to server to address %s...", c.config.Address)
	var err error
	c.conn, err = grpc.NewClient(c.config.Address, c.dialOptions...)
	if err != nil {
		c.logger.Infof("connection could not be established, retrying...")
		time.Sleep(time.Second)
		return errors.New(err).
			WithTrace("grpc.DialContext")
	}
	c.conn.Connect()
	c.logger.Infof("state: %s", c.conn.GetState().String())
	c.client = file.NewServiceClient(c.conn)
	return nil
}
