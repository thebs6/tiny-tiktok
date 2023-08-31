package logic

import (
	"context"
	"time"

	// ctxdata "tiny-tiktok/common/ctxData"
	"tiny-tiktok/service/relation/internal/model"
	"tiny-tiktok/service/relation/internal/svc"
	"tiny-tiktok/service/relation/relation"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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

func (l *ActionLogic) Action(in *relation.ActionRequest) (*relation.ActionResponse, error) {
	// todo: add your logic here and delete this line
	// userId := ctxdata.GetUidFromCtx(l.ctx)
	var userId int64
	userId = 6
	fromUser, err := l.svcCtx.UserModel.FindOne(l.ctx, userId) 
	if err != nil {
		return nil, err
	}
	// 先查关注的用户是否存在
	toUserId := in.ToUserId
	toUser, err := l.svcCtx.UserModel.FindOne(l.ctx, toUserId)
	// 如果关注的用户不存在直接返回
	if err != nil && err != sqlx.ErrNotFound {
		return nil, err
	}

	if err != nil && err == sqlx.ErrNotFound {
		return &relation.ActionResponse{
			StatusCode: 4401,
			StatusMsg: "用户不存在",
		}, err
	}

	switch in.ActionType {
	// 关注
	case 1 : 
		resp, err := l.Follow(fromUser, toUser)
		if err != nil {
			return resp, err
		}
	// 取消关注
	case 2:
		resp, err := l.UnFollow(fromUser, toUser)
		if err != nil {
			return resp, err
		}
	default:
		return &relation.ActionResponse{
			StatusCode: 4004,
			StatusMsg: "错误操作",
		}, nil
	}
	
	return &relation.ActionResponse{}, nil
}

func (l *ActionLogic) Follow(fromUser *model.User, toUser *model.User) (*relation.ActionResponse, error) {
	// 判断是否已经关注
	fromUserId := fromUser.Id
	toUserId := toUser.Id
	relationResp, err := l.svcCtx.RelationModel.FindRelationByTwoId(l.ctx, toUserId, fromUserId)
	if err != nil && err != sqlx.ErrNotFound {
		return nil, err
	}
	
	// 已经关注直接返回
	if relationResp != nil {
		return &relation.ActionResponse{
			StatusCode: 4402,
			StatusMsg: "已关注",
		}, nil
	}

	// 未关注则关注同时更新双方关注数量 
	var newRelation model.Relation
	newRelation.FollowerId = fromUserId
	newRelation.FollowId = toUserId
	newRelation.CreatedAt = time.Now()
	newRelation.UpdatedAt = time.Now()
	_, err = l.svcCtx.RelationModel.Insert(l.ctx, &newRelation)
	if err != nil {
		return nil, err
	}

	fromUser.FollowCount += 1
	l.svcCtx.UserModel.Update(l.ctx, fromUser)

	toUser.FollowerCount += 1
	l.svcCtx.UserModel.Update(l.ctx, toUser)

	return &relation.ActionResponse{}, nil
}

func (l *ActionLogic) UnFollow(fromUser *model.User, toUser *model.User) (*relation.ActionResponse, error) {
	// 判断是否已经关注
	fromUserId := fromUser.Id
	toUserId := toUser.Id
	relationResp, err := l.svcCtx.RelationModel.FindRelationByTwoId(l.ctx, toUserId, fromUserId)
	if err != nil && err != sqlx.ErrNotFound {
		return nil, err
	}

	
	// 未关注直接返回
	if err != nil && err == sqlx.ErrNotFound || relationResp == nil {
		return &relation.ActionResponse{
			StatusCode: 4402,
			StatusMsg: "已关注",
		}, nil
	}

	// 关注则取关，同时更新双方关注数量 
	err = l.svcCtx.RelationModel.Delete(l.ctx, relationResp.Id)
	if err != nil {
		return nil, err
	}

	fromUser.FollowCount -= 1
	l.svcCtx.UserModel.Update(l.ctx, fromUser)

	toUser.FollowerCount -= 1
	l.svcCtx.UserModel.Update(l.ctx, toUser)

	return &relation.ActionResponse{}, nil
}
