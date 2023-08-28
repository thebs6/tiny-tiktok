package logic

import (
	"testing"
	"tiny-tiktok/service/user/pb/user"
)

func TestUserInfoList(t *testing.T) {
	type args struct {
		user_id_list []int64
	}
	tests := []struct {
		testName string
		args     args
		want     bool
		wantErr  bool
	}{
		{
			testName: "test 1",
			args: args{
				user_id_list: []int64{1, 2, 3},
			},
			want:    true,
			wantErr: false,
		},
		{
			testName: "test 2",
			args: args{
				user_id_list: make([]int64, 1000),
			},
			want:    true,
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			_, err := NewUserInfoListLogic(ctx, svcCtx).UserInfoList(&user.UserInfoListReq{
				UserIdList: test.args.user_id_list,
			})
			if (err != nil) != test.wantErr {
				t.Errorf("UserInfoList() error: %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}
