package handler

import (
	"go-service/restful/art/internal/logic"
	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func articleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetArticleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewArticleLogic(r.Context(), svcCtx)
		resp, err := l.Article(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
