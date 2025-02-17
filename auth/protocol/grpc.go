package protocol

import (
	"auth/conf"
	"auth/logger"
	"context"
	"net"

	"github.com/infraboard/mcube/ioc"
	"google.golang.org/grpc"
)

type Grpc struct {
	server *grpc.Server
}

func NewGrpc() *Grpc {
	return &Grpc{
		server: grpc.NewServer(),
	}
}

func (g *Grpc) Start() error {
	ioc.LoadGrpcController(g.server)
	logger.L().Debug().Msgf("loaded controllers: %s", ioc.ListControllerObjectNames())

	lis, err := net.Listen("tcp", conf.C().Grpc.Address())
	if err != nil {
		return err
	}

	logger.L().Debug().Msgf("lunch GRPC server: %s successfully!", conf.C().Grpc.Address())

	return g.server.Serve(lis)
}

func (g *Grpc) Stop(ctx context.Context) {
	g.server.GracefulStop()
}
