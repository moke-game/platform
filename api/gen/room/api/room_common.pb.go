// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: room/room_common.proto

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

// Room error code
type RoomErrorCode int32

const (
	RoomErrorCode_ROOM_ERROR_CODE_OK              RoomErrorCode = 0
	RoomErrorCode_ROOM_ERROR_CODE_INVALID         RoomErrorCode = 1
	RoomErrorCode_ROOM_ERROR_CODE_NOT_FOUND       RoomErrorCode = 2
	RoomErrorCode_ROOM_ERROR_CODE_FULL            RoomErrorCode = 3
	RoomErrorCode_ROOM_ERROR_CODE_ALREADY_IN      RoomErrorCode = 4
	RoomErrorCode_ROOM_ERROR_CODE_NOT_IN          RoomErrorCode = 5
	RoomErrorCode_ROOM_ERROR_CODE_NOT_READY       RoomErrorCode = 6
	RoomErrorCode_ROOM_ERROR_CODE_ALREADY_READY   RoomErrorCode = 7
	RoomErrorCode_ROOM_ERROR_CODE_NOT_STARTED     RoomErrorCode = 8
	RoomErrorCode_ROOM_ERROR_CODE_ALREADY_STARTED RoomErrorCode = 9
	RoomErrorCode_ROOM_ERROR_CODE_NOT_ENDED       RoomErrorCode = 10
	RoomErrorCode_ROOM_ERROR_CODE_ALREADY_ENDED   RoomErrorCode = 11
	RoomErrorCode_ROOM_ERROR_CODE_NOT_SYNC        RoomErrorCode = 12
	RoomErrorCode_ROOM_ERROR_CODE_ALREADY_SYNC    RoomErrorCode = 13
)

// Enum value maps for RoomErrorCode.
var (
	RoomErrorCode_name = map[int32]string{
		0:  "ROOM_ERROR_CODE_OK",
		1:  "ROOM_ERROR_CODE_INVALID",
		2:  "ROOM_ERROR_CODE_NOT_FOUND",
		3:  "ROOM_ERROR_CODE_FULL",
		4:  "ROOM_ERROR_CODE_ALREADY_IN",
		5:  "ROOM_ERROR_CODE_NOT_IN",
		6:  "ROOM_ERROR_CODE_NOT_READY",
		7:  "ROOM_ERROR_CODE_ALREADY_READY",
		8:  "ROOM_ERROR_CODE_NOT_STARTED",
		9:  "ROOM_ERROR_CODE_ALREADY_STARTED",
		10: "ROOM_ERROR_CODE_NOT_ENDED",
		11: "ROOM_ERROR_CODE_ALREADY_ENDED",
		12: "ROOM_ERROR_CODE_NOT_SYNC",
		13: "ROOM_ERROR_CODE_ALREADY_SYNC",
	}
	RoomErrorCode_value = map[string]int32{
		"ROOM_ERROR_CODE_OK":              0,
		"ROOM_ERROR_CODE_INVALID":         1,
		"ROOM_ERROR_CODE_NOT_FOUND":       2,
		"ROOM_ERROR_CODE_FULL":            3,
		"ROOM_ERROR_CODE_ALREADY_IN":      4,
		"ROOM_ERROR_CODE_NOT_IN":          5,
		"ROOM_ERROR_CODE_NOT_READY":       6,
		"ROOM_ERROR_CODE_ALREADY_READY":   7,
		"ROOM_ERROR_CODE_NOT_STARTED":     8,
		"ROOM_ERROR_CODE_ALREADY_STARTED": 9,
		"ROOM_ERROR_CODE_NOT_ENDED":       10,
		"ROOM_ERROR_CODE_ALREADY_ENDED":   11,
		"ROOM_ERROR_CODE_NOT_SYNC":        12,
		"ROOM_ERROR_CODE_ALREADY_SYNC":    13,
	}
)

func (x RoomErrorCode) Enum() *RoomErrorCode {
	p := new(RoomErrorCode)
	*p = x
	return p
}

func (x RoomErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_room_room_common_proto_enumTypes[0].Descriptor()
}

func (RoomErrorCode) Type() protoreflect.EnumType {
	return &file_room_room_common_proto_enumTypes[0]
}

func (x RoomErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomErrorCode.Descriptor instead.
func (RoomErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_room_room_common_proto_rawDescGZIP(), []int{0}
}

// Common response
// Common response message,
// all Response messages should be returned to the client as the message field of this message
// 通用响应消息，所有Response消息都应该作为此消息的message字段返回给客户端
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode RoomErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=room.v1.RoomErrorCode" json:"error_code,omitempty"`
	Message   []byte        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_room_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_room_room_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_room_room_common_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetErrorCode() RoomErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return RoomErrorCode_ROOM_ERROR_CODE_OK
}

func (x *Response) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

// Command data
// Operation frame data, used for client to send to the server, and the server broadcasts to all clients in the room
// 操作帧数据, 用于客户端发送给服务器，服务器广播给房间内所有客户端
type CmdData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// player uid
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"` // player uid
	// joystick x axis value
	X float32 `protobuf:"fixed32,2,opt,name=x,proto3" json:"x,omitempty"` // x position
	// joystick y axis value
	Y float32 `protobuf:"fixed32,3,opt,name=y,proto3" json:"y,omitempty"` // y position
	// player action
	Action int32 `protobuf:"varint,4,opt,name=action,proto3" json:"action,omitempty"` // action
	// custom data
	Custom []byte `protobuf:"bytes,5,opt,name=custom,proto3" json:"custom,omitempty"` // custom data
}

