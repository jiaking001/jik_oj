// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.19.4
// source: question.proto

package question

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// --------------------------------题目--------------------------------
type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                   //id
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`              //标题
	Content     string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`          //内容
	Tags        string `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`                //标签列表（json 数组）
	Answer      string `protobuf:"bytes,5,opt,name=answer,proto3" json:"answer,omitempty"`            //题目答案
	SubmitNum   int64  `protobuf:"varint,6,opt,name=submitNum,proto3" json:"submitNum,omitempty"`     //题目提交数
	AcceptedNum int64  `protobuf:"varint,7,opt,name=acceptedNum,proto3" json:"acceptedNum,omitempty"` //题目通过数
	JudgeCase   string `protobuf:"bytes,8,opt,name=judgeCase,proto3" json:"judgeCase,omitempty"`      //判题用例（json 数组）
	JudgeConfig string `protobuf:"bytes,9,opt,name=judgeConfig,proto3" json:"judgeConfig,omitempty"`  //判题配置（json 对象）
	ThumbNum    int64  `protobuf:"varint,10,opt,name=thumbNum,proto3" json:"thumbNum,omitempty"`      //点赞数
	FavourNum   int64  `protobuf:"varint,11,opt,name=favourNum,proto3" json:"favourNum,omitempty"`    //收藏数
	UserId      int64  `protobuf:"varint,12,opt,name=userId,proto3" json:"userId,omitempty"`          //创建用户 id
	CreateTime  int64  `protobuf:"varint,13,opt,name=createTime,proto3" json:"createTime,omitempty"`  //创建时间
	UpdateTime  int64  `protobuf:"varint,14,opt,name=updateTime,proto3" json:"updateTime,omitempty"`  //更新时间
	IsDelete    int64  `protobuf:"varint,15,opt,name=isDelete,proto3" json:"isDelete,omitempty"`      //是否删除
}

func (x *Question) Reset() {
	*x = Question{}
	mi := &file_question_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{0}
}

func (x *Question) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Question) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Question) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Question) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Question) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *Question) GetSubmitNum() int64 {
	if x != nil {
		return x.SubmitNum
	}
	return 0
}

func (x *Question) GetAcceptedNum() int64 {
	if x != nil {
		return x.AcceptedNum
	}
	return 0
}

func (x *Question) GetJudgeCase() string {
	if x != nil {
		return x.JudgeCase
	}
	return ""
}

func (x *Question) GetJudgeConfig() string {
	if x != nil {
		return x.JudgeConfig
	}
	return ""
}

func (x *Question) GetThumbNum() int64 {
	if x != nil {
		return x.ThumbNum
	}
	return 0
}

func (x *Question) GetFavourNum() int64 {
	if x != nil {
		return x.FavourNum
	}
	return 0
}

func (x *Question) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Question) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Question) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *Question) GetIsDelete() int64 {
	if x != nil {
		return x.IsDelete
	}
	return 0
}

type AddQuestionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`              //标题
	Content     string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`          //内容
	Tags        string `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`                //标签列表（json 数组）
	Answer      string `protobuf:"bytes,4,opt,name=answer,proto3" json:"answer,omitempty"`            //题目答案
	SubmitNum   int64  `protobuf:"varint,5,opt,name=submitNum,proto3" json:"submitNum,omitempty"`     //题目提交数
	AcceptedNum int64  `protobuf:"varint,6,opt,name=acceptedNum,proto3" json:"acceptedNum,omitempty"` //题目通过数
	JudgeCase   string `protobuf:"bytes,7,opt,name=judgeCase,proto3" json:"judgeCase,omitempty"`      //判题用例（json 数组）
	JudgeConfig string `protobuf:"bytes,8,opt,name=judgeConfig,proto3" json:"judgeConfig,omitempty"`  //判题配置（json 对象）
	ThumbNum    int64  `protobuf:"varint,9,opt,name=thumbNum,proto3" json:"thumbNum,omitempty"`       //点赞数
	FavourNum   int64  `protobuf:"varint,10,opt,name=favourNum,proto3" json:"favourNum,omitempty"`    //收藏数
	UserId      int64  `protobuf:"varint,11,opt,name=userId,proto3" json:"userId,omitempty"`          //创建用户 id
	CreateTime  int64  `protobuf:"varint,12,opt,name=createTime,proto3" json:"createTime,omitempty"`  //创建时间
	UpdateTime  int64  `protobuf:"varint,13,opt,name=updateTime,proto3" json:"updateTime,omitempty"`  //更新时间
	IsDelete    int64  `protobuf:"varint,14,opt,name=isDelete,proto3" json:"isDelete,omitempty"`      //是否删除
}

