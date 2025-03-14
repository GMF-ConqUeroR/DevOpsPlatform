package tools

import (
	"context"

	"auth/apps/endpoint"
	"auth/clients/rpc"
	"auth/logger"

	"github.com/emicklei/go-restful/v3"
)

// 业务
// 1. 先获取当前路由的所有条目
// 2. 转换成需要注册的 数据结构
// 3. 调用rpc 进行注册
type EndpointRegister struct {
	c       *rpc.Client
	service string
}

func NewEndpointRegister(service string) *EndpointRegister {
	return &EndpointRegister{
		c:       rpc.Cli(),
		service: service,
	}
}

// gorestful 的跟理由传递给过滤
func (r *EndpointRegister) Registry(ctx context.Context, c *restful.Container) error {
	set := endpoint.NewEndpointSpecSet()

	wss := c.RegisteredWebServices()
	for i := range wss {
		ws := wss[i]
		rs := ws.Routes()
		for i := range rs {
			route := rs[i]
			//
			e := endpoint.NewEndpointSpec(
				r.service,
				route.Method+"."+route.Path,
			)

			meta := route.Metadata
			e.Resoruce = endpoint.GetValueFromMeta(meta, "resource")
			e.Action = endpoint.GetValueFromMeta(meta, "action")
			set.Add(e)
			logger.L().Debug().Msgf("----%v", e)
		}
	}

	_, err := r.c.Endpoint().RegistryEndpoints(ctx, set)
	return err
}
