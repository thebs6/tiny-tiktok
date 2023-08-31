package logic

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"testing"
	"tiny-tiktok/service/feed/internal/config"
	"tiny-tiktok/service/feed/internal/svc"
	"tiny-tiktok/service/feed/pb/feed"
)

var configFile = flag.String("f", "D:\\code\\tiktok\\tiny-tiktok\\service\\feed\\etc\\feed.yaml", "the config file")

func BenchmarkFeedLogic_FeedTest(b *testing.B) {
	var c config.Config

	conf.MustLoad(*configFile, &c)

	req := feed.FeedRequest{
		LatestTime: "2023-08-29 17:59:26",
		Token:      "",
	}

	logic := NewFeedLogic(context.Background(), svc.NewServiceContext(c))
	logic.Feed(&req)
}
