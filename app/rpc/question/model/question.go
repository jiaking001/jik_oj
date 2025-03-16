package model

import (
	"app/rpc/question/model/dal/mysql"
	"app/rpc/question/types/question"
	"context"
)

// List 获取问题列表
func List(ctx context.Context, current int, pageSize int) ([]*question.Question, int64, error) {
	var questions []*question.Question
	offset := (current - 1) * pageSize // 偏移量
	var total int64                    // 总记录数
	if err := mysql.DB.Table("questions").Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := mysql.DB.Limit(pageSize).Offset(offset).Find(&questions).Error; err != nil {
		return nil, 0, err
	}
	return questions, total, nil
}
