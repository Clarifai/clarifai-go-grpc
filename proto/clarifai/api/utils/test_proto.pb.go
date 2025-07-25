// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/clarifai/api/utils/test_proto.proto

package utils

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

type TestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message    string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Value      float64 `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	ImageBytes []byte  `protobuf:"bytes,4,opt,name=image_bytes,json=imageBytes,proto3" json:"image_bytes,omitempty"`
	// Types that are assignable to OneOfField:
	//	*TestProto_StringOneof
	//	*TestProto_BoolOneof
	//	*TestProto_MessageOneof
	OneOfField isTestProto_OneOfField `protobuf_oneof:"one_of_field"`
}

func (x *TestProto) Reset() {
	*x = TestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clarifai_api_utils_test_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestProto) ProtoMessage() {}

func (x *TestProto) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clarifai_api_utils_test_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestProto.ProtoReflect.Descriptor instead.
func (*TestProto) Descriptor() ([]byte, []int) {
	return file_proto_clarifai_api_utils_test_proto_proto_rawDescGZIP(), []int{0}
}

func (x *TestProto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TestProto) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TestProto) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TestProto) GetImageBytes() []byte {
	if x != nil {
		return x.ImageBytes
	}
	return nil
}

func (m *TestProto) GetOneOfField() isTestProto_OneOfField {
	if m != nil {
		return m.OneOfField
	}
	return nil
}

func (x *TestProto) GetStringOneof() string {
	if x, ok := x.GetOneOfField().(*TestProto_StringOneof); ok {
		return x.StringOneof
	}
	return ""
}

func (x *TestProto) GetBoolOneof() bool {
	if x, ok := x.GetOneOfField().(*TestProto_BoolOneof); ok {
		return x.BoolOneof
	}
	return false
}

func (x *TestProto) GetMessageOneof() *TestProto2 {
	if x, ok := x.GetOneOfField().(*TestProto_MessageOneof); ok {
		return x.MessageOneof
	}
	return nil
}

type isTestProto_OneOfField interface {
	isTestProto_OneOfField()
}

type TestProto_StringOneof struct {
	StringOneof string `protobuf:"bytes,5,opt,name=string_oneof,json=stringOneof,proto3,oneof"`
}

type TestProto_BoolOneof struct {
	BoolOneof bool `protobuf:"varint,6,opt,name=bool_oneof,json=boolOneof,proto3,oneof"`
}

type TestProto_MessageOneof struct {
	MessageOneof *TestProto2 `protobuf:"bytes,7,opt,name=message_oneof,json=messageOneof,proto3,oneof"`
}

func (*TestProto_StringOneof) isTestProto_OneOfField() {}

func (*TestProto_BoolOneof) isTestProto_OneOfField() {}

func (*TestProto_MessageOneof) isTestProto_OneOfField() {}

type TestProto2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Flip bool   `protobuf:"varint,2,opt,name=flip,proto3" json:"flip,omitempty"`
}

func (x *TestProto2) Reset() {
	*x = TestProto2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clarifai_api_utils_test_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestProto2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestProto2) ProtoMessage() {}

func (x *TestProto2) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clarifai_api_utils_test_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestProto2.ProtoReflect.Descriptor instead.
func (*TestProto2) Descriptor() ([]byte, []int) {
	return file_proto_clarifai_api_utils_test_proto_proto_rawDescGZIP(), []int{1}
}

func (x *TestProto2) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TestProto2) GetFlip() bool {
	if x != nil {
		return x.Flip
	}
	return false
}

var File_proto_clarifai_api_utils_test_proto_proto protoreflect.FileDescriptor

var file_proto_clarifai_api_utils_test_proto_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6c, 0x61,
	0x72, 0x69, 0x66, 0x61, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x1a,
	0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x09, 0x54,
	0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x80, 0xb5, 0x18, 0x01, 0x8a,
	0xb5, 0x18, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x63, 0x6f,
	0x6f, 0x6c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x42, 0x07, 0xd5, 0xb5, 0x18, 0x00,
	0x00, 0x80, 0x3f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0c, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4f, 0x6e, 0x65, 0x6f, 0x66,
	0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x4f, 0x6e, 0x65, 0x6f,
	0x66, 0x12, 0x45, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6c, 0x61, 0x72, 0x69,
	0x66, 0x61, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x48, 0x00, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x0e, 0x0a, 0x0c, 0x6f, 0x6e, 0x65, 0x5f,
	0x6f, 0x66, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x30, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x69, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x6c, 0x69, 0x70, 0x42, 0x65, 0x0a, 0x1b, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69,
	0x2f, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0xa2, 0x02, 0x04, 0x43, 0x41, 0x49,
	0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_clarifai_api_utils_test_proto_proto_rawDescOnce sync.Once
	file_proto_clarifai_api_utils_test_proto_proto_rawDescData = file_proto_clarifai_api_utils_test_proto_proto_rawDesc
)

func file_proto_clarifai_api_utils_test_proto_proto_rawDescGZIP() []byte {
	file_proto_clarifai_api_utils_test_proto_proto_rawDescOnce.Do(func() {
		file_proto_clarifai_api_utils_test_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_clarifai_api_utils_test_proto_proto_rawDescData)
	})
	return file_proto_clarifai_api_utils_test_proto_proto_rawDescData
}

var file_proto_clarifai_api_utils_test_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_clarifai_api_utils_test_proto_proto_goTypes = []interface{}{
	(*TestProto)(nil),  // 0: clarifai.api.utils.TestProto
	(*TestProto2)(nil), // 1: clarifai.api.utils.TestProto2
}
var file_proto_clarifai_api_utils_test_proto_proto_depIdxs = []int32{
	1, // 0: clarifai.api.utils.TestProto.message_oneof:type_name -> clarifai.api.utils.TestProto2
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_clarifai_api_utils_test_proto_proto_init() }
func file_proto_clarifai_api_utils_test_proto_proto_init() {
	if File_proto_clarifai_api_utils_test_proto_proto != nil {
		return
	}
	file_proto_clarifai_api_utils_extensions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_clarifai_api_utils_test_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestProto); i {
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
		file_proto_clarifai_api_utils_test_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestProto2); i {
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
	file_proto_clarifai_api_utils_test_proto_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TestProto_StringOneof)(nil),
		(*TestProto_BoolOneof)(nil),
		(*TestProto_MessageOneof)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_clarifai_api_utils_test_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_clarifai_api_utils_test_proto_proto_goTypes,
		DependencyIndexes: file_proto_clarifai_api_utils_test_proto_proto_depIdxs,
		MessageInfos:      file_proto_clarifai_api_utils_test_proto_proto_msgTypes,
	}.Build()
	File_proto_clarifai_api_utils_test_proto_proto = out.File
	file_proto_clarifai_api_utils_test_proto_proto_rawDesc = nil
	file_proto_clarifai_api_utils_test_proto_proto_goTypes = nil
	file_proto_clarifai_api_utils_test_proto_proto_depIdxs = nil
}
