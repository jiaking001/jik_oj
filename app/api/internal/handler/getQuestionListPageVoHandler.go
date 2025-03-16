package handler

import (
	"app/api/internal/utils"
	"net/http"

	"app/api/internal/logic"
	"app/api/internal/svc"
	"app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取题目列表
func getQuestionListPageVoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QuestionQueryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetQuestionListPageVoLogic(r.Context(), svcCtx)
		resp, err := l.GetQuestionListPageVo(&req)
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
