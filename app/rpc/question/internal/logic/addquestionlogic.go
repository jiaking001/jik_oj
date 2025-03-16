package logic

import (
	"context"

	"app/rpc/question/internal/svc"
	"app/rpc/question/types/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddQuestionLogic {
	return &AddQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------题目-----------------------
func (l *AddQuestionLogic) AddQuestion(in *question.AddQuestionReq) (*question.AddQuestionResp, error) {
	// todo: add your logic here and delete this line

	return &question.AddQuestionResp{}, nil
}
