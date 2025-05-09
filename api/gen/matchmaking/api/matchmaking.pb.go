// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: matchmaking/matchmaking.proto

// MatchService is a service for match，need custom match function
// 匹配服务，用于匹配玩家，需要自定义匹配函数

package matchmaking

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

type MatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	GameId string `protobuf:"bytes,2,opt,name=gameId,proto3" json:"gameId,omitempty"`
}

func (x *MatchRequest) Reset() {
	*x = MatchRequest{}
	mi := &file_matchmaking_matchmaking_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchRequest) ProtoMessage() {}

func (x *MatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_matchmaking_matchmaking_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchRequest.ProtoReflect.Descriptor instead.
func (*MatchRequest) Descriptor() ([]byte, []int) {
	return file_matchmaking_matchmaking_proto_rawDescGZIP(), []int{0}
}

func (x *MatchRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *MatchRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

type MatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId string    `protobuf:"bytes,1,opt,name=matchId,proto3" json:"matchId,omitempty"`
	UserIds []*Ticket `protobuf:"bytes,2,rep,name=userIds,proto3" json:"userIds,omitempty"`
	GameId  string    `protobuf:"bytes,3,opt,name=gameId,proto3" json:"gameId,omitempty"`
}

func (x *MatchResponse) Reset() {
	*x = MatchResponse{}
	mi := &file_matchmaking_matchmaking_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchResponse) ProtoMessage() {}

func (x *MatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_matchmaking_matchmaking_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchResponse.ProtoReflect.Descriptor instead.
func (*MatchResponse) Descriptor() ([]byte, []int) {
	return file_matchmaking_matchmaking_proto_rawDescGZIP(), []int{1}
}

func (x *MatchResponse) GetMatchId() string {
	if x != nil {
		return x.MatchId
	}
	return ""
}

func (x *MatchResponse) GetUserIds() []*Ticket {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *MatchResponse) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId string `protobuf:"bytes,1,opt,name=ticketId,proto3" json:"ticketId,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	GameId   string `protobuf:"bytes,3,opt,name=gameId,proto3" json:"gameId,omitempty"`
	MatchId  string `protobuf:"bytes,4,opt,name=matchId,proto3" json:"matchId,omitempty"`
	Status   string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	mi := &file_matchmaking_matchmaking_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_matchmaking_matchmaking_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_matchmaking_matchmaking_proto_rawDescGZIP(), []int{2}
}

func (x *Ticket) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

func (x *Ticket) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Ticket) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *Ticket) GetMatchId() string {
	if x != nil {
		return x.MatchId
	}
	return ""
}

func (x *Ticket) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_matchmaking_matchmaking_proto protoreflect.FileDescriptor

var file_matchmaking_matchmaking_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x22,
	0x3e, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x22,
	0x73, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x49, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x56, 0x0a,
	0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a,
	0x05, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x9c, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x1b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x3b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0xa2, 0x02, 0x03,
	0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_matchmaking_matchmaking_proto_rawDescOnce sync.Once
	file_matchmaking_matchmaking_proto_rawDescData = file_matchmaking_matchmaking_proto_rawDesc
)

func file_matchmaking_matchmaking_proto_rawDescGZIP() []byte {
	file_matchmaking_matchmaking_proto_rawDescOnce.Do(func() {
		file_matchmaking_matchmaking_proto_rawDescData = protoimpl.X.CompressGZIP(file_matchmaking_matchmaking_proto_rawDescData)
	})
	return file_matchmaking_matchmaking_proto_rawDescData
}

var file_matchmaking_matchmaking_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_matchmaking_matchmaking_proto_goTypes = []any{
	(*MatchRequest)(nil),  // 0: matchmaking.v1.MatchRequest
	(*MatchResponse)(nil), // 1: matchmaking.v1.MatchResponse
	(*Ticket)(nil),        // 2: matchmaking.v1.Ticket
}
var file_matchmaking_matchmaking_proto_depIdxs = []int32{
	2, // 0: matchmaking.v1.MatchResponse.userIds:type_name -> matchmaking.v1.Ticket
	0, // 1: matchmaking.v1.MatchService.Match:input_type -> matchmaking.v1.MatchRequest
	1, // 2: matchmaking.v1.MatchService.Match:output_type -> matchmaking.v1.MatchResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_matchmaking_matchmaking_proto_init() }
func file_matchmaking_matchmaking_proto_init() {
	if File_matchmaking_matchmaking_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_matchmaking_matchmaking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_matchmaking_matchmaking_proto_goTypes,
		DependencyIndexes: file_matchmaking_matchmaking_proto_depIdxs,
		MessageInfos:      file_matchmaking_matchmaking_proto_msgTypes,
	}.Build()
	File_matchmaking_matchmaking_proto = out.File
	file_matchmaking_matchmaking_proto_rawDesc = nil
	file_matchmaking_matchmaking_proto_goTypes = nil
	file_matchmaking_matchmaking_proto_depIdxs = nil
}
