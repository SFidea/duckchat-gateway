// Code generated by protoc-gen-go. DO NOT EDIT.
// source: site/api_tmpchat_delete.proto

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

// *
//
// action: api.tmpchat.delete
//
type ApiTmpchatDeleteRequest struct {
	ToUserId             string   `protobuf:"bytes,1,opt,name=toUserId,proto3" json:"toUserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiTmpchatDeleteRequest) Reset()         { *m = ApiTmpchatDeleteRequest{} }
func (m *ApiTmpchatDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*ApiTmpchatDeleteRequest) ProtoMessage()    {}
func (*ApiTmpchatDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_tmpchat_delete_04f83583176bb0cc, []int{0}
}
func (m *ApiTmpchatDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiTmpchatDeleteRequest.Unmarshal(m, b)
}
func (m *ApiTmpchatDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiTmpchatDeleteRequest.Marshal(b, m, deterministic)
}
func (dst *ApiTmpchatDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiTmpchatDeleteRequest.Merge(dst, src)
}
func (m *ApiTmpchatDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_ApiTmpchatDeleteRequest.Size(m)
}
func (m *ApiTmpchatDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiTmpchatDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApiTmpchatDeleteRequest proto.InternalMessageInfo

func (m *ApiTmpchatDeleteRequest) GetToUserId() string {
	if m != nil {
		return m.ToUserId
	}
	return ""
}

type ApiTmpchatDeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiTmpchatDeleteResponse) Reset()         { *m = ApiTmpchatDeleteResponse{} }
func (m *ApiTmpchatDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*ApiTmpchatDeleteResponse) ProtoMessage()    {}
func (*ApiTmpchatDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_tmpchat_delete_04f83583176bb0cc, []int{1}
}
func (m *ApiTmpchatDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiTmpchatDeleteResponse.Unmarshal(m, b)
}
func (m *ApiTmpchatDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiTmpchatDeleteResponse.Marshal(b, m, deterministic)
}
func (dst *ApiTmpchatDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiTmpchatDeleteResponse.Merge(dst, src)
}
func (m *ApiTmpchatDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_ApiTmpchatDeleteResponse.Size(m)
}
func (m *ApiTmpchatDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiTmpchatDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApiTmpchatDeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ApiTmpchatDeleteRequest)(nil), "site.ApiTmpchatDeleteRequest")
	proto.RegisterType((*ApiTmpchatDeleteResponse)(nil), "site.ApiTmpchatDeleteResponse")
}

func init() {
	proto.RegisterFile("site/api_tmpchat_delete.proto", fileDescriptor_api_tmpchat_delete_04f83583176bb0cc)
}

var fileDescriptor_api_tmpchat_delete_04f83583176bb0cc = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0xce, 0x2c, 0x49,
	0xd5, 0x4f, 0x2c, 0xc8, 0x8c, 0x2f, 0xc9, 0x2d, 0x48, 0xce, 0x48, 0x2c, 0x89, 0x4f, 0x49, 0xcd,
	0x49, 0x2d, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x49, 0x2b, 0x99, 0x72,
	0x89, 0x3b, 0x16, 0x64, 0x86, 0x40, 0x14, 0xb8, 0x80, 0xe5, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0xa4, 0xb8, 0x38, 0x4a, 0xf2, 0x43, 0x8b, 0x53, 0x8b, 0x3c, 0x53, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x25, 0x29, 0x2e, 0x09, 0x4c, 0x6d, 0xc5, 0x05, 0xf9, 0x79,
	0xc5, 0xa9, 0x4e, 0x11, 0x5c, 0xc2, 0xc9, 0xf9, 0xb9, 0x7a, 0x55, 0x89, 0x39, 0x95, 0x10, 0xab,
	0xf4, 0x40, 0x36, 0x45, 0xe9, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0xa7, 0x94, 0x26, 0x67, 0x83, 0x74, 0xc2, 0x19, 0xba, 0xe9, 0x89, 0x25, 0xa9, 0xe5, 0x89, 0x95,
	0xfa, 0x60, 0x0d, 0xfa, 0x20, 0x0d, 0xa7, 0x98, 0xf8, 0xa3, 0x12, 0x73, 0x2a, 0x63, 0x02, 0x40,
	0x22, 0x31, 0xc1, 0x99, 0x25, 0xa9, 0x49, 0x6c, 0x60, 0x59, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x91, 0x9d, 0x05, 0x5c, 0xda, 0x00, 0x00, 0x00,
}
