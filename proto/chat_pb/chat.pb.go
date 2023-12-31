// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: chat.proto

package chat_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of google/protobuf/timestamp.proto.

type Timestamp = timestamppb.Timestamp

type MessageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId int64                  `protobuf:"varint,1,opt,name=MessageId,proto3" json:"MessageId,omitempty"`
	Content   string                 `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	MsgType   int32                  `protobuf:"varint,3,opt,name=MsgType,proto3" json:"MsgType,omitempty"`
	SendTime  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=SendTime,proto3" json:"SendTime,omitempty"`
	ReadTime  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=ReadTime,proto3" json:"ReadTime,omitempty"`
}

func (x *MessageInfo) Reset() {
	*x = MessageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageInfo) ProtoMessage() {}

func (x *MessageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageInfo.ProtoReflect.Descriptor instead.
func (*MessageInfo) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *MessageInfo) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *MessageInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MessageInfo) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *MessageInfo) GetSendTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SendTime
	}
	return nil
}

func (x *MessageInfo) GetReadTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReadTime
	}
	return nil
}

type SendClientReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemId string       `protobuf:"bytes,1,opt,name=SystemId,proto3" json:"SystemId,omitempty"`
	UserId   int64        `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	TargetId int64        `protobuf:"varint,3,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
	Data     *MessageInfo `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *SendClientReq) Reset() {
	*x = SendClientReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendClientReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendClientReq) ProtoMessage() {}

func (x *SendClientReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendClientReq.ProtoReflect.Descriptor instead.
func (*SendClientReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *SendClientReq) GetSystemId() string {
	if x != nil {
		return x.SystemId
	}
	return ""
}

func (x *SendClientReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SendClientReq) GetTargetId() int64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *SendClientReq) GetData() *MessageInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type CloseClientReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemId string `protobuf:"bytes,1,opt,name=SystemId,proto3" json:"SystemId,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *CloseClientReq) Reset() {
	*x = CloseClientReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseClientReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseClientReq) ProtoMessage() {}

func (x *CloseClientReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseClientReq.ProtoReflect.Descriptor instead.
func (*CloseClientReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *CloseClientReq) GetSystemId() string {
	if x != nil {
		return x.SystemId
	}
	return ""
}

func (x *CloseClientReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type SendClientReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64        `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string       `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Data *MessageInfo `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *SendClientReply) Reset() {
	*x = SendClientReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendClientReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendClientReply) ProtoMessage() {}

func (x *SendClientReply) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendClientReply.ProtoReflect.Descriptor instead.
func (*SendClientReply) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *SendClientReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SendClientReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SendClientReply) GetData() *MessageInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type CloseClientReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *CloseClientReply) Reset() {
	*x = CloseClientReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseClientReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseClientReply) ProtoMessage() {}

func (x *CloseClientReply) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseClientReply.ProtoReflect.Descriptor instead.
func (*CloseClientReply) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *CloseClientReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CloseClientReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type CreateCommunityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunityId int64                  `protobuf:"varint,1,opt,name=CommunityId,proto3" json:"CommunityId,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	OwnerId     int64                  `protobuf:"varint,3,opt,name=OwnerId,proto3" json:"OwnerId,omitempty"`
	LimitNum    int64                  `protobuf:"varint,4,opt,name=LimitNum,proto3" json:"LimitNum,omitempty"`
	Desc        string                 `protobuf:"bytes,5,opt,name=Desc,proto3" json:"Desc,omitempty"`
	CreateTime  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
}

func (x *CreateCommunityReq) Reset() {
	*x = CreateCommunityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommunityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommunityReq) ProtoMessage() {}

func (x *CreateCommunityReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommunityReq.ProtoReflect.Descriptor instead.
func (*CreateCommunityReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCommunityReq) GetCommunityId() int64 {
	if x != nil {
		return x.CommunityId
	}
	return 0
}

func (x *CreateCommunityReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCommunityReq) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *CreateCommunityReq) GetLimitNum() int64 {
	if x != nil {
		return x.LimitNum
	}
	return 0
}

func (x *CreateCommunityReq) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CreateCommunityReq) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type CreateCommunityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunityId int64  `protobuf:"varint,1,opt,name=CommunityId,proto3" json:"CommunityId,omitempty"`
	Code        int64  `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg         string `protobuf:"bytes,3,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *CreateCommunityReply) Reset() {
	*x = CreateCommunityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommunityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommunityReply) ProtoMessage() {}

func (x *CreateCommunityReply) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommunityReply.ProtoReflect.Descriptor instead.
func (*CreateCommunityReply) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{6}
}

func (x *CreateCommunityReply) GetCommunityId() int64 {
	if x != nil {
		return x.CommunityId
	}
	return 0
}

func (x *CreateCommunityReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateCommunityReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type QuitCommunityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	CommunityId int64 `protobuf:"varint,2,opt,name=CommunityId,proto3" json:"CommunityId,omitempty"`
}

func (x *QuitCommunityReq) Reset() {
	*x = QuitCommunityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuitCommunityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuitCommunityReq) ProtoMessage() {}

