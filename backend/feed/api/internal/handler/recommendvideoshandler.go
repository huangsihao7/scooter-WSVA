package handler

import (
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/logic"
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RecommendVideosHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendVideosListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRecommendVideosLogic(r.Context(), svcCtx)
		resp, err := l.RecommendVideos(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
