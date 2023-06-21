package article

import (
	"go-service/internal/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-service/restful/art/internal/logic/article"
	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"
)

func PostArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Article
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewPostArticleLogic(r.Context(), svcCtx)
		resp, err := l.PostArticle(&req)
		response.Response(w, resp, err)
	}
}
