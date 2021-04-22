//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pb/event/event.proto

package event

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/infraboard/protoc-gen-go-ext/extension/tag"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type Level int32

const (
	Level_Trace    Level = 0
	Level_Debug    Level = 1
	Level_Info     Level = 2
	Level_Warn     Level = 3
	Level_Error    Level = 4
	Level_Critical Level = 5
	Level_Disaster Level = 6
)

// Enum value maps for Level.
var (
	Level_name = map[int32]string{
		0: "Trace",
		1: "Debug",
		2: "Info",
		3: "Warn",
		4: "Error",
		5: "Critical",
		6: "Disaster",
	}
	Level_value = map[string]int32{
		"Trace":    0,
		"Debug":    1,
		"Info":     2,
		"Warn":     3,
		"Error":    4,
		"Critical": 5,
		"Disaster": 6,
	}
)

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Level) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_event_event_proto_enumTypes[0].Descriptor()
}

func (Level) Type() protoreflect.EnumType {
	return &file_pb_event_event_proto_enumTypes[0]
}

func (x Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Level.Descriptor instead.
func (Level) EnumDescriptor() ([]byte, []int) {
	return file_pb_event_event_proto_rawDescGZIP(), []int{0}
}

type Type int32

const (
	// 全局配置, 所有服务可以读取
	Type_Operate Type = 0
	// 组内配置, 组里面的服务可以读取
	Type_Status Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "Operate",
		1: "Status",
	}
	Type_value = map[string]int32{
		"Operate": 0,
		"Status":  1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_event_event_proto_enumTypes[1].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_pb_event_event_proto_enumTypes[1]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_pb_event_event_proto_rawDescGZIP(), []int{1}
}

// Event to be used by controllers.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 事件头
	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header" bson:"header"`
	// data 具体的数据
	Body *anypb.Any `protobuf:"bytes,2,opt,name=body,proto3" json:"body" bson:"-"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_event_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_pb_event_event_proto_msgTypes[0]
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
	return file_pb_event_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Event) GetBody() *anypb.Any {
	if x != nil {
		return x.Body
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 事件ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 事件发送时间
	Time int64 `protobuf:"varint,2,opt,name=time,proto3" json:"time" bson:"time"`
	// 事件类型
	Type Type `protobuf:"varint,3,opt,name=type,proto3,enum=mcube.event.Type" json:"type" bson:"type"`
	// 事件来源
	Source string `protobuf:"bytes,4,opt,name=source,proto3" json:"source" bson:"source"`
	// 事件等级
	Level Level `protobuf:"varint,5,opt,name=level,proto3,enum=mcube.event.Level" json:"level" bson:"level"`
	// 数据metas
	Meta map[string]string `protobuf:"bytes,6,rep,name=meta,proto3" json:"meta" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"meta"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_event_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_pb_event_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_pb_event_event_proto_rawDescGZIP(), []int{1}
}

func (x *Header) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Header) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Header) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_Operate
}

func (x *Header) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Header) GetLevel() Level {
	if x != nil {
		return x.Level
	}
	return Level_Trace
}

func (x *Header) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

// OperateEvent 事件具体数据
type OperateEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 回话ID
	Session string `protobuf:"bytes,1,opt,name=session,proto3" json:"session" bson:"session"`
	// 操作人
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account" bson:"account"`
	// 请求ID
	RequestId string `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id" bson:"request_id"`
	// 操作者IP
	IpAddress string `protobuf:"bytes,4,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address" bson:"ip_address"`
	// 用户UA
	UserAgent string `protobuf:"bytes,5,opt,name=user_agent,json=userAgent,proto3" json:"user_agent" bson:"user_agent"`
	// 用户名称
	UserName string `protobuf:"bytes,6,opt,name=user_name,json=userName,proto3" json:"user_name" bson:"user_name"`
	// 用户类型
	UserType string `protobuf:"bytes,7,opt,name=user_type,json=userType,proto3" json:"user_type" bson:"user_type"`
	// 操作的域
	UserDomain string `protobuf:"bytes,8,opt,name=user_domain,json=userDomain,proto3" json:"user_domain" bson:"user_domain"`
	// 服务名称
	ServiceName string `protobuf:"bytes,9,opt,name=service_name,json=serviceName,proto3" json:"service_name" bson:"service_name"`
	// 功能路径
	FeaturePath string `protobuf:"bytes,10,opt,name=feature_path,json=featurePath,proto3" json:"feature_path" bson:"feature_path"`
	// 资源类型
	ResourceType string `protobuf:"bytes,11,opt,name=resource_type,json=resourceType,proto3" json:"resource_type" bson:"resource_type"`
	// 操作动作
	Action string `protobuf:"bytes,12,opt,name=action,proto3" json:"action" bson:"action"`
	// 事件数据
	Request *anypb.Any `protobuf:"bytes,13,opt,name=request,proto3" json:"request,omitempty" bson:"request"`
	// 事件数据
	Response *anypb.Any `protobuf:"bytes,14,opt,name=response,proto3" json:"response,omitempty" bson:"response"`
}

