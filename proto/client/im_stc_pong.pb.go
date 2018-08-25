// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client/im_stc_pong.proto

package client // import "github.com/duckchat/duckchat-gateway/proto/client"

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
type ImStcPongRequest struct {
	Random               string   `protobuf:"bytes,1,opt,name=random,proto3" json:"random,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImStcPongRequest) Reset()         { *m = ImStcPongRequest{} }
func (m *ImStcPongRequest) String() string { return proto.CompactTextString(m) }
func (*ImStcPongRequest) ProtoMessage()    {}
func (*ImStcPongRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_im_stc_pong_5d4c001564ca06cd, []int{0}
}
func (m *ImStcPongRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImStcPongRequest.Unmarshal(m, b)
}
func (m *ImStcPongRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImStcPongRequest.Marshal(b, m, deterministic)
}
func (dst *ImStcPongRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImStcPongRequest.Merge(dst, src)
}
func (m *ImStcPongRequest) XXX_Size() int {
	return xxx_messageInfo_ImStcPongRequest.Size(m)
}
func (m *ImStcPongRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImStcPongRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImStcPongRequest proto.InternalMessageInfo

func (m *ImStcPongRequest) GetRandom() string {
	if m != nil {
		return m.Random
	}
	return ""
}

func init() {
	proto.RegisterType((*ImStcPongRequest)(nil), "client.ImStcPongRequest")
}

func init() {
	proto.RegisterFile("client/im_stc_pong.proto", fileDescriptor_im_stc_pong_5d4c001564ca06cd)
}

var fileDescriptor_im_stc_pong_5d4c001564ca06cd = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0xcf, 0xcc, 0x8d, 0x2f, 0x2e, 0x49, 0x8e, 0x2f, 0xc8, 0xcf, 0x4b, 0xd7, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xc8, 0x28, 0x69, 0x71, 0x09, 0x78, 0xe6, 0x06, 0x97,
	0x24, 0x07, 0xe4, 0xe7, 0xa5, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x71, 0xb1,
	0x15, 0x25, 0xe6, 0xa5, 0xe4, 0xe7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x4e,
	0x71, 0x5c, 0xa2, 0xc9, 0xf9, 0xb9, 0x7a, 0x55, 0x89, 0x39, 0x95, 0x10, 0x53, 0xf4, 0x20, 0x86,
	0x44, 0x19, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0x94, 0x26,
	0x67, 0x27, 0x67, 0x24, 0x96, 0xc0, 0x19, 0xba, 0xe9, 0x89, 0x25, 0xa9, 0xe5, 0x89, 0x95, 0xfa,
	0x60, 0x2d, 0xfa, 0x10, 0x2d, 0xa7, 0x98, 0x04, 0xa3, 0x12, 0x73, 0x2a, 0x63, 0x02, 0x40, 0x62,
	0x31, 0xce, 0x60, 0xb1, 0x24, 0x36, 0xb0, 0x0a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d,
	0xba, 0x79, 0xbc, 0xb6, 0x00, 0x00, 0x00,
}
