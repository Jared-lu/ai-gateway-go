// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: ai/v1/ai.proto

package aiv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StreamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	mi := &file_ai_v1_ai_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ai_v1_ai_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_ai_v1_ai_proto_rawDescGZIP(), []int{0}
}

func (x *StreamRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StreamRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type StreamResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Final            bool                   `protobuf:"varint,1,opt,name=final,proto3" json:"final,omitempty"`
	ReasoningContent string                 `protobuf:"bytes,2,opt,name=reasoningContent,proto3" json:"reasoningContent,omitempty"`
	Content          string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Err              string                 `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	mi := &file_ai_v1_ai_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ai_v1_ai_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_ai_v1_ai_proto_rawDescGZIP(), []int{1}
}

func (x *StreamResponse) GetFinal() bool {
	if x != nil {
		return x.Final
	}
	return false
}

func (x *StreamResponse) GetReasoningContent() string {
	if x != nil {
		return x.ReasoningContent
	}
	return ""
}

func (x *StreamResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *StreamResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_ai_v1_ai_proto protoreflect.FileDescriptor

const file_ai_v1_ai_proto_rawDesc = "" +
	"\n" +
	"\x0eai/v1/ai.proto\x12\x05ai.v1\"3\n" +
	"\rStreamRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04text\x18\x02 \x01(\tR\x04text\"~\n" +
	"\x0eStreamResponse\x12\x14\n" +
	"\x05final\x18\x01 \x01(\bR\x05final\x12*\n" +
	"\x10reasoningContent\x18\x02 \x01(\tR\x10reasoningContent\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\x12\x10\n" +
	"\x03err\x18\x04 \x01(\tR\x03err2D\n" +
	"\tAIService\x127\n" +
	"\x06Stream\x12\x14.ai.v1.StreamRequest\x1a\x15.ai.v1.StreamResponse0\x01B\x80\x01\n" +
	"\tcom.ai.v1B\aAiProtoP\x01Z5github.com/ecodeclub/ai-gateway-go/api/gen/ai/v1;aiv1\xa2\x02\x03AXX\xaa\x02\x05Ai.V1\xca\x02\x05Ai\\V1\xe2\x02\x11Ai\\V1\\GPBMetadata\xea\x02\x06Ai::V1b\x06proto3"

var (
	file_ai_v1_ai_proto_rawDescOnce sync.Once
	file_ai_v1_ai_proto_rawDescData []byte
)

func file_ai_v1_ai_proto_rawDescGZIP() []byte {
	file_ai_v1_ai_proto_rawDescOnce.Do(func() {
		file_ai_v1_ai_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ai_v1_ai_proto_rawDesc), len(file_ai_v1_ai_proto_rawDesc)))
	})
	return file_ai_v1_ai_proto_rawDescData
}

var file_ai_v1_ai_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ai_v1_ai_proto_goTypes = []any{
	(*StreamRequest)(nil),  // 0: ai.v1.StreamRequest
	(*StreamResponse)(nil), // 1: ai.v1.StreamResponse
}
var file_ai_v1_ai_proto_depIdxs = []int32{
	0, // 0: ai.v1.AIService.Stream:input_type -> ai.v1.StreamRequest
	1, // 1: ai.v1.AIService.Stream:output_type -> ai.v1.StreamResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ai_v1_ai_proto_init() }
func file_ai_v1_ai_proto_init() {
	if File_ai_v1_ai_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ai_v1_ai_proto_rawDesc), len(file_ai_v1_ai_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ai_v1_ai_proto_goTypes,
		DependencyIndexes: file_ai_v1_ai_proto_depIdxs,
		MessageInfos:      file_ai_v1_ai_proto_msgTypes,
	}.Build()
	File_ai_v1_ai_proto = out.File
	file_ai_v1_ai_proto_goTypes = nil
	file_ai_v1_ai_proto_depIdxs = nil
}