func (x *AddQuestionReq) Reset() {
	*x = AddQuestionReq{}
	mi := &file_question_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddQuestionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddQuestionReq) ProtoMessage() {}

func (x *AddQuestionReq) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddQuestionReq.ProtoReflect.Descriptor instead.
func (*AddQuestionReq) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{1}
}

func (x *AddQuestionReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddQuestionReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AddQuestionReq) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *AddQuestionReq) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *AddQuestionReq) GetSubmitNum() int64 {
	if x != nil {
		return x.SubmitNum
	}
	return 0
}

func (x *AddQuestionReq) GetAcceptedNum() int64 {
	if x != nil {
		return x.AcceptedNum
	}
	return 0
}

func (x *AddQuestionReq) GetJudgeCase() string {
	if x != nil {
		return x.JudgeCase
	}
	return ""
}

func (x *AddQuestionReq) GetJudgeConfig() string {
	if x != nil {
		return x.JudgeConfig
	}
	return ""
}

func (x *AddQuestionReq) GetThumbNum() int64 {
	if x != nil {
		return x.ThumbNum
	}
	return 0
}

func (x *AddQuestionReq) GetFavourNum() int64 {
	if x != nil {
		return x.FavourNum
	}
	return 0
}

func (x *AddQuestionReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddQuestionReq) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *AddQuestionReq) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *AddQuestionReq) GetIsDelete() int64 {
	if x != nil {
		return x.IsDelete
	}
	return 0
}

type AddQuestionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddQuestionResp) Reset() {
	*x = AddQuestionResp{}
	mi := &file_question_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddQuestionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddQuestionResp) ProtoMessage() {}

func (x *AddQuestionResp) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddQuestionResp.ProtoReflect.Descriptor instead.
func (*AddQuestionResp) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{2}
}

type UpdateQuestionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                   //id
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`              //标题
	Content     string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`          //内容
	Tags        string `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`                //标签列表（json 数组）
	Answer      string `protobuf:"bytes,5,opt,name=answer,proto3" json:"answer,omitempty"`            //题目答案
	SubmitNum   int64  `protobuf:"varint,6,opt,name=submitNum,proto3" json:"submitNum,omitempty"`     //题目提交数
	AcceptedNum int64  `protobuf:"varint,7,opt,name=acceptedNum,proto3" json:"acceptedNum,omitempty"` //题目通过数
	JudgeCase   string `protobuf:"bytes,8,opt,name=judgeCase,proto3" json:"judgeCase,omitempty"`      //判题用例（json 数组）
	JudgeConfig string `protobuf:"bytes,9,opt,name=judgeConfig,proto3" json:"judgeConfig,omitempty"`  //判题配置（json 对象）
	ThumbNum    int64  `protobuf:"varint,10,opt,name=thumbNum,proto3" json:"thumbNum,omitempty"`      //点赞数
	FavourNum   int64  `protobuf:"varint,11,opt,name=favourNum,proto3" json:"favourNum,omitempty"`    //收藏数
	UserId      int64  `protobuf:"varint,12,opt,name=userId,proto3" json:"userId,omitempty"`          //创建用户 id
	CreateTime  int64  `protobuf:"varint,13,opt,name=createTime,proto3" json:"createTime,omitempty"`  //创建时间
	UpdateTime  int64  `protobuf:"varint,14,opt,name=updateTime,proto3" json:"updateTime,omitempty"`  //更新时间
	IsDelete    int64  `protobuf:"varint,15,opt,name=isDelete,proto3" json:"isDelete,omitempty"`      //是否删除
}

