package extra_second

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
)

type RelationFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFriendListLogic {
	return &RelationFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFriendListLogic) RelationFriendList(req *types.RelationFriendListReq) (resp *types.RelationFriendListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
