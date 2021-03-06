// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	ReqHeartbeat
	ResHeartbeat
	Chat
	HistroyChat
	ReqLogin
	ResLogin
*/
package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MSGID int32

const (
	MSGID_Undefine_ MSGID = 0
	// 心跳
	MSGID_ReqHeartbeatE   MSGID = 1001
	MSGID_ResHeartbeatE   MSGID = 1002
	MSGID_ReqLoginE       MSGID = 1003
	MSGID_ResLoginE       MSGID = 1004
	MSGID_ReqChatE        MSGID = 1005
	MSGID_ResChatE        MSGID = 1006
	MSGID_ResHistoryChatE MSGID = 1007
	MSGID_MAX_COUNT       MSGID = 65536
)

var MSGID_name = map[int32]string{
	0:     "Undefine_",
	1001:  "ReqHeartbeatE",
	1002:  "ResHeartbeatE",
	1003:  "ReqLoginE",
	1004:  "ResLoginE",
	1005:  "ReqChatE",
	1006:  "ResChatE",
	1007:  "ResHistoryChatE",
	65536: "MAX_COUNT",
}
var MSGID_value = map[string]int32{
	"Undefine_":       0,
	"ReqHeartbeatE":   1001,
	"ResHeartbeatE":   1002,
	"ReqLoginE":       1003,
	"ResLoginE":       1004,
	"ReqChatE":        1005,
	"ResChatE":        1006,
	"ResHistoryChatE": 1007,
	"MAX_COUNT":       65536,
}

func (x MSGID) String() string {
	return proto.EnumName(MSGID_name, int32(x))
}
func (MSGID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ResLogin_Result int32

const (
	ResLogin_unknown ResLogin_Result = 0
	ResLogin_success ResLogin_Result = 1
	ResLogin_fail    ResLogin_Result = 2
)

var ResLogin_Result_name = map[int32]string{
	0: "unknown",
	1: "success",
	2: "fail",
}
var ResLogin_Result_value = map[string]int32{
	"unknown": 0,
	"success": 1,
	"fail":    2,
}

func (x ResLogin_Result) String() string {
	return proto.EnumName(ResLogin_Result_name, int32(x))
}
func (ResLogin_Result) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type ReqHeartbeat struct {
	Uid int64 `protobuf:"varint,1,opt,name=uid" json:"uid"`
}

func (m *ReqHeartbeat) Reset()                    { *m = ReqHeartbeat{} }
func (m *ReqHeartbeat) String() string            { return proto.CompactTextString(m) }
func (*ReqHeartbeat) ProtoMessage()               {}
func (*ReqHeartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReqHeartbeat) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type ResHeartbeat struct {
	Uid            int64 `protobuf:"varint,1,opt,name=uid" json:"uid"`
	ServerUnixTime int64 `protobuf:"varint,2,opt,name=serverUnixTime" json:"serverUnixTime"`
}

func (m *ResHeartbeat) Reset()                    { *m = ResHeartbeat{} }
func (m *ResHeartbeat) String() string            { return proto.CompactTextString(m) }
func (*ResHeartbeat) ProtoMessage()               {}
func (*ResHeartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResHeartbeat) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *ResHeartbeat) GetServerUnixTime() int64 {
	if m != nil {
		return m.ServerUnixTime
	}
	return 0
}

type Chat struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name"`
	Content string `protobuf:"bytes,2,opt,name=content" json:"content"`
}

func (m *Chat) Reset()                    { *m = Chat{} }
func (m *Chat) String() string            { return proto.CompactTextString(m) }
func (*Chat) ProtoMessage()               {}
func (*Chat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Chat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Chat) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type HistroyChat struct {
	History []*Chat `protobuf:"bytes,1,rep,name=history" json:"history"`
}

func (m *HistroyChat) Reset()                    { *m = HistroyChat{} }
func (m *HistroyChat) String() string            { return proto.CompactTextString(m) }
func (*HistroyChat) ProtoMessage()               {}
func (*HistroyChat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HistroyChat) GetHistory() []*Chat {
	if m != nil {
		return m.History
	}
	return nil
}

type ReqLogin struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name"`
	RoomId int32  `protobuf:"varint,2,opt,name=roomId" json:"roomId"`
	UserId int64  `protobuf:"varint,3,opt,name=userId" json:"userId"`
}

func (m *ReqLogin) Reset()                    { *m = ReqLogin{} }
func (m *ReqLogin) String() string            { return proto.CompactTextString(m) }
func (*ReqLogin) ProtoMessage()               {}
func (*ReqLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReqLogin) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqLogin) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *ReqLogin) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ResLogin struct {
	Result ResLogin_Result `protobuf:"varint,1,opt,name=result,enum=Message.ResLogin_Result" json:"result"`
}

func (m *ResLogin) Reset()                    { *m = ResLogin{} }
func (m *ResLogin) String() string            { return proto.CompactTextString(m) }
func (*ResLogin) ProtoMessage()               {}
func (*ResLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ResLogin) GetResult() ResLogin_Result {
	if m != nil {
		return m.Result
	}
	return ResLogin_unknown
}

