package protocol

import (
	authClient "auth/clients/rpc/middlewares/auth"
	"auth/clients/rpc/tools"
	"cmdb/conf"
	"cmdb/logger"
	"cmdb/swagger"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/apps/apidoc"
)

type Http struct {
	server     *http.Server
	r          *restful.Container
	apiDocPath string
}

func NewHttp() *Http {
	r := restful.DefaultContainer

	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)

	// CORS中间件
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"*"},
		AllowedDomains: []string{"*"},
		AllowedMethods: []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"},
		CookiesAllowed: false,
		Container:      r,
	}

	// 使用Filter加载 GoRestful中间件
	r.Filter(cors.Filter)

	// 接入用户中心中间件, 加载全局变量 auth client
	// 需要提前加载该client, 由于有默认配置:
	r.Filter(authClient.NewAuthFilter("cmdb"))

	return &Http{
		r:          r,
		apiDocPath: "/apidocs.json",
		server: &http.Server{
			// 服务器监听地址
			Addr: conf.C().Http.Address(),
			// 配置http server使用的路由
			Handler: r,
			//  client ---> server
			ReadTimeout: 5 * time.Second,
			// server ---> client
			WriteTimeout: 5 * time.Second,
		},
	}
}

// 设置API的 URL 前缀
func (h *Http) PathPrefix() string {
	return fmt.Sprintf("/%s/api", conf.C().AppName.Name)
}

// 启动 HTTP Server
func (h *Http) Start() error {

	// LoadGinApi 基于 gin 的restful handler 实例的托管对象, 需要通过gin router来暴露restful接口
	// LoadGoRestfulApi 基于 gorestful http框架的托管对象
	// LoadGrpcController 实现了Grpc Server的对象 需要注册到中央的协议服务 grpc server

	// 加载子服务路由
	ioc.LoadGoRestfulApi(h.PathPrefix(), h.r)
	logger.L().Debug().Msgf("loaded apis: %s successfully!", ioc.ListApiObjectNames())
	logger.L().Debug().Msgf("launch HTTP server: %s successfully!", conf.C().Http.Address())

	h.r.Add(apidoc.APIDocs(h.apiDocPath, swagger.Docs))
	logger.L().Debug().Msgf("Swagger API Doc访问地址: http://%s%s", conf.C().Http.Address(), h.apiDocPath)

	// 所有的路由都已经加载好了, 获取当前DefaultContainer 里面注册的所有的WebService(一个service 就是一个模块)
	r := tools.NewEndpointRegister("cmdb")
	err := r.Registry(context.Background(), h.r)
	if err != nil {
		logger.L().Error().Msgf("registry endpoint error, %s", err)
	}

	return h.server.ListenAndServe()
}

func (h *Http) Stop(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
