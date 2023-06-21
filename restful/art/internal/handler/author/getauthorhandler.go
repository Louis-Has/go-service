package author

import (
	"go-service/internal/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-service/restful/art/internal/logic/author"
	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"
)

func GetAuthorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PathID
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := author.NewGetAuthorLogic(r.Context(), svcCtx)
		resp, err := l.GetAuthor(&req)
		response.Response(w, resp, err)
	}
}
