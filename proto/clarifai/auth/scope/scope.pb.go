// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: proto/clarifai/auth/scope/scope.proto

package scope

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// Next index: 41
// NOTE: When updating the list of "clarifai_exposed" scopes, please also
// update the TestScopes function in main_key_test.go and TestGetExposedScopes function in
// scope_test.go
//
// The dependencies listed for each scope are simply recommendations so that a user
// cannot make a key that would be useless. Beyond the key creation they are not enforced
// but rather the scopes are enforce when data is accessed.
//
// There is the following conventions in place, make sure you add them to the scopes for all new
// resource types:
//
// 1. *_Add requires the corresponding _Get.
// 2. *_Delete requires the corresponding _Add and _Get.
// 3. *_Patch is deprecated and not check anywhere.
//
type S int32

const (
	// introduce undef so that the zero (default/unset) value of the enum is not a real
	// permission.  undef is only present for this purpose and should not be used
	// to indicate any "real" value.
	S_undef S = 0
	S_All   S = 1
	// Make an rpc to our prediction services.
	S_Predict S = 2
	// Make an rpc to our search services.
	S_Search S = 3
	// Write to the inputs table in the DB.
	S_Inputs_Add S = 4
	// Read from the inputs table in the DB.
	S_Inputs_Get S = 5
	// To patch we need read/write.
	// Deprecated.
	// Optionally needs Concepts_Add.
	//
	// Deprecated: Do not use.
	S_Inputs_Patch S = 7
	// To delete we need read/write
	S_Inputs_Delete S = 8
	// Deprecated.
	//
	// Deprecated: Do not use.
	S_Outputs_Patch S = 9
	// Write to the concepts DB tables.
	S_Concepts_Add S = 10
	// Read from the concepts DB tables.
	S_Concepts_Get S = 11
	// To patch we need read/write.
	// Deprecated
	//
	// Deprecated: Do not use.
	S_Concepts_Patch S = 12
	// To delete we need read/write.
	// Note: not implemented.
	S_Concepts_Delete S = 13
	// Write to the models DB tables.
	S_Models_Add S = 14
	// Read from the models and models versions DB tables.
	S_Models_Get S = 15
	// To patch we need read/write.
	// Deprecated.
	//
	// Deprecated: Do not use.
	S_Models_Patch S = 16
	// To delete we need read/write.
	S_Models_Delete S = 17
	// Note: Models_Train is effectively doing POST /models/{models_id}/versions so it's treated that
	// way in terms of data access under the hood.
	// Write to the model versions DB table.
	S_Models_Train S = 26
	// Internal only model syncing.
	S_Models_Sync S = 27
	// Write to the workflows DB table.
	S_Workflows_Add S = 18
	// Read from the workflows DB table.
	S_Workflows_Get S = 19
	// To patch we need read/write.
	// Deprecated.
	//
	// Deprecated: Do not use.
	S_Workflows_Patch S = 20
	// To delete we need read/write.
	S_Workflows_Delete       S = 21
	S_WorkflowMetrics_Get    S = 96
	S_WorkflowMetrics_Add    S = 97
	S_WorkflowMetrics_Delete S = 98
	// Write to the visualizations DB table.
	// Deprecated
	//
	// Deprecated: Do not use.
	S_TSNEVisualizations_Add S = 24
	// Read from the visualizations DB table.
	// Deprecated
	//
	// Deprecated: Do not use.
	S_TSNEVisualizations_Get S = 25
	// Write to the annotations DB table.
	S_Annotations_Add S = 37
	// Read from the annotations DB table.
	S_Annotations_Get S = 38
	// To patch we need read/write.
	// Deprecated.
	//
	// Deprecated: Do not use.
	S_Annotations_Patch S = 39
	// To delete we need read/write.
	S_Annotations_Delete S = 40
	// Write to the collectors DB table.
	S_Collectors_Add S = 41
	// Read from the collectors DB table.
	S_Collectors_Get S = 42
	// To delete we need read/write.
	S_Collectors_Delete S = 43
	// Write to the apps DB table.
	S_Apps_Add S = 44
	// Read from the apps DB table.
	S_Apps_Get S = 45
	// To delete we need read/write.
	S_Apps_Delete S = 46
	// Write to the keys DB table.
	S_Keys_Add S = 47
	// Read from the keys DB table.
	S_Keys_Get S = 48
	// To delete we need read/write.
	S_Keys_Delete S = 49
	// Write to the app sharing DB table
	S_Collaborators_Add S = 51
	// Read from the app sharing DB table
	S_Collaborators_Get S = 50
	// To delete we need read/write
	S_Collaborators_Delete S = 52
	// Write to the metrics table
	S_Metrics_Add S = 54
	// Read from metrics table
	S_Metrics_Get S = 53
	// To delete we need read/write
	S_Metrics_Delete S = 63
	// Write to tasks DB table.
	S_Tasks_Add S = 55
	// Read from the tasks DB table.
	S_Tasks_Get S = 56
	// To delete we need read/write
	S_Tasks_Delete S = 70
	// Write to the password_policies DB table
	S_PasswordPolicies_Add S = 57
	// Read from the password_policies DB table
	S_PasswordPolicies_Get S = 58
	// To delete password_policies we need read/write
	S_PasswordPolicies_Delete S = 59
	// Read from label orders table
	S_LabelOrders_Get S = 67
	// Write to label orders table
	S_LabelOrders_Add S = 68
	// To delete label orders we need read/write
	S_LabelOrders_Delete S = 69
	// Read from user_feature_configs table
	S_UserFeatureConfigs_Get S = 71
	// CRUD on FindDuplicateAnnotationsJobs table
	S_FindDuplicateAnnotationsJobs_Add    S = 102
	S_FindDuplicateAnnotationsJobs_Get    S = 103
	S_FindDuplicateAnnotationsJobs_Delete S = 104
)

