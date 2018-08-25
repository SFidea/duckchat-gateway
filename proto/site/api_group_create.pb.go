// Code generated by protoc-gen-go. DO NOT EDIT.
// source: site/api_group_create.proto

package site // import "github.com/duckchat/duckchat-gateway/proto/site"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/duckchat/duckchat-gateway/proto/core"

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
// action: api.group.create
//
type ApiGroupCreateRequest struct {
	GroupName            string   `protobuf:"bytes,1,opt,name=groupName,proto3" json:"groupName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiGroupCreateRequest) Reset()         { *m = ApiGroupCreateRequest{} }
func (m *ApiGroupCreateRequest) String() string { return proto.CompactTextString(m) }
func (*ApiGroupCreateRequest) ProtoMessage()    {}
func (*ApiGroupCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_group_create_63b3a180b58569f2, []int{0}
}
func (m *ApiGroupCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiGroupCreateRequest.Unmarshal(m, b)
}
func (m *ApiGroupCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiGroupCreateRequest.Marshal(b, m, deterministic)
}
func (dst *ApiGroupCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiGroupCreateRequest.Merge(dst, src)
}
func (m *ApiGroupCreateRequest) XXX_Size() int {
	return xxx_messageInfo_ApiGroupCreateRequest.Size(m)
}
func (m *ApiGroupCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiGroupCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApiGroupCreateRequest proto.InternalMessageInfo

func (m *ApiGroupCreateRequest) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type ApiGroupCreateResponse struct {
	Profile              *core.AllGroupProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ApiGroupCreateResponse) Reset()         { *m = ApiGroupCreateResponse{} }
func (m *ApiGroupCreateResponse) String() string { return proto.CompactTextString(m) }
func (*ApiGroupCreateResponse) ProtoMessage()    {}
func (*ApiGroupCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_group_create_63b3a180b58569f2, []int{1}
}
func (m *ApiGroupCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiGroupCreateResponse.Unmarshal(m, b)
}
func (m *ApiGroupCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiGroupCreateResponse.Marshal(b, m, deterministic)
}
func (dst *ApiGroupCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiGroupCreateResponse.Merge(dst, src)
}
func (m *ApiGroupCreateResponse) XXX_Size() int {
	return xxx_messageInfo_ApiGroupCreateResponse.Size(m)
}
func (m *ApiGroupCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiGroupCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApiGroupCreateResponse proto.InternalMessageInfo

func (m *ApiGroupCreateResponse) GetProfile() *core.AllGroupProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func init() {
	proto.RegisterType((*ApiGroupCreateRequest)(nil), "site.ApiGroupCreateRequest")
	proto.RegisterType((*ApiGroupCreateResponse)(nil), "site.ApiGroupCreateResponse")
}

func init() {
	proto.RegisterFile("site/api_group_create.proto", fileDescriptor_api_group_create_63b3a180b58569f2)
}

var fileDescriptor_api_group_create_63b3a180b58569f2 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0xce, 0x2c, 0x49,
	0xd5, 0x4f, 0x2c, 0xc8, 0x8c, 0x4f, 0x2f, 0xca, 0x2f, 0x2d, 0x88, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c,
	0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x49, 0x4a, 0x09, 0x24, 0xe7, 0x17,
	0xa5, 0xea, 0x83, 0xa5, 0x21, 0xe2, 0x4a, 0xa6, 0x5c, 0xa2, 0x8e, 0x05, 0x99, 0xee, 0x20, 0x11,
	0x67, 0xb0, 0xfa, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x19, 0x2e, 0x4e, 0xb0, 0x3a,
	0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x84, 0x80, 0x92, 0x27, 0x97,
	0x18, 0xba, 0xb6, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x7d, 0x2e, 0xf6, 0x82, 0xa2, 0xfc,
	0xb4, 0xcc, 0x1c, 0x88, 0x2e, 0x6e, 0x23, 0x51, 0x3d, 0x90, 0xa5, 0x7a, 0x8e, 0x39, 0x39, 0x60,
	0xe5, 0x01, 0x10, 0xc9, 0x20, 0x98, 0x2a, 0xa7, 0x08, 0x2e, 0xe1, 0xe4, 0xfc, 0x5c, 0xbd, 0xaa,
	0xc4, 0x9c, 0x4a, 0x88, 0x9b, 0xf4, 0x40, 0x4e, 0x8d, 0xd2, 0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x29, 0x4d, 0xce, 0x4e, 0xce, 0x48, 0x2c, 0x81, 0x33, 0x74,
	0xd3, 0x13, 0x4b, 0x52, 0xcb, 0x13, 0x2b, 0xf5, 0xc1, 0x1a, 0xf4, 0x41, 0x1a, 0x4e, 0x31, 0xf1,
	0x47, 0x25, 0xe6, 0x54, 0xc6, 0x04, 0x80, 0x44, 0x62, 0x82, 0x33, 0x4b, 0x52, 0x93, 0xd8, 0xc0,
	0xb2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xfb, 0xa6, 0x16, 0x19, 0x01, 0x00, 0x00,
}
