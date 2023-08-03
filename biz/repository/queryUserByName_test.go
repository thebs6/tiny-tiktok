package repository

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := Init(); err != nil {
		os.Exit(1)
	}
	// if err := util.InitLogger(); err != nil {
	// 	os.Exit(1)
	// }
	m.Run()
}

func TestQueryUserByName(t *testing.T) {

	tests := []struct {
		testName string
		userName string
		wantUser bool
		wantErr  bool
	}{
		{
			testName: "not existing name and right password",
			userName: "Alex",
			wantUser: false,
			wantErr:  false,
		},
		{
			testName: "exsiting name",
			userName: "Jerry",
			wantUser: true,
			wantErr:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			user, err := NewUserDaoInstance().QueryUserByName(test.userName)
			if (err != nil) != test.wantErr {
				t.Errorf("do not want err but err occur")
				return
			}
			if (user != nil) != test.wantUser {
				t.Errorf("username do not exist but return user")
				return
			}
		})
	}
}
