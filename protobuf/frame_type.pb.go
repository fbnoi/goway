// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: protobuf/frame_type.proto

package protobuf

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

type FrameType int32

const (
	// control frame
	FrameType_HEARTBEAT FrameType = 0
	FrameType_CLOSE     FrameType = 1
	FrameType_MAINTAIN  FrameType = 2
	// game message
	FrameType_DROP_DOT           FrameType = 3
	FrameType_DROP_DOT_CONFIRM   FrameType = 4
	FrameType_DROP_CHECK         FrameType = 5
	FrameType_DROP_CHECK_CONFIRM FrameType = 6
	FrameType_WIN                FrameType = 7
	FrameType_SURRENDER          FrameType = 8
)

// Enum value maps for FrameType.
var (
	FrameType_name = map[int32]string{
		0: "HEARTBEAT",
		1: "CLOSE",
		2: "MAINTAIN",
		3: "DROP_DOT",
		4: "DROP_DOT_CONFIRM",
		5: "DROP_CHECK",
		6: "DROP_CHECK_CONFIRM",
		7: "WIN",
		8: "SURRENDER",
	}
	FrameType_value = map[string]int32{
		"HEARTBEAT":          0,
		"CLOSE":              1,
		"MAINTAIN":           2,
		"DROP_DOT":           3,
		"DROP_DOT_CONFIRM":   4,
		"DROP_CHECK":         5,
		"DROP_CHECK_CONFIRM": 6,
		"WIN":                7,
		"SURRENDER":          8,
	}
)

func (x FrameType) Enum() *FrameType {
	p := new(FrameType)
	*p = x
	return p
}

func (x FrameType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FrameType) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_frame_type_proto_enumTypes[0].Descriptor()
}

func (FrameType) Type() protoreflect.EnumType {
	return &file_protobuf_frame_type_proto_enumTypes[0]
}

func (x FrameType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FrameType.Descriptor instead.
func (FrameType) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_frame_type_proto_rawDescGZIP(), []int{0}
}

var File_protobuf_frame_type_proto protoreflect.FileDescriptor

var file_protobuf_frame_type_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2a, 0x97, 0x01, 0x0a, 0x09, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x4d, 0x41, 0x49, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x52, 0x4f, 0x50, 0x5f, 0x44, 0x4f, 0x54, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x52, 0x4f,
	0x50, 0x5f, 0x44, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x10, 0x04, 0x12,
	0x0e, 0x0a, 0x0a, 0x44, 0x52, 0x4f, 0x50, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x10, 0x05, 0x12,
	0x16, 0x0a, 0x12, 0x44, 0x52, 0x4f, 0x50, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x43, 0x4f,
	0x4e, 0x46, 0x49, 0x52, 0x4d, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x49, 0x4e, 0x10, 0x07,
	0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x08, 0x42,
	0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_frame_type_proto_rawDescOnce sync.Once
	file_protobuf_frame_type_proto_rawDescData = file_protobuf_frame_type_proto_rawDesc
)

func file_protobuf_frame_type_proto_rawDescGZIP() []byte {
	file_protobuf_frame_type_proto_rawDescOnce.Do(func() {
		file_protobuf_frame_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_frame_type_proto_rawDescData)
	})
	return file_protobuf_frame_type_proto_rawDescData
}

var file_protobuf_frame_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_frame_type_proto_goTypes = []interface{}{
	(FrameType)(0), // 0: protobuf.FrameType
}
var file_protobuf_frame_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_frame_type_proto_init() }
func file_protobuf_frame_type_proto_init() {
	if File_protobuf_frame_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_frame_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_frame_type_proto_goTypes,
		DependencyIndexes: file_protobuf_frame_type_proto_depIdxs,
		EnumInfos:         file_protobuf_frame_type_proto_enumTypes,
	}.Build()
	File_protobuf_frame_type_proto = out.File
	file_protobuf_frame_type_proto_rawDesc = nil
	file_protobuf_frame_type_proto_goTypes = nil
	file_protobuf_frame_type_proto_depIdxs = nil
}
