// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: buddy/buddy_common.proto

package buddy

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

// Nothing is used when there is no data to be sent.
type Nothing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Nothing) Reset() {
	*x = Nothing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buddy_buddy_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nothing) ProtoMessage() {}

func (x *Nothing) ProtoReflect() protoreflect.Message {
	mi := &file_buddy_buddy_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nothing.ProtoReflect.Descriptor instead.
func (*Nothing) Descriptor() ([]byte, []int) {
	return file_buddy_buddy_common_proto_rawDescGZIP(), []int{0}
}

type ProfileId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileId string `protobuf:"bytes,1,opt,name=profileId,proto3" json:"profileId,omitempty"`
}

func (x *ProfileId) Reset() {
	*x = ProfileId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buddy_buddy_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileId) ProtoMessage() {}

func (x *ProfileId) ProtoReflect() protoreflect.Message {
	mi := &file_buddy_buddy_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileId.ProtoReflect.Descriptor instead.
func (*ProfileId) Descriptor() ([]byte, []int) {
	return file_buddy_buddy_common_proto_rawDescGZIP(), []int{1}
}

func (x *ProfileId) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

// Buddy contains state associated with a buddy.
type Buddy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid           string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ReceiveReward int32  `protobuf:"varint,2,opt,name=receiveReward,proto3" json:"receiveReward,omitempty"`
	IsFavorite    bool   `protobuf:"varint,3,opt,name=isFavorite,proto3" json:"isFavorite,omitempty"`
	Remark        string `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	ActTime       int64  `protobuf:"varint,5,opt,name=actTime,proto3" json:"actTime,omitempty"`
}

func (x *Buddy) Reset() {
	*x = Buddy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buddy_buddy_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Buddy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Buddy) ProtoMessage() {}

func (x *Buddy) ProtoReflect() protoreflect.Message {
	mi := &file_buddy_buddy_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Buddy.ProtoReflect.Descriptor instead.
func (*Buddy) Descriptor() ([]byte, []int) {
	return file_buddy_buddy_common_proto_rawDescGZIP(), []int{2}
}

func (x *Buddy) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Buddy) GetReceiveReward() int32 {
	if x != nil {
		return x.ReceiveReward
	}
	return 0
}

func (x *Buddy) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Buddy) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Buddy) GetActTime() int64 {
	if x != nil {
		return x.ActTime
	}
	return 0
}

type Inviter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ReqInfo string `protobuf:"bytes,2,opt,name=reqInfo,proto3" json:"reqInfo,omitempty"`
	ReqTime int64  `protobuf:"varint,3,opt,name=reqTime,proto3" json:"reqTime,omitempty"`
}

func (x *Inviter) Reset() {
	*x = Inviter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buddy_buddy_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inviter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inviter) ProtoMessage() {}

func (x *Inviter) ProtoReflect() protoreflect.Message {
	mi := &file_buddy_buddy_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inviter.ProtoReflect.Descriptor instead.
func (*Inviter) Descriptor() ([]byte, []int) {
	return file_buddy_buddy_common_proto_rawDescGZIP(), []int{3}
}

func (x *Inviter) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Inviter) GetReqInfo() string {
	if x != nil {
		return x.ReqInfo
	}
	return ""
}

func (x *Inviter) GetReqTime() int64 {
	if x != nil {
		return x.ReqTime
	}
	return 0
}

type Blocked struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	AddTime int64  `protobuf:"varint,2,opt,name=addTime,proto3" json:"addTime,omitempty"`
}

func (x *Blocked) Reset() {
	*x = Blocked{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buddy_buddy_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blocked) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blocked) ProtoMessage() {}

func (x *Blocked) ProtoReflect() protoreflect.Message {
	mi := &file_buddy_buddy_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blocked.ProtoReflect.Descriptor instead.
func (*Blocked) Descriptor() ([]byte, []int) {
	return file_buddy_buddy_common_proto_rawDescGZIP(), []int{4}
}

func (x *Blocked) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Blocked) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

type Buddies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Buddies      map[string]*Buddy   `protobuf:"bytes,1,rep,name=buddies,proto3" json:"buddies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Inviters     map[string]*Inviter `protobuf:"bytes,2,rep,name=inviters,proto3" json:"inviters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	InviterSends map[string]*Inviter `protobuf:"bytes,3,rep,name=inviterSends,proto3" json:"inviterSends,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Blocked      map[string]*Blocked `protobuf:"bytes,4,rep,name=blocked,proto3" json:"blocked,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Buddies) Reset() {
	*x = Buddies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buddy_buddy_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Buddies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Buddies) ProtoMessage() {}

