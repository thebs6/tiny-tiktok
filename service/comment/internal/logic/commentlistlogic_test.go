package logic

import (
	"testing"
	"tiny-tiktok/service/comment/pb/comment"
)

func TestCommentList(t *testing.T) {
	type args struct {
		videoId int64
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
				videoId: 1,
			},
			wantErr: false,
		},
		{
			testName: "valid comment",
			args: args{
				videoId: 2,
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			_, err := NewCommentListLogic(ctx, svcCtx).CommentList(&comment.CommentListReq{
				VideoId: test.args.videoId,
			})
			if (err != nil) != test.wantErr {
				t.Errorf("CommentAction() error: %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}
