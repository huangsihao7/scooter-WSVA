// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/danmu/listv3",
				Handler: GetDanmuListHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/comment/action",
				Handler: CommentActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/comment/list",
				Handler: GetCommentListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/danmu/action",
				Handler: DanmuActionHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
