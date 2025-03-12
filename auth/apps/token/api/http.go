package api

import (
	"auth/apps/token"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/ioc"
)

type Handler struct {
	svc token.Service
	ioc.ObjectImpl
}

func (h *Handler) Init() error {
	h.svc = ioc.GetController(token.AppName).(token.Service)
	return nil
}

func (h *Handler) Name() string {
	return token.AppName
}

func (h *Handler) Registry(r *restful.WebService) {
	// api doc中的分类标签
	tags := []string{"登录"}

	r.Route(r.POST("/").To(h.IssueToken).Doc("颁发令牌").
		// 作为OpenApi的值作为展示
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.IssueTokenRequest{}).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))

}

func init() {
	ioc.RegistryApi(&Handler{})
}
