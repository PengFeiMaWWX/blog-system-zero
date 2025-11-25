// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"gateway/internal/logic"
	"gateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPostsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetPostsLogic(r.Context(), svcCtx)
		resp, err := l.GetPosts()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
