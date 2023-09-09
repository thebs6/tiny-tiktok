package logic

import (
	"context"
	"flag"
	"testing"
	"tiny-tiktok/service/favorite/internal/config"
	"tiny-tiktok/service/favorite/internal/svc"
	"tiny-tiktok/service/favorite/pb/favorite"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
)

var configFile = flag.String("f", "../../etc/favorite.yaml", "the config file")
var ctx = context.Background()
var svcCtx *svc.ServiceContext

func TestMain(m *testing.M) {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx = svc.NewServiceContext(c)

	logconf := logc.LogConf{
		ServiceName: c.LogConf.ServiceName,
		Mode:        c.LogConf.Mode,
		Path:        c.LogConf.Path,
		Level:       "debug",
	}
	logc.MustSetup(logconf)

	m.Run()
}
func TestFavoriteActionLogic(t *testing.T) {
	type args struct {
		userId     int64
		videoId    int64
		actionType int32
	}

	tests := []struct {
		testName string
		args     args
		wantMsg  string
		wantErr  bool
	}{
		{
			testName: "valid favor",
			args: args{
				userId:     1,
				videoId:    2,
				actionType: 1,
			},
			wantMsg: "success to favor",
			wantErr: false,
		},
		{
			testName: "invalid favor",
			args: args{
				userId:     1,
				videoId:    1,
				actionType: 1,
			},
			wantMsg: "you have favored this video already",
			wantErr: false,
		},
		{
			testName: "valid cancel favor",
			args: args{
				userId:     1,
				videoId:    1,
				actionType: 2,
			},
			wantMsg: "success to cancel favor",
			wantErr: false,
		},
		{
			testName: "invalid cancel favor",
			args: args{
				userId:     1,
				videoId:    1,
				actionType: 2,
			},
			wantMsg: "you have not favored this video yet",
			wantErr: false,
		},
	}

	for _, test := range tests {
		// t.Run(test.testName, func(t *testing.T) {
		resp, err := NewFavoriteActionLogic(ctx, svcCtx).FavoriteAction(&favorite.FavoriteActionReq{
			UserId:     test.args.userId,
			VideoId:    test.args.videoId,
			ActionType: test.args.actionType,
		})
		if (err != nil) != test.wantErr {
			t.Errorf("FavoriteAction() error: %v", err)
			return
		}
		if resp.StatusMsg != test.wantMsg {
			// t.Errorf("wrong Msf: want: %s, actual: %s", test.wantMsg, resp.StatusMsg)
			// return
		}
		// })
	}
}
