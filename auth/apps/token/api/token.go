package api

import (
	"auth/apps/token"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"
)

func (h *Handler) IssueToken(r *restful.Request, w *restful.Response) {
	// gin Bind()
	in := token.NewIssueTokenRequest()
	err := r.ReadEntity(in)
	if err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.svc.IssueToken(r.Request.Context(), in)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 用隐式接口 Desense() 来完成自动脱敏
	response.Success(w, ins)
}