func (x *CmdData) Reset() {
	*x = CmdData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_room_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdData) ProtoMessage() {}

func (x *CmdData) ProtoReflect() protoreflect.Message {
	mi := &file_room_room_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdData.ProtoReflect.Descriptor instead.
func (*CmdData) Descriptor() ([]byte, []int) {
	return file_room_room_common_proto_rawDescGZIP(), []int{1}
}

func (x *CmdData) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *CmdData) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *CmdData) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *CmdData) GetAction() int32 {
	if x != nil {
		return x.Action
	}
	return 0
}

func (x *CmdData) GetCustom() []byte {
	if x != nil {
		return x.Custom
	}
	return nil
}

// Frame data
type FrameData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// frame index
	FrameIndex uint32 `protobuf:"varint,1,opt,name=frame_index,json=frameIndex,proto3" json:"frame_index,omitempty"`
	// operation data
	Cmds []*CmdData `protobuf:"bytes,2,rep,name=cmds,proto3" json:"cmds,omitempty"`
}

func (x *FrameData) Reset() {
	*x = FrameData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_room_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrameData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrameData) ProtoMessage() {}

func (x *FrameData) ProtoReflect() protoreflect.Message {
	mi := &file_room_room_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrameData.ProtoReflect.Descriptor instead.
func (*FrameData) Descriptor() ([]byte, []int) {
	return file_room_room_common_proto_rawDescGZIP(), []int{2}
}

func (x *FrameData) GetFrameIndex() uint32 {
	if x != nil {
		return x.FrameIndex
	}
	return 0
}

func (x *FrameData) GetCmds() []*CmdData {
	if x != nil {
		return x.Cmds
	}
	return nil
}

// Player info
type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar   string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_room_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_room_room_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_room_room_common_proto_rawDescGZIP(), []int{3}
}

func (x *Player) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Player) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Player) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

var File_room_room_common_proto protoreflect.FileDescriptor

var file_room_room_common_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76,
	0x31, 0x22, 0x5b, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x67,
	0x0a, 0x07, 0x43, 0x6d, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x22, 0x52, 0x0a, 0x09, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x6d, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6d,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x63, 0x6d, 0x64, 0x73, 0x22, 0x4e, 0x0a, 0x06, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2a, 0xc3, 0x03, 0x0a, 0x0d,
	0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x12, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x02, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x52,
	0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41,
	0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x49, 0x4e, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x52,
	0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x4f, 0x4f, 0x4d, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52,
	0x45, 0x41, 0x44, 0x59, 0x10, 0x06, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x07, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x4f, 0x4f,
	0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x08, 0x12, 0x23, 0x0a, 0x1f, 0x52, 0x4f,
	0x4f, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x4c,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x09, 0x12,
	0x1d, 0x0a, 0x19, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x21,
	0x0a, 0x1d, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10,
	0x0b, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x0c, 0x12,
	0x20, 0x0a, 0x1c, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x10,
	0x0d, 0x42, 0x6a, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31,
	0x42, 0x0f, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x0d, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x72, 0x6f,
	0x6f, 0x6d, 0xa2, 0x02, 0x03, 0x52, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x52, 0x6f, 0x6f, 0x6d, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x07, 0x52, 0x6f, 0x6f, 0x6d, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x52,
	0x6f, 0x6f, 0x6d, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_room_room_common_proto_rawDescOnce sync.Once
	file_room_room_common_proto_rawDescData = file_room_room_common_proto_rawDesc
)

func file_room_room_common_proto_rawDescGZIP() []byte {
	file_room_room_common_proto_rawDescOnce.Do(func() {
		file_room_room_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_room_room_common_proto_rawDescData)
	})
	return file_room_room_common_proto_rawDescData
}

var file_room_room_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_room_room_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_room_room_common_proto_goTypes = []any{
	(RoomErrorCode)(0), // 0: room.v1.RoomErrorCode
	(*Response)(nil),   // 1: room.v1.Response
	(*CmdData)(nil),    // 2: room.v1.CmdData
	(*FrameData)(nil),  // 3: room.v1.FrameData
	(*Player)(nil),     // 4: room.v1.Player
}
var file_room_room_common_proto_depIdxs = []int32{
	0, // 0: room.v1.Response.error_code:type_name -> room.v1.RoomErrorCode
	2, // 1: room.v1.FrameData.cmds:type_name -> room.v1.CmdData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_room_room_common_proto_init() }
func file_room_room_common_proto_init() {
	if File_room_room_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_room_room_common_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Response); i {
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
		file_room_room_common_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CmdData); i {
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
		file_room_room_common_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FrameData); i {
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
		file_room_room_common_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Player); i {
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
			RawDescriptor: file_room_room_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_room_room_common_proto_goTypes,
		DependencyIndexes: file_room_room_common_proto_depIdxs,
		EnumInfos:         file_room_room_common_proto_enumTypes,
		MessageInfos:      file_room_room_common_proto_msgTypes,
	}.Build()
	File_room_room_common_proto = out.File
	file_room_room_common_proto_rawDesc = nil
	file_room_room_common_proto_goTypes = nil
	file_room_room_common_proto_depIdxs = nil
}