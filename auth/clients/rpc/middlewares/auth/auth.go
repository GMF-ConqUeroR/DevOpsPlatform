package auth

import (
	"auth/apps/token"
	"auth/clients/rpc"
	"strings"

	"auth/logger"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/response"
)

// 中间件构造函数
func NewAuthFilter() restful.FilterFunction {
	return NewAuth(rpc.Cli()).Filter
}

type Auth struct {
	c *rpc.Client
}

func NewAuth(c *rpc.Client) *Auth {
	return &Auth{
		c: c,
	}
}

func (a *Auth) auth(req *restful.Request) (*token.Token, error) {
	tk := req.HeaderParameter(token.HeaderAuthKey)

	if tk != "" {
		tkFormat := strings.Split(tk, " ")
		if len(tkFormat) > 1 {
			tk = tkFormat[1]
		}
	}

	// 兼容 Cookie
	if tk == "" {
		ck, err := req.Request.Cookie(token.CookieAuthKey)
		if err != nil {
			logger.L().Error().Msg(err.Error())
		} else {
			tk = ck.Value
		}
	}

	return a.c.Token().ValidateToken(req.Request.Context(), token.NewValidateTokenRequest(tk))

}

// 业务功能
//  1. 需要rpc客户端, 采用全局变量来进行设计
//  2. rpc.ValidateToken 来进行鉴权
//     2.1 token: cookie(web)/header(Authtication)/query(websocket)参数
//     2.2 校验 成功 pass, 失败: 中断流程 返回认证错误(401)
func (a *Auth) Filter(req *restful.Request, res *restful.Response, next *restful.FilterChain) {
	authEnable := false
	// 判断路由是否开启了认证
	// 获取当前匹配到路由, 读取路由的Meta信息
	authv := req.SelectedRoute().Metadata()["auth"]
	if authv != nil {
		if v, ok := authv.(bool); ok {
			authEnable = v
		}
	}

	if authEnable {
		tk, err := a.auth(req)
		if err != nil {
			response.Failed(res, exception.NewUnauthorized("未认证:%s", err))
			return
		}
		// 认证通过后需要把认证通过的信息放到上下文中去
		req.SetAttribute(token.CONTEXT_ATTRIBUTE_KEY, tk)
	}
	next.ProcessFilter(req, res)

}
