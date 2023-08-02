package repository

import (
	"sync"
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId     int64     `gorm:"column:user_id"`
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
	var user User
	err := db.Where("uid = ?", uid).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		// util.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) QueryUserByName(name string) (*User, error) {
	var user User
	err := db.Where("name = ?", name).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		// util.Logger.Error("batch find user by id err:" + err.Error())
		return nil, err
	}

	return &user, nil
}
