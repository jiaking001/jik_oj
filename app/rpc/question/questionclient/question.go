// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: question.proto

package questionclient

import (
	"context"

	"app/rpc/question/types/question"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddQuestionReq      = question.AddQuestionReq
	AddQuestionResp     = question.AddQuestionResp
	DelQuestionReq      = question.DelQuestionReq
	DelQuestionResp     = question.DelQuestionResp
	GetQuestionByIdReq  = question.GetQuestionByIdReq
	GetQuestionByIdResp = question.GetQuestionByIdResp
	Question            = question.Question
	SearchQuestionReq   = question.SearchQuestionReq
	SearchQuestionResp  = question.SearchQuestionResp
	UpdateQuestionReq   = question.UpdateQuestionReq
	UpdateQuestionResp  = question.UpdateQuestionResp

	QuestionZrpcClient interface {
		// -----------------------题目-----------------------
		AddQuestion(ctx context.Context, in *AddQuestionReq, opts ...grpc.CallOption) (*AddQuestionResp, error)
		UpdateQuestion(ctx context.Context, in *UpdateQuestionReq, opts ...grpc.CallOption) (*UpdateQuestionResp, error)
		DelQuestion(ctx context.Context, in *DelQuestionReq, opts ...grpc.CallOption) (*DelQuestionResp, error)
		GetQuestionById(ctx context.Context, in *GetQuestionByIdReq, opts ...grpc.CallOption) (*GetQuestionByIdResp, error)
		SearchQuestion(ctx context.Context, in *SearchQuestionReq, opts ...grpc.CallOption) (*SearchQuestionResp, error)
	}

	defaultQuestionZrpcClient struct {
		cli zrpc.Client
	}
)

func NewQuestionZrpcClient(cli zrpc.Client) QuestionZrpcClient {
	return &defaultQuestionZrpcClient{
		cli: cli,
	}
}

// -----------------------题目-----------------------
func (m *defaultQuestionZrpcClient) AddQuestion(ctx context.Context, in *AddQuestionReq, opts ...grpc.CallOption) (*AddQuestionResp, error) {
	client := question.NewQuestionClient(m.cli.Conn())
	return client.AddQuestion(ctx, in, opts...)
}

func (m *defaultQuestionZrpcClient) UpdateQuestion(ctx context.Context, in *UpdateQuestionReq, opts ...grpc.CallOption) (*UpdateQuestionResp, error) {
	client := question.NewQuestionClient(m.cli.Conn())
	return client.UpdateQuestion(ctx, in, opts...)
}

func (m *defaultQuestionZrpcClient) DelQuestion(ctx context.Context, in *DelQuestionReq, opts ...grpc.CallOption) (*DelQuestionResp, error) {
	client := question.NewQuestionClient(m.cli.Conn())
	return client.DelQuestion(ctx, in, opts...)
}

func (m *defaultQuestionZrpcClient) GetQuestionById(ctx context.Context, in *GetQuestionByIdReq, opts ...grpc.CallOption) (*GetQuestionByIdResp, error) {
	client := question.NewQuestionClient(m.cli.Conn())
	return client.GetQuestionById(ctx, in, opts...)
}

func (m *defaultQuestionZrpcClient) SearchQuestion(ctx context.Context, in *SearchQuestionReq, opts ...grpc.CallOption) (*SearchQuestionResp, error) {
	client := question.NewQuestionClient(m.cli.Conn())
	return client.SearchQuestion(ctx, in, opts...)
}
