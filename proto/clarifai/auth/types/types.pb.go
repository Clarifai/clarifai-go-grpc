// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: proto/clarifai/auth/types/types.proto

package types

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Authorization type for endpoints.
type AuthType int32

const (
	// introduce undef so that the zero (default/unset) value of the enum is not a real
	// permission.  undef is only present for this purpose and should not be used
	// to indicate any "real" value.
	AuthType_undef AuthType = 0
	// No authorization need for this endpoint.
	AuthType_NoAuth AuthType = 1
	// This authorization requires API keys (both app-spceific keys and personal access tokens).
	// The endpoints that use this AuthType may also include a list of
	// clarifai.auth.utils.cl_depending_scopes.
	AuthType_KeyAuth AuthType = 2
	// This uses a session token from your web browser. This is reserved for users/account level APIs
	// that are only needed in a browser.
	AuthType_SessionTokenAuth AuthType = 3
	// This uses a special token for admin access to the APIs.
	AuthType_AdminAuth AuthType = 4
	// This authorization requires personal access tokens. This is used for endpoints such as
	// /users/{user_id}/apps which are not specific. An app-specific API key will not work
	// when PATAuth is used.
	AuthType_PATAuth AuthType = 5
)

// Enum value maps for AuthType.
var (
	AuthType_name = map[int32]string{
		0: "undef",
		1: "NoAuth",
		2: "KeyAuth",
		3: "SessionTokenAuth",
		4: "AdminAuth",
		5: "PATAuth",
	}
	AuthType_value = map[string]int32{
		"undef":            0,
		"NoAuth":           1,
		"KeyAuth":          2,
		"SessionTokenAuth": 3,
		"AdminAuth":        4,
		"PATAuth":          5,
	}
)

func (x AuthType) Enum() *AuthType {
	p := new(AuthType)
	*p = x
	return p
}

func (x AuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_clarifai_auth_types_types_proto_enumTypes[0].Descriptor()
}

func (AuthType) Type() protoreflect.EnumType {
	return &file_proto_clarifai_auth_types_types_proto_enumTypes[0]
}

func (x AuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthType.Descriptor instead.
func (AuthType) EnumDescriptor() ([]byte, []int) {
	return file_proto_clarifai_auth_types_types_proto_rawDescGZIP(), []int{0}
}

var File_proto_clarifai_auth_types_types_proto protoreflect.FileDescriptor

var file_proto_clarifai_auth_types_types_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61,
	0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2a, 0x60, 0x0a, 0x08,
	0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x75, 0x6e, 0x64, 0x65,
	0x66, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68,
	0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x10,
	0x04, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x41, 0x54, 0x41, 0x75, 0x74, 0x68, 0x10, 0x05, 0x42, 0x67,
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x50, 0x01,
	0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x61,
	0x72, 0x69, 0x66, 0x61, 0x69, 0x2f, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x2d, 0x67,
	0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61,
	0x72, 0x69, 0x66, 0x61, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0xa2, 0x02, 0x04, 0x43, 0x41, 0x49, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_clarifai_auth_types_types_proto_rawDescOnce sync.Once
	file_proto_clarifai_auth_types_types_proto_rawDescData = file_proto_clarifai_auth_types_types_proto_rawDesc
)

func file_proto_clarifai_auth_types_types_proto_rawDescGZIP() []byte {
	file_proto_clarifai_auth_types_types_proto_rawDescOnce.Do(func() {
		file_proto_clarifai_auth_types_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_clarifai_auth_types_types_proto_rawDescData)
	})
	return file_proto_clarifai_auth_types_types_proto_rawDescData
}

var file_proto_clarifai_auth_types_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_clarifai_auth_types_types_proto_goTypes = []interface{}{
	(AuthType)(0), // 0: clarifai.auth.types.AuthType
}
var file_proto_clarifai_auth_types_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_clarifai_auth_types_types_proto_init() }
func file_proto_clarifai_auth_types_types_proto_init() {
	if File_proto_clarifai_auth_types_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_clarifai_auth_types_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_clarifai_auth_types_types_proto_goTypes,
		DependencyIndexes: file_proto_clarifai_auth_types_types_proto_depIdxs,
		EnumInfos:         file_proto_clarifai_auth_types_types_proto_enumTypes,
	}.Build()
	File_proto_clarifai_auth_types_types_proto = out.File
	file_proto_clarifai_auth_types_types_proto_rawDesc = nil
	file_proto_clarifai_auth_types_types_proto_goTypes = nil
	file_proto_clarifai_auth_types_types_proto_depIdxs = nil
}
