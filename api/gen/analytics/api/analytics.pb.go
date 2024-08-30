// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: analytics/analytics.proto

// Analytics service for sending analytics events,
// support multi delivery type: local,thinkingdata,clickhouse,mixpanel etc.
// 分析服务用于发送分析事件
// 支持多种投递方式: local,thinkingdata,clickhouse,mixpanel等

package v1

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

type DeliveryType int32

const (
	// Deliver to the local file
	DeliveryType_Local DeliveryType = 0
	// Deliver to ThinkingData
	// https://www.thinkingdata.cn/
	DeliveryType_ThinkingData DeliveryType = 1
	//Deliver to clickhouse
	DeliveryType_ClickHouse DeliveryType = 2
	// Deliver to Mixpanel
	// https://mixpanel.com/
	DeliveryType_Mixpanel DeliveryType = 3
)

// Enum value maps for DeliveryType.
var (
	DeliveryType_name = map[int32]string{
		0: "Local",
		1: "ThinkingData",
		2: "ClickHouse",
		3: "Mixpanel",
	}
	DeliveryType_value = map[string]int32{
		"Local":        0,
		"ThinkingData": 1,
		"ClickHouse":   2,
		"Mixpanel":     3,
	}
)

func (x DeliveryType) Enum() *DeliveryType {
	p := new(DeliveryType)
	*p = x
	return p
}

func (x DeliveryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeliveryType) Descriptor() protoreflect.EnumDescriptor {
	return file_analytics_analytics_proto_enumTypes[0].Descriptor()
}

func (DeliveryType) Type() protoreflect.EnumType {
	return &file_analytics_analytics_proto_enumTypes[0]
}

func (x DeliveryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeliveryType.Descriptor instead.
func (DeliveryType) EnumDescriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{0}
}

type AnalyticsEvents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *AnalyticsEvents) Reset() {
	*x = AnalyticsEvents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyticsEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyticsEvents) ProtoMessage() {}

func (x *AnalyticsEvents) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyticsEvents.ProtoReflect.Descriptor instead.
func (*AnalyticsEvents) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{0}
}

func (x *AnalyticsEvents) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

// AnalyticsEvent is a single analytics event to capture.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique name for this event.  Be pragmatic with event names and store additional properties in the
	// properties field.
	//NOTE: only contain: number,letter(ignoring case) and underscore"_" ,no spaces in the configuration
	Event string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	// Generic JSON property key/value pairs. {"id":"fun","age":10}
	Properties []byte `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
	// Where to deliver this event to, defaults to Local.
	DeliverTo DeliveryType `protobuf:"varint,3,opt,name=deliver_to,json=deliverTo,proto3,enum=analytics.v1.DeliveryType" json:"deliver_to,omitempty"`
	// user_id is the unique identifier for the user.
	// if use thinkingdata ,distinct_id /user_id is required
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// distinct_id is the unique identifier for the user/visitor.
	// if use thinkingdata ,distinct_id /user_id is required
	DistinctId string `protobuf:"bytes,5,opt,name=distinct_id,json=distinctId,proto3" json:"distinct_id,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *Event) GetProperties() []byte {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Event) GetDeliverTo() DeliveryType {
	if x != nil {
		return x.DeliverTo
	}
	return DeliveryType_Local
}

func (x *Event) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Event) GetDistinctId() string {
	if x != nil {
		return x.DistinctId
	}
	return ""
}

// Nothing is an empty message.  Used when there's nothing to send.
type Nothing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Nothing) Reset() {
	*x = Nothing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nothing) ProtoMessage() {}

func (x *Nothing) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[2]
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
	return file_analytics_analytics_proto_rawDescGZIP(), []int{2}
}

var File_analytics_analytics_proto protoreflect.FileDescriptor

var file_analytics_analytics_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0f, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x54, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x49, 0x64, 0x22, 0x09, 0x0a, 0x07,
	0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2a, 0x49, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x48, 0x6f, 0x75,
	0x73, 0x65, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x69, 0x78, 0x70, 0x61, 0x6e, 0x65, 0x6c,
	0x10, 0x03, 0x32, 0x6f, 0x0a, 0x10, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x09, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x12, 0x1d, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x1a, 0x15, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x42, 0x85, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x10, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41,
	0x58, 0x58, 0xaa, 0x02, 0x0c, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0c, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x18, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_analytics_analytics_proto_rawDescOnce sync.Once
	file_analytics_analytics_proto_rawDescData = file_analytics_analytics_proto_rawDesc
)

func file_analytics_analytics_proto_rawDescGZIP() []byte {
	file_analytics_analytics_proto_rawDescOnce.Do(func() {
		file_analytics_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_analytics_analytics_proto_rawDescData)
	})
	return file_analytics_analytics_proto_rawDescData
}

var file_analytics_analytics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_analytics_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_analytics_analytics_proto_goTypes = []any{
	(DeliveryType)(0),       // 0: analytics.v1.DeliveryType
	(*AnalyticsEvents)(nil), // 1: analytics.v1.AnalyticsEvents
	(*Event)(nil),           // 2: analytics.v1.Event
	(*Nothing)(nil),         // 3: analytics.v1.Nothing
}
var file_analytics_analytics_proto_depIdxs = []int32{
	2, // 0: analytics.v1.AnalyticsEvents.events:type_name -> analytics.v1.Event
	0, // 1: analytics.v1.Event.deliver_to:type_name -> analytics.v1.DeliveryType
	1, // 2: analytics.v1.AnalyticsService.Analytics:input_type -> analytics.v1.AnalyticsEvents
	3, // 3: analytics.v1.AnalyticsService.Analytics:output_type -> analytics.v1.Nothing
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_analytics_analytics_proto_init() }
func file_analytics_analytics_proto_init() {
	if File_analytics_analytics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_analytics_analytics_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AnalyticsEvents); i {
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
		file_analytics_analytics_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Event); i {
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
		file_analytics_analytics_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_analytics_analytics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analytics_analytics_proto_goTypes,
		DependencyIndexes: file_analytics_analytics_proto_depIdxs,
		EnumInfos:         file_analytics_analytics_proto_enumTypes,
		MessageInfos:      file_analytics_analytics_proto_msgTypes,
	}.Build()
	File_analytics_analytics_proto = out.File
	file_analytics_analytics_proto_rawDesc = nil
	file_analytics_analytics_proto_goTypes = nil
	file_analytics_analytics_proto_depIdxs = nil
}