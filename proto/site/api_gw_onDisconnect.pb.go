// Code generated by protoc-gen-go. DO NOT EDIT.
// source: site/api_gw_onDisconnect.proto

package site // import "github.com/duckchat/duckchat-gateway/proto/site"

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

// No parameters are required.
type ApiGwOnDisconnectRequest struct {
	SocketId             string   `protobuf:"bytes,1,opt,name=socketId,proto3" json:"socketId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiGwOnDisconnectRequest) Reset()         { *m = ApiGwOnDisconnectRequest{} }
func (m *ApiGwOnDisconnectRequest) String() string { return proto.CompactTextString(m) }
func (*ApiGwOnDisconnectRequest) ProtoMessage()    {}
func (*ApiGwOnDisconnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_gw_onDisconnect_aafcb7c97e89e924, []int{0}
}
func (m *ApiGwOnDisconnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiGwOnDisconnectRequest.Unmarshal(m, b)
}
func (m *ApiGwOnDisconnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiGwOnDisconnectRequest.Marshal(b, m, deterministic)
}
func (dst *ApiGwOnDisconnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiGwOnDisconnectRequest.Merge(dst, src)
}
func (m *ApiGwOnDisconnectRequest) XXX_Size() int {
	return xxx_messageInfo_ApiGwOnDisconnectRequest.Size(m)
}
func (m *ApiGwOnDisconnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiGwOnDisconnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApiGwOnDisconnectRequest proto.InternalMessageInfo

func (m *ApiGwOnDisconnectRequest) GetSocketId() string {
	if m != nil {
		return m.SocketId
	}
	return ""
}

type ApiGwOnDisconnectResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiGwOnDisconnectResponse) Reset()         { *m = ApiGwOnDisconnectResponse{} }
func (m *ApiGwOnDisconnectResponse) String() string { return proto.CompactTextString(m) }
func (*ApiGwOnDisconnectResponse) ProtoMessage()    {}
func (*ApiGwOnDisconnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_gw_onDisconnect_aafcb7c97e89e924, []int{1}
}
func (m *ApiGwOnDisconnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiGwOnDisconnectResponse.Unmarshal(m, b)
}
func (m *ApiGwOnDisconnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiGwOnDisconnectResponse.Marshal(b, m, deterministic)
}
func (dst *ApiGwOnDisconnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiGwOnDisconnectResponse.Merge(dst, src)
}
func (m *ApiGwOnDisconnectResponse) XXX_Size() int {
	return xxx_messageInfo_ApiGwOnDisconnectResponse.Size(m)
}
func (m *ApiGwOnDisconnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiGwOnDisconnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApiGwOnDisconnectResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ApiGwOnDisconnectRequest)(nil), "site.ApiGwOnDisconnectRequest")
	proto.RegisterType((*ApiGwOnDisconnectResponse)(nil), "site.ApiGwOnDisconnectResponse")
}

func init() {
	proto.RegisterFile("site/api_gw_onDisconnect.proto", fileDescriptor_api_gw_onDisconnect_aafcb7c97e89e924)
}

var fileDescriptor_api_gw_onDisconnect_aafcb7c97e89e924 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0xce, 0x2c, 0x49,
	0xd5, 0x4f, 0x2c, 0xc8, 0x8c, 0x4f, 0x2f, 0x8f, 0xcf, 0xcf, 0x73, 0xc9, 0x2c, 0x4e, 0xce, 0xcf,
	0xcb, 0x4b, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xc9, 0x2b, 0x99,
	0x71, 0x49, 0x38, 0x16, 0x64, 0xba, 0x97, 0xfb, 0x23, 0x29, 0x08, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x92, 0xe2, 0xe2, 0x28, 0xce, 0x4f, 0xce, 0x4e, 0x2d, 0xf1, 0x4c, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x95, 0xa4, 0xb9, 0x24, 0xb1, 0xe8, 0x2b, 0x2e, 0xc8, 0xcf,
	0x2b, 0x4e, 0x75, 0x8a, 0xe0, 0x12, 0x4e, 0xce, 0xcf, 0xd5, 0xab, 0x4a, 0xcc, 0xa9, 0x84, 0x58,
	0xa6, 0x07, 0xb2, 0x2b, 0x4a, 0x3f, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0x3f, 0xa5, 0x34, 0x39, 0x3b, 0x39, 0x23, 0xb1, 0x04, 0xce, 0xd0, 0x4d, 0x4f, 0x2c, 0x49, 0x2d,
	0x4f, 0xac, 0xd4, 0x07, 0x6b, 0xd0, 0x07, 0x69, 0x38, 0xc5, 0xc4, 0x1f, 0x95, 0x98, 0x53, 0x19,
	0x13, 0x00, 0x12, 0x89, 0x09, 0xce, 0x2c, 0x49, 0x4d, 0x62, 0x03, 0xcb, 0x1a, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x56, 0xc5, 0x6c, 0xb9, 0xdd, 0x00, 0x00, 0x00,
}
