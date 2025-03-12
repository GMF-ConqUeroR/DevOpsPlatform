package api

import (
	"fmt"

	"auth/apps/token"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"
)

func (h *Handler) QuerySecret(r *restful.Request, w *restful.Response) {
	tk := r.Attribute(token.CONTEXT_ATTRIBUTE_KEY)
	fmt.Println(tk)
	response.Success(w, tk)
}
