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
		// delete the cache first time
		err := l.svcCtx.CommentRedis.Del(l.ctx, in.VideoId)
		if err != nil {
			logc.Alert(l.ctx, err.Error())
			return &comment.CommentActionResp{
				StatusMsg: "Failed to comment",
				Comment:   nil,
			}, err
		}
		var comment_id int64
		err = l.svcCtx.CommentModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			// insert the record into mysql
			res, err := l.svcCtx.CommentModel.TransInsert(l.ctx, session, data)
			if err != nil {
				return err
			}

			comment_id, err = res.LastInsertId()
			if err == nil {
				return err
			}

			// delete the cache second time
			time.Sleep(time.Duration(500 * time.Millisecond))
			err = l.svcCtx.CommentRedis.Del(l.ctx, in.VideoId)
			return err
		})

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

	// delete the cache first time
	l.svcCtx.CommentRedis.Del(l.ctx, in.VideoId)
	err := l.svcCtx.CommentModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// Get the comment from db before soft deleting it.
		// Because once we set the deleted_at, the record in db will be different from the record in redis,
		// and we will be unable to remove the record from redis.
		_, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.CommentId)
		if err != nil {
			return err
		}

		// Soft delete the record in db
		err = l.svcCtx.CommentModel.TransSoftDel(l.ctx, session, in.CommentId)
		if err != nil {
			return err
		}

		// delete the cache second time
		time.Sleep(time.Duration(500 * time.Millisecond))
		err = l.svcCtx.CommentRedis.Del(l.ctx, in.VideoId)
		return err
	})

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
