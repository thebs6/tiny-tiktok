package logic

import (
	"context"

	"tiny-tiktok/service/comment/internal/svc"
	"tiny-tiktok/service/comment/model"
	"tiny-tiktok/service/comment/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentActionLogic) CommentAction(in *comment.CommentActionReq) (*comment.CommentActionResp, error) {
	if in.ActionType == 1 {
		// publish a comment
		data := model.Comment{
			Video:   in.VideoId,
			User:    in.UserId,
			Content: in.CommentText,
		}
		res, err := l.svcCtx.CommentModel.Insert(l.ctx, &data)
		if err != nil {
			return &comment.CommentActionResp{
				StatusMsg: "Failed to comment",
				Comment:   nil,
			}, nil
		}

		id, err := res.LastInsertId()
		if err != nil {
			return &comment.CommentActionResp{
				StatusMsg: "Failed to comment",
				Comment:   nil,
			}, nil
		}

		// TODO(gcx): change to Microservice api
		user := queryUserById(in.UserId)
		return &comment.CommentActionResp{
			StatusMsg: "Comment successfully",
			Comment: &comment.Comment{
				Id:      id,
				User:    user,
				Content: in.CommentText,
			},
		}, nil
	} else {
		// delete a comment
		err := l.svcCtx.CommentModel.Delete(l.ctx, in.CommentId)
		if err != nil {
			return &comment.CommentActionResp{
				StatusMsg: "Failed to delete the comment",
				Comment:   nil,
			}, nil
		}

		return &comment.CommentActionResp{
			StatusMsg: "Delete the comment successfully",
			Comment:   nil,
		}, nil
	}
}
