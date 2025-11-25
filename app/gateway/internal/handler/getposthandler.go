// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"gateway/internal/logic"
	"gateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetPostLogic(r.Context(), svcCtx)
		resp, err := l.GetPost()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
