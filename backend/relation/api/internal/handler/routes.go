// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/relation/action",
				Handler: FavoriteActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/relation/favoriteList",
				Handler: FavoriteListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/relation/followerList",
				Handler: FollowerListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/relation/friendList",
				Handler: FriendListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
