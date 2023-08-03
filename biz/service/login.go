package service

import (
	"errors"

	"github.com/justGoRun/tinyTiktok/biz/repository"
)

func Login(name, password string) (int64, error) {
	return NewLoginFlow(name, password).Do()
}

func NewLoginFlow(name, password string) *LoginFlow {
	return &LoginFlow{
		name:     name,
		password: password,
	}
}

type LoginFlow struct {
	name     string
	password string
	userId   int64
}

func (f *LoginFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return -1, err
	}
	return f.userId, nil
}

func (f *LoginFlow) checkParam() error {
	user, err := repository.NewUserDaoInstance().QueryUserByName(f.name)
	if user != nil {
		// found
		if f.password != user.Password {
			return errors.New("the password is wrong")
		}
		f.userId = user.UserId
		return nil
	} else if user == nil && err == nil {
		// not found
		return errors.New("this username does not exist")
	}
	return err
}
