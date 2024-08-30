// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: mail/mail.proto

// MailService is a service for mail，support GM custom mail, system template mail
// 邮件服务 用于邮件管理，支持GM自定义邮件，系统模板邮件

package mail

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type MailStatus int32

const (
	MailStatus_UNREAD   MailStatus = 0
	MailStatus_READ     MailStatus = 1
	MailStatus_REWARDED MailStatus = 2
	MailStatus_DELETED  MailStatus = 3
)

// Enum value maps for MailStatus.
var (
	MailStatus_name = map[int32]string{
		0: "UNREAD",
		1: "READ",
		2: "REWARDED",
		3: "DELETED",
	}
	MailStatus_value = map[string]int32{
		"UNREAD":   0,
		"READ":     1,
		"REWARDED": 2,
		"DELETED":  3,
	}
)

func (x MailStatus) Enum() *MailStatus {
	p := new(MailStatus)
	*p = x
	return p
}

func (x MailStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MailStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_mail_mail_proto_enumTypes[0].Descriptor()
}

func (MailStatus) Type() protoreflect.EnumType {
	return &file_mail_mail_proto_enumTypes[0]
}

func (x MailStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MailStatus.Descriptor instead.
func (MailStatus) EnumDescriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{0}
}

type SendMailRequest_SendType int32

const (
	SendMailRequest_NONE SendMailRequest_SendType = 0
	SendMailRequest_ALL  SendMailRequest_SendType = 1
	SendMailRequest_ROLE SendMailRequest_SendType = 2
)

// Enum value maps for SendMailRequest_SendType.
var (
	SendMailRequest_SendType_name = map[int32]string{
		0: "NONE",
		1: "ALL",
		2: "ROLE",
	}
	SendMailRequest_SendType_value = map[string]int32{
		"NONE": 0,
		"ALL":  1,
		"ROLE": 2,
	}
)

func (x SendMailRequest_SendType) Enum() *SendMailRequest_SendType {
	p := new(SendMailRequest_SendType)
	*p = x
	return p
}

func (x SendMailRequest_SendType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SendMailRequest_SendType) Descriptor() protoreflect.EnumDescriptor {
	return file_mail_mail_proto_enumTypes[1].Descriptor()
}

func (SendMailRequest_SendType) Type() protoreflect.EnumType {
	return &file_mail_mail_proto_enumTypes[1]
}

func (x SendMailRequest_SendType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SendMailRequest_SendType.Descriptor instead.
func (SendMailRequest_SendType) EnumDescriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{2, 0}
}

type MailReward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`   // reward id
	Num int32 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"` // reward num
	// if expire>0, id must be unique eg.timestamp
	Expire int64 `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"` // reward expire time (0 means no expire)
	Type   int32 `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`     // reward type
}

func (x *MailReward) Reset() {
	*x = MailReward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailReward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailReward) ProtoMessage() {}

func (x *MailReward) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailReward.ProtoReflect.Descriptor instead.
func (*MailReward) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{0}
}

func (x *MailReward) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MailReward) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *MailReward) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

func (x *MailReward) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type Mail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // mail uid (optional, if not set, will be generated)
	// mail title (required,key: language, value: title)
	// if template_id is set, title will be ignored
	Title map[string]string `protobuf:"bytes,2,rep,name=title,proto3" json:"title,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// mail content (required key: language, value: content),
	// if template_id is set, content will be ignored
	Body     map[string]string `protobuf:"bytes,3,rep,name=body,proto3" json:"body,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Date     int64             `protobuf:"varint,4,opt,name=date,proto3" json:"date,omitempty"`                         // mail  send time (optional, if not set, will be now)
	ExpireAt int64             `protobuf:"varint,5,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"` // mail expire time (optional, if not set, will be now+90 days)
	// mail sender (required)
	// if template_id is set, sender will be ignored
	From         string        `protobuf:"bytes,6,opt,name=from,proto3" json:"from,omitempty"`
	Rewards      []*MailReward `protobuf:"bytes,7,rep,name=rewards,proto3" json:"rewards,omitempty"`                                //mail rewards (optional)
	Status       MailStatus    `protobuf:"varint,8,opt,name=status,proto3,enum=mail.v1.MailStatus" json:"status,omitempty"`         // mail status (0:unread,1:read,2:rewarded,3:deleted)
	TemplateId   int32         `protobuf:"varint,9,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`       // mail template id (optional)
	TemplateArgs []string      `protobuf:"bytes,10,rep,name=template_args,json=templateArgs,proto3" json:"template_args,omitempty"` // mail template args (optional)
	Filters      *Mail_Filter  `protobuf:"bytes,11,opt,name=filters,proto3" json:"filters,omitempty"`                               // mail filters
}

