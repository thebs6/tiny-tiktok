package logic

import (
	"testing"
	"tiny-tiktok/service/user/pb/user"
)

func TestRegister(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		testName string
		args     args
		want     bool
		wantErr  bool
	}{
		{
			testName: "not exsiting name",
			args: args{
				username: "Alex",
				password: "123456",
			},
			want:    true,
			wantErr: false,
		},
		{
			testName: "existing name",
			args: args{
				username: "gao",
				password: "123456",
			},
			want:    false,
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			resp, err := NewRegisterLogic(ctx, svcCtx).Register(&user.RegisterReq{
				Username: test.args.username,
				Password: test.args.password,
			})
			if (err != nil) != test.wantErr {
				t.Errorf("Register() error: %v, wantErr %v", err, test.wantErr)
				return
			}
			if (resp.UserId != -1) != test.want {
				t.Errorf("Register() error: want %v", test.want)
				return
			}
		})
	}
}