func (x *QuitCommunityReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuitCommunityReq.ProtoReflect.Descriptor instead.
func (*QuitCommunityReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{7}
}

func (x *QuitCommunityReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *QuitCommunityReq) GetCommunityId() int64 {
	if x != nil {
		return x.CommunityId
	}
	return 0
}

type QuitCommunityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *QuitCommunityReply) Reset() {
	*x = QuitCommunityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuitCommunityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuitCommunityReply) ProtoMessage() {}

func (x *QuitCommunityReply) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuitCommunityReply.ProtoReflect.Descriptor instead.
func (*QuitCommunityReply) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{8}
}

func (x *QuitCommunityReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *QuitCommunityReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type JoinCommunityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64                  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	CommunityId int64                  `protobuf:"varint,2,opt,name=CommunityId,proto3" json:"CommunityId,omitempty"`
	JoinTime    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=JoinTime,proto3" json:"JoinTime,omitempty"`
}

func (x *JoinCommunityReq) Reset() {
	*x = JoinCommunityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinCommunityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinCommunityReq) ProtoMessage() {}

func (x *JoinCommunityReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinCommunityReq.ProtoReflect.Descriptor instead.
func (*JoinCommunityReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{9}
}

func (x *JoinCommunityReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *JoinCommunityReq) GetCommunityId() int64 {
	if x != nil {
		return x.CommunityId
	}
	return 0
}

func (x *JoinCommunityReq) GetJoinTime() *timestamppb.Timestamp {
	if x != nil {
		return x.JoinTime
	}
	return nil
}

type JoinCommunityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *JoinCommunityReply) Reset() {
	*x = JoinCommunityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinCommunityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinCommunityReply) ProtoMessage() {}

func (x *JoinCommunityReply) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinCommunityReply.ProtoReflect.Descriptor instead.
func (*JoinCommunityReply) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{10}
}

func (x *JoinCommunityReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *JoinCommunityReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01,
	0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a,
	0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x36, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x52, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x81, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x44, 0x0a, 0x0e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x0f, 0x53, 0x65, 0x6e,
	0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d,
	0x73, 0x67, 0x12, 0x20, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x38, 0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0xd0,
	0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4e, 0x75,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4e, 0x75,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x44, 0x65, 0x73, 0x63, 0x12, 0x3a, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x5e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73,
	0x67, 0x22, 0x4c, 0x0a, 0x10, 0x51, 0x75, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22,
	0x3a, 0x0a, 0x12, 0x51, 0x75, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x84, 0x01, 0x0a, 0x10,
	0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x43,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x4a, 0x6f,
	0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x12, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x32, 0xa2,
	0x02, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f,
	0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x32, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x31, 0x0a, 0x0b, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0f,
	0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x11, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x3d, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x37, 0x0a, 0x0d, 0x51, 0x75, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x12, 0x11, 0x2e, 0x51, 0x75, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x51, 0x75, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x37, 0x0a, 0x0d, 0x4a, 0x6f,
	0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x11, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x0a, 0x5a, 0x08, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x70, 0x62, 0x2f, 0x50,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_chat_proto_goTypes = []interface{}{
	(*MessageInfo)(nil),           // 0: MessageInfo
	(*SendClientReq)(nil),         // 1: SendClientReq
	(*CloseClientReq)(nil),        // 2: CloseClientReq
	(*SendClientReply)(nil),       // 3: SendClientReply
	(*CloseClientReply)(nil),      // 4: CloseClientReply
	(*CreateCommunityReq)(nil),    // 5: CreateCommunityReq
	(*CreateCommunityReply)(nil),  // 6: CreateCommunityReply
	(*QuitCommunityReq)(nil),      // 7: QuitCommunityReq
	(*QuitCommunityReply)(nil),    // 8: QuitCommunityReply
	(*JoinCommunityReq)(nil),      // 9: JoinCommunityReq
	(*JoinCommunityReply)(nil),    // 10: JoinCommunityReply
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_chat_proto_depIdxs = []int32{
	11, // 0: MessageInfo.SendTime:type_name -> google.protobuf.Timestamp
	11, // 1: MessageInfo.ReadTime:type_name -> google.protobuf.Timestamp
	0,  // 2: SendClientReq.Data:type_name -> MessageInfo
	0,  // 3: SendClientReply.Data:type_name -> MessageInfo
	11, // 4: CreateCommunityReq.CreateTime:type_name -> google.protobuf.Timestamp
	11, // 5: JoinCommunityReq.JoinTime:type_name -> google.protobuf.Timestamp
	1,  // 6: ChatService.Send2Client:input_type -> SendClientReq
	2,  // 7: ChatService.CloseClient:input_type -> CloseClientReq
	5,  // 8: ChatService.CreateCommunity:input_type -> CreateCommunityReq
	7,  // 9: ChatService.QuitCommunity:input_type -> QuitCommunityReq
	9,  // 10: ChatService.JoinCommunity:input_type -> JoinCommunityReq
	3,  // 11: ChatService.Send2Client:output_type -> SendClientReply
	4,  // 12: ChatService.CloseClient:output_type -> CloseClientReply
	6,  // 13: ChatService.CreateCommunity:output_type -> CreateCommunityReply
	8,  // 14: ChatService.QuitCommunity:output_type -> QuitCommunityReply
	10, // 15: ChatService.JoinCommunity:output_type -> JoinCommunityReply
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendClientReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseClientReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendClientReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseClientReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommunityReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommunityReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuitCommunityReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuitCommunityReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinCommunityReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinCommunityReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
