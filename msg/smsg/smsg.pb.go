// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg/smsg/smsg.proto

/*
Package smsg is a generated protocol buffer package.

It is generated from these files:
	msg/smsg/smsg.proto

It has these top-level messages:
	Server2AllSession
	GaCeReqLogin
	GaCeRespLogin
	CeGaBindGameServer
	CeGamReqJoinGame
	CeGamRespJoinGame
	GamCeNoticeGameStart
	GamCeNoticeGameEnd
	CeGameUserDisconnect
	CeGameUserReconnect
	GaCeUserDisconnect
	CeGaCloseSession
*/
package smsg

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

type GaCeRespLogin_Error int32

const (
	GaCeRespLogin_Invalid GaCeRespLogin_Error = 0
)

var GaCeRespLogin_Error_name = map[int32]string{
	0: "Invalid",
}
var GaCeRespLogin_Error_value = map[string]int32{
	"Invalid": 0,
}

func (x GaCeRespLogin_Error) String() string {
	return proto.EnumName(GaCeRespLogin_Error_name, int32(x))
}
func (GaCeRespLogin_Error) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type CeGamRespJoinGame_Error int32

const (
	CeGamRespJoinGame_Invalid      CeGamRespJoinGame_Error = 0
	CeGamRespJoinGame_GameNotExist CeGamRespJoinGame_Error = 1
)

var CeGamRespJoinGame_Error_name = map[int32]string{
	0: "Invalid",
	1: "GameNotExist",
}
var CeGamRespJoinGame_Error_value = map[string]int32{
	"Invalid":      0,
	"GameNotExist": 1,
}

func (x CeGamRespJoinGame_Error) String() string {
	return proto.EnumName(CeGamRespJoinGame_Error_name, int32(x))
}
func (CeGamRespJoinGame_Error) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type Server2AllSession struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Server2AllSession) Reset()                    { *m = Server2AllSession{} }
func (m *Server2AllSession) String() string            { return proto.CompactTextString(m) }
func (*Server2AllSession) ProtoMessage()               {}
func (*Server2AllSession) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Server2AllSession) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GaCeReqLogin struct {
	Sesid int32  `protobuf:"varint,1,opt,name=sesid" json:"sesid,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *GaCeReqLogin) Reset()                    { *m = GaCeReqLogin{} }
func (m *GaCeReqLogin) String() string            { return proto.CompactTextString(m) }
func (*GaCeReqLogin) ProtoMessage()               {}
func (*GaCeReqLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GaCeReqLogin) GetSesid() int32 {
	if m != nil {
		return m.Sesid
	}
	return 0
}

func (m *GaCeReqLogin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GaCeRespLogin struct {
	UserID uint64              `protobuf:"varint,1,opt,name=userID" json:"userID,omitempty"`
	Err    GaCeRespLogin_Error `protobuf:"varint,2,opt,name=err,enum=smsg.GaCeRespLogin_Error" json:"err,omitempty"`
	Token  string              `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *GaCeRespLogin) Reset()                    { *m = GaCeRespLogin{} }
func (m *GaCeRespLogin) String() string            { return proto.CompactTextString(m) }
func (*GaCeRespLogin) ProtoMessage()               {}
func (*GaCeRespLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GaCeRespLogin) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *GaCeRespLogin) GetErr() GaCeRespLogin_Error {
	if m != nil {
		return m.Err
	}
	return GaCeRespLogin_Invalid
}

