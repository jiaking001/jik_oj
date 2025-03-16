package logic

import (
	"app/api/internal/svc"
	"app/api/internal/types"
	"app/rpc/question/types/question"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据题目id获取题目详情
func NewGetQuestionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionByIdLogic {
	return &GetQuestionByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQuestionByIdLogic) GetQuestionById(req *types.GetQuestionByIdReq) (resp *types.QuestionVO, err error) {
	idReq := &question.GetQuestionByIdReq{
		Id: int64(req.Id),
	}
	q, err := l.svcCtx.Question.GetQuestionById(l.ctx, idReq)
	if err != nil {
		return
	}
	resp = &types.QuestionVO{
		Id:         int(q.Question.Id),
		Title:      q.Question.Title,
		Content:    q.Question.Content,
		CreateTime: strconv.FormatInt(q.Question.CreateTime, 10),
	}

	return
}
