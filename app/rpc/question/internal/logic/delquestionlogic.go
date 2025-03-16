package logic

import (
	"context"

	"app/rpc/question/internal/svc"
	"app/rpc/question/types/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelQuestionLogic {
	return &DelQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelQuestionLogic) DelQuestion(in *question.DelQuestionReq) (*question.DelQuestionResp, error) {
	// todo: add your logic here and delete this line

	return &question.DelQuestionResp{}, nil
}
