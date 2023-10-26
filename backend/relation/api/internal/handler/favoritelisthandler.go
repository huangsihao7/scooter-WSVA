package handler

import (
	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/types"
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/logic"
	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewFavoriteListLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
