package core

import (
	"context"
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
	// 2. 可以将视频上传、截图和publishrpc服务调用改成并发执行？
	// 		改成并发后，如果其中一个失败了，另一个怎么处理？
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
		}, err
	}

	user_id, _ := l.ctx.Value("payload").(int64)

	_, err = l.svcCtx.Publish.PublishAction(l.ctx, &publish.PublishActionReq{
		UserId:   user_id,
		PlayUrl:  l.svcCtx.Config.Cos.URL + videoKey,
		CoverUrl: l.svcCtx.Config.Cos.URL + coverKey,
		Title:    req.Title,
	})
	if err != nil {
		return &types.PublishActionResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "Publish failed!",
		}, err
	}

	return &types.PublishActionResp{
		StatusCode: http.StatusOK,
		StatusMsg:  "Publish success!",
	}, nil
}

func (l *PublishActionLogic) uploadVideo(videoKey string) error {
	client := l.svcCtx.CosClient
	file, err := l.File.Open()
	if err != nil {
		logc.Info(l.ctx, "File.Open failed", err)
		return err
	}
	defer file.Close()

	// Put() can only upload file which is less than 5GB
	_, err = client.Object.Put(l.ctx, videoKey, file, nil)
	if err != nil {
		logc.Info(l.ctx, "Object.Put failed", err)
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
