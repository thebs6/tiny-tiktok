package repository

import (
	"fmt"
	"testing"
)

func TestQueryUserById(t *testing.T) {

	tests := []struct {
		testName string
		userId   int64
		wantUser bool
		wantErr  bool
	}{
		{
			testName: "exsiting id",
			userId:   1,
			wantUser: true,
			wantErr:  false,
		},
		{
			testName: "not exsiting id",
			userId:   4,
			wantUser: false,
			wantErr:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			user, err := NewUserDaoInstance().QueryUserById(test.userId)
			fmt.Println(user.UserId, user.Name)
			if (user != nil) != test.wantUser {
				t.Errorf("Register() error: %v, wantUser %v, wantErr %v", err, test.wantUser, test.wantErr)
				return
			}
			if (err != nil) != test.wantErr {
				t.Errorf("Register() error: %v, wantUser %v, wantErr %v", err, test.wantUser, test.wantErr)
				return
			}
		})
	}
}
