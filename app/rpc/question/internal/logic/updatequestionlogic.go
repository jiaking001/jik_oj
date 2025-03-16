package logic

import (
	"context"

	"app/rpc/question/internal/svc"
	"app/rpc/question/types/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateQuestionLogic {
	return &UpdateQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateQuestionLogic) UpdateQuestion(in *question.UpdateQuestionReq) (*question.UpdateQuestionResp, error) {
	// todo: add your logic here and delete this line

	return &question.UpdateQuestionResp{}, nil
}
