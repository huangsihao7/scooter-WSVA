package handler

import (
	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusUnauthorized, &types.ActionResp{401, "鉴权失败"})
}