func (m *GaCeRespLogin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CeGaBindGameServer struct {
	Sesid        int32 `protobuf:"varint,1,opt,name=sesid" json:"sesid,omitempty"`
	Gameserverid int32 `protobuf:"varint,2,opt,name=gameserverid" json:"gameserverid,omitempty"`
}

func (m *CeGaBindGameServer) Reset()                    { *m = CeGaBindGameServer{} }
func (m *CeGaBindGameServer) String() string            { return proto.CompactTextString(m) }
func (*CeGaBindGameServer) ProtoMessage()               {}
func (*CeGaBindGameServer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CeGaBindGameServer) GetSesid() int32 {
	if m != nil {
		return m.Sesid
	}
	return 0
}

func (m *CeGaBindGameServer) GetGameserverid() int32 {
	if m != nil {
		return m.Gameserverid
	}
	return 0
}

type CeGamReqJoinGame struct {
	Userid       uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Sesid        int32  `protobuf:"varint,2,opt,name=sesid" json:"sesid,omitempty"`
	Nickname     string `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
	GateServerid int32  `protobuf:"varint,4,opt,name=gateServerid" json:"gateServerid,omitempty"`
}

func (m *CeGamReqJoinGame) Reset()                    { *m = CeGamReqJoinGame{} }
func (m *CeGamReqJoinGame) String() string            { return proto.CompactTextString(m) }
func (*CeGamReqJoinGame) ProtoMessage()               {}
func (*CeGamReqJoinGame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CeGamReqJoinGame) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *CeGamReqJoinGame) GetSesid() int32 {
	if m != nil {
		return m.Sesid
	}
	return 0
}

func (m *CeGamReqJoinGame) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *CeGamReqJoinGame) GetGateServerid() int32 {
	if m != nil {
		return m.GateServerid
	}
	return 0
}

type CeGamRespJoinGame struct {
	Err    CeGamRespJoinGame_Error `protobuf:"varint,1,opt,name=err,enum=smsg.CeGamRespJoinGame_Error" json:"err,omitempty"`
	Gameid int64                   `protobuf:"varint,2,opt,name=gameid" json:"gameid,omitempty"`
}

func (m *CeGamRespJoinGame) Reset()                    { *m = CeGamRespJoinGame{} }
func (m *CeGamRespJoinGame) String() string            { return proto.CompactTextString(m) }
func (*CeGamRespJoinGame) ProtoMessage()               {}
func (*CeGamRespJoinGame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CeGamRespJoinGame) GetErr() CeGamRespJoinGame_Error {
	if m != nil {
		return m.Err
	}
	return CeGamRespJoinGame_Invalid
}

func (m *CeGamRespJoinGame) GetGameid() int64 {
	if m != nil {
		return m.Gameid
	}
	return 0
}

type GamCeNoticeGameStart struct {
	Gameid int64 `protobuf:"varint,1,opt,name=gameid" json:"gameid,omitempty"`
}

func (m *GamCeNoticeGameStart) Reset()                    { *m = GamCeNoticeGameStart{} }
func (m *GamCeNoticeGameStart) String() string            { return proto.CompactTextString(m) }
func (*GamCeNoticeGameStart) ProtoMessage()               {}
func (*GamCeNoticeGameStart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GamCeNoticeGameStart) GetGameid() int64 {
	if m != nil {
		return m.Gameid
	}
	return 0
}

type GamCeNoticeGameEnd struct {
	Gameid int64 `protobuf:"varint,1,opt,name=gameid" json:"gameid,omitempty"`
}

func (m *GamCeNoticeGameEnd) Reset()                    { *m = GamCeNoticeGameEnd{} }
func (m *GamCeNoticeGameEnd) String() string            { return proto.CompactTextString(m) }
func (*GamCeNoticeGameEnd) ProtoMessage()               {}
func (*GamCeNoticeGameEnd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GamCeNoticeGameEnd) GetGameid() int64 {
	if m != nil {
		return m.Gameid
	}
	return 0
}

type CeGameUserDisconnect struct {
	Userid uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
}

func (m *CeGameUserDisconnect) Reset()                    { *m = CeGameUserDisconnect{} }
func (m *CeGameUserDisconnect) String() string            { return proto.CompactTextString(m) }
func (*CeGameUserDisconnect) ProtoMessage()               {}
func (*CeGameUserDisconnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CeGameUserDisconnect) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type CeGameUserReconnect struct {
	Userid    uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	GateID    int32  `protobuf:"varint,2,opt,name=gateID" json:"gateID,omitempty"`
	SessionID int32  `protobuf:"varint,3,opt,name=sessionID" json:"sessionID,omitempty"`
}

func (m *CeGameUserReconnect) Reset()                    { *m = CeGameUserReconnect{} }
func (m *CeGameUserReconnect) String() string            { return proto.CompactTextString(m) }
func (*CeGameUserReconnect) ProtoMessage()               {}
func (*CeGameUserReconnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CeGameUserReconnect) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *CeGameUserReconnect) GetGateID() int32 {
	if m != nil {
		return m.GateID
	}
	return 0
}

func (m *CeGameUserReconnect) GetSessionID() int32 {
	if m != nil {
		return m.SessionID
	}
	return 0
}

type GaCeUserDisconnect struct {
	SessionID int32 `protobuf:"varint,1,opt,name=sessionID" json:"sessionID,omitempty"`
}

func (m *GaCeUserDisconnect) Reset()                    { *m = GaCeUserDisconnect{} }
func (m *GaCeUserDisconnect) String() string            { return proto.CompactTextString(m) }
func (*GaCeUserDisconnect) ProtoMessage()               {}
func (*GaCeUserDisconnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GaCeUserDisconnect) GetSessionID() int32 {
	if m != nil {
		return m.SessionID
	}
	return 0
}

type CeGaCloseSession struct {
	SessionID int32 `protobuf:"varint,1,opt,name=sessionID" json:"sessionID,omitempty"`
}

func (m *CeGaCloseSession) Reset()                    { *m = CeGaCloseSession{} }
func (m *CeGaCloseSession) String() string            { return proto.CompactTextString(m) }
func (*CeGaCloseSession) ProtoMessage()               {}
func (*CeGaCloseSession) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CeGaCloseSession) GetSessionID() int32 {
	if m != nil {
		return m.SessionID
	}
	return 0
}

func init() {
	proto.RegisterType((*Server2AllSession)(nil), "smsg.Server2AllSession")
	proto.RegisterType((*GaCeReqLogin)(nil), "smsg.GaCeReqLogin")
	proto.RegisterType((*GaCeRespLogin)(nil), "smsg.GaCeRespLogin")
	proto.RegisterType((*CeGaBindGameServer)(nil), "smsg.CeGaBindGameServer")
	proto.RegisterType((*CeGamReqJoinGame)(nil), "smsg.CeGamReqJoinGame")
	proto.RegisterType((*CeGamRespJoinGame)(nil), "smsg.CeGamRespJoinGame")
	proto.RegisterType((*GamCeNoticeGameStart)(nil), "smsg.GamCeNoticeGameStart")
	proto.RegisterType((*GamCeNoticeGameEnd)(nil), "smsg.GamCeNoticeGameEnd")
	proto.RegisterType((*CeGameUserDisconnect)(nil), "smsg.CeGameUserDisconnect")
	proto.RegisterType((*CeGameUserReconnect)(nil), "smsg.CeGameUserReconnect")
	proto.RegisterType((*GaCeUserDisconnect)(nil), "smsg.GaCeUserDisconnect")
	proto.RegisterType((*CeGaCloseSession)(nil), "smsg.CeGaCloseSession")
	proto.RegisterEnum("smsg.GaCeRespLogin_Error", GaCeRespLogin_Error_name, GaCeRespLogin_Error_value)
	proto.RegisterEnum("smsg.CeGamRespJoinGame_Error", CeGamRespJoinGame_Error_name, CeGamRespJoinGame_Error_value)
}

func init() { proto.RegisterFile("msg/smsg/smsg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xc5, 0x4d, 0x36, 0xd0, 0x21, 0xa0, 0xd4, 0x8d, 0xaa, 0x80, 0x40, 0xaa, 0x7c, 0x80, 0x4a,
	0xa0, 0x2d, 0x0a, 0x37, 0x6e, 0xb0, 0x89, 0xa2, 0x20, 0xd4, 0x83, 0x2b, 0x7e, 0x80, 0xd9, 0x1d,
	0x45, 0x56, 0xb3, 0x76, 0x6a, 0x9b, 0x8a, 0x23, 0x07, 0x7e, 0x38, 0xf2, 0xc7, 0x76, 0x77, 0x03,
	0x81, 0xcb, 0x6a, 0x67, 0xf4, 0xde, 0xcc, 0x7b, 0xcf, 0x36, 0x9c, 0xd6, 0x76, 0x73, 0x69, 0x9b,
	0x4f, 0xbe, 0x33, 0xda, 0x69, 0x3a, 0xf4, 0xff, 0xec, 0x35, 0x9c, 0x5c, 0xa3, 0xb9, 0x43, 0x33,
	0xff, 0xb8, 0xdd, 0x5e, 0xa3, 0xb5, 0x52, 0x2b, 0x4a, 0x61, 0x58, 0x09, 0x27, 0x66, 0xe4, 0x9c,
	0x5c, 0x8c, 0x79, 0xf8, 0x67, 0x1f, 0x60, 0xbc, 0x12, 0x05, 0x72, 0xbc, 0xfd, 0xa2, 0x37, 0x52,
	0xd1, 0x29, 0x64, 0x16, 0xad, 0xac, 0x02, 0x28, 0xe3, 0xb1, 0xf0, 0x5d, 0xa7, 0x6f, 0x50, 0xcd,
	0x8e, 0xce, 0xc9, 0xc5, 0x31, 0x8f, 0x05, 0xfb, 0x49, 0xe0, 0x49, 0x24, 0xdb, 0x5d, 0x64, 0x9f,
	0xc1, 0xe8, 0xbb, 0x45, 0xb3, 0x5e, 0x04, 0xfa, 0x90, 0xa7, 0x8a, 0xbe, 0x81, 0x01, 0x1a, 0x13,
	0xd8, 0x4f, 0xe7, 0xcf, 0xf2, 0x20, 0xb7, 0xc7, 0xcc, 0x97, 0xc6, 0x68, 0xc3, 0x3d, 0xaa, 0x5d,
	0x36, 0xe8, 0x2e, 0x9b, 0x42, 0x16, 0x30, 0xf4, 0x31, 0x3c, 0x5c, 0xab, 0x3b, 0xb1, 0x95, 0xd5,
	0xe4, 0x01, 0xbb, 0x02, 0x5a, 0xe0, 0x4a, 0x7c, 0x92, 0xaa, 0x5a, 0x89, 0x1a, 0xa3, 0xe7, 0x03,
	0x26, 0x18, 0x8c, 0x37, 0xa2, 0x46, 0x1b, 0x30, 0xb2, 0x0a, 0x6a, 0x32, 0xde, 0xeb, 0x79, 0x4b,
	0x13, 0x3f, 0xb0, 0xe6, 0x78, 0xfb, 0x59, 0x4b, 0xe5, 0x87, 0x36, 0xae, 0xd2, 0xbc, 0xe4, 0x2a,
	0xa6, 0x12, 0xd7, 0x1c, 0x75, 0xd7, 0x3c, 0x87, 0x47, 0x4a, 0x96, 0x37, 0x4a, 0xd4, 0x98, 0x1c,
	0xdc, 0xd7, 0x51, 0x82, 0x4b, 0x32, 0x65, 0x35, 0x1b, 0x36, 0x12, 0xda, 0x1e, 0xfb, 0x45, 0xe0,
	0x24, 0x49, 0xb0, 0xbb, 0x7b, 0x0d, 0x97, 0x31, 0x41, 0x12, 0x12, 0x7c, 0x19, 0x13, 0xfc, 0x03,
	0xd5, 0x4d, 0xf1, 0x0c, 0x46, 0xde, 0x59, 0x52, 0x37, 0xe0, 0xa9, 0x62, 0xaf, 0xfe, 0x96, 0x23,
	0x9d, 0xf8, 0x6b, 0x50, 0xe3, 0x95, 0x76, 0xcb, 0x1f, 0xd2, 0xba, 0x09, 0x61, 0x39, 0x4c, 0x57,
	0xa2, 0x2e, 0x7c, 0x4b, 0x96, 0x18, 0xc2, 0x75, 0xc2, 0xb8, 0xce, 0x5c, 0xd2, 0x9b, 0xfb, 0x16,
	0xe8, 0x1e, 0x7e, 0xa9, 0xaa, 0x83, 0xe8, 0x1c, 0xa6, 0x41, 0x3d, 0x7e, 0xb5, 0x68, 0x16, 0xd2,
	0x96, 0x5a, 0x29, 0x2c, 0xdd, 0xa1, 0xa8, 0x59, 0x09, 0xa7, 0x2d, 0x9e, 0xe3, 0x7f, 0xe0, 0x71,
	0xad, 0xc3, 0xf5, 0x22, 0x1d, 0x4d, 0xaa, 0xe8, 0x0b, 0x38, 0xb6, 0xf1, 0x31, 0xac, 0x17, 0xe1,
	0x70, 0x32, 0xde, 0x36, 0xd8, 0xdc, 0x5b, 0x28, 0xf6, 0x25, 0xf5, 0x38, 0x64, 0x9f, 0xf3, 0x2e,
	0xde, 0x97, 0x62, 0xab, 0x2d, 0x36, 0xef, 0xec, 0x9f, 0x8c, 0x6f, 0xa3, 0xf0, 0x4e, 0xdf, 0xff,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x77, 0x6e, 0x6e, 0x31, 0xbe, 0x03, 0x00, 0x00,
}