// Enum value maps for S.
var (
	S_name = map[int32]string{
		0:   "undef",
		1:   "All",
		2:   "Predict",
		3:   "Search",
		4:   "Inputs_Add",
		5:   "Inputs_Get",
		7:   "Inputs_Patch",
		8:   "Inputs_Delete",
		9:   "Outputs_Patch",
		10:  "Concepts_Add",
		11:  "Concepts_Get",
		12:  "Concepts_Patch",
		13:  "Concepts_Delete",
		14:  "Models_Add",
		15:  "Models_Get",
		16:  "Models_Patch",
		17:  "Models_Delete",
		26:  "Models_Train",
		27:  "Models_Sync",
		18:  "Workflows_Add",
		19:  "Workflows_Get",
		20:  "Workflows_Patch",
		21:  "Workflows_Delete",
		96:  "WorkflowMetrics_Get",
		97:  "WorkflowMetrics_Add",
		98:  "WorkflowMetrics_Delete",
		24:  "TSNEVisualizations_Add",
		25:  "TSNEVisualizations_Get",
		37:  "Annotations_Add",
		38:  "Annotations_Get",
		39:  "Annotations_Patch",
		40:  "Annotations_Delete",
		41:  "Collectors_Add",
		42:  "Collectors_Get",
		43:  "Collectors_Delete",
		44:  "Apps_Add",
		45:  "Apps_Get",
		46:  "Apps_Delete",
		47:  "Keys_Add",
		48:  "Keys_Get",
		49:  "Keys_Delete",
		51:  "Collaborators_Add",
		50:  "Collaborators_Get",
		52:  "Collaborators_Delete",
		54:  "Metrics_Add",
		53:  "Metrics_Get",
		63:  "Metrics_Delete",
		55:  "Tasks_Add",
		56:  "Tasks_Get",
		70:  "Tasks_Delete",
		57:  "PasswordPolicies_Add",
		58:  "PasswordPolicies_Get",
		59:  "PasswordPolicies_Delete",
		67:  "LabelOrders_Get",
		68:  "LabelOrders_Add",
		69:  "LabelOrders_Delete",
		71:  "UserFeatureConfigs_Get",
		102: "FindDuplicateAnnotationsJobs_Add",
		103: "FindDuplicateAnnotationsJobs_Get",
		104: "FindDuplicateAnnotationsJobs_Delete",
	}
	S_value = map[string]int32{
		"undef":                               0,
		"All":                                 1,
		"Predict":                             2,
		"Search":                              3,
		"Inputs_Add":                          4,
		"Inputs_Get":                          5,
		"Inputs_Patch":                        7,
		"Inputs_Delete":                       8,
		"Outputs_Patch":                       9,
		"Concepts_Add":                        10,
		"Concepts_Get":                        11,
		"Concepts_Patch":                      12,
		"Concepts_Delete":                     13,
		"Models_Add":                          14,
		"Models_Get":                          15,
		"Models_Patch":                        16,
		"Models_Delete":                       17,
		"Models_Train":                        26,
		"Models_Sync":                         27,
		"Workflows_Add":                       18,
		"Workflows_Get":                       19,
		"Workflows_Patch":                     20,
		"Workflows_Delete":                    21,
		"WorkflowMetrics_Get":                 96,
		"WorkflowMetrics_Add":                 97,
		"WorkflowMetrics_Delete":              98,
		"TSNEVisualizations_Add":              24,
		"TSNEVisualizations_Get":              25,
		"Annotations_Add":                     37,
		"Annotations_Get":                     38,
		"Annotations_Patch":                   39,
		"Annotations_Delete":                  40,
		"Collectors_Add":                      41,
		"Collectors_Get":                      42,
		"Collectors_Delete":                   43,
		"Apps_Add":                            44,
		"Apps_Get":                            45,
		"Apps_Delete":                         46,
		"Keys_Add":                            47,
		"Keys_Get":                            48,
		"Keys_Delete":                         49,
		"Collaborators_Add":                   51,
		"Collaborators_Get":                   50,
		"Collaborators_Delete":                52,
		"Metrics_Add":                         54,
		"Metrics_Get":                         53,
		"Metrics_Delete":                      63,
		"Tasks_Add":                           55,
		"Tasks_Get":                           56,
		"Tasks_Delete":                        70,
		"PasswordPolicies_Add":                57,
		"PasswordPolicies_Get":                58,
		"PasswordPolicies_Delete":             59,
		"LabelOrders_Get":                     67,
		"LabelOrders_Add":                     68,
		"LabelOrders_Delete":                  69,
		"UserFeatureConfigs_Get":              71,
		"FindDuplicateAnnotationsJobs_Add":    102,
		"FindDuplicateAnnotationsJobs_Get":    103,
		"FindDuplicateAnnotationsJobs_Delete": 104,
	}
)

