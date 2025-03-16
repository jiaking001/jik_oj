package logic

import (
	"app/rpc/question/model"
	"context"

	"app/rpc/question/internal/svc"
	"app/rpc/question/types/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchQuestionLogic {
	return &SearchQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchQuestionLogic) SearchQuestion(in *question.SearchQuestionReq) (*question.SearchQuestionResp, error) {
	questions, total, err := model.List(l.ctx, int(in.Page), int(in.Limit))
	if err != nil {
		return nil, err
	}
	return &question.SearchQuestionResp{
		Total:    total,
		Question: questions,
	}, nil
}