func (x *UpdateQuestionReq) Reset() {
	*x = UpdateQuestionReq{}
	mi := &file_question_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateQuestionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuestionReq) ProtoMessage() {}

func (x *UpdateQuestionReq) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuestionReq.ProtoReflect.Descriptor instead.
func (*UpdateQuestionReq) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateQuestionReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateQuestionReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateQuestionReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateQuestionReq) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *UpdateQuestionReq) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *UpdateQuestionReq) GetSubmitNum() int64 {
	if x != nil {
		return x.SubmitNum
	}
	return 0
}

func (x *UpdateQuestionReq) GetAcceptedNum() int64 {
	if x != nil {
		return x.AcceptedNum
	}
	return 0
}

func (x *UpdateQuestionReq) GetJudgeCase() string {
	if x != nil {
		return x.JudgeCase
	}
	return ""
}

func (x *UpdateQuestionReq) GetJudgeConfig() string {
	if x != nil {
		return x.JudgeConfig
	}
	return ""
}

func (x *UpdateQuestionReq) GetThumbNum() int64 {
	if x != nil {
		return x.ThumbNum
	}
	return 0
}

func (x *UpdateQuestionReq) GetFavourNum() int64 {
	if x != nil {
		return x.FavourNum
	}
	return 0
}

func (x *UpdateQuestionReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateQuestionReq) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *UpdateQuestionReq) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *UpdateQuestionReq) GetIsDelete() int64 {
	if x != nil {
		return x.IsDelete
	}
	return 0
}

type UpdateQuestionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateQuestionResp) Reset() {
	*x = UpdateQuestionResp{}
	mi := &file_question_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateQuestionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuestionResp) ProtoMessage() {}

func (x *UpdateQuestionResp) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuestionResp.ProtoReflect.Descriptor instead.
func (*UpdateQuestionResp) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{4}
}

type DelQuestionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //id
}

func (x *DelQuestionReq) Reset() {
	*x = DelQuestionReq{}
	mi := &file_question_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelQuestionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelQuestionReq) ProtoMessage() {}

func (x *DelQuestionReq) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelQuestionReq.ProtoReflect.Descriptor instead.
func (*DelQuestionReq) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{5}
}

func (x *DelQuestionReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DelQuestionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DelQuestionResp) Reset() {
	*x = DelQuestionResp{}
	mi := &file_question_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelQuestionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelQuestionResp) ProtoMessage() {}

func (x *DelQuestionResp) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelQuestionResp.ProtoReflect.Descriptor instead.
func (*DelQuestionResp) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{6}
}

type GetQuestionByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //id
}

func (x *GetQuestionByIdReq) Reset() {
	*x = GetQuestionByIdReq{}
	mi := &file_question_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuestionByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionByIdReq) ProtoMessage() {}

func (x *GetQuestionByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionByIdReq.ProtoReflect.Descriptor instead.
func (*GetQuestionByIdReq) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{7}
}

func (x *GetQuestionByIdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetQuestionByIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question *Question `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"` //question
}

func (x *GetQuestionByIdResp) Reset() {
	*x = GetQuestionByIdResp{}
	mi := &file_question_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuestionByIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionByIdResp) ProtoMessage() {}

func (x *GetQuestionByIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionByIdResp.ProtoReflect.Descriptor instead.
func (*GetQuestionByIdResp) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{8}
}

func (x *GetQuestionByIdResp) GetQuestion() *Question {
	if x != nil {
		return x.Question
	}
	return nil
}

type SearchQuestionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`   //page
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"` //limit
}

func (x *SearchQuestionReq) Reset() {
	*x = SearchQuestionReq{}
	mi := &file_question_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchQuestionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchQuestionReq) ProtoMessage() {}

func (x *SearchQuestionReq) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchQuestionReq.ProtoReflect.Descriptor instead.
func (*SearchQuestionReq) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{9}
}