func init() {
	proto.RegisterType((*ReqHeartbeat)(nil), "Message.ReqHeartbeat")
	proto.RegisterType((*ResHeartbeat)(nil), "Message.ResHeartbeat")
	proto.RegisterType((*Chat)(nil), "Message.Chat")
	proto.RegisterType((*HistroyChat)(nil), "Message.HistroyChat")
	proto.RegisterType((*ReqLogin)(nil), "Message.ReqLogin")
	proto.RegisterType((*ResLogin)(nil), "Message.ResLogin")
	proto.RegisterEnum("Message.MSGID", MSGID_name, MSGID_value)
	proto.RegisterEnum("Message.ResLogin_Result", ResLogin_Result_name, ResLogin_Result_value)
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x97, 0x25, 0x8b, 0x97, 0x53, 0xd2, 0x59, 0x16, 0x42, 0xb9, 0xac, 0x72, 0x01, 0x13,
	0x42, 0x15, 0x1a, 0x88, 0x7b, 0x18, 0x15, 0xad, 0x44, 0x87, 0x64, 0x56, 0x09, 0x71, 0x33, 0x65,
	0xcd, 0xd9, 0x66, 0x68, 0x6c, 0x66, 0x3b, 0x40, 0xef, 0xfa, 0x2e, 0x3c, 0x1d, 0xff, 0x5f, 0x01,
	0xd9, 0x71, 0x50, 0x85, 0xd0, 0xee, 0xce, 0xef, 0x3b, 0x9f, 0x3f, 0xe7, 0x9c, 0x18, 0xf2, 0x06,
	0x8d, 0xa9, 0x2e, 0x71, 0xfc, 0x41, 0x2b, 0xab, 0x18, 0x99, 0x77, 0x58, 0x8e, 0xe0, 0x16, 0xc7,
	0xeb, 0x29, 0x56, 0xda, 0x9e, 0x63, 0x65, 0x19, 0x85, 0xb8, 0x15, 0x75, 0x11, 0x8d, 0xa2, 0xc3,
	0x98, 0xbb, 0xb2, 0x9c, 0x3a, 0x87, 0xb9, 0xc1, 0xc1, 0xee, 0xc2, 0xd0, 0xa0, 0xfe, 0x88, 0x7a,
	0x21, 0xc5, 0xe7, 0x53, 0xd1, 0x60, 0xb1, 0xeb, 0x9b, 0xff, 0xa8, 0xe5, 0x63, 0x48, 0x8e, 0xaf,
	0x2a, 0xcb, 0x18, 0x24, 0xb2, 0x6a, 0xd0, 0x47, 0x64, 0xdc, 0xd7, 0xac, 0x00, 0xb2, 0x54, 0xd2,
	0xa2, 0xb4, 0xfe, 0x70, 0xc6, 0x7b, 0x2c, 0x9f, 0xc0, 0x60, 0x2a, 0x8c, 0xd5, 0x6a, 0xed, 0x0f,
	0xdf, 0x03, 0x72, 0x25, 0x8c, 0x55, 0x7a, 0x5d, 0x44, 0xa3, 0xf8, 0x70, 0x70, 0x94, 0x8f, 0xc3,
	0x2c, 0x63, 0xd7, 0xe7, 0x7d, 0xb7, 0x3c, 0x81, 0x7d, 0x8e, 0xd7, 0x2f, 0xd5, 0xa5, 0x90, 0xff,
	0xbd, 0xf1, 0x0e, 0xa4, 0x5a, 0xa9, 0x66, 0x56, 0xfb, 0x0b, 0xf7, 0x78, 0x20, 0xa7, 0xb7, 0x06,
	0xf5, 0xac, 0x2e, 0x62, 0x3f, 0x45, 0xa0, 0xf2, 0x9d, 0xcb, 0x33, 0x5d, 0xde, 0x43, 0x48, 0x35,
	0x9a, 0x76, 0x65, 0x7d, 0xe2, 0xf0, 0xa8, 0xf8, 0xfb, 0x0d, 0xbd, 0xc5, 0x15, 0xed, 0xca, 0xf2,
	0xe0, 0x2b, 0x1f, 0x40, 0xda, 0x29, 0x6c, 0x00, 0xa4, 0x95, 0xef, 0xa5, 0xfa, 0x24, 0xe9, 0x8e,
	0x03, 0xd3, 0x2e, 0x97, 0x68, 0x0c, 0x8d, 0xd8, 0x3e, 0x24, 0x17, 0x95, 0x58, 0xd1, 0xdd, 0xfb,
	0x5f, 0x22, 0xd8, 0x9b, 0xbf, 0x7e, 0x31, 0x7b, 0xce, 0x72, 0xc8, 0x16, 0xb2, 0xc6, 0x0b, 0x21,
	0xf1, 0x8c, 0xee, 0x30, 0x06, 0xf9, 0xf6, 0xef, 0x9a, 0xd0, 0xaf, 0xa4, 0xd3, 0xcc, 0x96, 0xf6,
	0x8d, 0xb0, 0x21, 0x64, 0xfd, 0xf0, 0x13, 0xfa, 0x3d, 0xb0, 0x09, 0xfc, 0x83, 0xb0, 0xdc, 0x2f,
	0xc7, 0x2d, 0x6c, 0x42, 0x7f, 0x06, 0x34, 0x1d, 0xfe, 0x22, 0xec, 0x36, 0x1c, 0xb8, 0xc4, 0x6e,
	0x91, 0x9d, 0xfa, 0x9b, 0xb0, 0x03, 0xc8, 0xe6, 0x4f, 0xdf, 0x9c, 0x1d, 0xbf, 0x5a, 0x9c, 0x9c,
	0xd2, 0xcd, 0x26, 0x79, 0x96, 0xbd, 0x25, 0xe1, 0x55, 0x9d, 0xa7, 0xfe, 0x59, 0x3d, 0xfa, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xdf, 0xef, 0x75, 0x5c, 0x67, 0x02, 0x00, 0x00,
}