func (x S) Enum() *S {
	p := new(S)
	*p = x
	return p
}

func (x S) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (S) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_clarifai_auth_scope_scope_proto_enumTypes[0].Descriptor()
}

func (S) Type() protoreflect.EnumType {
	return &file_proto_clarifai_auth_scope_scope_proto_enumTypes[0]
}

func (x S) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use S.Descriptor instead.
func (S) EnumDescriptor() ([]byte, []int) {
	return file_proto_clarifai_auth_scope_scope_proto_rawDescGZIP(), []int{0}
}

type ScopeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// These are the list of low-level scopes to check from the enum below.
	Scopes []S `protobuf:"varint,1,rep,packed,name=scopes,proto3,enum=clarifai.auth.scope.S" json:"scopes,omitempty"`
	// This is a list of fully qualified grpc names to check.
	Endpoints []string `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *ScopeList) Reset() {
	*x = ScopeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clarifai_auth_scope_scope_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeList) ProtoMessage() {}

func (x *ScopeList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clarifai_auth_scope_scope_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeList.ProtoReflect.Descriptor instead.
func (*ScopeList) Descriptor() ([]byte, []int) {
	return file_proto_clarifai_auth_scope_scope_proto_rawDescGZIP(), []int{0}
}

func (x *ScopeList) GetScopes() []S {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *ScopeList) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

var file_proto_clarifai_auth_scope_scope_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.EnumValueOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         80318,
		Name:          "clarifai.auth.scope.clarfai_exposed",
		Tag:           "varint,80318,opt,name=clarfai_exposed",
		Filename:      "proto/clarifai/auth/scope/scope.proto",
	},
	{
		ExtendedType:  (*descriptor.EnumValueOptions)(nil),
		ExtensionType: ([]S)(nil),
		Field:         80319,
		Name:          "clarifai.auth.scope.clarifai_depending_scopes",
		Tag:           "varint,80319,rep,name=clarifai_depending_scopes,enum=clarifai.auth.scope.S",
		Filename:      "proto/clarifai/auth/scope/scope.proto",
	},
}

// Extension fields to descriptor.EnumValueOptions.
var (
	// indicates whether the given scope should be returned by the Get /scopes/ call
	// or any other call that returns list of available perms.
	//
	// optional bool clarfai_exposed = 80318;
	E_ClarfaiExposed = &file_proto_clarifai_auth_scope_scope_proto_extTypes[0]
	// TODO: We have no way of picking extension field numbers within clarifai to guarantee
	// uniqueness.  Note: 50000-99999 are for organizational use (like this)
	//
	// repeated clarifai.auth.scope.S clarifai_depending_scopes = 80319;
	E_ClarifaiDependingScopes = &file_proto_clarifai_auth_scope_scope_proto_extTypes[1]
)

var File_proto_clarifai_auth_scope_scope_proto protoreflect.FileDescriptor

var file_proto_clarifai_auth_scope_scope_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2f, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61,
	0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59,
	0x0a, 0x09, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6c,
	0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x2e, 0x53, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2a, 0xc8, 0x0e, 0x0a, 0x01, 0x53, 0x12,
	0x09, 0x0a, 0x05, 0x75, 0x6e, 0x64, 0x65, 0x66, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x03, 0x41, 0x6c,
	0x6c, 0x10, 0x01, 0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01, 0x12, 0x11, 0x0a, 0x07, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x10, 0x02, 0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01, 0x12, 0x10, 0x0a, 0x06,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x10, 0x03, 0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01, 0x12, 0x18,
	0x0a, 0x0a, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x5f, 0x41, 0x64, 0x64, 0x10, 0x04, 0x1a, 0x08,
	0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x05, 0x12, 0x14, 0x0a, 0x0a, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x10, 0x05, 0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01, 0x12, 0x20,
	0x0a, 0x0c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x5f, 0x50, 0x61, 0x74, 0x63, 0x68, 0x10, 0x07,
	0x1a, 0x0e, 0x08, 0x01, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x04, 0xf8, 0x9b, 0x27, 0x05,
	0x12, 0x1f, 0x0a, 0x0d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x10, 0x08, 0x1a, 0x0c, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x04, 0xf8, 0x9b, 0x27,
	0x05, 0x12, 0x1d, 0x0a, 0x0d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x5f, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x10, 0x09, 0x1a, 0x0a, 0x08, 0x01, 0xf8, 0x9b, 0x27, 0x05, 0xf8, 0x9b, 0x27, 0x02,
	0x12, 0x1a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x73, 0x5f, 0x41, 0x64, 0x64,
	0x10, 0x0a, 0x1a, 0x08, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x0b, 0x12, 0x16, 0x0a, 0x0c,
	0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x10, 0x0b, 0x1a, 0x04,
	0xf0, 0x9b, 0x27, 0x01, 0x12, 0x22, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x73,
	0x5f, 0x50, 0x61, 0x74, 0x63, 0x68, 0x10, 0x0c, 0x1a, 0x0e, 0x08, 0x01, 0xf0, 0x9b, 0x27, 0x01,
	0xf8, 0x9b, 0x27, 0x0a, 0xf8, 0x9b, 0x27, 0x0b, 0x12, 0x1d, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x63,
	0x65, 0x70, 0x74, 0x73, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x0d, 0x1a, 0x08, 0xf8,
	0x9b, 0x27, 0x0a, 0xf8, 0x9b, 0x27, 0x0b, 0x12, 0x18, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x5f, 0x41, 0x64, 0x64, 0x10, 0x0e, 0x1a, 0x08, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27,
	0x0f, 0x12, 0x14, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x10,
	0x0f, 0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01, 0x12, 0x24, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x5f, 0x50, 0x61, 0x74, 0x63, 0x68, 0x10, 0x10, 0x1a, 0x12, 0x08, 0x01, 0xf0, 0x9b, 0x27,
	0x01, 0xf8, 0x9b, 0x27, 0x0e, 0xf8, 0x9b, 0x27, 0x0f, 0xf8, 0x9b, 0x27, 0x1a, 0x12, 0x1f, 0x0a,
	0x0d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x11,
	0x1a, 0x0c, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x0e, 0xf8, 0x9b, 0x27, 0x0f, 0x12, 0x1a,
	0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5f, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x10, 0x1a,
	0x1a, 0x08, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x0f, 0x12, 0x15, 0x0a, 0x0b, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x5f, 0x53, 0x79, 0x6e, 0x63, 0x10, 0x1b, 0x1a, 0x04, 0xf8, 0x9b, 0x27,
	0x0f, 0x12, 0x1b, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x5f, 0x41,
	0x64, 0x64, 0x10, 0x12, 0x1a, 0x08, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x13, 0x12, 0x17,
	0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x10,
	0x13, 0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01, 0x12, 0x23, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x73, 0x5f, 0x50, 0x61, 0x74, 0x63, 0x68, 0x10, 0x14, 0x1a, 0x0e, 0x08, 0x01,
	0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x12, 0xf8, 0x9b, 0x27, 0x13, 0x12, 0x22, 0x0a, 0x10,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x10, 0x15, 0x1a, 0x0c, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x12, 0xf8, 0x9b, 0x27, 0x13,
	0x12, 0x1d, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x10, 0x60, 0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01, 0x12,
	0x21, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x5f, 0x41, 0x64, 0x64, 0x10, 0x61, 0x1a, 0x08, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b,
	0x27, 0x60, 0x12, 0x28, 0x0a, 0x16, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x62, 0x1a, 0x0c,
	0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x61, 0xf8, 0x9b, 0x27, 0x60, 0x12, 0x22, 0x0a, 0x16,
	0x54, 0x53, 0x4e, 0x45, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x5f, 0x41, 0x64, 0x64, 0x10, 0x18, 0x1a, 0x06, 0x08, 0x01, 0xf8, 0x9b, 0x27, 0x19,
	0x12, 0x1e, 0x0a, 0x16, 0x54, 0x53, 0x4e, 0x45, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x10, 0x19, 0x1a, 0x02, 0x08, 0x01,
	0x12, 0x1d, 0x0a, 0x0f, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f,
	0x41, 0x64, 0x64, 0x10, 0x25, 0x1a, 0x08, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x26, 0x12,
	0x19, 0x0a, 0x0f, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x47,
	0x65, 0x74, 0x10, 0x26, 0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01, 0x12, 0x25, 0x0a, 0x11, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x50, 0x61, 0x74, 0x63, 0x68, 0x10,
	0x27, 0x1a, 0x0e, 0x08, 0x01, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x25, 0xf8, 0x9b, 0x27,
	0x26, 0x12, 0x24, 0x0a, 0x12, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x28, 0x1a, 0x0c, 0xf0, 0x9b, 0x27, 0x01, 0xf8,
	0x9b, 0x27, 0x25, 0xf8, 0x9b, 0x27, 0x26, 0x12, 0x1c, 0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x41, 0x64, 0x64, 0x10, 0x29, 0x1a, 0x08, 0xf0, 0x9b, 0x27,
	0x01, 0xf8, 0x9b, 0x27, 0x2a, 0x12, 0x18, 0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x10, 0x2a, 0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01, 0x12,
	0x23, 0x0a, 0x11, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x10, 0x2b, 0x1a, 0x0c, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x29,
	0xf8, 0x9b, 0x27, 0x2a, 0x12, 0x16, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x73, 0x5f, 0x41, 0x64, 0x64,
	0x10, 0x2c, 0x1a, 0x08, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x2d, 0x12, 0x12, 0x0a, 0x08,
	0x41, 0x70, 0x70, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x10, 0x2d, 0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01,
	0x12, 0x1d, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x73, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10,
	0x2e, 0x1a, 0x0c, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x2c, 0xf8, 0x9b, 0x27, 0x2d, 0x12,
	0x16, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x73, 0x5f, 0x41, 0x64, 0x64, 0x10, 0x2f, 0x1a, 0x08, 0xf0,
	0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x30, 0x12, 0x12, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x73, 0x5f,
	0x47, 0x65, 0x74, 0x10, 0x30, 0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01, 0x12, 0x1d, 0x0a, 0x0b, 0x4b,
	0x65, 0x79, 0x73, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x31, 0x1a, 0x0c, 0xf0, 0x9b,
	0x27, 0x01, 0xf8, 0x9b, 0x27, 0x2f, 0xf8, 0x9b, 0x27, 0x30, 0x12, 0x1f, 0x0a, 0x11, 0x43, 0x6f,
	0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x41, 0x64, 0x64, 0x10,
	0x33, 0x1a, 0x08, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x32, 0x12, 0x1b, 0x0a, 0x11, 0x43,
	0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x47, 0x65, 0x74,
	0x10, 0x32, 0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01, 0x12, 0x26, 0x0a, 0x14, 0x43, 0x6f, 0x6c, 0x6c,
	0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x10, 0x34, 0x1a, 0x0c, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x33, 0xf8, 0x9b, 0x27, 0x32,
	0x12, 0x19, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x41, 0x64, 0x64, 0x10,
	0x36, 0x1a, 0x08, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x35, 0x12, 0x15, 0x0a, 0x0b, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x10, 0x35, 0x1a, 0x04, 0xf0, 0x9b,
	0x27, 0x01, 0x12, 0x20, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x10, 0x3f, 0x1a, 0x0c, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x36,
	0xf8, 0x9b, 0x27, 0x35, 0x12, 0x17, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x41, 0x64,
	0x64, 0x10, 0x37, 0x1a, 0x08, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x38, 0x12, 0x13, 0x0a,
	0x09, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x10, 0x38, 0x1a, 0x04, 0xf0, 0x9b,
	0x27, 0x01, 0x12, 0x1e, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x10, 0x46, 0x1a, 0x0c, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x37, 0xf8, 0x9b,
	0x27, 0x38, 0x12, 0x22, 0x0a, 0x14, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x41, 0x64, 0x64, 0x10, 0x39, 0x1a, 0x08, 0xf0, 0x9b,
	0x27, 0x01, 0xf8, 0x9b, 0x27, 0x3a, 0x12, 0x1e, 0x0a, 0x14, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x10, 0x3a,
	0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01, 0x12, 0x29, 0x0a, 0x17, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x10, 0x3b, 0x1a, 0x0c, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x39, 0xf8, 0x9b, 0x27,
	0x3a, 0x12, 0x19, 0x0a, 0x0f, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x5f, 0x47, 0x65, 0x74, 0x10, 0x43, 0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01, 0x12, 0x1d, 0x0a, 0x0f,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x41, 0x64, 0x64, 0x10,
	0x44, 0x1a, 0x08, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x43, 0x12, 0x24, 0x0a, 0x12, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x10, 0x45, 0x1a, 0x0c, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b, 0x27, 0x44, 0xf8, 0x9b, 0x27,
	0x43, 0x12, 0x20, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x10, 0x47, 0x1a, 0x04, 0xf0,
	0x9b, 0x27, 0x01, 0x12, 0x2e, 0x0a, 0x20, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x75, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a,
	0x6f, 0x62, 0x73, 0x5f, 0x41, 0x64, 0x64, 0x10, 0x66, 0x1a, 0x08, 0xf0, 0x9b, 0x27, 0x01, 0xf8,
	0x9b, 0x27, 0x67, 0x12, 0x2a, 0x0a, 0x20, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x75, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a,
	0x6f, 0x62, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x10, 0x67, 0x1a, 0x04, 0xf0, 0x9b, 0x27, 0x01, 0x12,
	0x35, 0x0a, 0x23, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x6f, 0x62, 0x73, 0x5f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x68, 0x1a, 0x0c, 0xf0, 0x9b, 0x27, 0x01, 0xf8, 0x9b,
	0x27, 0x66, 0xf8, 0x9b, 0x27, 0x67, 0x22, 0x04, 0x08, 0x1e, 0x10, 0x1e, 0x22, 0x04, 0x08, 0x1f,
	0x10, 0x1f, 0x22, 0x04, 0x08, 0x20, 0x10, 0x20, 0x22, 0x04, 0x08, 0x21, 0x10, 0x21, 0x22, 0x04,
	0x08, 0x22, 0x10, 0x22, 0x3a, 0x4c, 0x0a, 0x0f, 0x63, 0x6c, 0x61, 0x72, 0x66, 0x61, 0x69, 0x5f,
	0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbe, 0xf3, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x63, 0x6c, 0x61, 0x72, 0x66, 0x61, 0x69, 0x45, 0x78, 0x70, 0x6f, 0x73,
	0x65, 0x64, 0x3a, 0x77, 0x0a, 0x19, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x5f, 0x64,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12,
	0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xbf, 0xf3, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6c, 0x61,
	0x72, 0x69, 0x66, 0x61, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x2e, 0x53, 0x52, 0x17, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x44, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x42, 0x67, 0x0a, 0x1c, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x50, 0x01, 0x5a, 0x3e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x61, 0x72, 0x69, 0x66,
	0x61, 0x69, 0x2f, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x61, 0x69, 0x2d, 0x67, 0x6f, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66,
	0x61, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0xa2, 0x02, 0x04,
	0x43, 0x41, 0x49, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_clarifai_auth_scope_scope_proto_rawDescOnce sync.Once
	file_proto_clarifai_auth_scope_scope_proto_rawDescData = file_proto_clarifai_auth_scope_scope_proto_rawDesc
)

func file_proto_clarifai_auth_scope_scope_proto_rawDescGZIP() []byte {
	file_proto_clarifai_auth_scope_scope_proto_rawDescOnce.Do(func() {
		file_proto_clarifai_auth_scope_scope_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_clarifai_auth_scope_scope_proto_rawDescData)
	})
	return file_proto_clarifai_auth_scope_scope_proto_rawDescData
}

var file_proto_clarifai_auth_scope_scope_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_clarifai_auth_scope_scope_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_clarifai_auth_scope_scope_proto_goTypes = []interface{}{
	(S)(0),                              // 0: clarifai.auth.scope.S
	(*ScopeList)(nil),                   // 1: clarifai.auth.scope.ScopeList
	(*descriptor.EnumValueOptions)(nil), // 2: google.protobuf.EnumValueOptions
}
var file_proto_clarifai_auth_scope_scope_proto_depIdxs = []int32{
	0, // 0: clarifai.auth.scope.ScopeList.scopes:type_name -> clarifai.auth.scope.S
	2, // 1: clarifai.auth.scope.clarfai_exposed:extendee -> google.protobuf.EnumValueOptions
	2, // 2: clarifai.auth.scope.clarifai_depending_scopes:extendee -> google.protobuf.EnumValueOptions
	0, // 3: clarifai.auth.scope.clarifai_depending_scopes:type_name -> clarifai.auth.scope.S
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	3, // [3:4] is the sub-list for extension type_name
	1, // [1:3] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_clarifai_auth_scope_scope_proto_init() }
func file_proto_clarifai_auth_scope_scope_proto_init() {
	if File_proto_clarifai_auth_scope_scope_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_clarifai_auth_scope_scope_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopeList); i {
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
			RawDescriptor: file_proto_clarifai_auth_scope_scope_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_proto_clarifai_auth_scope_scope_proto_goTypes,
		DependencyIndexes: file_proto_clarifai_auth_scope_scope_proto_depIdxs,
		EnumInfos:         file_proto_clarifai_auth_scope_scope_proto_enumTypes,
		MessageInfos:      file_proto_clarifai_auth_scope_scope_proto_msgTypes,
		ExtensionInfos:    file_proto_clarifai_auth_scope_scope_proto_extTypes,
	}.Build()
	File_proto_clarifai_auth_scope_scope_proto = out.File
	file_proto_clarifai_auth_scope_scope_proto_rawDesc = nil
	file_proto_clarifai_auth_scope_scope_proto_goTypes = nil
	file_proto_clarifai_auth_scope_scope_proto_depIdxs = nil
}
