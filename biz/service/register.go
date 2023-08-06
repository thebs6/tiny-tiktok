package service

import (
	"errors"
	"sync"
	"time"

	"github.com/justGoRun/tinyTiktok/biz/repository"
)

func Register(name, password string) (int64, error) {
	return NewRegisterFlow(name, password).Do()
}
func NewRegisterFlow(name, password string) *RegisterFlow {
	return &RegisterFlow{
		name:     name,
		password: password,
	}
}

type RegisterFlow struct {
	name     string
	password string
	userId   int64
	nameMu   sync.Mutex
}

func (f *RegisterFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return -1, err
	}
	if err := f.register(); err != nil {
		return -1, err
	}
	return f.userId, nil
}

func (f *RegisterFlow) checkParam() error {
	if len(f.password) < 5 {
		return errors.New("this password is too short")
	}

	f.nameMu.Lock()
	user, err := repository.NewUserDaoInstance().QueryUserByName(f.name)
	f.nameMu.Unlock()

	if user != nil {
		return errors.New("this user name has been used")
	}

	return err
}

func (f *RegisterFlow) register() error {
	user := &repository.User{
		Name:       f.name,
		Password:   f.password,
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
	}
	if err := repository.NewUserDaoInstance().CreateNewUser(user); err != nil {
		return err
	}
	f.userId = user.UserId
	return nil
}
