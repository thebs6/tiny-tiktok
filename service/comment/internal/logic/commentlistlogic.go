package logic

import (
	"context"

	"tiny-tiktok/service/comment/internal/svc"
	"tiny-tiktok/service/comment/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *comment.CommentListReq) (*comment.CommentListResp, error) {
	respComments, err := l.svcCtx.CommentModel.List(l.ctx, in.VideoId)
	if err != nil {
		return nil, err
	} else {
		userIdList := make([]int64, len(respComments))
		for i, respComment := range respComments {
			userIdList[i] = respComment.UserId
		}
		userList := make([]*comment.User, len(respComments))
		// TODO(gcx): change to Microservice api
		queryUsersByIds(userIdList, userList)

		var comments []*comment.Comment
		for i, respComment := range respComments {
			// TODO(gcx): change to Microservice api(drop)
			// user := queryUserById(respComment.UserId)

			comments = append(comments, &comment.Comment{
				Id: respComment.Id,
				// User:       user,
				User:       userList[i],
				Content:    respComment.Content,
				CreateDate: respComment.CreatedAt.Format("01-02"),
			})
		}

		return &comment.CommentListResp{
			StatusMsg:   "Get comment list succesfully",
			CommentList: comments,
		}, nil
	}
}

// stub code
func queryUserById(user_id int64) *comment.User {
	return &comment.User{
		Id:            user_id,
		Name:          "gao",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
}

// stub code
func queryUsersByIds(userIds []int64, users []*comment.User) {
	for i, userId := range userIds {
		users[i] = &comment.User{
			Id:            userId,
			Name:          "gao",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}
	}
}
