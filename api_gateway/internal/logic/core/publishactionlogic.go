package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"path"
	"time"

	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/publish/pb/publish"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	"golang.org/x/sync/errgroup"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	// File       multipart.File
	FileHeader *multipart.FileHeader
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishActionLogic) PublishAction(req *types.PublishActionReq) (resp *types.PublishActionResp, err error) {
	// TODO(gcx)：
	// 1.视频截图及上传，视频发布微服务，二者之一失败，但前面的操作已经完成，怎么处理？
	// 2.消息队列
	videoCosId, err := l.getVideoKey(req.Title)
	if err != nil {
		logc.Alert(l.ctx, err.Error())
		return &types.PublishActionResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "Publish failed!",
		}, nil
	}

	videoKey := "video/" + videoCosId + path.Ext(l.FileHeader.Filename)
	coverKey := "cover/" + videoCosId + ".jpeg"

	// DEPRECATED upload video asynchronously now
	// err = l.uploadVideo(videoKey)
	// if err != nil {
	// 	return &types.PublishActionResp{
	// 		StatusCode: http.StatusOK,
	// 		StatusMsg:  "Publish failed!",
	// 	}, nil
	// }

	// err = l.snapshotAndUpload(coverKey, videoKey)
	// if err != nil {
	// 	return &types.PublishActionResp{
	// 		StatusCode: http.StatusOK,
	// 		StatusMsg:  "Publish failed!",
	// 	}, nil
	// }

	// upload video and snapshot in goroutine
	group := new(errgroup.Group)
	group.Go(func() error {
		err = l.uploadVideo(videoKey)
		if err != nil {
			return err
		}

		// DEPRECATED use cos to snapshot instead now
		// err = l.snapshotAndUpload(coverKey, videoKey)
		// if err != nil {
		// 	return err
		// }

		return nil
	})

	uid, err := l.ctx.Value("payload").(json.Number).Int64()
	if err != nil {
		logc.Debugf(l.ctx, "payload.(string) failed")
		return &types.PublishActionResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "Publish failed!",
		}, nil
	}

	_, err = l.svcCtx.PublishRpc.PublishAction(l.ctx, &publish.PublishActionReq{
		UserId:   uid,
		PlayUrl:  l.svcCtx.Config.Cos.URL + videoKey,
		CoverUrl: l.svcCtx.Config.Cos.URL + coverKey,
		Title:    req.Title,
	})

	if err != nil || group.Wait() != nil {
		// note: error of group.Wait() has logged in uploadVideo() or snapshotAndUpload()
		if err != nil {
			logc.Alert(l.ctx, err.Error())
		}
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
	file, err := l.FileHeader.Open()
	if err != nil {
		logc.Alert(l.ctx, "uploadVideo() "+err.Error())
		return err
	}
	defer file.Close()

	client := l.svcCtx.CosClient
	// Put() can only upload file which is less than 5GB
	_, err = client.Object.Put(l.ctx, videoKey, file, nil)
	if err != nil {
		logc.Alert(l.ctx, "uploadVideo() "+err.Error())
		return err
	}

	return nil
}

func (l *PublishActionLogic) snapshotAndUpload(coverKey, videoKey string) error {
	file, err := l.FileHeader.Open()
	if err != nil {
		logc.Alert(l.ctx, "snapshotAndUpload() fileopen "+err.Error())
		return err
	}
	defer file.Close()

	// snapshot the video at framNum
	frameNum := 1
	buf := bytes.NewBuffer(nil)

	// method 1. use ffmpeg to snapshot from io.Reader(fastest)
	// newFile, _ := os.Create("./cos_test_com.mp4")
	// defer newFile.Close()
	// b, _ := io.ReadAll(file)
	// newFile.Write(b)
	// err := ffmpeg.Input("pipe:", ffmpeg.KwArgs{}).WithInput(file).
	// 	// err := ffmpeg.Input("./cos_test_com.mp4").
	// 	Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
	// 	Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
	// 	WithOutput(buf, os.Stdout).
	// 	Run()

	// method 2. use ffmpeg to snapshot from playurl and put it to cos
	err = ffmpeg.Input(l.svcCtx.Config.Cos.URL+"/"+videoKey).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		// WithOutput(buf, os.Stdout).
		WithOutput(buf).
		Run()

	// method 3. use ffmpeg to snapshot and put it to cos directly by s3
	// TODO(gcx): upload to cos directly by s3
	// https://cloud.tencent.com/developer/article/1814657
	// err = ffmpeg.Input("pipe:", ffmpeg.KwArgs{}).
	// 	Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
	// 	Output("s3://data-1251825869/test_out.ts", ffmpeg.KwArgs{
	// 		"aws_config": &aws.Config{
	// 			Credentials: credentials.NewStaticCredentials("xx", "yyy", ""),
	// 			Endpoint:    aws.String("xx"),
	// 			// Region: aws.String("yyy"),
	// 		},
	// 		"format": "image2"}).
	// 	WithOutput(buf).WithInput(file).
	// 	Run()

	if err != nil {
		logc.Alert(l.ctx, "snapshotAndUpload() ffmpeg "+err.Error())
		return err
	}

	client := l.svcCtx.CosClient
	_, err = client.Object.Put(l.ctx, coverKey, buf, nil)
	if err != nil {
		logc.Alert(l.ctx, "snapshotAndUpload() put to cos "+err.Error())
		return err
	}
	return nil
}

func (l *PublishActionLogic) getVideoKey(title string) (string, error) {
	now := time.Now()
	// format: type(2) + year(2) + month(2) + day(2) + hour(2) + minute(2) + second(2) + id_in_one_second(4)
	id, err := l.svcCtx.Redis.INCR(l.ctx)
	if err != nil {
		return "", err
	}
	return title + "_01" + now.Format("060102150405") + fmt.Sprintf("%04d", id), nil
}
