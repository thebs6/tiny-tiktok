package logic

import (
	"context"
	"flag"
	"testing"
	"tiny-tiktok/service/publish/internal/config"
	"tiny-tiktok/service/publish/internal/svc"
	"tiny-tiktok/service/publish/pb/publish"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "../../etc/publish.yaml", "the config file")
var ctx = context.Background()
var svcCtx *svc.ServiceContext

func TestMain(m *testing.M) {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx = svc.NewServiceContext(c)

	m.Run()
}
func TestCommentList(t *testing.T) {
	type args struct {
		userId int64
	}
	tests := []struct {
		testName string
		args     args
		wantAct  string
		wantErr  bool
	}{
		{
			testName: "valid comment",
			args: args{
				userId: 1,
			},
			wantErr: false,
		},
		{
			testName: "valid comment",
			args: args{
				userId: 2,
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			_, err := NewPublishListLogic(ctx, svcCtx).PublishList(&publish.PublishListReq{
				UserId: test.args.userId,
			})
			if (err != nil) != test.wantErr {
				t.Errorf("CommentList() error: %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}
