// Code generated by protoc-gen-go. DO NOT EDIT.
// source: site/api_group_profile.proto

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
// action: api.group.profile
//
type ApiGroupProfileRequest struct {
	GroupId              string   `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiGroupProfileRequest) Reset()         { *m = ApiGroupProfileRequest{} }
func (m *ApiGroupProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ApiGroupProfileRequest) ProtoMessage()    {}
func (*ApiGroupProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_group_profile_1a2345c06a6813a7, []int{0}
}
func (m *ApiGroupProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiGroupProfileRequest.Unmarshal(m, b)
}
func (m *ApiGroupProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiGroupProfileRequest.Marshal(b, m, deterministic)
}
func (dst *ApiGroupProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiGroupProfileRequest.Merge(dst, src)
}
func (m *ApiGroupProfileRequest) XXX_Size() int {
	return xxx_messageInfo_ApiGroupProfileRequest.Size(m)
}
func (m *ApiGroupProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiGroupProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApiGroupProfileRequest proto.InternalMessageInfo

func (m *ApiGroupProfileRequest) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

type ApiGroupProfileResponse struct {
	Profile              *core.PublicGroupProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	IsMute               bool                     `protobuf:"varint,2,opt,name=isMute,proto3" json:"isMute,omitempty"`
	MemberType           core.GroupMemberType     `protobuf:"varint,3,opt,name=memberType,proto3,enum=core.GroupMemberType" json:"memberType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ApiGroupProfileResponse) Reset()         { *m = ApiGroupProfileResponse{} }
func (m *ApiGroupProfileResponse) String() string { return proto.CompactTextString(m) }
func (*ApiGroupProfileResponse) ProtoMessage()    {}
func (*ApiGroupProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_group_profile_1a2345c06a6813a7, []int{1}
}
func (m *ApiGroupProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiGroupProfileResponse.Unmarshal(m, b)
}
func (m *ApiGroupProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiGroupProfileResponse.Marshal(b, m, deterministic)
}
func (dst *ApiGroupProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiGroupProfileResponse.Merge(dst, src)
}
func (m *ApiGroupProfileResponse) XXX_Size() int {
	return xxx_messageInfo_ApiGroupProfileResponse.Size(m)
}
func (m *ApiGroupProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiGroupProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApiGroupProfileResponse proto.InternalMessageInfo

func (m *ApiGroupProfileResponse) GetProfile() *core.PublicGroupProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *ApiGroupProfileResponse) GetIsMute() bool {
	if m != nil {
		return m.IsMute
	}
	return false
}

func (m *ApiGroupProfileResponse) GetMemberType() core.GroupMemberType {
	if m != nil {
		return m.MemberType
	}
	return core.GroupMemberType_GroupMemberGuest
}

func init() {
	proto.RegisterType((*ApiGroupProfileRequest)(nil), "site.ApiGroupProfileRequest")
	proto.RegisterType((*ApiGroupProfileResponse)(nil), "site.ApiGroupProfileResponse")
}

func init() {
	proto.RegisterFile("site/api_group_profile.proto", fileDescriptor_api_group_profile_1a2345c06a6813a7)
}

var fileDescriptor_api_group_profile_1a2345c06a6813a7 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x99, 0x2a, 0xad, 0x46, 0x50, 0x89, 0x58, 0x83, 0xb8, 0x18, 0xba, 0x9a, 0x8d, 0x09,
	0x8c, 0xf8, 0x00, 0xba, 0x11, 0x17, 0x85, 0x61, 0x74, 0x21, 0xa5, 0x50, 0x32, 0xe9, 0x75, 0x1a,
	0x9c, 0x31, 0x31, 0x3f, 0xc8, 0xf8, 0x2c, 0x3e, 0x91, 0x4f, 0x25, 0xc9, 0xb4, 0x45, 0xe9, 0xee,
	0xde, 0x73, 0xce, 0x77, 0x2f, 0x07, 0x5d, 0x59, 0xe9, 0x80, 0x71, 0x2d, 0x17, 0xb5, 0x51, 0x5e,
	0x2f, 0xb4, 0x51, 0xaf, 0xb2, 0x01, 0xaa, 0x8d, 0x72, 0x0a, 0xef, 0x07, 0xf7, 0xf2, 0x54, 0x28,
	0x03, 0x2c, 0xfa, 0xbd, 0x3e, 0xc9, 0xd1, 0xf8, 0x4e, 0xcb, 0x87, 0xa0, 0x14, 0x3d, 0x50, 0xc2,
	0x87, 0x07, 0xeb, 0x30, 0x41, 0xa3, 0x18, 0x7c, 0x5c, 0x92, 0x24, 0x4d, 0xb2, 0xc3, 0x72, 0xb3,
	0x4e, 0xbe, 0x13, 0x74, 0xb1, 0x03, 0x59, 0xad, 0xde, 0x2d, 0xe0, 0x1c, 0x8d, 0xd6, 0x8f, 0x23,
	0x75, 0x94, 0x13, 0x1a, 0x7e, 0xd2, 0xc2, 0x57, 0x8d, 0x14, 0xff, 0x90, 0x4d, 0x10, 0x8f, 0xd1,
	0x50, 0xda, 0xa9, 0x77, 0x40, 0x06, 0x69, 0x92, 0x1d, 0x94, 0xeb, 0x0d, 0xdf, 0x22, 0xd4, 0x42,
	0x5b, 0x81, 0x79, 0xee, 0x34, 0x90, 0xbd, 0x34, 0xc9, 0x8e, 0xf3, 0xf3, 0xfe, 0x5c, 0x3c, 0x34,
	0xdd, 0x9a, 0xe5, 0x9f, 0xe0, 0xfd, 0x0b, 0x3a, 0x13, 0xaa, 0xa5, 0x5f, 0xbc, 0xe9, 0xfa, 0x92,
	0x34, 0x74, 0x9f, 0xb1, 0x5a, 0xba, 0x95, 0xaf, 0xa8, 0x50, 0x2d, 0x5b, 0x7a, 0xf1, 0x26, 0x56,
	0xdc, 0x6d, 0x87, 0xeb, 0x9a, 0x3b, 0xf8, 0xe4, 0x1d, 0x8b, 0x00, 0x0b, 0xc0, 0xcf, 0xe0, 0x64,
	0xc6, 0x9b, 0x6e, 0x5e, 0x04, 0x65, 0xfe, 0x24, 0x1d, 0x54, 0xc3, 0xe8, 0xde, 0xfc, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x91, 0xe8, 0x3a, 0xf9, 0x6b, 0x01, 0x00, 0x00,
}
