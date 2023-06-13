package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Success bool        `json:"success"`
	Err     string      `json:"err"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Err = err.Error()
	} else {
		body.Success = true
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
