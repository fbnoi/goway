// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: protobuf/ctrl.proto

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

type Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type FrameType `protobuf:"varint,1,opt,name=type,proto3,enum=protobuf.FrameType" json:"type,omitempty"`
	Body []byte    `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Frame) Reset() {
	*x = Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_ctrl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frame) ProtoMessage() {}

func (x *Frame) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_ctrl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frame.ProtoReflect.Descriptor instead.
func (*Frame) Descriptor() ([]byte, []int) {
	return file_protobuf_ctrl_proto_rawDescGZIP(), []int{0}
}

func (x *Frame) GetType() FrameType {
	if x != nil {
		return x.Type
	}
	return FrameType_HEARTBEAT
}

func (x *Frame) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type HeartBeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpTimestamp   int64 `protobuf:"varint,1,opt,name=upTimestamp,proto3" json:"upTimestamp,omitempty"`
	DownTimestamp int64 `protobuf:"varint,2,opt,name=downTimestamp,proto3" json:"downTimestamp,omitempty"`
}

func (x *HeartBeat) Reset() {
	*x = HeartBeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_ctrl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeat) ProtoMessage() {}

func (x *HeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_ctrl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeat.ProtoReflect.Descriptor instead.
func (*HeartBeat) Descriptor() ([]byte, []int) {
	return file_protobuf_ctrl_proto_rawDescGZIP(), []int{1}
}

func (x *HeartBeat) GetUpTimestamp() int64 {
	if x != nil {
		return x.UpTimestamp
	}
	return 0
}

func (x *HeartBeat) GetDownTimestamp() int64 {
	if x != nil {
		return x.DownTimestamp
	}
	return 0
}

type Close struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Close) Reset() {
	*x = Close{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_ctrl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Close) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Close) ProtoMessage() {}

func (x *Close) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_ctrl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Close.ProtoReflect.Descriptor instead.
func (*Close) Descriptor() ([]byte, []int) {
	return file_protobuf_ctrl_proto_rawDescGZIP(), []int{2}
}

type Maintain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Maintain) Reset() {
	*x = Maintain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_ctrl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Maintain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Maintain) ProtoMessage() {}

func (x *Maintain) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_ctrl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Maintain.ProtoReflect.Descriptor instead.
func (*Maintain) Descriptor() ([]byte, []int) {
	return file_protobuf_ctrl_proto_rawDescGZIP(), []int{3}
}

var File_protobuf_ctrl_proto protoreflect.FileDescriptor

var file_protobuf_ctrl_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x74, 0x72, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a,
	0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x05, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x22, 0x53, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x24, 0x0a, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x07, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x22, 0x0a,
	0x0a, 0x08, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_ctrl_proto_rawDescOnce sync.Once
	file_protobuf_ctrl_proto_rawDescData = file_protobuf_ctrl_proto_rawDesc
)

func file_protobuf_ctrl_proto_rawDescGZIP() []byte {
	file_protobuf_ctrl_proto_rawDescOnce.Do(func() {
		file_protobuf_ctrl_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_ctrl_proto_rawDescData)
	})
	return file_protobuf_ctrl_proto_rawDescData
}

var file_protobuf_ctrl_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protobuf_ctrl_proto_goTypes = []interface{}{
	(*Frame)(nil),     // 0: protobuf.Frame
	(*HeartBeat)(nil), // 1: protobuf.HeartBeat
	(*Close)(nil),     // 2: protobuf.Close
	(*Maintain)(nil),  // 3: protobuf.Maintain
	(FrameType)(0),    // 4: protobuf.FrameType
}
var file_protobuf_ctrl_proto_depIdxs = []int32{
	4, // 0: protobuf.Frame.type:type_name -> protobuf.FrameType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_ctrl_proto_init() }
func file_protobuf_ctrl_proto_init() {
	if File_protobuf_ctrl_proto != nil {
		return
	}
	file_protobuf_frame_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protobuf_ctrl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frame); i {
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
		file_protobuf_ctrl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeat); i {
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
		file_protobuf_ctrl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Close); i {
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
		file_protobuf_ctrl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Maintain); i {
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
			RawDescriptor: file_protobuf_ctrl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_ctrl_proto_goTypes,
		DependencyIndexes: file_protobuf_ctrl_proto_depIdxs,
		MessageInfos:      file_protobuf_ctrl_proto_msgTypes,
	}.Build()
	File_protobuf_ctrl_proto = out.File
	file_protobuf_ctrl_proto_rawDesc = nil
	file_protobuf_ctrl_proto_goTypes = nil
	file_protobuf_ctrl_proto_depIdxs = nil
}
