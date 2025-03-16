package model

import (
	"app/rpc/question/model/dal/mysql"
	"app/rpc/question/types/question"
	"context"
	"time"
)

// Question 表示题目表
type Question struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(512);comment:标题" json:"title"`
	Content     string    `gorm:"type:text;comment:内容" json:"content"`
	Tags        string    `gorm:"type:varchar(1024);comment:标签列表（json 数组）" json:"tags"`
	Answer      string    `gorm:"type:text;comment:题目答案" json:"answer"`
	SubmitNum   int       `gorm:"default:0;comment:题目提交数" json:"submitNum"`
	AcceptedNum int       `gorm:"default:0;comment:题目通过数" json:"acceptedNum"`
	JudgeCase   string    `gorm:"type:text;comment:判题用例（json 数组）" json:"judgeCase"`
	JudgeConfig string    `gorm:"type:text;comment:判题配置（json 对象）" json:"judgeConfig"`
	ThumbNum    int       `gorm:"default:0;comment:点赞数" json:"thumbNum"`
	FavourNum   int       `gorm:"default:0;comment:收藏数" json:"favourNum"`
	UserId      uint64    `gorm:"notNull;comment:创建用户 id" json:"userId"`
	CreateTime  time.Time `gorm:"default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"`
	UpdateTime  time.Time `gorm:"default:CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"`
	IsDelete    int8      `gorm:"default:0;comment:是否删除" json:"isDelete"`
}

// TableName 返回表名
func (Question) TableName() string {
	return "questions"
}

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

func GetById(ctx context.Context, id int) (Question, error) {
	var questions Question
	if err := mysql.DB.Where("id = ?", id).Find(&questions).Error; err != nil {
		return Question{}, err
	}
	return questions, nil
}
