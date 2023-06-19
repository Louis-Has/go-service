package handler

import (
	"go-service/internal/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-service/restful/art/internal/logic"
	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"
)

func articleGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PathID
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewArticleGetLogic(r.Context(), svcCtx)
		resp, err := l.ArticleGet(&req)
		response.Response(w, resp, err)
	}
}
