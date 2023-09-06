package logic

import (
	"context"

	"tiny-tiktok/service/comment/internal/svc"
	"tiny-tiktok/service/comment/pb/comment"

	"github.com/zeromicro/go-zero/core/logc"
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
	var comments []*comment.Comment

	exist, errRds := l.svcCtx.CommentRedis.Exists(l.ctx, in.VideoId)

	// redis error and get comment list from db
	if errRds != nil || exist == 0 {
		if errRds != nil {
			logc.Alert(l.ctx, errRds.Error())
		}

		commentList, err := l.svcCtx.CommentModel.List(l.ctx, in.VideoId)
		if err != nil {
			logc.Alert(l.ctx, err.Error())
			return &comment.CommentListResp{
				StatusMsg:   "Fail to get comment list",
				CommentList: comments,
			}, err
		}

		userIdList := make([]int64, len(commentList))
		for i, respComment := range commentList {
			userIdList[i] = respComment.UserId
		}
		userList := make([]*comment.User, len(commentList))
		// TODO(gcx): change to Microservice api
		queryUsersByIds(userIdList, userList)

		if errRds == nil && exist == 0 {
			err = l.svcCtx.CommentRedis.ZAddList(l.ctx, in.VideoId, commentList)
			logc.Alert(l.ctx, "Failed to write commentList in redis"+err.Error())
		}

		for i, respComment := range commentList {
			comments = append(comments, &comment.Comment{
				Id:         respComment.Id,
				User:       userList[i],
				Content:    respComment.Content,
				CreateDate: respComment.CreatedAt.Format("01-02"),
			})
		}

	} else if exist == 1 {
		commentList, err := l.svcCtx.CommentRedis.ZRevRangeWithScores(l.ctx, in.VideoId)
		if err != nil {
			logc.Alert(l.ctx, err.Error())
			return nil, err
		}

		userIdList := make([]int64, len(commentList))
		for i, respComment := range commentList {
			userIdList[i] = respComment.UserId
		}
		userList := make([]*comment.User, len(commentList))
		// TODO(gcx): change to Microservice api
		queryUsersByIds(userIdList, userList)

		for i, c := range commentList {
			// TODO(gcx): change to Microservice api
			// user := queryUserById(c.UserId)
			comments = append(comments, &comment.Comment{
				Id: c.Id,
				// User:       user,
				User:       userList[i],
				Content:    c.Content,
				CreateDate: c.CreatedAt.Format("01-02"),
			})
		}
	}

	return &comment.CommentListResp{
		StatusMsg:   "Get comment list succesfully",
		CommentList: comments,
	}, nil
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
