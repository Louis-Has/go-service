package article

import (
	"go-service/internal/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-service/restful/art/internal/logic/article"
	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"
)

func GetArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PathID
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewGetArticleLogic(r.Context(), svcCtx)
		resp, err := l.GetArticle(&req)
		response.Response(w, resp, err)
	}
}