func (x *Buddies) ProtoReflect() protoreflect.Message {
	mi := &file_buddy_buddy_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Buddies.ProtoReflect.Descriptor instead.
func (*Buddies) Descriptor() ([]byte, []int) {
	return file_buddy_buddy_common_proto_rawDescGZIP(), []int{5}
}

func (x *Buddies) GetBuddies() map[string]*Buddy {
	if x != nil {
		return x.Buddies
	}
	return nil
}

func (x *Buddies) GetInviters() map[string]*Inviter {
	if x != nil {
		return x.Inviters
	}
	return nil
}

func (x *Buddies) GetInviterSends() map[string]*Inviter {
	if x != nil {
		return x.InviterSends
	}
	return nil
}

func (x *Buddies) GetBlocked() map[string]*Blocked {
	if x != nil {
		return x.Blocked
	}
	return nil
}

type ProfileIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileIds []*ProfileId `protobuf:"bytes,1,rep,name=profileIds,proto3" json:"profileIds,omitempty"`
}

func (x *ProfileIds) Reset() {
	*x = ProfileIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buddy_buddy_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileIds) ProtoMessage() {}

func (x *ProfileIds) ProtoReflect() protoreflect.Message {
	mi := &file_buddy_buddy_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileIds.ProtoReflect.Descriptor instead.
func (*ProfileIds) Descriptor() ([]byte, []int) {
	return file_buddy_buddy_common_proto_rawDescGZIP(), []int{6}
}

func (x *ProfileIds) GetProfileIds() []*ProfileId {
	if x != nil {
		return x.ProfileIds
	}
	return nil
}

var File_buddy_buddy_common_proto protoreflect.FileDescriptor

var file_buddy_buddy_common_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x75, 0x64, 0x64, 0x79, 0x2f, 0x62, 0x75, 0x64, 0x64, 0x79, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x75, 0x64, 0x64,
	0x79, 0x2e, 0x70, 0x62, 0x22, 0x09, 0x0a, 0x07, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22,
	0x29, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x05, 0x42,
	0x75, 0x64, 0x64, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x63, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4f,
	0x0a, 0x07, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x71, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x71, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x35, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xc3, 0x04, 0x0a, 0x07, 0x42, 0x75, 0x64, 0x64, 0x69,
	0x65, 0x73, 0x12, 0x38, 0x0a, 0x07, 0x62, 0x75, 0x64, 0x64, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x75, 0x64, 0x64, 0x79, 0x2e, 0x70, 0x62, 0x2e, 0x42,
	0x75, 0x64, 0x64, 0x69, 0x65, 0x73, 0x2e, 0x42, 0x75, 0x64, 0x64, 0x69, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x62, 0x75, 0x64, 0x64, 0x69, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x08,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x62, 0x75, 0x64, 0x64, 0x79, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x64, 0x64, 0x69, 0x65,
	0x73, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x73, 0x12, 0x47, 0x0a, 0x0c, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x62, 0x75, 0x64, 0x64, 0x79, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x64, 0x64, 0x69,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x53, 0x65, 0x6e,
	0x64, 0x73, 0x12, 0x38, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x75, 0x64, 0x64, 0x79, 0x2e, 0x70, 0x62, 0x2e, 0x42,
	0x75, 0x64, 0x64, 0x69, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x1a, 0x4b, 0x0a, 0x0c,
	0x42, 0x75, 0x64, 0x64, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x62, 0x75, 0x64, 0x64, 0x79, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x64, 0x64, 0x79, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4e, 0x0a, 0x0d, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x75,
	0x64, 0x64, 0x79, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x52, 0x0a, 0x11, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x62, 0x75, 0x64, 0x64, 0x79, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4d, 0x0a,
	0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x62, 0x75, 0x64, 0x64, 0x79, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x41, 0x0a, 0x0a,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x12, 0x33, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x62, 0x75, 0x64, 0x64, 0x79, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x42,
	0x11, 0x5a, 0x0f, 0x62, 0x75, 0x64, 0x64, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x62, 0x75, 0x64,
	0x64, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buddy_buddy_common_proto_rawDescOnce sync.Once
	file_buddy_buddy_common_proto_rawDescData = file_buddy_buddy_common_proto_rawDesc
)

