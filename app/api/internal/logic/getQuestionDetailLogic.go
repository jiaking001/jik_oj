package logic

import (
	"app/rpc/question/types/question"
	"context"
	"strconv"

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
	searchQuestionReq := &question.SearchQuestionReq{
		Limit: int64(req.PageSize),
		Page:  int64(req.Current),
	}
	questions, err := l.svcCtx.Question.SearchQuestion(l.ctx, searchQuestionReq)
	if err != nil {
		return nil, err
	}
	var q []types.Question
	for _, v := range questions.Question {
		q = append(q, types.Question{
			AcceptedNum: int(v.AcceptedNum),                  // 通过次数
			Content:     v.Content,                           // 问题内容
			CreateTime:  strconv.FormatInt(v.CreateTime, 10), // 创建时间
			Id:          int(v.Id),                           // 问题ID
			SubmitNum:   int(v.SubmitNum),                    // 提交次数
			Tags:        v.Tags,                              // 问题标签
			Title:       v.Title,                             // 问题标题
		})
	}
	resp = &types.PageQuestionResq{
		Current: req.Current,
		Pages:   (int(questions.Total) + req.PageSize - 1) / req.PageSize,
		Records: q,
		Size:    req.PageSize,
		Total:   int(questions.Total),
	}
	return
}
