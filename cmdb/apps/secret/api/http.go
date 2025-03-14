package api

import (
	"cmdb/apps/secret"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/ioc"
)

type Handler struct {
	ioc.ObjectImpl
}

func (h *Handler) Init() error {
	return nil
}

func (h *Handler) Name() string {
	return secret.AppName
}

func (h *Handler) Registry(r *restful.WebService) {
	// api doc中的分类标签
	tags := []string{"凭证管理"}
	// 怎么控制哪些接口认证，哪些接口不认证
	r.Route(r.GET("/").To(h.QuerySecret).Doc("查询凭证列表").Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("auth", true).
		Metadata("permission", true).
		Metadata("resource", "Secret").
		Metadata("action", "list"))
}

func init() {
	ioc.RegistryApi(&Handler{})
}
