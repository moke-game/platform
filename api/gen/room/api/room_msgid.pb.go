// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: room/room_msgid.proto

package room

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

// MsgID
type MsgID int32

const (
	MsgID_MSG_ID_INVALID   MsgID = 0
	MsgID_MSG_ID_HEARTBEAT MsgID = 1
	MsgID_MSG_ID_ROOM_JOIN MsgID = 1000
	MsgID_MSG_ID_ROOM_EXIT MsgID = 1001
	MsgID_MSG_ID_ROOM_SYNC MsgID = 1002
)

// Enum value maps for MsgID.
var (
	MsgID_name = map[int32]string{
		0:    "MSG_ID_INVALID",
		1:    "MSG_ID_HEARTBEAT",
		1000: "MSG_ID_ROOM_JOIN",
		1001: "MSG_ID_ROOM_EXIT",
		1002: "MSG_ID_ROOM_SYNC",
	}
	MsgID_value = map[string]int32{
		"MSG_ID_INVALID":   0,
		"MSG_ID_HEARTBEAT": 1,
		"MSG_ID_ROOM_JOIN": 1000,
		"MSG_ID_ROOM_EXIT": 1001,
		"MSG_ID_ROOM_SYNC": 1002,
	}
)

func (x MsgID) Enum() *MsgID {
	p := new(MsgID)
	*p = x
	return p
}

func (x MsgID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgID) Descriptor() protoreflect.EnumDescriptor {
	return file_room_room_msgid_proto_enumTypes[0].Descriptor()
}

func (MsgID) Type() protoreflect.EnumType {
	return &file_room_room_msgid_proto_enumTypes[0]
}

func (x MsgID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgID.Descriptor instead.
func (MsgID) EnumDescriptor() ([]byte, []int) {
	return file_room_room_msgid_proto_rawDescGZIP(), []int{0}
}

// NoticeID
type NoticeID int32

const (
	NoticeID_NOTICE_ID_INVALID     NoticeID = 0
	NoticeID_NOTICE_ID_ROOM_JOINED NoticeID = 100
	NoticeID_NOTICE_ID_ROOM_EXIT   NoticeID = 101
	NoticeID_NOTICE_ID_ROOM_SYNC   NoticeID = 102
)

// Enum value maps for NoticeID.
var (
	NoticeID_name = map[int32]string{
		0:   "NOTICE_ID_INVALID",
		100: "NOTICE_ID_ROOM_JOINED",
		101: "NOTICE_ID_ROOM_EXIT",
		102: "NOTICE_ID_ROOM_SYNC",
	}
	NoticeID_value = map[string]int32{
		"NOTICE_ID_INVALID":     0,
		"NOTICE_ID_ROOM_JOINED": 100,
		"NOTICE_ID_ROOM_EXIT":   101,
		"NOTICE_ID_ROOM_SYNC":   102,
	}
)

func (x NoticeID) Enum() *NoticeID {
	p := new(NoticeID)
	*p = x
	return p
}

func (x NoticeID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NoticeID) Descriptor() protoreflect.EnumDescriptor {
	return file_room_room_msgid_proto_enumTypes[1].Descriptor()
}

func (NoticeID) Type() protoreflect.EnumType {
	return &file_room_room_msgid_proto_enumTypes[1]
}

func (x NoticeID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NoticeID.Descriptor instead.
func (NoticeID) EnumDescriptor() ([]byte, []int) {
	return file_room_room_msgid_proto_rawDescGZIP(), []int{1}
}

var File_room_room_msgid_proto protoreflect.FileDescriptor

var file_room_room_msgid_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6d, 0x73, 0x67, 0x69,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31,
	0x2a, 0x76, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x53, 0x47,
	0x5f, 0x49, 0x44, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x4d, 0x53, 0x47, 0x5f, 0x49, 0x44, 0x5f, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41,
	0x54, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x10, 0x4d, 0x53, 0x47, 0x5f, 0x49, 0x44, 0x5f, 0x52, 0x4f,
	0x4f, 0x4d, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x10, 0xe8, 0x07, 0x12, 0x15, 0x0a, 0x10, 0x4d, 0x53,
	0x47, 0x5f, 0x49, 0x44, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x58, 0x49, 0x54, 0x10, 0xe9,
	0x07, 0x12, 0x15, 0x0a, 0x10, 0x4d, 0x53, 0x47, 0x5f, 0x49, 0x44, 0x5f, 0x52, 0x4f, 0x4f, 0x4d,
	0x5f, 0x53, 0x59, 0x4e, 0x43, 0x10, 0xea, 0x07, 0x2a, 0x6e, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x49, 0x44, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x4f, 0x54, 0x49, 0x43, 0x45, 0x5f, 0x49,
	0x44, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4e,
	0x4f, 0x54, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x4a, 0x4f,
	0x49, 0x4e, 0x45, 0x44, 0x10, 0x64, 0x12, 0x17, 0x0a, 0x13, 0x4e, 0x4f, 0x54, 0x49, 0x43, 0x45,
	0x5f, 0x49, 0x44, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x58, 0x49, 0x54, 0x10, 0x65, 0x12,
	0x17, 0x0a, 0x13, 0x4e, 0x4f, 0x54, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x52, 0x4f, 0x4f,
	0x4d, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x66, 0x42, 0x69, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e,
	0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x73, 0x67,
	0x69, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0d, 0x72, 0x6f, 0x6f, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x3b, 0x72, 0x6f, 0x6f, 0x6d, 0xa2, 0x02, 0x03, 0x52, 0x58, 0x58, 0xaa, 0x02,
	0x07, 0x52, 0x6f, 0x6f, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x52, 0x6f, 0x6f, 0x6d, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x13, 0x52, 0x6f, 0x6f, 0x6d, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_room_room_msgid_proto_rawDescOnce sync.Once
	file_room_room_msgid_proto_rawDescData = file_room_room_msgid_proto_rawDesc
)

func file_room_room_msgid_proto_rawDescGZIP() []byte {
	file_room_room_msgid_proto_rawDescOnce.Do(func() {
		file_room_room_msgid_proto_rawDescData = protoimpl.X.CompressGZIP(file_room_room_msgid_proto_rawDescData)
	})
	return file_room_room_msgid_proto_rawDescData
}

var file_room_room_msgid_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_room_room_msgid_proto_goTypes = []any{
	(MsgID)(0),    // 0: room.v1.MsgID
	(NoticeID)(0), // 1: room.v1.NoticeID
}
var file_room_room_msgid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_room_room_msgid_proto_init() }
func file_room_room_msgid_proto_init() {
	if File_room_room_msgid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_room_room_msgid_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_room_room_msgid_proto_goTypes,
		DependencyIndexes: file_room_room_msgid_proto_depIdxs,
		EnumInfos:         file_room_room_msgid_proto_enumTypes,
	}.Build()
	File_room_room_msgid_proto = out.File
	file_room_room_msgid_proto_rawDesc = nil
	file_room_room_msgid_proto_goTypes = nil
	file_room_room_msgid_proto_depIdxs = nil
}
