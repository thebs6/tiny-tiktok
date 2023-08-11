package logic

import (
	"context"

	"tiny-tiktok/service/comment/internal/model"
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
	// TODO(gcx): Add redis cache here
	var comments []*comment.Comment
	if l.svcCtx.CommentRedis.Exists(l.ctx, in.VideoId) == 1 {
		// 1. query redis
		respList, err := l.svcCtx.CommentRedis.ZRevRangeWithScores(l.ctx, in.VideoId)
		if err != nil {
			return nil, err
		}
		for _, resp := range respList {
			r, _ := resp.Member.(model.Comment)
			// TODO(gcx): change to Microservice api
			user := queryUserById(r.UserId)

			comments = append(comments, &comment.Comment{
				Id:         r.Id,
				User:       user,
				Content:    r.Content,
				CreateDate: r.CreatedAt.Format("01-02"),
			})
		}
		return &comment.CommentListResp{
			StatusMsg:   "Get comment list succesfully",
			CommentList: comments,
		}, nil
	}
	// 2. if not exist, query mysql and insert into redis
	respComments, err := l.svcCtx.CommentModel.List(l.ctx, in.VideoId)
	if err != nil {
		return nil, err
	}

	for _, respComment := range respComments {
		// TODO(gcx): change to Microservice api
		user := queryUserById(respComment.UserId)

		comments = append(comments, &comment.Comment{
			Id:         respComment.Id,
			User:       user,
			Content:    respComment.Content,
			CreateDate: respComment.CreatedAt.Format("01-02"),
		})
	}
	return &comment.CommentListResp{
		StatusMsg:   "Get comment list succesfully",
		CommentList: comments,
	}, nil

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
