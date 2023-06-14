package handler

import (
	"go-service/internal/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-service/restful/art/internal/logic"
	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"
)

func articlePutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleRes
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewArticlePutLogic(r.Context(), svcCtx)
		resp, err := l.ArticlePut(&req)
		response.Response(w, resp, err)
	}
}
