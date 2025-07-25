// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: proto/kushal.proto

package proto

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

type NoParam struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NoParam) Reset() {
	*x = NoParam{}
	mi := &file_proto_kushal_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoParam) ProtoMessage() {}

func (x *NoParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kushal_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoParam.ProtoReflect.Descriptor instead.
func (*NoParam) Descriptor() ([]byte, []int) {
	return file_proto_kushal_proto_rawDescGZIP(), []int{0}
}

type NameList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Names         []string               `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NameList) Reset() {
	*x = NameList{}
	mi := &file_proto_kushal_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NameList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameList) ProtoMessage() {}

func (x *NameList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kushal_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameList.ProtoReflect.Descriptor instead.
func (*NameList) Descriptor() ([]byte, []int) {
	return file_proto_kushal_proto_rawDescGZIP(), []int{1}
}

func (x *NameList) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type HelloRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Names         string                 `protobuf:"bytes,1,opt,name=names,proto3" json:"names,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	mi := &file_proto_kushal_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kushal_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_kushal_proto_rawDescGZIP(), []int{2}
}

func (x *HelloRequest) GetNames() string {
	if x != nil {
		return x.Names
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	mi := &file_proto_kushal_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kushal_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_proto_kushal_proto_rawDescGZIP(), []int{3}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MessageList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Messages      []string               `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageList) Reset() {
	*x = MessageList{}
	mi := &file_proto_kushal_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageList) ProtoMessage() {}

func (x *MessageList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kushal_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageList.ProtoReflect.Descriptor instead.
func (*MessageList) Descriptor() ([]byte, []int) {
	return file_proto_kushal_proto_rawDescGZIP(), []int{4}
}

func (x *MessageList) GetMessages() []string {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_proto_kushal_proto protoreflect.FileDescriptor

const file_proto_kushal_proto_rawDesc = "" +
	"\n" +
	"\x12proto/kushal.proto\x12\x0ekushal_service\"\t\n" +
	"\aNoParam\" \n" +
	"\bNameList\x12\x14\n" +
	"\x05names\x18\x01 \x03(\tR\x05names\"$\n" +
	"\fHelloRequest\x12\x14\n" +
	"\x05names\x18\x01 \x01(\tR\x05names\")\n" +
	"\rHelloResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\")\n" +
	"\vMessageList\x12\x1a\n" +
	"\bmessages\x18\x01 \x03(\tR\bmessages2\xe4\x02\n" +
	"\rKushalService\x12B\n" +
	"\bSayHello\x12\x17.kushal_service.NoParam\x1a\x1d.kushal_service.HelloResponse\x12T\n" +
	"\x17SayHelloServerStreaming\x12\x18.kushal_service.NameList\x1a\x1d.kushal_service.HelloResponse0\x01\x12V\n" +
	"\x17SayHelloClientStreaming\x12\x1c.kushal_service.HelloRequest\x1a\x1b.kushal_service.MessageList(\x01\x12a\n" +
	"\x1eSayHelloBidirectionalStreaming\x12\x1c.kushal_service.HelloRequest\x1a\x1d.kushal_service.HelloResponse(\x010\x01B+Z)github.com/kushal88053/go_with_grpc/protob\x06proto3"

var (
	file_proto_kushal_proto_rawDescOnce sync.Once
	file_proto_kushal_proto_rawDescData []byte
)

func file_proto_kushal_proto_rawDescGZIP() []byte {
	file_proto_kushal_proto_rawDescOnce.Do(func() {
		file_proto_kushal_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_kushal_proto_rawDesc), len(file_proto_kushal_proto_rawDesc)))
	})
	return file_proto_kushal_proto_rawDescData
}

var file_proto_kushal_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_kushal_proto_goTypes = []any{
	(*NoParam)(nil),       // 0: kushal_service.NoParam
	(*NameList)(nil),      // 1: kushal_service.NameList
	(*HelloRequest)(nil),  // 2: kushal_service.HelloRequest
	(*HelloResponse)(nil), // 3: kushal_service.HelloResponse
	(*MessageList)(nil),   // 4: kushal_service.MessageList
}
var file_proto_kushal_proto_depIdxs = []int32{
	0, // 0: kushal_service.KushalService.SayHello:input_type -> kushal_service.NoParam
	1, // 1: kushal_service.KushalService.SayHelloServerStreaming:input_type -> kushal_service.NameList
	2, // 2: kushal_service.KushalService.SayHelloClientStreaming:input_type -> kushal_service.HelloRequest
	2, // 3: kushal_service.KushalService.SayHelloBidirectionalStreaming:input_type -> kushal_service.HelloRequest
	3, // 4: kushal_service.KushalService.SayHello:output_type -> kushal_service.HelloResponse
	3, // 5: kushal_service.KushalService.SayHelloServerStreaming:output_type -> kushal_service.HelloResponse
	4, // 6: kushal_service.KushalService.SayHelloClientStreaming:output_type -> kushal_service.MessageList
	3, // 7: kushal_service.KushalService.SayHelloBidirectionalStreaming:output_type -> kushal_service.HelloResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_kushal_proto_init() }
func file_proto_kushal_proto_init() {
	if File_proto_kushal_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_kushal_proto_rawDesc), len(file_proto_kushal_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_kushal_proto_goTypes,
		DependencyIndexes: file_proto_kushal_proto_depIdxs,
		MessageInfos:      file_proto_kushal_proto_msgTypes,
	}.Build()
	File_proto_kushal_proto = out.File
	file_proto_kushal_proto_goTypes = nil
	file_proto_kushal_proto_depIdxs = nil
}