func (x *Mail) Reset() {
	*x = Mail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mail) ProtoMessage() {}

func (x *Mail) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mail.ProtoReflect.Descriptor instead.
func (*Mail) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{1}
}

func (x *Mail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Mail) GetTitle() map[string]string {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *Mail) GetBody() map[string]string {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Mail) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *Mail) GetExpireAt() int64 {
	if x != nil {
		return x.ExpireAt
	}
	return 0
}

func (x *Mail) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Mail) GetRewards() []*MailReward {
	if x != nil {
		return x.Rewards
	}
	return nil
}

func (x *Mail) GetStatus() MailStatus {
	if x != nil {
		return x.Status
	}
	return MailStatus_UNREAD
}

func (x *Mail) GetTemplateId() int32 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

func (x *Mail) GetTemplateArgs() []string {
	if x != nil {
		return x.TemplateArgs
	}
	return nil
}

func (x *Mail) GetFilters() *Mail_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

// send mail(private)
type SendMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlatformId string                   `protobuf:"bytes,1,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`                                  // platform id (optional, if not set, will be all platforms)
	ChannelId  string                   `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`                                     // channel id (optional, if not set, will be all channels)
	SendType   SendMailRequest_SendType `protobuf:"varint,4,opt,name=send_type,json=sendType,proto3,enum=mail.v1.SendMailRequest_SendType" json:"send_type,omitempty"` // send type (0:none,1:all,2:role)
	RoleIds    []string                 `protobuf:"bytes,5,rep,name=role_ids,json=roleIds,proto3" json:"role_ids,omitempty"`                                           // role id (optional, if not set, will be all roles)
	Mail       *Mail                    `protobuf:"bytes,6,opt,name=mail,proto3" json:"mail,omitempty"`                                                                // mail content
}

func (x *SendMailRequest) Reset() {
	*x = SendMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailRequest) ProtoMessage() {}

func (x *SendMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailRequest.ProtoReflect.Descriptor instead.
func (*SendMailRequest) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{2}
}

func (x *SendMailRequest) GetPlatformId() string {
	if x != nil {
		return x.PlatformId
	}
	return ""
}

func (x *SendMailRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *SendMailRequest) GetSendType() SendMailRequest_SendType {
	if x != nil {
		return x.SendType
	}
	return SendMailRequest_NONE
}

func (x *SendMailRequest) GetRoleIds() []string {
	if x != nil {
		return x.RoleIds
	}
	return nil
}

func (x *SendMailRequest) GetMail() *Mail {
	if x != nil {
		return x.Mail
	}
	return nil
}

// send mail response
type SendMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendMailResponse) Reset() {
	*x = SendMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailResponse) ProtoMessage() {}

func (x *SendMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailResponse.ProtoReflect.Descriptor instead.
func (*SendMailResponse) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{3}
}

// watch mail changes request
type WatchMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel      string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`                                // channel
	Language     string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`                              // language
	RegisterTime int64  `protobuf:"varint,3,opt,name=register_time,json=registerTime,proto3" json:"register_time,omitempty"` // register time
}

func (x *WatchMailRequest) Reset() {
	*x = WatchMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchMailRequest) ProtoMessage() {}

func (x *WatchMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchMailRequest.ProtoReflect.Descriptor instead.
func (*WatchMailRequest) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{4}
}

func (x *WatchMailRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *WatchMailRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *WatchMailRequest) GetRegisterTime() int64 {
	if x != nil {
		return x.RegisterTime
	}
	return 0
}

