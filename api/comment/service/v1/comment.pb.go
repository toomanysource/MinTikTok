// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: api/comment/service/v1/comment.proto

package v1

import (
	v1 "Atreus/api/entity/v1"
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

type CommentListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                     // 用户鉴权token
	VideoId int64  `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"` // 视频id
}

func (x *CommentListRequest) Reset() {
	*x = CommentListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comment_service_v1_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListRequest) ProtoMessage() {}

func (x *CommentListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_comment_service_v1_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListRequest.ProtoReflect.Descriptor instead.
func (*CommentListRequest) Descriptor() ([]byte, []int) {
	return file_api_comment_service_v1_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CommentListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CommentListRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type CommentListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode  int32       `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg   string      `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`       // 返回状态描述
	CommentList *v1.Comment `protobuf:"bytes,3,opt,name=comment_list,json=commentList,proto3" json:"comment_list,omitempty"` // 评论列表
}

func (x *CommentListReply) Reset() {
	*x = CommentListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comment_service_v1_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListReply) ProtoMessage() {}

func (x *CommentListReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_comment_service_v1_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListReply.ProtoReflect.Descriptor instead.
func (*CommentListReply) Descriptor() ([]byte, []int) {
	return file_api_comment_service_v1_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentListReply) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CommentListReply) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CommentListReply) GetCommentList() *v1.Comment {
	if x != nil {
		return x.CommentList
	}
	return nil
}

type CommentActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                                // 用户鉴权token
	VideoId     int64  `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`            // 视频id
	ActionType  int32  `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`   // 1-发布评论，2-删除评论
	CommentText string `protobuf:"bytes,4,opt,name=comment_text,json=commentText,proto3" json:"comment_text,omitempty"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   int64  `protobuf:"varint,5,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`      // 要删除的评论id，在action_type=2的时候使用
}

func (x *CommentActionRequest) Reset() {
	*x = CommentActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comment_service_v1_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentActionRequest) ProtoMessage() {}

func (x *CommentActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_comment_service_v1_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentActionRequest.ProtoReflect.Descriptor instead.
func (*CommentActionRequest) Descriptor() ([]byte, []int) {
	return file_api_comment_service_v1_comment_proto_rawDescGZIP(), []int{2}
}

func (x *CommentActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CommentActionRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *CommentActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

func (x *CommentActionRequest) GetCommentText() string {
	if x != nil {
		return x.CommentText
	}
	return ""
}

func (x *CommentActionRequest) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

type CommentActionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32       `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string      `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	Comment    *v1.Comment `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`                          // 评论成功返回评论内容，不需要重新拉取整个列表
}

func (x *CommentActionReply) Reset() {
	*x = CommentActionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comment_service_v1_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentActionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentActionReply) ProtoMessage() {}

func (x *CommentActionReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_comment_service_v1_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentActionReply.ProtoReflect.Descriptor instead.
func (*CommentActionReply) Descriptor() ([]byte, []int) {
	return file_api_comment_service_v1_comment_proto_rawDescGZIP(), []int{3}
}

func (x *CommentActionReply) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CommentActionReply) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CommentActionReply) GetComment() *v1.Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

var File_api_comment_service_v1_comment_proto protoreflect.FileDescriptor

var file_api_comment_service_v1_comment_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x89, 0x01,
	0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x73, 0x67, 0x12, 0x35, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xaa, 0x01, 0x0a, 0x14, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x2c, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0xd0, 0x01, 0x0a, 0x07,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x60, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x0d, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x22,
	0x5a, 0x20, 0x41, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_comment_service_v1_comment_proto_rawDescOnce sync.Once
	file_api_comment_service_v1_comment_proto_rawDescData = file_api_comment_service_v1_comment_proto_rawDesc
)

func file_api_comment_service_v1_comment_proto_rawDescGZIP() []byte {
	file_api_comment_service_v1_comment_proto_rawDescOnce.Do(func() {
		file_api_comment_service_v1_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_comment_service_v1_comment_proto_rawDescData)
	})
	return file_api_comment_service_v1_comment_proto_rawDescData
}

var file_api_comment_service_v1_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_comment_service_v1_comment_proto_goTypes = []interface{}{
	(*CommentListRequest)(nil),   // 0: comment.service.v1.CommentListRequest
	(*CommentListReply)(nil),     // 1: comment.service.v1.CommentListReply
	(*CommentActionRequest)(nil), // 2: comment.service.v1.CommentActionRequest
	(*CommentActionReply)(nil),   // 3: comment.service.v1.CommentActionReply
	(*v1.Comment)(nil),           // 4: entity.v1.Comment
}
var file_api_comment_service_v1_comment_proto_depIdxs = []int32{
	4, // 0: comment.service.v1.CommentListReply.comment_list:type_name -> entity.v1.Comment
	4, // 1: comment.service.v1.CommentActionReply.comment:type_name -> entity.v1.Comment
	0, // 2: comment.service.v1.Comment.GetCommentList:input_type -> comment.service.v1.CommentListRequest
	2, // 3: comment.service.v1.Comment.CommentAction:input_type -> comment.service.v1.CommentActionRequest
	1, // 4: comment.service.v1.Comment.GetCommentList:output_type -> comment.service.v1.CommentListReply
	3, // 5: comment.service.v1.Comment.CommentAction:output_type -> comment.service.v1.CommentActionReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_comment_service_v1_comment_proto_init() }
func file_api_comment_service_v1_comment_proto_init() {
	if File_api_comment_service_v1_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_comment_service_v1_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListRequest); i {
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
		file_api_comment_service_v1_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListReply); i {
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
		file_api_comment_service_v1_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentActionRequest); i {
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
		file_api_comment_service_v1_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentActionReply); i {
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
			RawDescriptor: file_api_comment_service_v1_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_comment_service_v1_comment_proto_goTypes,
		DependencyIndexes: file_api_comment_service_v1_comment_proto_depIdxs,
		MessageInfos:      file_api_comment_service_v1_comment_proto_msgTypes,
	}.Build()
	File_api_comment_service_v1_comment_proto = out.File
	file_api_comment_service_v1_comment_proto_rawDesc = nil
	file_api_comment_service_v1_comment_proto_goTypes = nil
	file_api_comment_service_v1_comment_proto_depIdxs = nil
}