func (x *SearchQuestionReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchQuestionReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type SearchQuestionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int64       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`      // 总记录数
	Question []*Question `protobuf:"bytes,2,rep,name=question,proto3" json:"question,omitempty"` //question
}

func (x *SearchQuestionResp) Reset() {
	*x = SearchQuestionResp{}
	mi := &file_question_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchQuestionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchQuestionResp) ProtoMessage() {}

func (x *SearchQuestionResp) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchQuestionResp.ProtoReflect.Descriptor instead.
func (*SearchQuestionResp) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{10}
}

func (x *SearchQuestionResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SearchQuestionResp) GetQuestion() []*Question {
	if x != nil {
		return x.Question
	}
	return nil
}

var File_question_proto protoreflect.FileDescriptor

var file_question_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa4, 0x03, 0x0a, 0x08, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x75, 0x6d,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x75,
	0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4e, 0x75, 0x6d,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43, 0x61, 0x73, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43, 0x61, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x4e, 0x75, 0x6d, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x4e, 0x75, 0x6d, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x22, 0x9a, 0x03, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4e, 0x75, 0x6d,
	0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43, 0x61, 0x73, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43, 0x61, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x4e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x11,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x22, 0xad, 0x03, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x75, 0x6d,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x75,
	0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4e, 0x75, 0x6d,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43, 0x61, 0x73, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43, 0x61, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x4e, 0x75, 0x6d, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x4e, 0x75, 0x6d, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x24, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x45, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x08, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x11, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5a, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x32, 0xfc, 0x02, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x42, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4b, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x42, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4b, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_question_proto_rawDescOnce sync.Once
	file_question_proto_rawDescData = file_question_proto_rawDesc
)

func file_question_proto_rawDescGZIP() []byte {
	file_question_proto_rawDescOnce.Do(func() {
		file_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_question_proto_rawDescData)
	})
	return file_question_proto_rawDescData
}

var file_question_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_question_proto_goTypes = []any{
	(*Question)(nil),            // 0: question.Question
	(*AddQuestionReq)(nil),      // 1: question.AddQuestionReq
	(*AddQuestionResp)(nil),     // 2: question.AddQuestionResp
	(*UpdateQuestionReq)(nil),   // 3: question.UpdateQuestionReq
	(*UpdateQuestionResp)(nil),  // 4: question.UpdateQuestionResp
	(*DelQuestionReq)(nil),      // 5: question.DelQuestionReq
	(*DelQuestionResp)(nil),     // 6: question.DelQuestionResp
	(*GetQuestionByIdReq)(nil),  // 7: question.GetQuestionByIdReq
	(*GetQuestionByIdResp)(nil), // 8: question.GetQuestionByIdResp
	(*SearchQuestionReq)(nil),   // 9: question.SearchQuestionReq
	(*SearchQuestionResp)(nil),  // 10: question.SearchQuestionResp
}
var file_question_proto_depIdxs = []int32{
	0,  // 0: question.GetQuestionByIdResp.question:type_name -> question.Question
	0,  // 1: question.SearchQuestionResp.question:type_name -> question.Question
	1,  // 2: question.question.AddQuestion:input_type -> question.AddQuestionReq
	3,  // 3: question.question.UpdateQuestion:input_type -> question.UpdateQuestionReq
	5,  // 4: question.question.DelQuestion:input_type -> question.DelQuestionReq
	7,  // 5: question.question.GetQuestionById:input_type -> question.GetQuestionByIdReq
	9,  // 6: question.question.SearchQuestion:input_type -> question.SearchQuestionReq
	2,  // 7: question.question.AddQuestion:output_type -> question.AddQuestionResp
	4,  // 8: question.question.UpdateQuestion:output_type -> question.UpdateQuestionResp
	6,  // 9: question.question.DelQuestion:output_type -> question.DelQuestionResp
	8,  // 10: question.question.GetQuestionById:output_type -> question.GetQuestionByIdResp
	10, // 11: question.question.SearchQuestion:output_type -> question.SearchQuestionResp
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_question_proto_init() }
func file_question_proto_init() {
	if File_question_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_question_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_question_proto_goTypes,
		DependencyIndexes: file_question_proto_depIdxs,
		MessageInfos:      file_question_proto_msgTypes,
	}.Build()
	File_question_proto = out.File
	file_question_proto_rawDesc = nil
	file_question_proto_goTypes = nil
	file_question_proto_depIdxs = nil
}
