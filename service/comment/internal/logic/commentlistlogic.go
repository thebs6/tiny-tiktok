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
		var comments []*comment.Comment
		for _, respComment := range respComments {
			// TODO(gcx): change to Microservice api
			user := queryUserById(respComment.UserId)

			comments = append(comments, &comment.Comment{
				Id:         respComment.Id,
				User:       user,
				Content:    respComment.Content,
				CreateDate: respComment.CreatedAt.String(),
			})
		}
		return &comment.CommentListResp{
			StatusMsg:   "Failed to get comment list",
			CommentList: comments,
		}, nil
	} else {

		return &comment.CommentListResp{}, nil
	}
}

func queryUserById(user_id int64) *comment.User {
	return &comment.User{
		Id:            user_id,
		Name:          "gao",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
}
