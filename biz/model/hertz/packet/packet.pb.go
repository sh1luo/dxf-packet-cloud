// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: packet.proto

package packet

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "packet_cloud/biz/model/api"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Code int32

const (
	Code_Success      Code = 0
	Code_ParamInvalid Code = 1
	Code_DBErr        Code = 2
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0: "Success",
		1: "ParamInvalid",
		2: "DBErr",
	}
	Code_value = map[string]int32{
		"Success":      0,
		"ParamInvalid": 1,
		"DBErr":        2,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_packet_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_packet_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{0}
}

type Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty" form:"ID" query:"ID"`
	Region   string `protobuf:"bytes,2,opt,name=Region,proto3" json:"Region,omitempty" form:"Region" query:"Region"`
	Name     string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty" form:"Name" query:"Name"`
	Content  string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty" form:"Content" query:"Content"`
	Channel  string `protobuf:"bytes,5,opt,name=Channel,proto3" json:"Channel,omitempty" form:"Channel" query:"Channel"`
	Uploader string `protobuf:"bytes,6,opt,name=Uploader,proto3" json:"Uploader,omitempty" form:"Uploader" query:"Uploader"`
	Time     string `protobuf:"bytes,7,opt,name=Time,proto3" json:"Time,omitempty" form:"Time" query:"Time"`
}

func (x *Packet) Reset() {
	*x = Packet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{0}
}

func (x *Packet) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Packet) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Packet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Packet) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Packet) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Packet) GetUploader() string {
	if x != nil {
		return x.Uploader
	}
	return ""
}

func (x *Packet) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type CreatePacketReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"id,omitempty" form:"id" vd:"$>0 && $<500"`
	Region   string `protobuf:"bytes,2,opt,name=Region,proto3" json:"region,omitempty" form:"region" vd:"(len($) > 0 && len($) < 30)"`
	Name     string `protobuf:"bytes,3,opt,name=Name,proto3" json:"name,omitempty" form:"name" vd:"(len($) > 0 && len($) < 100)"`
	Content  string `protobuf:"bytes,4,opt,name=Content,proto3" form:"content" form:"content" json:"content,omitempty" vd:"(len($) > 0 && len($) < 3000)"`
	Channel  string `protobuf:"bytes,5,opt,name=Channel,proto3" form:"channel" form:"channel" json:"channel,omitempty" vd:"(len($) > 0 && len($) < 50)"`
	Uploader string `protobuf:"bytes,6,opt,name=Uploader,proto3" json:"uploader,omitempty" form:"uploader" vd:"(len($) > 0 && len($) < 30)"`
	Time     string `protobuf:"bytes,7,opt,name=Time,proto3" json:"time,omitempty" form:"time" vd:"(len($) > 0 && len($) < 30)"`
}

func (x *CreatePacketReq) Reset() {
	*x = CreatePacketReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePacketReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePacketReq) ProtoMessage() {}

func (x *CreatePacketReq) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePacketReq.ProtoReflect.Descriptor instead.
func (*CreatePacketReq) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePacketReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CreatePacketReq) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CreatePacketReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePacketReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreatePacketReq) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *CreatePacketReq) GetUploader() string {
	if x != nil {
		return x.Uploader
	}
	return ""
}

func (x *CreatePacketReq) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type CreatePacketResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Code   `protobuf:"varint,1,opt,name=Code,proto3,enum=user.Code" json:"Code,omitempty" form:"Code" query:"Code"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty" form:"Msg" query:"Msg"`
}

func (x *CreatePacketResp) Reset() {
	*x = CreatePacketResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePacketResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePacketResp) ProtoMessage() {}

func (x *CreatePacketResp) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePacketResp.ProtoReflect.Descriptor instead.
func (*CreatePacketResp) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePacketResp) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_Success
}

