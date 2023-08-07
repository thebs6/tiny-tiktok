package extra_second

import (
	"context"

	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLogic {
	return &MessageChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageChatLogic) MessageChat(req *types.MessageChatReq) (resp *types.MessageChatResp, err error) {
	// todo: add your logic here and delete this line

	return
}
