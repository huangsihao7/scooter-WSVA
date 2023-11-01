package handler

import (
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/user/api/internal/logic"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadImgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadImgLogic(r.Context(), svcCtx)
		resp, err := l.UploadImg(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
