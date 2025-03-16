// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: question.proto

package question

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Question_AddQuestion_FullMethodName     = "/question.question/AddQuestion"
	Question_UpdateQuestion_FullMethodName  = "/question.question/UpdateQuestion"
	Question_DelQuestion_FullMethodName     = "/question.question/DelQuestion"
	Question_GetQuestionById_FullMethodName = "/question.question/GetQuestionById"
	Question_SearchQuestion_FullMethodName  = "/question.question/SearchQuestion"
)

// QuestionClient is the client API for Question service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionClient interface {
	// -----------------------题目-----------------------
	AddQuestion(ctx context.Context, in *AddQuestionReq, opts ...grpc.CallOption) (*AddQuestionResp, error)
	UpdateQuestion(ctx context.Context, in *UpdateQuestionReq, opts ...grpc.CallOption) (*UpdateQuestionResp, error)
	DelQuestion(ctx context.Context, in *DelQuestionReq, opts ...grpc.CallOption) (*DelQuestionResp, error)
	GetQuestionById(ctx context.Context, in *GetQuestionByIdReq, opts ...grpc.CallOption) (*GetQuestionByIdResp, error)
	SearchQuestion(ctx context.Context, in *SearchQuestionReq, opts ...grpc.CallOption) (*SearchQuestionResp, error)
}

type questionClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionClient(cc grpc.ClientConnInterface) QuestionClient {
	return &questionClient{cc}
}

func (c *questionClient) AddQuestion(ctx context.Context, in *AddQuestionReq, opts ...grpc.CallOption) (*AddQuestionResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddQuestionResp)
	err := c.cc.Invoke(ctx, Question_AddQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionClient) UpdateQuestion(ctx context.Context, in *UpdateQuestionReq, opts ...grpc.CallOption) (*UpdateQuestionResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateQuestionResp)
	err := c.cc.Invoke(ctx, Question_UpdateQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionClient) DelQuestion(ctx context.Context, in *DelQuestionReq, opts ...grpc.CallOption) (*DelQuestionResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelQuestionResp)
	err := c.cc.Invoke(ctx, Question_DelQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionClient) GetQuestionById(ctx context.Context, in *GetQuestionByIdReq, opts ...grpc.CallOption) (*GetQuestionByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetQuestionByIdResp)
	err := c.cc.Invoke(ctx, Question_GetQuestionById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionClient) SearchQuestion(ctx context.Context, in *SearchQuestionReq, opts ...grpc.CallOption) (*SearchQuestionResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchQuestionResp)
	err := c.cc.Invoke(ctx, Question_SearchQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionServer is the server API for Question service.
// All implementations must embed UnimplementedQuestionServer
// for forward compatibility.
type QuestionServer interface {
	// -----------------------题目-----------------------
	AddQuestion(context.Context, *AddQuestionReq) (*AddQuestionResp, error)
	UpdateQuestion(context.Context, *UpdateQuestionReq) (*UpdateQuestionResp, error)
	DelQuestion(context.Context, *DelQuestionReq) (*DelQuestionResp, error)
	GetQuestionById(context.Context, *GetQuestionByIdReq) (*GetQuestionByIdResp, error)
	SearchQuestion(context.Context, *SearchQuestionReq) (*SearchQuestionResp, error)
	mustEmbedUnimplementedQuestionServer()
}

// UnimplementedQuestionServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQuestionServer struct{}

func (UnimplementedQuestionServer) AddQuestion(context.Context, *AddQuestionReq) (*AddQuestionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddQuestion not implemented")
}
func (UnimplementedQuestionServer) UpdateQuestion(context.Context, *UpdateQuestionReq) (*UpdateQuestionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuestion not implemented")
}
func (UnimplementedQuestionServer) DelQuestion(context.Context, *DelQuestionReq) (*DelQuestionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelQuestion not implemented")
}
func (UnimplementedQuestionServer) GetQuestionById(context.Context, *GetQuestionByIdReq) (*GetQuestionByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionById not implemented")
}
func (UnimplementedQuestionServer) SearchQuestion(context.Context, *SearchQuestionReq) (*SearchQuestionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchQuestion not implemented")
}
func (UnimplementedQuestionServer) mustEmbedUnimplementedQuestionServer() {}
func (UnimplementedQuestionServer) testEmbeddedByValue()                  {}

// UnsafeQuestionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionServer will
// result in compilation errors.
type UnsafeQuestionServer interface {
	mustEmbedUnimplementedQuestionServer()
}

func RegisterQuestionServer(s grpc.ServiceRegistrar, srv QuestionServer) {
	// If the following call pancis, it indicates UnimplementedQuestionServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Question_ServiceDesc, srv)
}

func _Question_AddQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddQuestionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServer).AddQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Question_AddQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServer).AddQuestion(ctx, req.(*AddQuestionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Question_UpdateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuestionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServer).UpdateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Question_UpdateQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServer).UpdateQuestion(ctx, req.(*UpdateQuestionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Question_DelQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelQuestionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServer).DelQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Question_DelQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServer).DelQuestion(ctx, req.(*DelQuestionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Question_GetQuestionById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServer).GetQuestionById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Question_GetQuestionById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServer).GetQuestionById(ctx, req.(*GetQuestionByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Question_SearchQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchQuestionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServer).SearchQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Question_SearchQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServer).SearchQuestion(ctx, req.(*SearchQuestionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Question_ServiceDesc is the grpc.ServiceDesc for Question service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Question_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "question.question",
	HandlerType: (*QuestionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddQuestion",
			Handler:    _Question_AddQuestion_Handler,
		},
		{
			MethodName: "UpdateQuestion",
			Handler:    _Question_UpdateQuestion_Handler,
		},
		{
			MethodName: "DelQuestion",
			Handler:    _Question_DelQuestion_Handler,
		},
		{
			MethodName: "GetQuestionById",
			Handler:    _Question_GetQuestionById_Handler,
		},
		{
			MethodName: "SearchQuestion",
			Handler:    _Question_SearchQuestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "question.proto",
}
