package logic

import (
	"app/api/internal/utils"
	"app/rpc/question/types/question"
	"context"
	"strconv"

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
	searchQuestionReq := &question.SearchQuestionReq{
		Limit: int64(req.PageSize),
		Page:  int64(req.Current),
	}
	questions, err := l.svcCtx.Question.SearchQuestion(l.ctx, searchQuestionReq)
	if err != nil {
		return nil, err
	}
	var questionVO []types.QuestionVO
	for _, q := range questions.Question {
		tags, err := utils.StringToStrings(q.Tags)
		if err != nil {
			return nil, err
		}
		questionVO = append(questionVO, types.QuestionVO{
			AcceptedNum: int(q.AcceptedNum),                  // 通过次数
			Content:     q.Content,                           // 问题内容
			CreateTime:  strconv.FormatInt(q.CreateTime, 10), // 创建时间
			Id:          int(q.Id),                           // 问题ID
			SubmitNum:   int(q.SubmitNum),                    // 提交次数
			Tags:        tags,                                // 问题标签
			Title:       q.Title,                             // 问题标题
		})
	}
	resp = &types.QuestionListResp{
		Current: req.Current,
		Pages:   (int(questions.Total) + req.PageSize - 1) / req.PageSize,
		Records: questionVO,
		Size:    req.PageSize,
		Total:   int(questions.Total),
	}
	return
}
