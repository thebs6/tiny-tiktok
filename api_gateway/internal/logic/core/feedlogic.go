package core

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"strconv"
	"time"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/feed/pb/feed"
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
	// todo: add your logic here and delete this line
	var dateTime string
	if req.LatestTime == "" || req.LatestTime == "0" {
		dateTime = time.Now().Format("2006-01-02 15:04:05")
	}
	LatestTime, _ := strconv.ParseInt(req.LatestTime, 10, 64)

	// 将时间戳转换为时间
	t := time.Unix(LatestTime, 0)

	// 将时间格式化为字符串
	dateTime = t.Format("2006-01-02 15:04:05")

	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "feed.rpc",
		},
	})

	client := feed.NewFeedServiceClient(conn.Conn())
	respRpc, err := client.Feed(context.Background(), &feed.FeedRequest{
		LatestTime: dateTime,
		Token:      req.Token,
	})
	if err != nil {
		logx.Info("出错了")
	}
	resp = &types.FeedResp{}
	copier.Copy(&resp, &respRpc)

	return resp, nil
}
