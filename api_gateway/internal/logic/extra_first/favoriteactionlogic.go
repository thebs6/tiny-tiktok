package extra_first

import (
	"context"
	"encoding/json"
	"net/http"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/favorite/pb/favorite"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.FavoriteActionReq) (resp *types.FavoriteActionResp, err error) {
	uid, err := l.ctx.Value("payload").(json.Number).Int64()
	if err != nil {
		logc.Debugf(l.ctx, "payload.(string) failed")
		return &types.FavoriteActionResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "点赞失败!",
		}, nil
	}

	respRpc, err := l.svcCtx.FavoriteRpc.FavoriteAction(l.ctx, &favorite.FavoriteActionReq{
		VideoId:    req.VideoID,
		UserId:     uid,
		ActionType: req.ActionType,
	})
	if err != nil {
		logc.Alert(l.ctx, "rpc favoriteaction failed, err: "+err.Error())
		return &types.FavoriteActionResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "点赞失败!",
		}, nil
	}
	return &types.FavoriteActionResp{
		StatusCode: respRpc.StatusCode,
		StatusMsg:  "点赞成功!",
	}, nil
}
