package extra_second

import (
	"context"

	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/relation/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowerListLogic {
	return &RelationFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowerListLogic) RelationFollowerList(req *types.RelationFollowerListReq) (resp *types.RelationFollowerListResp, err error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.RelationRpc.FollowerList(l.ctx, &relation.FollowerListRequest{
		UserId: req.UserID,
		Token: req.Token,
	})

	var userList []types.User
	if err != nil {
		return &types.RelationFollowerListResp{
			StatusCode: "4401",
			StatusMsg: "粉丝查询rpc错误",
			UserList: userList,
		}, nil
	}

	for _, user := range(rpcResp.UserList) {
		userList = append(userList, types.User{
			Id: user.Id,
			Avatar: user.Avatar,
			BackgroundImage: user.BackgroundImage,
			FavoriteCount: user.FavoriteCount,
			FollowCount: user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow: user.IsFollow,
			Name: user.Name,
			Signature: user.Signature,
			TotalFavorited: user.TotalFavorited,
			WorkCount: user.WorkCount,
		})
	}

	return &types.RelationFollowerListResp{
		StatusCode: "4000",
		StatusMsg: "查询成功",
		UserList: userList,
	}, nil
}
