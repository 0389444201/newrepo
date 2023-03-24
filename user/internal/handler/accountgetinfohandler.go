package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/user/internal/logic"
	"user/user/internal/svc"
	"user/user/internal/types"
)

func AccountGetInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAccountGetInfoLogic(r.Context(), svcCtx)
		resp, err := l.AccountGetInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
