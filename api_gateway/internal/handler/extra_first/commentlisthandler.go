package extra_first

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tiny-tiktok/api_gateway/internal/logic/extra_first"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
)

func CommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := extra_first.NewCommentListLogic(r.Context(), svcCtx)
		resp, err := l.CommentList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
