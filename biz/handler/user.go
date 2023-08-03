package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gocx/tinyDouyin/biz/service"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	user_id, err := service.Register(username, password)

	if err != nil {
		// c.JSON(http.StatusOK, UserLoginResponse{
		// 	Response: Response{
		// 		StatusCode: 1,
		// 		StatusMsg:  err.Error()},
		// })
		c.Set("status_msg", err.Error())
	} else {
		// c.JSON(http.StatusOK, UserLoginResponse{
		// 	Response: Response{
		// 		StatusCode: 0,
		// 		StatusMsg:  "Register success"},
		// 	UserId: user_id,
		// })
		c.Set("user_id", user_id)
	}
}

func Login(ctx context.Context, c *app.RequestContext) {
	fmt.Println("handler login1")
	username := c.Query("username")
	password := c.Query("password")

	user_id, err := service.Login(username, password)
	if err != nil {
		// c.JSON(http.StatusOK, UserLoginResponse{
		// 	UserId: -1,
		// })
		c.Set("status_msg", err.Error())
	} else {
		// c.JSON(http.StatusOK, UserLoginResponse{
		// 	UserId: user_id,
		// })
		// return user_id, ni
		c.Set("user_id", user_id)
	}
	fmt.Println("handler login2")
}

func UserInfo(ctx context.Context, c *app.RequestContext) {

	// if  {
	// 	c.JSON(http.StatusOK, UserResponse{
	// 		Response: Response{StatusCode: 0},
	// 		User:     user,
	// 	})
	// } else {
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	})
	// }
}
