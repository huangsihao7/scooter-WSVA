package handler

import (
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/live/api/internal/logic"
	"github.com/huangsihao7/scooter-WSVA/live/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewStartLogic(r.Context(), svcCtx)
		resp, err := l.Start()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
