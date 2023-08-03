package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gocx/tinyDouyin/biz/handler"
	"github.com/hertz-contrib/jwt"
)

type login struct {
	Username string `form:"username,required" json:"username,required"`
	Password string `form:"password,required" json:"password,required"`
}

var identityKey = "id"

// User demo
type User struct {
	UserId    int64
	UserName  string
	FirstName string
	LastName  string
}

func authMidwareInit() *jwt.HertzJWTMiddleware {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:            "test zone",
		Key:              []byte("secret key"),
		SigningAlgorithm: "HS256",
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		IdentityKey:      identityKey,
		TokenLookup:      "query: token",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{

					identityKey: v.UserId,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var user_id int64
			if _, exsit := c.Get("type"); !exsit {
				// login logic
				// var loginVals login
				// if err := c.BindAndValidate(&loginVals); err != nil {
				// 	return "", jwt.ErrMissingLoginValues
				// }
				// userID := loginVals.Username
				// password := loginVals.Password

				handler.Login(ctx, c)
			}
			uid_inter, _ := c.Get("user_id")
			user_id, _ = uid_inter.(int64)
			// if _, sta_msg_exsit := c.Get("status_msg"); sta_msg_exsit {
			// 	return nil, jwt.ErrFailedAuthentication
			// }

			return &User{
				UserId: user_id,
			}, nil
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			uid, _ := c.Get("user_id")
			status_msg, sta_msg_exsit := c.Get("status_msg")

			if _, exsit := c.Get("type"); exsit {
				// register logic
				if !sta_msg_exsit {
					c.Set("status_msg", "Register successfully")
				} else {
					c.JSON(http.StatusOK, map[string]interface{}{
						"status_msg": status_msg,
					})
					return
				}
			} else {
				// login logic
				if !sta_msg_exsit {
					c.Set("status_msg", "Login successfully")
				} else {
					c.JSON(http.StatusOK, map[string]interface{}{
						"status_msg": status_msg,
					})
					return
				}
			}
			status_msg, _ = c.Get("status_msg")
			c.JSON(http.StatusOK, map[string]interface{}{
				"status_msg": status_msg,
				"user_id":    uid,
				"token":      token,
			})
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware
}
