package logic

import (
	"context"
	"time"

	"tiny-tiktok/service/comment/internal/model"
	"tiny-tiktok/service/comment/internal/svc"
	"tiny-tiktok/service/comment/pb/comment"

	"github.com/zeromicro/go-zero/core/logc"
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
		data := &model.Comment{
			VideoId: in.VideoId,
			UserId:  in.UserId,
			Content: in.CommentText,
		}
		var comment_id int64
		err := l.svcCtx.CommentModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			// insert the record into mysql
			res, err := l.svcCtx.CommentModel.TransInsert(l.ctx, session, data)
			if err != nil {
				return err
			}

			comment_id, err = res.LastInsertId()
			if err != nil {
				return err
			}

			// // 1.Inset into redis after successing to insert into mysql
			// // 2.The created_at is a little bit later than the created_at in mysql.
			// // "commentlist" api just response the month-day to the client so it does not matter at most time.
			// // However, if the comment is published at 23:59:59, the "commentlist" api maybe response the next day
			// // Futhermore, this implementation will have us can not delete the record from the redis.
			// data.CreatedAt = time.Now()
			// err = l.svcCtx.CommentRedis.ZAdd(l.ctx, in.VideoId, time.Now().Unix(), data)

			// To solve the problems which are led by the different time, we query mysql.
			// Unluckly, this implementation will hurt the performance of our system
			data, err = l.svcCtx.CommentModel.TransFindone(l.ctx, session, comment_id)
			if err != nil {
				return err
			}

			// insert the record into redis
			err = l.svcCtx.CommentRedis.ZAdd(l.ctx, in.VideoId, time.Now().Unix(), data)
			// mysql rollback if failed to insert into redis
			if err != nil {
				return err
			}

			return nil
		})

		// transaction fails
		if err != nil {
			logc.Alert(l.ctx, err.Error())

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
				Id:         comment_id,
				User:       user,
				Content:    in.CommentText,
				CreateDate: time.Now().Format("01-02"),
			},
		}, nil
	}

	// delete a comment
	// TODO(gcx): whether we should judge the comment which is going to be deleted
	// is publish by the user who try to delete it?

	// get the comment from mysql and delete it from redis
	resp, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.CommentId)
	if err != nil {
		return &comment.CommentActionResp{
			StatusMsg: "Failed to delete the comment",
			Comment:   nil,
		}, err
	}

	err = l.svcCtx.CommentRedis.ZRem(l.ctx, in.VideoId, resp)
	if err != nil {
		return &comment.CommentActionResp{
			StatusMsg: "Failed to delete the comment",
			Comment:   nil,
		}, err
	}

	// soft delete the record in mysql after deleting the record in redis
	// or the record will be different from the one in redis
	// and fail to delete the record in redis
	err = l.svcCtx.CommentModel.SoftDel(l.ctx, in.CommentId)
	if err != nil {
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