func (x *CreatePacketResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type QueryPacketsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time    string `protobuf:"bytes,1,opt,name=Time,proto3" json:"time,omitempty" form:"time" vd:"len($)>0"`
	Account string `protobuf:"bytes,2,opt,name=Account,proto3" json:"account,omitempty" form:"account" vd:"len($)>0 && len($)<20"`
}

func (x *QueryPacketsReq) Reset() {
	*x = QueryPacketsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPacketsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPacketsReq) ProtoMessage() {}

func (x *QueryPacketsReq) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPacketsReq.ProtoReflect.Descriptor instead.
func (*QueryPacketsReq) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{3}
}

func (x *QueryPacketsReq) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *QueryPacketsReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type QueryPacketsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   Code      `protobuf:"varint,1,opt,name=Code,proto3,enum=user.Code" json:"Code,omitempty" form:"Code" query:"Code"`
	Msg    string    `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty" form:"Msg" query:"Msg"`
	Packet []*Packet `protobuf:"bytes,3,rep,name=Packet,proto3" json:"Packet" form:"Packet" query:"Packet"`
	Total  int64     `protobuf:"varint,4,opt,name=Total,proto3" json:"Total,omitempty" form:"Total" query:"Total"`
}

func (x *QueryPacketsResp) Reset() {
	*x = QueryPacketsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPacketsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPacketsResp) ProtoMessage() {}

func (x *QueryPacketsResp) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPacketsResp.ProtoReflect.Descriptor instead.
func (*QueryPacketsResp) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{4}
}

func (x *QueryPacketsResp) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_Success
}

func (x *QueryPacketsResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *QueryPacketsResp) GetPacket() []*Packet {
	if x != nil {
		return x.Packet
	}
	return nil
}

func (x *QueryPacketsResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_packet_proto protoreflect.FileDescriptor

var file_packet_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa8, 0x01, 0x0a, 0x06, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x90, 0x04, 0x0a, 0x0f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2c,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1c, 0xca, 0xbb, 0x18, 0x02,
	0x69, 0x64, 0xda, 0xbb, 0x18, 0x0c, 0x24, 0x3e, 0x30, 0x20, 0x26, 0x26, 0x20, 0x24, 0x3c, 0x35,
	0x30, 0x30, 0xe2, 0xbb, 0x18, 0x02, 0x69, 0x64, 0x52, 0x02, 0x49, 0x44, 0x12, 0x4b, 0x0a, 0x06,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xca, 0xbb,
	0x18, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0xda, 0xbb, 0x18, 0x1b, 0x28, 0x6c, 0x65, 0x6e,
	0x28, 0x24, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x20, 0x26, 0x26, 0x20, 0x6c, 0x65, 0x6e, 0x28, 0x24,
	0x29, 0x20, 0x3c, 0x20, 0x33, 0x30, 0x29, 0xe2, 0xbb, 0x18, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xca, 0xbb, 0x18, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0xda, 0xbb, 0x18, 0x1c, 0x28, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3e, 0x20, 0x30,
	0x20, 0x26, 0x26, 0x20, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3c, 0x20, 0x31, 0x30, 0x30,
	0x29, 0xe2, 0xbb, 0x18, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x51, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x37, 0xca, 0xbb, 0x18, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0xda, 0xbb, 0x18,
	0x1d, 0x28, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x20, 0x26, 0x26, 0x20,
	0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3c, 0x20, 0x33, 0x30, 0x30, 0x30, 0x29, 0xe2, 0xbb,
	0x18, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x4f, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x35, 0xca, 0xbb, 0x18, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0xda, 0xbb, 0x18, 0x1b, 0x28, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x20,
	0x26, 0x26, 0x20, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3c, 0x20, 0x35, 0x30, 0x29, 0xe2,
	0xbb, 0x18, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x07, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x53, 0x0a, 0x08, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x37, 0xca, 0xbb, 0x18, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0xda, 0xbb, 0x18, 0x1b, 0x28, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3e,
	0x20, 0x30, 0x20, 0x26, 0x26, 0x20, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3c, 0x20, 0x33,
	0x30, 0x29, 0xe2, 0xbb, 0x18, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x08,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xca, 0xbb, 0x18, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0xda, 0xbb, 0x18, 0x1b, 0x28, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x20,
	0x26, 0x26, 0x20, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3c, 0x20, 0x33, 0x30, 0x29, 0xe2,
	0xbb, 0x18, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x44, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1e, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4d, 0x73, 0x67, 0x22, 0x8e, 0x01, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x30, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xca, 0xbb, 0x18, 0x04, 0x74, 0x69, 0x6d, 0x65, 0xda,
	0xbb, 0x18, 0x08, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x3e, 0x30, 0xe2, 0xbb, 0x18, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xca, 0xbb, 0x18, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0xda, 0xbb, 0x18, 0x15, 0x6c, 0x65, 0x6e, 0x28, 0x24,
	0x29, 0x3e, 0x30, 0x20, 0x26, 0x26, 0x20, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x3c, 0x32, 0x30,
	0xe2, 0xbb, 0x18, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x24, 0x0a, 0x06, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x2a, 0x30, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x44, 0x42, 0x45, 0x72, 0x72, 0x10, 0x02, 0x32, 0xc9, 0x01, 0x0a, 0x0d, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x15, 0xd2, 0xc1, 0x18, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x5a, 0x0a, 0x13, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x14, 0xd2, 0xc1, 0x18, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x25, 0x5a, 0x23, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x68, 0x65, 0x72, 0x74, 0x7a, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packet_proto_rawDescOnce sync.Once
	file_packet_proto_rawDescData = file_packet_proto_rawDesc
)

