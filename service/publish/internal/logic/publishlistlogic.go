package logic

import (
	"context"

	"tiny-tiktok/service/publish/internal/svc"
	"tiny-tiktok/service/publish/pb/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *publish.PublishListReq) (*publish.PublishListResp, error) {
	// todo: add your logic here and delete this line

	return &publish.PublishListResp{}, nil
}
