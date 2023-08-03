package service

import (
	"os"
	"testing"

	"github.com/gocx/tinyDouyin/biz/repository"
)

func TestMain(m *testing.M) {
	if err := repository.Init(); err != nil {
		os.Exit(1)
	}
	// if err := util.InitLogger(); err != nil {
	// 	os.Exit(1)
	// }
	m.Run()
}
func TestLogin(t *testing.T) {
	type args struct {
		name     string
		password string
	}
	tests := []struct {
		testName string
		args     args
		wantErr  bool
	}{
		{
			testName: "existing name and right password",
			args: args{
				name:     "Jerry",
				password: "password",
			},
			wantErr: false,
		},
		{
			testName: "not existing name",
			args: args{
				name:     "Alex",
				password: "123456",
			},
			wantErr: true,
		},
		{
			testName: "existing name and wrong password",
			args: args{
				name:     "Jerry",
				password: "wrong_password",
			},
			wantErr: true,
		},
		{
			testName: "not existing name",
			args: args{
				name:     "J",
				password: "wrong_password",
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			_, err := Login(test.args.name, test.args.password)
			if (err != nil) != test.wantErr {
				t.Errorf("Login() error: %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}
