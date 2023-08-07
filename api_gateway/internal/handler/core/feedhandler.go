package core

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tiny-tiktok/api_gateway/internal/logic/core"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
)

func FeedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FeedReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := core.NewFeedLogic(r.Context(), svcCtx)
		resp, err := l.Feed(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
