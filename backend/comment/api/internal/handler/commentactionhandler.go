package handler

import (
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/logic"
	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ActionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCommentActionLogic(r.Context(), svcCtx)
		resp, err := l.CommentAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