func file_buddy_buddy_common_proto_rawDescGZIP() []byte {
	file_buddy_buddy_common_proto_rawDescOnce.Do(func() {
		file_buddy_buddy_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_buddy_buddy_common_proto_rawDescData)
	})
	return file_buddy_buddy_common_proto_rawDescData
}

var file_buddy_buddy_common_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_buddy_buddy_common_proto_goTypes = []any{
	(*Nothing)(nil),    // 0: buddy.pb.Nothing
	(*ProfileId)(nil),  // 1: buddy.pb.ProfileId
	(*Buddy)(nil),      // 2: buddy.pb.Buddy
	(*Inviter)(nil),    // 3: buddy.pb.Inviter
	(*Blocked)(nil),    // 4: buddy.pb.Blocked
	(*Buddies)(nil),    // 5: buddy.pb.Buddies
	(*ProfileIds)(nil), // 6: buddy.pb.ProfileIds
	nil,                // 7: buddy.pb.Buddies.BuddiesEntry
	nil,                // 8: buddy.pb.Buddies.InvitersEntry
	nil,                // 9: buddy.pb.Buddies.InviterSendsEntry
	nil,                // 10: buddy.pb.Buddies.BlockedEntry
}
var file_buddy_buddy_common_proto_depIdxs = []int32{
	7,  // 0: buddy.pb.Buddies.buddies:type_name -> buddy.pb.Buddies.BuddiesEntry
	8,  // 1: buddy.pb.Buddies.inviters:type_name -> buddy.pb.Buddies.InvitersEntry
	9,  // 2: buddy.pb.Buddies.inviterSends:type_name -> buddy.pb.Buddies.InviterSendsEntry
	10, // 3: buddy.pb.Buddies.blocked:type_name -> buddy.pb.Buddies.BlockedEntry
	1,  // 4: buddy.pb.ProfileIds.profileIds:type_name -> buddy.pb.ProfileId
	2,  // 5: buddy.pb.Buddies.BuddiesEntry.value:type_name -> buddy.pb.Buddy
	3,  // 6: buddy.pb.Buddies.InvitersEntry.value:type_name -> buddy.pb.Inviter
	3,  // 7: buddy.pb.Buddies.InviterSendsEntry.value:type_name -> buddy.pb.Inviter
	4,  // 8: buddy.pb.Buddies.BlockedEntry.value:type_name -> buddy.pb.Blocked
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_buddy_buddy_common_proto_init() }
func file_buddy_buddy_common_proto_init() {
	if File_buddy_buddy_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buddy_buddy_common_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Nothing); i {
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
		file_buddy_buddy_common_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ProfileId); i {
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
		file_buddy_buddy_common_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Buddy); i {
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
		file_buddy_buddy_common_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Inviter); i {
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
		file_buddy_buddy_common_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Blocked); i {
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
		file_buddy_buddy_common_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Buddies); i {
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
		file_buddy_buddy_common_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ProfileIds); i {
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
			RawDescriptor: file_buddy_buddy_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buddy_buddy_common_proto_goTypes,
		DependencyIndexes: file_buddy_buddy_common_proto_depIdxs,
		MessageInfos:      file_buddy_buddy_common_proto_msgTypes,
	}.Build()
	File_buddy_buddy_common_proto = out.File
	file_buddy_buddy_common_proto_rawDesc = nil
	file_buddy_buddy_common_proto_goTypes = nil
	file_buddy_buddy_common_proto_depIdxs = nil
}
