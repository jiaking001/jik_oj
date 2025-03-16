package logic

import (
	"context"

	"app/api/internal/svc"
	"app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionQuestionSubmitListPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询题目提交记录
func NewGetQuestionQuestionSubmitListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionQuestionSubmitListPageLogic {
	return &GetQuestionQuestionSubmitListPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQuestionQuestionSubmitListPageLogic) GetQuestionQuestionSubmitListPage(req *types.QuestionSubmitQueryRequest) (resp *types.PageQuestionSubmitVOResp, err error) {
	// todo: 实现从数据库中获取数据
	resp = &types.PageQuestionSubmitVOResp{
		CountId:          "",
		Current:          0,
		MaxLimit:         0,
		OptimizeCountSql: false,
		Orders:           nil,
		Pages:            0,
		Records:          nil,
		SearchCount:      false,
		Size:             0,
		Total:            0,
	}
	return
}
