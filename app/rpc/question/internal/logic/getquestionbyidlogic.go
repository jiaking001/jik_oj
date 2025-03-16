package logic

import (
	"app/rpc/question/model"
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
	questions, err := model.GetById(l.ctx, int(in.Id))
	if err != nil {
		return nil, err
	}
	return &question.GetQuestionByIdResp{
		Question: &question.Question{
			Id:          int64(questions.ID),
			Title:       questions.Title,
			Content:     questions.Content,
			Tags:        questions.Tags,
			Answer:      questions.Answer,
			SubmitNum:   int64(questions.SubmitNum),
			AcceptedNum: int64(questions.AcceptedNum),
			JudgeCase:   questions.JudgeCase,
			JudgeConfig: questions.JudgeConfig,
			ThumbNum:    int64(questions.ThumbNum),
			FavourNum:   int64(questions.FavourNum),
			UserId:      int64(questions.UserId),
			IsDelete:    int64(questions.IsDelete),
		},
	}, nil
}
