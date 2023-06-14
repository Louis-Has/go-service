package handler

import (
	"go-service/internal/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-service/restful/art/internal/logic"
	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"
)

func articlePostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Article
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewArticlePostLogic(r.Context(), svcCtx)
		resp, err := l.ArticlePost(&req)
		response.Response(w, resp, err)
	}
}
