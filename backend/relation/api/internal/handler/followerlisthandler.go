package handler

import (
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/logic"
	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FollowerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowerListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFollowerListLogic(r.Context(), svcCtx)
		resp, err := l.FollowerList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
