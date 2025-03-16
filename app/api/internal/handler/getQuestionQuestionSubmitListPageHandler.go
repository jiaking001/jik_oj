package handler

import (
	"app/api/internal/utils"
	"net/http"

	"app/api/internal/logic"
	"app/api/internal/svc"
	"app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 查询题目提交记录
func getQuestionQuestionSubmitListPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QuestionSubmitQueryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetQuestionQuestionSubmitListPageLogic(r.Context(), svcCtx)
		resp, err := l.GetQuestionQuestionSubmitListPage(&req)
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
