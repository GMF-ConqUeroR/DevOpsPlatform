package auth

import (
	"auth/apps/token"
	"auth/clients/rpc"
	"fmt"
	"strings"

	"auth/logger"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/response"
)

// 中间件构造函数
func NewAuthFilter(svc string) restful.FilterFunction {
	return NewAuth(rpc.Cli(), svc).Filter
}

type Auth struct {
	c       *rpc.Client
	service string
}

func NewAuth(cli *rpc.Client, svc string) *Auth {
	return &Auth{
		c:       cli,
		service: svc,
	}
}

func (a *Auth) auth(req *restful.Request) (*token.Token, error) {
	tk := req.HeaderParameter(token.HeaderAuthKey)

	if tk != "" {
		tkFormat := strings.Split(tk, " ")
		if len(tkFormat) > 1 {
			tk = tkFormat[1]
			fmt.Printf("checking:%v", tk)
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
	fmt.Printf("checking:%v", req.HeaderParameter(token.HeaderAuthKey))
	t, err := a.c.Token().ValidateToken(req.Request.Context(), token.NewValidateTokenRequest(tk))
	if err != nil {
		return nil, err
	}
	return t, nil

}

// 业务功能
//  1. 需要rpc客户端, 采用全局变量来进行设计
//  2. rpc.ValidateToken 来进行鉴权
//     2.1 token: cookie(web)/header(Authtication)/query(websocket)参数
//     2.2 校验 成功 pass, 失败: 中断流程 返回认证错误(401)
func (a *Auth) Filter(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {
	authEnable := false
	// 判断路由是否开启了认证
	// 获取当前匹配到路由, 读取路由的Meta信息
	meta := req.SelectedRoute().Metadata()
	authv := meta["auth"]
	if authv != nil {
		if v, ok := authv.(bool); ok {
			authEnable = v
		}
	}

	if authEnable {
		tk, err := a.auth(req)
		if err != nil {
			response.Failed(resp, exception.NewUnauthorized("未认证:%s", err))
			return
		}
		logger.L().Info().Msgf("welcome %v ", tk.Username)
		// 认证通过后 再判断是否是否有权限
		permissionEnable := false
		permv := meta["permission"]
		if permv != nil {
			if v, ok := permv.(bool); ok {
				permissionEnable = v
			}
		}
		logger.L().Debug().Msgf("check the access:%v", permissionEnable)
		if permissionEnable {
			resource, action := "", ""

			resourceValue := meta["resource"]
			if resourceValue != nil {
				resource = resourceValue.(string)
			}

			actionValue := meta["action"]
			if actionValue != nil {
				action = actionValue.(string)
			}
			// 根据用户角色判断 其收否有权限访问 相应的API
			logger.L().Debug().Msgf("check %v is role %v or not", tk.Username, tk.Roles)
			err := tk.HavePermissionOrNot(a.service, resource, action)
			if err != nil {
				response.Failed(resp, exception.NewPermissionDeny(err.Error()))
				return
			}

		}

		// 认证，授权通过的信息放到上下文中去
		req.SetAttribute(token.CONTEXT_ATTRIBUTE_KEY, tk)
	}
	next.ProcessFilter(req, resp)

}
