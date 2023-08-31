package extra_second

import (
	"context"

	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/relation/relation"

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

func (l *RelationActionLogic) RelationAction(req *types.RelationActionReq) (resp *types.RelationActionResp, err error) {
	// todo: add your logic here and delete this line

	RpcResp, err := l.svcCtx.RelationRpc.Action(l.ctx, &relation.ActionRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
	})

	if err != nil {
		return &types.RelationActionResp{
			StatusMsg:  "RPC调用错误",
			StatusCode: 1401,
		}, err
	}

	return &types.RelationActionResp{
		StatusMsg:  RpcResp.StatusMsg,
		StatusCode: 200,
	}, nil
}
