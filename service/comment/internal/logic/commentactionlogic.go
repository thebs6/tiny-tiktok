package logic

import (
	"context"
	"fmt"
	"time"

	"tiny-tiktok/service/comment/internal/model"
	"tiny-tiktok/service/comment/internal/svc"
	"tiny-tiktok/service/comment/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
			VideoId: in.VideoId,
			UserId:  in.UserId,
			Content: in.CommentText,
		}
		var user_id int64
		if err := l.svcCtx.CommentModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			res, err := l.svcCtx.CommentModel.Insert(l.ctx, &data)
			if err != nil {
				return err
			}
			user_id, err = res.LastInsertId()
			if err != nil {
				return err
			}
			// 1.inset into redis after successing to insert into mysql
			// 2.the created_at is a little bit later than the created_at in mysql.
			// "commentlist" api just response the month-day to the client so it does not matter at most time.
			// However, if the comment is published at 23:59:59, the "commentlist" api maybe response the next day
			data.CreatedAt = time.Now()
			err = l.svcCtx.CommentRedis.ZAdd(l.ctx, in.VideoId, time.Now().Unix(), &data)
			if err != nil {
				// mysql rollback if failed to insert into redis
				return err
			}
			return nil
		}); err != nil {
			// transaction fails
			return &comment.CommentActionResp{
				StatusMsg: "Failed to comment",
				Comment:   nil,
			}, err
		}

		// TODO(gcx): change to Microservice api
		user := queryUserById(in.UserId)
		return &comment.CommentActionResp{
			StatusMsg: "Comment successfully",
			Comment: &comment.Comment{
				Id:         user_id,
				User:       user,
				Content:    in.CommentText,
				CreateDate: time.Now().Format("01-02"),
			},
		}, nil
	} else {
		// delete a comment
		// TODO(gcx): whether we should judge the comment which is going to be deleted
		// is publish by the user who try to delete it?
		err := l.svcCtx.CommentModel.Delete(l.ctx, in.CommentId)

		// Attention: error will not occur when the commentid does not exsit
		if err != nil {
			fmt.Printf("error %s", err.Error())
			return &comment.CommentActionResp{
				StatusMsg: "Failed to delete the comment",
				Comment:   nil,
			}, err
		}

		return &comment.CommentActionResp{
			StatusMsg: "Delete the comment successfully",
			Comment:   nil,
		}, nil
	}
}
