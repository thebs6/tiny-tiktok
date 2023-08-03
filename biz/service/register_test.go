package service

import (
	"testing"
)

func TestRegister(t *testing.T) {
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
			testName: "not existing name",
			args: args{
				name:     "Alex",
				password: "123456",
			},
			wantErr: false,
		},
		{
			testName: "exsiting name",
			args: args{
				name:     "Jerry",
				password: "123456",
			},
			wantErr: true,
		},
		{
			testName: "too short password",
			args: args{
				name:     "Ben",
				password: "123",
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			_, err := Register(test.args.name, test.args.password)
			if (err != nil) != test.wantErr {
				t.Errorf("Register() error: %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}
