package logic

import (
	"context"

	"app/api/internal/svc"
	"app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询题目详情
func NewGetQuestionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionDetailLogic {
	return &GetQuestionDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQuestionDetailLogic) GetQuestionDetail(req *types.QuestionQueryRequest) (resp *types.PageQuestionResq, err error) {
	// todo: 从数据库中获取数据
	resp = &types.PageQuestionResq{
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
