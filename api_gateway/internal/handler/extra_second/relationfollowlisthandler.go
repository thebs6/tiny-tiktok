package extra_second

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"tiny-tiktok/api_gateway/internal/logic/extra_second"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
)

func RelationFollowListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RelationFollowListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := extra_second.NewRelationFollowListLogic(r.Context(), svcCtx)
		resp, err := l.RelationFollowList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
