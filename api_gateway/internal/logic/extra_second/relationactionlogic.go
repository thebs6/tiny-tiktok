package extra_second

import (
	"context"

	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationActionLogic) RelationAction(req *types.RelationActionReq) (resp *types.RelationFollowListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
