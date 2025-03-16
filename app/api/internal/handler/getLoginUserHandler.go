package handler

import (
	"app/api/internal/utils"
	"net/http"

	"app/api/internal/logic"
	"app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取用户登录信息
func getLoginUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetLoginUserLogic(r.Context(), svcCtx)
		resp, err := l.GetLoginUser()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, utils.Response{
				Code:    0,
				Message: "",
				Data:    resp,
			})
		}
	}
}
