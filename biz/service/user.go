package service

import "github.com/gocx/tinyDouyin/biz/repository"

type LoginData struct {
}

func Register(name, password string) {
	user, err := repository.NewUserDaoInstance().QueryUserByName(name)

	return user, err
}

func Login(name, password string) (*repository.User, error) {
	user, err := repository.NewUserDaoInstance().QueryUserByName(name)
	return user, err
}
