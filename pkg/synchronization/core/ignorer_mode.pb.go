// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: pkg/synchronization/core/ignorer_mode.proto

package core

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

// IgnorerMode specifies the mode for processing scan ignores.
type IgnorerMode int32

const (
	// IgnorerMode_IgnorerModeDefault represents the default mode for
	// using ignores during the scan.  It matches how .gitignore
	// rules are processed.
	IgnorerMode_IgnorerModeDefault IgnorerMode = 0
	// IgnorerMode_IgnorerModeDocker can be used during scans so that
	// ignores follow the Docker ignore (.dockerignore) rules
	IgnorerMode_IgnorerModeDocker IgnorerMode = 1
)

// Enum value maps for IgnorerMode.
var (
	IgnorerMode_name = map[int32]string{
		0: "IgnorerModeDefault",
		1: "IgnorerModeDocker",
	}
	IgnorerMode_value = map[string]int32{
		"IgnorerModeDefault": 0,
		"IgnorerModeDocker":  1,
	}
)

func (x IgnorerMode) Enum() *IgnorerMode {
	p := new(IgnorerMode)
	*p = x
	return p
}

func (x IgnorerMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IgnorerMode) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_synchronization_core_ignorer_mode_proto_enumTypes[0].Descriptor()
}

func (IgnorerMode) Type() protoreflect.EnumType {
	return &file_pkg_synchronization_core_ignorer_mode_proto_enumTypes[0]
}

func (x IgnorerMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IgnorerMode.Descriptor instead.
func (IgnorerMode) EnumDescriptor() ([]byte, []int) {
	return file_pkg_synchronization_core_ignorer_mode_proto_rawDescGZIP(), []int{0}
}

var File_pkg_synchronization_core_ignorer_mode_proto protoreflect.FileDescriptor

var file_pkg_synchronization_core_ignorer_mode_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63,
	0x6f, 0x72, 0x65, 0x2a, 0x3c, 0x0a, 0x0b, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x72, 0x4d, 0x6f, 0x64,
	0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x67,
	0x6e, 0x6f, 0x72, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x10,
	0x01, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_synchronization_core_ignorer_mode_proto_rawDescOnce sync.Once
	file_pkg_synchronization_core_ignorer_mode_proto_rawDescData = file_pkg_synchronization_core_ignorer_mode_proto_rawDesc
)

func file_pkg_synchronization_core_ignorer_mode_proto_rawDescGZIP() []byte {
	file_pkg_synchronization_core_ignorer_mode_proto_rawDescOnce.Do(func() {
		file_pkg_synchronization_core_ignorer_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_synchronization_core_ignorer_mode_proto_rawDescData)
	})
	return file_pkg_synchronization_core_ignorer_mode_proto_rawDescData
}

var file_pkg_synchronization_core_ignorer_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_synchronization_core_ignorer_mode_proto_goTypes = []interface{}{
	(IgnorerMode)(0), // 0: core.IgnorerMode
}
var file_pkg_synchronization_core_ignorer_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_synchronization_core_ignorer_mode_proto_init() }
func file_pkg_synchronization_core_ignorer_mode_proto_init() {
	if File_pkg_synchronization_core_ignorer_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_synchronization_core_ignorer_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_synchronization_core_ignorer_mode_proto_goTypes,
		DependencyIndexes: file_pkg_synchronization_core_ignorer_mode_proto_depIdxs,
		EnumInfos:         file_pkg_synchronization_core_ignorer_mode_proto_enumTypes,
	}.Build()
	File_pkg_synchronization_core_ignorer_mode_proto = out.File
	file_pkg_synchronization_core_ignorer_mode_proto_rawDesc = nil
	file_pkg_synchronization_core_ignorer_mode_proto_goTypes = nil
	file_pkg_synchronization_core_ignorer_mode_proto_depIdxs = nil
}