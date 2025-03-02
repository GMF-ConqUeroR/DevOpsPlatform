package rpc_test

import (
	"auth/apps/token"
	"auth/clients/rpc"
	"auth/conf"
	"context"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
)

var (
	ctx     = context.Background()
	clients = rpc.NewClient(*conf.NewDefaultGrpc())
)

// SDK 使用方式， 这部分代码是在其他服务里面引入使用的
func TestValidateToken(t *testing.T) {
	in := &token.ValidateTokenRequest{
		AccessToken: "cv1fv187j4glpff8p1cg",
	}
	ins, err := clients.Token().ValidateToken(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	// debug 的开关
	zap.DevelopmentSetup()
}
