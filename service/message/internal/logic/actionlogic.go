package logic

import (
	"context"
	model "tiny-tiktok/service/message/internal/genModel"

	"tiny-tiktok/service/message/internal/svc"
	"tiny-tiktok/service/message/pb/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLogic {
	return &ActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActionLogic) Action(in *message.DouyinRelationActionRequest) (*message.DouyinRelationActionResponse, error) {
	// todo: add your logic here and delete this line
	logx.Info(in.Content)
	resp := message.DouyinRelationActionResponse{
		StatusCode: 0,
		StatusMsg:  "返回成功",
	}
	data := model.Message{
		ToUserId:   in.ToUserId,
		FromUserId: in.FromUserId,
		Content:    in.Content,
	}
	res, err := l.svcCtx.Model.Insert(l.ctx, &data)
	if err != nil || err == model.ErrNotFound {
		logx.Error(err)
		resp.StatusMsg = "插入失败"
		return &resp, err
	}
	if res == nil {
		logx.Error(err)
		resp.StatusMsg = "未知错误"
		return &resp, err
	}

	return &resp, nil
}