func file_packet_proto_rawDescGZIP() []byte {
	file_packet_proto_rawDescOnce.Do(func() {
		file_packet_proto_rawDescData = protoimpl.X.CompressGZIP(file_packet_proto_rawDescData)
	})
	return file_packet_proto_rawDescData
}

var file_packet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_packet_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_packet_proto_goTypes = []interface{}{
	(Code)(0),                // 0: user.Code
	(*Packet)(nil),           // 1: user.Packet
	(*CreatePacketReq)(nil),  // 2: user.CreatePacketReq
	(*CreatePacketResp)(nil), // 3: user.CreatePacketResp
	(*QueryPacketsReq)(nil),  // 4: user.QueryPacketsReq
	(*QueryPacketsResp)(nil), // 5: user.QueryPacketsResp
}
var file_packet_proto_depIdxs = []int32{
	0, // 0: user.CreatePacketResp.Code:type_name -> user.Code
	0, // 1: user.QueryPacketsResp.Code:type_name -> user.Code
	1, // 2: user.QueryPacketsResp.Packet:type_name -> user.Packet
	2, // 3: user.PacketService.CreatePacketResponse:input_type -> user.CreatePacketReq
	4, // 4: user.PacketService.QueryPacketResponse:input_type -> user.QueryPacketsReq
	3, // 5: user.PacketService.CreatePacketResponse:output_type -> user.CreatePacketResp
	5, // 6: user.PacketService.QueryPacketResponse:output_type -> user.QueryPacketsResp
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_packet_proto_init() }
func file_packet_proto_init() {
	if File_packet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_packet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Packet); i {
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
		file_packet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePacketReq); i {
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
		file_packet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePacketResp); i {
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
		file_packet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPacketsReq); i {
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
		file_packet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPacketsResp); i {
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
			RawDescriptor: file_packet_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_packet_proto_goTypes,
		DependencyIndexes: file_packet_proto_depIdxs,
		EnumInfos:         file_packet_proto_enumTypes,
		MessageInfos:      file_packet_proto_msgTypes,
	}.Build()
	File_packet_proto = out.File
	file_packet_proto_rawDesc = nil
	file_packet_proto_goTypes = nil
	file_packet_proto_depIdxs = nil
}
