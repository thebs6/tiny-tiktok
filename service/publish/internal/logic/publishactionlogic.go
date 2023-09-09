package logic

import (
	"context"

	"tiny-tiktok/service/publish/internal/svc"
	"tiny-tiktok/service/publish/model"
	"tiny-tiktok/service/publish/pb/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishActionLogic) PublishAction(in *publish.PublishActionReq) (*publish.PublishActionResp, error) {
	video := &model.Video{
		Author:   in.UserId,
		PlayUrl:  in.PlayUrl,
		CoverUrl: in.CoverUrl,
		Title:    in.Title,
	}
	_, err := l.svcCtx.VideoModel.Insert(l.ctx, video)
	if err != nil {
		return nil, err
	}

	return &publish.PublishActionResp{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}
