package logic

import (
	"context"

	"tiny-tiktok/service/favorite/internal/svc"
	"tiny-tiktok/service/favorite/pb/favorite"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *favorite.FavoriteActionReq) (*favorite.FavoriteActionResp, error) {
	err := l.svcCtx.Redis.XADD(l.ctx, in.UserId, in.VideoId, in.ActionType)
	var statusMsg string
	if err != nil {
		if in.ActionType == 1 {
			statusMsg = "fail to favor"
		} else {
			statusMsg = "fail to cancel favor"
		}
	} else {
		if in.ActionType == 1 {
			statusMsg = "success to favor"
		} else {
			statusMsg = "success to cancel favor"
		}
	}

	return &favorite.FavoriteActionResp{
		StatusCode: 0,
		StatusMsg:  statusMsg,
	}, err
}