// watch mail changes response
type WatchMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mails map[int64]*Mail `protobuf:"bytes,1,rep,name=mails,proto3" json:"mails,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // mail changes
}

func (x *WatchMailResponse) Reset() {
	*x = WatchMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchMailResponse) ProtoMessage() {}

func (x *WatchMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchMailResponse.ProtoReflect.Descriptor instead.
func (*WatchMailResponse) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{5}
}

func (x *WatchMailResponse) GetMails() map[int64]*Mail {
	if x != nil {
		return x.Mails
	}
	return nil
}

// update mail status
type UpdateMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updates map[int64]MailStatus `protobuf:"bytes,1,rep,name=updates,proto3" json:"updates,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=mail.v1.MailStatus"` //  <mailId,mailStatus> mailId=0 means update all mails
}

func (x *UpdateMailRequest) Reset() {
	*x = UpdateMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMailRequest) ProtoMessage() {}

func (x *UpdateMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMailRequest.ProtoReflect.Descriptor instead.
func (*UpdateMailRequest) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateMailRequest) GetUpdates() map[int64]MailStatus {
	if x != nil {
		return x.Updates
	}
	return nil
}

// update mail status response
type UpdateMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rewards map[int64]*MailReward `protobuf:"bytes,1,rep,name=rewards,proto3" json:"rewards,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // if update mail status to REWARDED, rewards will be returned
}

func (x *UpdateMailResponse) Reset() {
	*x = UpdateMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMailResponse) ProtoMessage() {}

func (x *UpdateMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMailResponse.ProtoReflect.Descriptor instead.
func (*UpdateMailResponse) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateMailResponse) GetRewards() map[int64]*MailReward {
	if x != nil {
		return x.Rewards
	}
	return nil
}

type Mail_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// role register time (optional, if not set, will be now, if -1 means all register time)
	RegisterTime int64 `protobuf:"varint,1,opt,name=register_time,json=registerTime,proto3" json:"register_time,omitempty"`
}

func (x *Mail_Filter) Reset() {
	*x = Mail_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mail_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mail_Filter) ProtoMessage() {}

func (x *Mail_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mail_Filter.ProtoReflect.Descriptor instead.
func (*Mail_Filter) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Mail_Filter) GetRegisterTime() int64 {
	if x != nil {
		return x.RegisterTime
	}
	return 0
}

var File_mail_mail_proto protoreflect.FileDescriptor

var file_mail_mail_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x0a, 0x4d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0xac, 0x04, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2b, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x2d, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x2b,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x41, 0x72, 0x67,
	0x73, 0x12, 0x2e, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69,
	0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x1a, 0x2d, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x1a, 0x38, 0x0a, 0x0a, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a, 0x09, 0x42, 0x6f,
	0x64, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xf8, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x73,
	0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x49,
	0x64, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x52,
	0x04, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x27, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41,
	0x4c, 0x4c, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x4f, 0x4c, 0x45, 0x10, 0x02, 0x22, 0x12,
	0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x6d, 0x0a, 0x10, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x99, 0x01, 0x0a, 0x11, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x6d, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6d,
	0x61, 0x69, 0x6c, 0x73, 0x1a, 0x47, 0x0a, 0x0a, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x69, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa7, 0x01,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x4f, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa9, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42,
	0x0a, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x1a, 0x4f, 0x0a, 0x0c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x2a, 0x3d, 0x0a, 0x0a, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x52, 0x45, 0x41, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x57, 0x41, 0x52,
	0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44,
	0x10, 0x03, 0x32, 0xb2, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x12, 0x61, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x69, 0x6c,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x32, 0x57, 0x0a, 0x12, 0x4d, 0x61, 0x69, 0x6c, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a,
	0x08, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x64, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x42,
	0x09, 0x4d, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0d, 0x6d, 0x61,
	0x69, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x6d, 0x61, 0x69, 0x6c, 0xa2, 0x02, 0x03, 0x4d, 0x58,
	0x58, 0xaa, 0x02, 0x07, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x4d, 0x61,
	0x69, 0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x4d, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x4d, 0x61,
	0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mail_mail_proto_rawDescOnce sync.Once
	file_mail_mail_proto_rawDescData = file_mail_mail_proto_rawDesc
)

func file_mail_mail_proto_rawDescGZIP() []byte {
	file_mail_mail_proto_rawDescOnce.Do(func() {
		file_mail_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_mail_mail_proto_rawDescData)
	})
	return file_mail_mail_proto_rawDescData
}

