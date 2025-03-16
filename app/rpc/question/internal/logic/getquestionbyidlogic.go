package logic

import (
	"context"

	"app/rpc/question/internal/svc"
	"app/rpc/question/types/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetQuestionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionByIdLogic {
	return &GetQuestionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetQuestionByIdLogic) GetQuestionById(in *question.GetQuestionByIdReq) (*question.GetQuestionByIdResp, error) {
	// todo: add your logic here and delete this line

	return &question.GetQuestionByIdResp{}, nil
}
