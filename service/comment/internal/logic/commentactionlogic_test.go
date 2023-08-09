package logic

import (
	"context"
	"flag"
	"testing"
	"tiny-tiktok/service/comment/internal/config"
	"tiny-tiktok/service/comment/internal/svc"
	"tiny-tiktok/service/comment/pb/comment"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "../../etc/comment.yaml", "the config file")
var ctx = context.Background()
var svcCtx *svc.ServiceContext

func TestMain(m *testing.M) {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx = svc.NewServiceContext(c)

	m.Run()
}

func TestCommentAction(t *testing.T) {
	type args struct {
		userId      int64
		videoId     int64
		actionType  int32
		commentText string
		commentId   int64
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
				userId:      1,
				videoId:     1,
				actionType:  1,
				commentText: "comment test1",
			},
			wantAct: "Comment successfully",
			wantErr: false,
		},
		{
			testName: "valid comment",
			args: args{
				userId:      1,
				videoId:     1,
				actionType:  1,
				commentText: "comment test2",
			},
			wantAct: "Comment successfully",
			wantErr: false,
		},
		{
			testName: "valid comment",
			args: args{
				userId:      1,
				videoId:     1,
				actionType:  1,
				commentText: "comment test3",
			},
			wantAct: "Comment successfully",
			wantErr: false,
		},
		{
			testName: "valid comment deletion",
			args: args{
				userId:     1,
				videoId:    1,
				actionType: 2,
				commentId:  1,
			},
			wantAct: "Delete the comment successfully",
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			resp, err := NewCommentActionLogic(ctx, svcCtx).CommentAction(&comment.CommentActionReq{
				UserId:      test.args.userId,
				VideoId:     test.args.videoId,
				ActionType:  test.args.actionType,
				CommentText: test.args.commentText,
				CommentId:   test.args.commentId,
			})
			if (err != nil) != test.wantErr {
				t.Errorf("CommentAction() error: %v, wantErr %v", err, test.wantErr)
				return
			}
			if resp.StatusMsg != test.wantAct {
				t.Errorf("wrong action: want %s; real %s", test.wantAct, resp.StatusMsg)
				return
			}
		})
	}
}
