package model

import (
	"context"
	"testing"
)

func TestTrans(t *testing.T) {
	NewDBDao()

	type args struct {
		userId     int64
		videoId    int64
		actionType int64
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
				userId:     2,
				videoId:    1,
				actionType: 1,
			},
			wantErr: false,
		},
		{
			testName: "valid favor",
			args: args{
				userId:     2,
				videoId:    1,
				actionType: 1,
			},
			wantErr: false,
		},
		{
			testName: "valid favor",
			args: args{
				userId:     3,
				videoId:    1,
				actionType: 1,
			},
			wantErr: false,
		},
		{
			testName: "valid cancel favor",
			args: args{
				userId:     3,
				videoId:    1,
				actionType: 2,
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		err := Trans(context.Background(), test.args.userId, test.args.videoId, test.args.actionType)
		if (err != nil) != test.wantErr {
			t.Errorf("FavoriteAction() error: %v", err)
			return
		}
	}
}
