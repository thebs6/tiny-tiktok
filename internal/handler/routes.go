// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	core "tiny-tiktok/internal/handler/core"
	extra_first "tiny-tiktok/internal/handler/extra_first"
	extra_second "tiny-tiktok/internal/handler/extra_second"
	"tiny-tiktok/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/feed",
				Handler: core.FeedHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: core.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: core.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user",
				Handler: core.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/publish/action",
				Handler: core.PublishActionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/publish/list",
				Handler: core.PublishListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/favorite/action",
				Handler: extra_first.FavoriteActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/favorite/list",
				Handler: extra_first.FavoriteListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/comment/action",
				Handler: extra_first.CommentActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/comment/list",
				Handler: extra_first.CommentListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/relation/action",
				Handler: extra_second.RelationActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/relation/follow/list",
				Handler: extra_second.RelationFollowListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/relation/follower/list",
				Handler: extra_second.RelationFollowerListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/relation/friend/list",
				Handler: extra_second.RelationFriendListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/message/action",
				Handler: extra_second.MessageActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/message/chat",
				Handler: extra_second.MessageChatHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin"),
	)
}
