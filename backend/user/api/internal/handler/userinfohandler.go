package handler

import (
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/types"
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/user/api/internal/logic"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
