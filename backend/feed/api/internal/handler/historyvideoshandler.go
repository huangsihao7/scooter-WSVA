package handler

import (
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/logic"
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HistoryVideosHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHistoryVideosLogic(r.Context(), svcCtx)
		resp, err := l.HistoryVideos()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
