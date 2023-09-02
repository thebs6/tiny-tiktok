package logic

// import (
// 	"context"
// 	"strconv"

// 	"tiny-tiktok/service/favorite/internal/svc"
// 	"tiny-tiktok/service/favorite/pb/favorite"

// 	"tiny-tiktok/service/favorite/internal/redis_model"

// 	"github.com/zeromicro/go-zero/core/logc"
// 	"github.com/zeromicro/go-zero/core/logx"
// )

// type FavoriteActionLogic struct {
// 	ctx    context.Context
// 	svcCtx *svc.ServiceContext
// 	logx.Logger
// }

// func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
// 	return &FavoriteActionLogic{
// 		ctx:    ctx,
// 		svcCtx: svcCtx,
// 		Logger: logx.WithContext(ctx),
// 	}
// }

// func (l *FavoriteActionLogic) FavoriteAction(in *favorite.FavoriteActionReq) (*favorite.FavoriteActionResp, error) {
// 	if in.ActionType == 1 {
// 		// l.svcCtx.FavoriteModel.Insert(l.ctx, &model.Favorite{
// 		// 	VideoId: in.VideoId,
// 		// 	UserId:  in.UserId,
// 		// })
// 		err := l.svcCtx.Redis.Favor(l.ctx, in.UserId, in.VideoId)
// 		if err == redis_model.ErrRecordRepeated {
// 			logging := "user_" + strconv.FormatInt(in.UserId, 10) + " try to favor video_" + strconv.FormatInt(in.VideoId, 10) + " again"
// 			logc.Info(l.ctx, logging)
// 			return &favorite.FavoriteActionResp{
// 				StatusCode: 1,
// 				StatusMsg:  "you have favored this video already",
// 			}, nil
// 		} else if err != nil {
// 			logc.Alert(l.ctx, err.Error())
// 			return &favorite.FavoriteActionResp{
// 				StatusCode: 1,
// 				StatusMsg:  "fail to favor",
// 			}, err
// 		}
// 		return &favorite.FavoriteActionResp{
// 			StatusCode: 0,
// 			StatusMsg:  "success to favor",
// 		}, nil
// 	}
// 	err := l.svcCtx.Redis.CancelFavor(l.ctx, in.UserId, in.VideoId)
// 	if err == redis_model.ErrRecordNonExist {
// 		logging := "user_" + strconv.FormatInt(in.UserId, 10) + " try to cancel favor video_" + strconv.FormatInt(in.VideoId, 10) + " he has not favored yet"
// 		logc.Info(l.ctx, logging)
// 		return &favorite.FavoriteActionResp{
// 			StatusCode: 1,
// 			StatusMsg:  "you have not favored this video yet",
// 		}, nil
// 	} else if err != nil {
// 		logc.Alert(l.ctx, err.Error())
// 		return &favorite.FavoriteActionResp{
// 			StatusCode: 1,
// 			StatusMsg:  "fail to cancel favor",
// 		}, err
// 	}
// 	return &favorite.FavoriteActionResp{
// 		StatusCode: 0,
// 		StatusMsg:  "success to cancel favor",
// 	}, nil
// }
