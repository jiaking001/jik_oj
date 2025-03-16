package svc

import (
	"app/api/internal/config"
	"app/rpc/question/questionclient"
	"app/rpc/question/types/question"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	Question question.QuestionClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Question: questionclient.NewQuestionZrpcClient(zrpc.MustNewClient(c.RPC)),
	}
}
