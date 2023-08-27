package core

import (
	"net/http"

	"tiny-tiktok/api_gateway/internal/logic/core"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	defaultMultipartMemory = 32 << 20 // 32 MB
)

func PublishActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishActionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if err := r.ParseMultipartForm(defaultMultipartMemory); err != nil {
			logc.Info(r.Context(), "ParseMultipartForm failed", err)
			httpx.Error(w, err)
			return
		}

		l := core.NewPublishActionLogic(r.Context(), svcCtx)
		l.File = r.MultipartForm.File["data"][0]

		resp, err := l.PublishAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