var file_mail_mail_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mail_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_mail_mail_proto_goTypes = []any{
	(MailStatus)(0),               // 0: mail.v1.MailStatus
	(SendMailRequest_SendType)(0), // 1: mail.v1.SendMailRequest.SendType
	(*MailReward)(nil),            // 2: mail.v1.MailReward
	(*Mail)(nil),                  // 3: mail.v1.Mail
	(*SendMailRequest)(nil),       // 4: mail.v1.SendMailRequest
	(*SendMailResponse)(nil),      // 5: mail.v1.SendMailResponse
	(*WatchMailRequest)(nil),      // 6: mail.v1.WatchMailRequest
	(*WatchMailResponse)(nil),     // 7: mail.v1.WatchMailResponse
	(*UpdateMailRequest)(nil),     // 8: mail.v1.UpdateMailRequest
	(*UpdateMailResponse)(nil),    // 9: mail.v1.UpdateMailResponse
	(*Mail_Filter)(nil),           // 10: mail.v1.Mail.Filter
	nil,                           // 11: mail.v1.Mail.TitleEntry
	nil,                           // 12: mail.v1.Mail.BodyEntry
	nil,                           // 13: mail.v1.WatchMailResponse.MailsEntry
	nil,                           // 14: mail.v1.UpdateMailRequest.UpdatesEntry
	nil,                           // 15: mail.v1.UpdateMailResponse.RewardsEntry
}
var file_mail_mail_proto_depIdxs = []int32{
	11, // 0: mail.v1.Mail.title:type_name -> mail.v1.Mail.TitleEntry
	12, // 1: mail.v1.Mail.body:type_name -> mail.v1.Mail.BodyEntry
	2,  // 2: mail.v1.Mail.rewards:type_name -> mail.v1.MailReward
	0,  // 3: mail.v1.Mail.status:type_name -> mail.v1.MailStatus
	10, // 4: mail.v1.Mail.filters:type_name -> mail.v1.Mail.Filter
	1,  // 5: mail.v1.SendMailRequest.send_type:type_name -> mail.v1.SendMailRequest.SendType
	3,  // 6: mail.v1.SendMailRequest.mail:type_name -> mail.v1.Mail
	13, // 7: mail.v1.WatchMailResponse.mails:type_name -> mail.v1.WatchMailResponse.MailsEntry
	14, // 8: mail.v1.UpdateMailRequest.updates:type_name -> mail.v1.UpdateMailRequest.UpdatesEntry
	15, // 9: mail.v1.UpdateMailResponse.rewards:type_name -> mail.v1.UpdateMailResponse.RewardsEntry
	3,  // 10: mail.v1.WatchMailResponse.MailsEntry.value:type_name -> mail.v1.Mail
	0,  // 11: mail.v1.UpdateMailRequest.UpdatesEntry.value:type_name -> mail.v1.MailStatus
	2,  // 12: mail.v1.UpdateMailResponse.RewardsEntry.value:type_name -> mail.v1.MailReward
	6,  // 13: mail.v1.MailService.Watch:input_type -> mail.v1.WatchMailRequest
	8,  // 14: mail.v1.MailService.UpdateMail:input_type -> mail.v1.UpdateMailRequest
	4,  // 15: mail.v1.MailPrivateService.SendMail:input_type -> mail.v1.SendMailRequest
	7,  // 16: mail.v1.MailService.Watch:output_type -> mail.v1.WatchMailResponse
	9,  // 17: mail.v1.MailService.UpdateMail:output_type -> mail.v1.UpdateMailResponse
	5,  // 18: mail.v1.MailPrivateService.SendMail:output_type -> mail.v1.SendMailResponse
	16, // [16:19] is the sub-list for method output_type
	13, // [13:16] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_mail_mail_proto_init() }
func file_mail_mail_proto_init() {
	if File_mail_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mail_mail_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MailReward); i {
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
		file_mail_mail_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Mail); i {
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
		file_mail_mail_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SendMailRequest); i {
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
		file_mail_mail_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SendMailResponse); i {
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
		file_mail_mail_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*WatchMailRequest); i {
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
		file_mail_mail_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*WatchMailResponse); i {
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
		file_mail_mail_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateMailRequest); i {
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
		file_mail_mail_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateMailResponse); i {
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
		file_mail_mail_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Mail_Filter); i {
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
			RawDescriptor: file_mail_mail_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_mail_mail_proto_goTypes,
		DependencyIndexes: file_mail_mail_proto_depIdxs,
		EnumInfos:         file_mail_mail_proto_enumTypes,
		MessageInfos:      file_mail_mail_proto_msgTypes,
	}.Build()
	File_mail_mail_proto = out.File
	file_mail_mail_proto_rawDesc = nil
	file_mail_mail_proto_goTypes = nil
	file_mail_mail_proto_depIdxs = nil
}