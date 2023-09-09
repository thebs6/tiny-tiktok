package extra_second

import (
	"context"
	"net/http"

	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/relation/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowListLogic {
	return &RelationFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowListLogic) RelationFollowList(req *types.RelationFollowListReq) (resp *types.RelationFollowListResp, err error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.RelationRpc.FollowerList(l.ctx, &relation.FollowerListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})

	var userList []types.User
	if err != nil {
		logx.Error(err)
		resp = &types.RelationFollowListResp{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "fail",
			UserList:   userList,
		}
		return
	}

	for _, user := range rpcResp.UserList {
		userList = append(userList, types.User{
			Id:              user.Id,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			FavoriteCount:   user.FavoriteCount,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        user.IsFollow,
			Name:            user.Name,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
		})
	}

	return &types.RelationFollowListResp{
		StatusCode: http.StatusOK,
		StatusMsg:  "success",
		UserList:   userList,
	}, nil
}
