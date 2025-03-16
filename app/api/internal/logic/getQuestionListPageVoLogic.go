package logic

import (
	"context"

	"app/api/internal/svc"
	"app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionListPageVoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取题目列表
func NewGetQuestionListPageVoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionListPageVoLogic {
	return &GetQuestionListPageVoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQuestionListPageVoLogic) GetQuestionListPageVo(req *types.QuestionQueryReq) (resp *types.QuestionListResp, err error) {
	// todo: 实现从数据库中获取数据
	resp = &types.QuestionListResp{
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
