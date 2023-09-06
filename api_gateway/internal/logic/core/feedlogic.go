package core

import (
	"context"
	"strconv"
	"time"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/feed/pb/feed"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedReq) (resp *types.FeedResp, err error) {
	var dateTime string

	if req.LatestTime == "0" || req.LatestTime == "" {

		dateTime = time.Now().Format("2006-01-02 15:04:05")
	}
	LatestTime, _ := strconv.ParseInt(req.LatestTime, 10, 64)

	// 将时间戳转换为时间
	t := time.Unix(LatestTime, 0)

	// 将时间格式化为字符串
	dateTime = t.Format("2006-01-02 15:04:05")

	// use etc/service.yaml instead
	// conn := zrpc.MustNewClient(zrpc.RpcClientConf{
	// 	Etcd: discov.EtcdConf{
	// 		Hosts: []string{"127.0.0.1:2379"},
	// 		Key:   "feed.rpc",
	// 	},
	// })

	// client := feed.NewFeedServiceClient(conn.Conn())
	client := l.svcCtx.FeedRpc
	respRpc, err := client.Feed(context.Background(), &feed.FeedRequest{
		LatestTime: dateTime,
		Token:      req.Token,
	})
	if err != nil {
		logx.Info("出错了")
	}
	resp = &types.FeedResp{}
	err = copier.Copy(&resp, &respRpc)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return resp, nil
}
