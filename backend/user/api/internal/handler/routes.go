// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/uploadImg",
				Handler: UploadImgHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/userinfo",
				Handler: UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/upload",
				Handler: UploadHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/update",
				Handler: UpdateHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
