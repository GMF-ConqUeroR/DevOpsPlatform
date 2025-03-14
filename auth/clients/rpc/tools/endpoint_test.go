package tools_test

import (
	"auth/clients/rpc/tools"
	"context"
	"testing"

	"github.com/emicklei/go-restful/v3"
)

func TestEndpointRegistry(t *testing.T) {
	r := tools.NewEndpointRegister("cmdb")
	r.Registry(context.Background(), restful.DefaultContainer)
}