func (x *OperateEventData) Reset() {
	*x = OperateEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_event_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateEventData) ProtoMessage() {}

func (x *OperateEventData) ProtoReflect() protoreflect.Message {
	mi := &file_pb_event_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateEventData.ProtoReflect.Descriptor instead.
func (*OperateEventData) Descriptor() ([]byte, []int) {
	return file_pb_event_event_proto_rawDescGZIP(), []int{2}
}

func (x *OperateEventData) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

func (x *OperateEventData) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *OperateEventData) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *OperateEventData) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *OperateEventData) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *OperateEventData) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *OperateEventData) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

func (x *OperateEventData) GetUserDomain() string {
	if x != nil {
		return x.UserDomain
	}
	return ""
}

func (x *OperateEventData) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *OperateEventData) GetFeaturePath() string {
	if x != nil {
		return x.FeaturePath
	}
	return ""
}

func (x *OperateEventData) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *OperateEventData) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *OperateEventData) GetRequest() *anypb.Any {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *OperateEventData) GetResponse() *anypb.Any {
	if x != nil {
		return x.Response
	}
	return nil
}

type OperateEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 事件头
	Header *Header           `protobuf:"bytes,1,opt,name=header,proto3" json:"header" bson:"header"`
	Body   *OperateEventData `protobuf:"bytes,2,opt,name=body,proto3" json:"body" bson:"body"`
}

func (x *OperateEvent) Reset() {
	*x = OperateEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_event_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateEvent) ProtoMessage() {}

func (x *OperateEvent) ProtoReflect() protoreflect.Message {
	mi := &file_pb_event_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateEvent.ProtoReflect.Descriptor instead.
func (*OperateEvent) Descriptor() ([]byte, []int) {
	return file_pb_event_event_proto_rawDescGZIP(), []int{3}
}

func (x *OperateEvent) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *OperateEvent) GetBody() *OperateEventData {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_pb_event_event_proto protoreflect.FileDescriptor

var file_pb_event_event_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9d, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x63, 0x75, 0x62,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x21,
	0xc2, 0xde, 0x1f, 0x1d, 0x0a, 0x1b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x22, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x1a, 0xc2,
	0xde, 0x1f, 0x16, 0x0a, 0x14, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2d, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22,
	0xbe, 0x03, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xc2, 0xde, 0x1f, 0x16, 0x0a, 0x14, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69,
	0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x1d, 0xc2, 0xde, 0x1f, 0x19,
	0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x39, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x21, 0xc2, 0xde, 0x1f, 0x1d, 0x0a, 0x1b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x22, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x63, 0x75, 0x62,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x1f, 0xc2,
	0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x52, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x50, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x6d, 0x65, 0x74, 0x61, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x65, 0x74, 0x61,
	0x22, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xbf, 0x08, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xc2, 0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x07, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xc2, 0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x48, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xc2, 0xde, 0x1f, 0x25, 0x0a, 0x23, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x22,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x22, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x48, 0x0a,
	0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x29, 0xc2, 0xde, 0x1f, 0x25, 0x0a, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69,
	0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x52, 0x09, 0x69, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x48, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xc2, 0xde, 0x1f,
	0x25, 0x0a, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x44, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23,
	0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4c, 0x0a,
	0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x2b, 0xc2, 0xde, 0x1f, 0x27, 0x0a, 0x25, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x52,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x50, 0x0a, 0x0c, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x2d, 0xc2, 0xde, 0x1f, 0x29, 0x0a, 0x27, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a,
	0x0c, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2d, 0xc2, 0xde, 0x1f, 0x29, 0x0a, 0x27, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x22, 0x52, 0x0b, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x54, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xc2, 0xde, 0x1f, 0x2b, 0x0a, 0x29, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xc2, 0xde, 0x1f, 0x1d, 0x0a, 0x1b, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x5d, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x2d, 0xc2, 0xde, 0x1f, 0x29, 0x0a, 0x27, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2c, 0x6f, 0x6d, 0x69, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x61, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x2f, 0xc2, 0xde, 0x1f, 0x2b, 0x0a, 0x29, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2c, 0x6f, 0x6d,
	0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x21, 0xc2, 0xde, 0x1f, 0x1d, 0x0a, 0x1b,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x50, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x62, 0x6f,
	0x64, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x2a, 0x58, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x09,
	0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x65, 0x62,
	0x75, 0x67, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x57, 0x61, 0x72, 0x6e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x10,
	0x05, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x61, 0x73, 0x74, 0x65, 0x72, 0x10, 0x06, 0x2a,
	0x1f, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0x01,
	0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f,
	0x62, 0x75, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_event_event_proto_rawDescOnce sync.Once
	file_pb_event_event_proto_rawDescData = file_pb_event_event_proto_rawDesc
)

