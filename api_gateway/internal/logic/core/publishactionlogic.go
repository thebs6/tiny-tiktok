package core

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"

	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/publish/pb/publish"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	File   *multipart.FileHeader
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishActionLogic) PublishAction(req *types.PublishActionReq) (resp *types.PublishActionResp, err error) {
	// TODO()：
	// 1. 标题重复怎么处理？
	videoKey := "video/" + req.Title + ".mp4"
	err = l.uploadVideo(videoKey)
	if err != nil {
		return &types.PublishActionResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "Publish failed!",
		}, err
	}

	coverKey := "cover/" + req.Title + ".mp4"
	err = l.snapshotAndUpload(coverKey)
	if err != nil {
		return &types.PublishActionResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "Publish failed!",
		}, nil
	}

	uid, err := l.ctx.Value("payload").(json.Number).Int64()
	if err != nil {
		logc.Debugf(l.ctx, "payload.(string) failed")
		return &types.PublishActionResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "Publish failed!",
		}, nil
	}

	_, err = l.svcCtx.Publish.PublishAction(l.ctx, &publish.PublishActionReq{
		UserId:   uid,
		PlayUrl:  l.svcCtx.Config.Cos.URL + videoKey,
		CoverUrl: l.svcCtx.Config.Cos.URL + coverKey,
		Title:    req.Title,
	})
	if err != nil {
		l.Logger.Error("svc.Publish.PublishAction failed", err)
		return &types.PublishActionResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "Publish failed!",
		}, nil
	}

	return &types.PublishActionResp{
		StatusCode: http.StatusOK,
		StatusMsg:  "Publish success!",
	}, nil
}

func (l *PublishActionLogic) uploadVideo(videoKey string) error {
	file, err := l.File.Open()
	if err != nil {
		l.Logger.Error("File.Open failed", err)
		return err
	}
	defer file.Close()

	client := l.svcCtx.CosClient
	// Put() can only upload file which is less than 5GB
	_, err = client.Object.Put(l.ctx, videoKey, file, nil)
	if err != nil {
		l.Logger.Error(l.ctx, "Object.Put failed", err)
		return err
	}

	return nil
}

func (l *PublishActionLogic) snapshotAndUpload(coverKey string) error {
	// TODO()： 使用ffmpeg截图

	// client := l.svcCtx.CosClient
	// _, err := client.Object.Put(l.ctx, coverKey, f, nil)
	// if err != nil {
	// 	return err
	// }
	return nil
}
