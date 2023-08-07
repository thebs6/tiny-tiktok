package repository

import (
	"context"
	"sync"
	"time"
)

type User struct {
	Id            int64     `gorm:"column:id"`
	Username      string    `gorm:"column:username"`
	Password      string    `gorm:"column:password"`
	FollowCount   string    `gorm:"column:follow_count"`
	FollowerCount string    `gorm:"column:follower_count"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at"`
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

func (*UserDao) QueryUserById(id int64) (*User, error) {
	user := &User{}
	query := "select id, username, password, follow_count, follower_count, created_at, updated_at, deleted_at from user where id=?"
	err := db.QueryRowCtx(context.Background(), user, query, id)

	return user, err
}

func (*UserDao) QueryUserByName(username string) (*User, error) {
	user := &User{}
	query := "select id, username, password, follow_count, follower_count, created_at, updated_at, deleted_at from user where username=?"
	err := db.QueryRowCtx(context.Background(), user, query, username)

	return user, err
}

func (*UserDao) CreateNewUser(user *User) (int64, error) {
	r, err := db.ExecCtx(context.Background(), "insert into user (username, password, created_at, updated_at) values (?, ?, ?)", user.Username)
	if err != nil {
		return -1, err
	}
	id, err := r.LastInsertId()

	return id, err
}
