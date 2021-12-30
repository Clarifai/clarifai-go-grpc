// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: proto/clarifai/auth/util/extension.proto

package util

import (
	scope "github.com/Clarifai/clarifai-go-grpc/proto/clarifai/auth/scope"
	types "github.com/Clarifai/clarifai-go-grpc/proto/clarifai/auth/types"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var file_proto_clarifai_auth_util_extension_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         80320,
		Name:          "clarifai.auth.util.cl_private_field",
		Tag:           "varint,80320,opt,name=cl_private_field",
		Filename:      "proto/clarifai/auth/util/extension.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         80321,
		Name:          "clarifai.auth.util.cl_private_rpc",
		Tag:           "varint,80321,opt,name=cl_private_rpc",
		Filename:      "proto/clarifai/auth/util/extension.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: ([]scope.S)(nil),
		Field:         80322,
		Name:          "clarifai.auth.util.cl_depending_scopes",
		Tag:           "varint,80322,rep,name=cl_depending_scopes,enum=clarifai.auth.scope.S",
		Filename:      "proto/clarifai/auth/util/extension.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*types.AuthType)(nil),
		Field:         80323,
		Name:          "clarifai.auth.util.cl_auth_type",
		Tag:           "varint,80323,opt,name=cl_auth_type,enum=clarifai.auth.types.AuthType",
		Filename:      "proto/clarifai/auth/util/extension.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         80324,
		Name:          "clarifai.auth.util.cl_private_message",
		Tag:           "varint,80324,opt,name=cl_private_message",
		Filename:      "proto/clarifai/auth/util/extension.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// indicates if this field should be private. If true, then internal use only. If False,
	// then publicly avaliable to world.
	//
	// optional bool cl_private_field = 80320;
	E_ClPrivateField = &file_proto_clarifai_auth_util_extension_proto_extTypes[0]
)

// Extension fields to descriptor.MethodOptions.
var (
	// indicates if this rpc should be private. If true, then internal use only. If False,
	// then publicly available to world.
	//
	// optional bool cl_private_rpc = 80321;
	E_ClPrivateRpc = &file_proto_clarifai_auth_util_extension_proto_extTypes[1]
	// For each grpc method we define a list of required low-level auth scopes that are needed
	// for the key that is issuing the request. These are checked at authorization time for the
	// request as one of the first things in the life of a request once it reaches out API
	// servers. There are additional checks for each scope when the resource is used throughout the
	// backend in order to protect that resource. For example, Inputs_Add protects the writes to the
	// inputs DB table. Since we know every PostInputs call needs to do that, we add Inputs_Add scope
	// to the cl_depending_scopes list for the PostInputs method.
	//
	// This is only checked valid when used with KeyAuth cl_auth_type.
	//
	// This should be the absolute minimum required scopes to make API calls with the method
	// that this options is used for. If there are some scopes that are needed some of the time
	// depending on the request, then leave those out as they will be checked at the tighest
	// possible portion of the codebase to protect that resource and won't effect every API call. For
	// example, the PostAnnotations call is often used to add concepts as well. If those concepts are
	// new, then the Concepts_Add scope will be needed. But since you don't always annotate with
	// concepts, it is not a hard requirements that Concepts_Add is used.
	//
	// repeated clarifai.auth.scope.S cl_depending_scopes = 80322;
	E_ClDependingScopes = &file_proto_clarifai_auth_util_extension_proto_extTypes[2]
	// This is the authorizer type for the endpoint our of multiple backend authorizers.
	//
	// optional clarifai.auth.types.AuthType cl_auth_type = 80323;
	E_ClAuthType = &file_proto_clarifai_auth_util_extension_proto_extTypes[3]
)

// Extension fields to descriptor.MessageOptions.
var (
	// indicates if this message should be private. If true, then internal use only. If False,
	// then publicly avaliable to world.
	//
	// optional bool cl_private_message = 80324;
	E_ClPrivateMessage = &file_proto_clarifai_auth_util_extension_proto_extTypes[4]
)

var File_proto_clarifai_auth_util_extension_proto protoreflect.FileDescriptor

var file_proto_clarifai_auth_util_extension_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6c, 0x61, 0x72,
	0x69, 0x66, 0x61, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x1a, 0x25,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61,
	0x72, 0x69, 0x66, 0x61, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x49,
	0x0a, 0x10, 0x63, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xc0, 0xf3, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x6c, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x46, 0x0a, 0x0e, 0x63, 0x6c, 0x5f,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x12, 0x1e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc1, 0xf3, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x63, 0x6c, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x70,
	0x63, 0x3a, 0x68, 0x0a, 0x13, 0x63, 0x6c, 0x5f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc2, 0xf3, 0x04, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x53, 0x52, 0x11, 0x63, 0x6c, 0x44, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x3a, 0x61, 0x0a, 0x0c, 0x63,
	0x6c, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc3, 0xf3, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0a, 0x63, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x4f,
	0x0a, 0x12, 0x63, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc4, 0xf3, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x63,
	0x6c, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x65, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x50, 0x01,
	0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x61,
	0x72, 0x69, 0x66, 0x61, 0x69, 0x2f, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x2d, 0x67,
	0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61,
	0x72, 0x69, 0x66, 0x61, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0xa2,
	0x02, 0x04, 0x43, 0x41, 0x49, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_clarifai_auth_util_extension_proto_goTypes = []interface{}{
	(*descriptor.FieldOptions)(nil),   // 0: google.protobuf.FieldOptions
	(*descriptor.MethodOptions)(nil),  // 1: google.protobuf.MethodOptions
	(*descriptor.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
	(scope.S)(0),                      // 3: clarifai.auth.scope.S
	(types.AuthType)(0),               // 4: clarifai.auth.types.AuthType
}
var file_proto_clarifai_auth_util_extension_proto_depIdxs = []int32{
	0, // 0: clarifai.auth.util.cl_private_field:extendee -> google.protobuf.FieldOptions
	1, // 1: clarifai.auth.util.cl_private_rpc:extendee -> google.protobuf.MethodOptions
	1, // 2: clarifai.auth.util.cl_depending_scopes:extendee -> google.protobuf.MethodOptions
	1, // 3: clarifai.auth.util.cl_auth_type:extendee -> google.protobuf.MethodOptions
	2, // 4: clarifai.auth.util.cl_private_message:extendee -> google.protobuf.MessageOptions
	3, // 5: clarifai.auth.util.cl_depending_scopes:type_name -> clarifai.auth.scope.S
	4, // 6: clarifai.auth.util.cl_auth_type:type_name -> clarifai.auth.types.AuthType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	5, // [5:7] is the sub-list for extension type_name
	0, // [0:5] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_clarifai_auth_util_extension_proto_init() }
func file_proto_clarifai_auth_util_extension_proto_init() {
	if File_proto_clarifai_auth_util_extension_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_clarifai_auth_util_extension_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_proto_clarifai_auth_util_extension_proto_goTypes,
		DependencyIndexes: file_proto_clarifai_auth_util_extension_proto_depIdxs,
		ExtensionInfos:    file_proto_clarifai_auth_util_extension_proto_extTypes,
	}.Build()
	File_proto_clarifai_auth_util_extension_proto = out.File
	file_proto_clarifai_auth_util_extension_proto_rawDesc = nil
	file_proto_clarifai_auth_util_extension_proto_goTypes = nil
	file_proto_clarifai_auth_util_extension_proto_depIdxs = nil
}
