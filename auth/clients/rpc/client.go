package rpc

import (
	"auth/apps/token"
	"auth/conf"
	"auth/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	cli *grpc.ClientConn
}

func NewClient(c conf.Grpc) *Client {
	logger.L().Debug().Msg(c.Address())
	conn, err := grpc.NewClient(c.Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &Client{
		cli: conn,
	}
}

func (c *Client) Token() token.RPCClient {
	return token.NewRPCClient(c.cli)
}
