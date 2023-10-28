package handler

import (
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/mq/api/internal/logic"
	"github.com/huangsihao7/scooter-WSVA/mq/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/mq/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func NullHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewNullLogic(r.Context(), svcCtx)
		resp, err := l.Null(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
