// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/feed/create",
				Handler: CreateVideoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/feed/UserVideosList",
				Handler: UserVideosListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/feed/VideosList",
				Handler: VideosListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/feed/CategoryVideosList",
				Handler: CategoryVideosListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/feed/recommends",
				Handler: RecommendVideosHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/feed/populars",
				Handler: PopularVideosHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/feed/createst",
				Handler: CreateVideoTestHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/feed/duration",
				Handler: DurationTestHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/feed/history",
				Handler: HistoryVideosHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/feed/neighbors",
				Handler: NeighborsVideosHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/feed/deleteViedo",
				Handler: DeleteVideoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
