package repository

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId     int64     `gorm:"column:uid"`
	Name       string    `gorm:"column:name"`
	Password   string    `gorm:"column:password"`
	Avatar     string    `gorm:"column:avatar"`
	CreateTime time.Time `gorm:"column:create_time"`
	ModifyTime time.Time `gorm:"column:modify_time"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) QueryUserById(uid int64) (*User, error) {
	// var user User
	user := &User{}
	err := db.Where("uid = ?", uid).First(user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		// util.Logger.Error("find user by id err:" + err.Error())
		return nil, fmt.Errorf("repo: %v", err)
	}
	return user, nil
}

func (*UserDao) QueryUserByName(name string) (*User, error) {
	user := &User{}
	err := db.Where("name = ?", name).First(user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		// util.Logger.Error("batch find user by id err:" + err.Error())
		return nil, fmt.Errorf("repo: %v", err)
	}
	return user, nil
}

func (*UserDao) CreateNewUser(user *User) error {
	if err := db.Create(user).Error; err != nil {
		return fmt.Errorf("repo: %v", err)
	}
	return nil
}