func file_pb_event_event_proto_rawDescGZIP() []byte {
	file_pb_event_event_proto_rawDescOnce.Do(func() {
		file_pb_event_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_event_event_proto_rawDescData)
	})
	return file_pb_event_event_proto_rawDescData
}

var file_pb_event_event_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pb_event_event_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_event_event_proto_goTypes = []interface{}{
	(Level)(0),               // 0: mcube.event.Level
	(Type)(0),                // 1: mcube.event.Type
	(*Event)(nil),            // 2: mcube.event.Event
	(*Header)(nil),           // 3: mcube.event.Header
	(*OperateEventData)(nil), // 4: mcube.event.OperateEventData
	(*OperateEvent)(nil),     // 5: mcube.event.OperateEvent
	nil,                      // 6: mcube.event.Header.MetaEntry
	(*anypb.Any)(nil),        // 7: google.protobuf.Any
}
var file_pb_event_event_proto_depIdxs = []int32{
	3, // 0: mcube.event.Event.header:type_name -> mcube.event.Header
	7, // 1: mcube.event.Event.body:type_name -> google.protobuf.Any
	1, // 2: mcube.event.Header.type:type_name -> mcube.event.Type
	0, // 3: mcube.event.Header.level:type_name -> mcube.event.Level
	6, // 4: mcube.event.Header.meta:type_name -> mcube.event.Header.MetaEntry
	7, // 5: mcube.event.OperateEventData.request:type_name -> google.protobuf.Any
	7, // 6: mcube.event.OperateEventData.response:type_name -> google.protobuf.Any
	3, // 7: mcube.event.OperateEvent.header:type_name -> mcube.event.Header
	4, // 8: mcube.event.OperateEvent.body:type_name -> mcube.event.OperateEventData
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_pb_event_event_proto_init() }
func file_pb_event_event_proto_init() {
	if File_pb_event_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_event_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pb_event_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_pb_event_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateEventData); i {
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
		file_pb_event_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateEvent); i {
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
			RawDescriptor: file_pb_event_event_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_event_event_proto_goTypes,
		DependencyIndexes: file_pb_event_event_proto_depIdxs,
		EnumInfos:         file_pb_event_event_proto_enumTypes,
		MessageInfos:      file_pb_event_event_proto_msgTypes,
	}.Build()
	File_pb_event_event_proto = out.File
	file_pb_event_event_proto_rawDesc = nil
	file_pb_event_event_proto_goTypes = nil
	file_pb_event_event_proto_depIdxs = nil
}
