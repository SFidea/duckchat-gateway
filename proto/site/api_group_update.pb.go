// Code generated by protoc-gen-go. DO NOT EDIT.
// source: site/api_group_update.proto

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

type ApiGroupUpdateType int32

const (
	ApiGroupUpdateType_ApiGroupUpdateInvalid             ApiGroupUpdateType = 0
	ApiGroupUpdateType_ApiGroupUpdateName                ApiGroupUpdateType = 1
	ApiGroupUpdateType_ApiGroupUpdatePermissionJoin      ApiGroupUpdateType = 2
	ApiGroupUpdateType_ApiGroupUpdateCanGuestReadMessage ApiGroupUpdateType = 3
	ApiGroupUpdateType_ApiGroupUpdateDescription         ApiGroupUpdateType = 4
	ApiGroupUpdateType_ApiGroupUpdateAdmin               ApiGroupUpdateType = 5
	ApiGroupUpdateType_ApiGroupUpdateSpeaker             ApiGroupUpdateType = 6
	ApiGroupUpdateType_ApiGroupUpdateIsMute              ApiGroupUpdateType = 7
)

var ApiGroupUpdateType_name = map[int32]string{
	0: "ApiGroupUpdateInvalid",
	1: "ApiGroupUpdateName",
	2: "ApiGroupUpdatePermissionJoin",
	3: "ApiGroupUpdateCanGuestReadMessage",
	4: "ApiGroupUpdateDescription",
	5: "ApiGroupUpdateAdmin",
	6: "ApiGroupUpdateSpeaker",
	7: "ApiGroupUpdateIsMute",
}
var ApiGroupUpdateType_value = map[string]int32{
	"ApiGroupUpdateInvalid":             0,
	"ApiGroupUpdateName":                1,
	"ApiGroupUpdatePermissionJoin":      2,
	"ApiGroupUpdateCanGuestReadMessage": 3,
	"ApiGroupUpdateDescription":         4,
	"ApiGroupUpdateAdmin":               5,
	"ApiGroupUpdateSpeaker":             6,
	"ApiGroupUpdateIsMute":              7,
}

func (x ApiGroupUpdateType) String() string {
	return proto.EnumName(ApiGroupUpdateType_name, int32(x))
}
func (ApiGroupUpdateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_api_group_update_e92a0520fb613eec, []int{0}
}

type ApiGroupUpdateValue struct {
	Type      ApiGroupUpdateType `protobuf:"varint,1,opt,name=type,proto3,enum=site.ApiGroupUpdateType" json:"type,omitempty"`
	WriteType core.DataWriteType `protobuf:"varint,2,opt,name=writeType,proto3,enum=core.DataWriteType" json:"writeType,omitempty"`
	// Types that are valid to be assigned to Fields:
	//	*ApiGroupUpdateValue_Name
	//	*ApiGroupUpdateValue_PermissionJoin
	//	*ApiGroupUpdateValue_CanGuestReadMessage
	//	*ApiGroupUpdateValue_Description
	//	*ApiGroupUpdateValue_IsMute
	Fields isApiGroupUpdateValue_Fields `protobuf_oneof:"fields"`
	// work with: writeType
	//
	// oneof doesnot support repeated.
	AdminUserIds         []string `protobuf:"bytes,7,rep,name=adminUserIds,proto3" json:"adminUserIds,omitempty"`
	SpeakerUserIds       []string `protobuf:"bytes,8,rep,name=speakerUserIds,proto3" json:"speakerUserIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiGroupUpdateValue) Reset()         { *m = ApiGroupUpdateValue{} }
func (m *ApiGroupUpdateValue) String() string { return proto.CompactTextString(m) }
func (*ApiGroupUpdateValue) ProtoMessage()    {}
func (*ApiGroupUpdateValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_group_update_e92a0520fb613eec, []int{0}
}
func (m *ApiGroupUpdateValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiGroupUpdateValue.Unmarshal(m, b)
}
func (m *ApiGroupUpdateValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiGroupUpdateValue.Marshal(b, m, deterministic)
}
func (dst *ApiGroupUpdateValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiGroupUpdateValue.Merge(dst, src)
}
func (m *ApiGroupUpdateValue) XXX_Size() int {
	return xxx_messageInfo_ApiGroupUpdateValue.Size(m)
}
func (m *ApiGroupUpdateValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiGroupUpdateValue.DiscardUnknown(m)
}

var xxx_messageInfo_ApiGroupUpdateValue proto.InternalMessageInfo

type isApiGroupUpdateValue_Fields interface {
	isApiGroupUpdateValue_Fields()
}

type ApiGroupUpdateValue_Name struct {
	Name string `protobuf:"bytes,3,opt,name=name,proto3,oneof"`
}
type ApiGroupUpdateValue_PermissionJoin struct {
	PermissionJoin core.GroupJoinPermissionType `protobuf:"varint,4,opt,name=permissionJoin,proto3,enum=core.GroupJoinPermissionType,oneof"`
}
type ApiGroupUpdateValue_CanGuestReadMessage struct {
	CanGuestReadMessage bool `protobuf:"varint,5,opt,name=canGuestReadMessage,proto3,oneof"`
}
type ApiGroupUpdateValue_Description struct {
	Description *core.GroupDescription `protobuf:"bytes,6,opt,name=description,proto3,oneof"`
}
type ApiGroupUpdateValue_IsMute struct {
	IsMute bool `protobuf:"varint,9,opt,name=isMute,proto3,oneof"`
}

func (*ApiGroupUpdateValue_Name) isApiGroupUpdateValue_Fields()                {}
func (*ApiGroupUpdateValue_PermissionJoin) isApiGroupUpdateValue_Fields()      {}
func (*ApiGroupUpdateValue_CanGuestReadMessage) isApiGroupUpdateValue_Fields() {}
func (*ApiGroupUpdateValue_Description) isApiGroupUpdateValue_Fields()         {}
func (*ApiGroupUpdateValue_IsMute) isApiGroupUpdateValue_Fields()              {}

func (m *ApiGroupUpdateValue) GetFields() isApiGroupUpdateValue_Fields {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *ApiGroupUpdateValue) GetType() ApiGroupUpdateType {
	if m != nil {
		return m.Type
	}
	return ApiGroupUpdateType_ApiGroupUpdateInvalid
}

func (m *ApiGroupUpdateValue) GetWriteType() core.DataWriteType {
	if m != nil {
		return m.WriteType
	}
	return core.DataWriteType_WriteUpdate
}

func (m *ApiGroupUpdateValue) GetName() string {
	if x, ok := m.GetFields().(*ApiGroupUpdateValue_Name); ok {
		return x.Name
	}
	return ""
}

func (m *ApiGroupUpdateValue) GetPermissionJoin() core.GroupJoinPermissionType {
	if x, ok := m.GetFields().(*ApiGroupUpdateValue_PermissionJoin); ok {
		return x.PermissionJoin
	}
	return core.GroupJoinPermissionType_GroupJoinPermissionPublic
}

func (m *ApiGroupUpdateValue) GetCanGuestReadMessage() bool {
	if x, ok := m.GetFields().(*ApiGroupUpdateValue_CanGuestReadMessage); ok {
		return x.CanGuestReadMessage
	}
	return false
}

func (m *ApiGroupUpdateValue) GetDescription() *core.GroupDescription {
	if x, ok := m.GetFields().(*ApiGroupUpdateValue_Description); ok {
		return x.Description
	}
	return nil
}

func (m *ApiGroupUpdateValue) GetIsMute() bool {
	if x, ok := m.GetFields().(*ApiGroupUpdateValue_IsMute); ok {
		return x.IsMute
	}
	return false
}

func (m *ApiGroupUpdateValue) GetAdminUserIds() []string {
	if m != nil {
		return m.AdminUserIds
	}
	return nil
}

func (m *ApiGroupUpdateValue) GetSpeakerUserIds() []string {
	if m != nil {
		return m.SpeakerUserIds
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ApiGroupUpdateValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ApiGroupUpdateValue_OneofMarshaler, _ApiGroupUpdateValue_OneofUnmarshaler, _ApiGroupUpdateValue_OneofSizer, []interface{}{
		(*ApiGroupUpdateValue_Name)(nil),
		(*ApiGroupUpdateValue_PermissionJoin)(nil),
		(*ApiGroupUpdateValue_CanGuestReadMessage)(nil),
		(*ApiGroupUpdateValue_Description)(nil),
		(*ApiGroupUpdateValue_IsMute)(nil),
	}
}

func _ApiGroupUpdateValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ApiGroupUpdateValue)
	// fields
	switch x := m.Fields.(type) {
	case *ApiGroupUpdateValue_Name:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Name)
	case *ApiGroupUpdateValue_PermissionJoin:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.PermissionJoin))
	case *ApiGroupUpdateValue_CanGuestReadMessage:
		t := uint64(0)
		if x.CanGuestReadMessage {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *ApiGroupUpdateValue_Description:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Description); err != nil {
			return err
		}
	case *ApiGroupUpdateValue_IsMute:
		t := uint64(0)
		if x.IsMute {
			t = 1
		}
		b.EncodeVarint(9<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("ApiGroupUpdateValue.Fields has unexpected type %T", x)
	}
	return nil
}

func _ApiGroupUpdateValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ApiGroupUpdateValue)
	switch tag {
	case 3: // fields.name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Fields = &ApiGroupUpdateValue_Name{x}
		return true, err
	case 4: // fields.permissionJoin
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Fields = &ApiGroupUpdateValue_PermissionJoin{core.GroupJoinPermissionType(x)}
		return true, err
	case 5: // fields.canGuestReadMessage
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Fields = &ApiGroupUpdateValue_CanGuestReadMessage{x != 0}
		return true, err
	case 6: // fields.description
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.GroupDescription)
		err := b.DecodeMessage(msg)
		m.Fields = &ApiGroupUpdateValue_Description{msg}
		return true, err
	case 9: // fields.isMute
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Fields = &ApiGroupUpdateValue_IsMute{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _ApiGroupUpdateValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ApiGroupUpdateValue)
	// fields
	switch x := m.Fields.(type) {
	case *ApiGroupUpdateValue_Name:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case *ApiGroupUpdateValue_PermissionJoin:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.PermissionJoin))
	case *ApiGroupUpdateValue_CanGuestReadMessage:
		n += 1 // tag and wire
		n += 1
	case *ApiGroupUpdateValue_Description:
		s := proto.Size(x.Description)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ApiGroupUpdateValue_IsMute:
		n += 1 // tag and wire
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// *
//
// action: api.group.update
//
type ApiGroupUpdateRequest struct {
	GroupId              string                 `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Values               []*ApiGroupUpdateValue `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ApiGroupUpdateRequest) Reset()         { *m = ApiGroupUpdateRequest{} }
func (m *ApiGroupUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*ApiGroupUpdateRequest) ProtoMessage()    {}
func (*ApiGroupUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_group_update_e92a0520fb613eec, []int{1}
}
func (m *ApiGroupUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiGroupUpdateRequest.Unmarshal(m, b)
}
func (m *ApiGroupUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiGroupUpdateRequest.Marshal(b, m, deterministic)
}
func (dst *ApiGroupUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiGroupUpdateRequest.Merge(dst, src)
}
func (m *ApiGroupUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_ApiGroupUpdateRequest.Size(m)
}
func (m *ApiGroupUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiGroupUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApiGroupUpdateRequest proto.InternalMessageInfo

func (m *ApiGroupUpdateRequest) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *ApiGroupUpdateRequest) GetValues() []*ApiGroupUpdateValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type ApiGroupUpdateResponse struct {
	Profile              *core.PublicGroupProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	IsMute               bool                     `protobuf:"varint,2,opt,name=isMute,proto3" json:"isMute,omitempty"`
	MemberType           core.GroupMemberType     `protobuf:"varint,3,opt,name=memberType,proto3,enum=core.GroupMemberType" json:"memberType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ApiGroupUpdateResponse) Reset()         { *m = ApiGroupUpdateResponse{} }
func (m *ApiGroupUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*ApiGroupUpdateResponse) ProtoMessage()    {}
func (*ApiGroupUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_group_update_e92a0520fb613eec, []int{2}
}
func (m *ApiGroupUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiGroupUpdateResponse.Unmarshal(m, b)
}
func (m *ApiGroupUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiGroupUpdateResponse.Marshal(b, m, deterministic)
}
func (dst *ApiGroupUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiGroupUpdateResponse.Merge(dst, src)
}
func (m *ApiGroupUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_ApiGroupUpdateResponse.Size(m)
}
func (m *ApiGroupUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiGroupUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApiGroupUpdateResponse proto.InternalMessageInfo

func (m *ApiGroupUpdateResponse) GetProfile() *core.PublicGroupProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *ApiGroupUpdateResponse) GetIsMute() bool {
	if m != nil {
		return m.IsMute
	}
	return false
}

func (m *ApiGroupUpdateResponse) GetMemberType() core.GroupMemberType {
	if m != nil {
		return m.MemberType
	}
	return core.GroupMemberType_GroupMemberGuest
}

func init() {
	proto.RegisterType((*ApiGroupUpdateValue)(nil), "site.ApiGroupUpdateValue")
	proto.RegisterType((*ApiGroupUpdateRequest)(nil), "site.ApiGroupUpdateRequest")
	proto.RegisterType((*ApiGroupUpdateResponse)(nil), "site.ApiGroupUpdateResponse")
	proto.RegisterEnum("site.ApiGroupUpdateType", ApiGroupUpdateType_name, ApiGroupUpdateType_value)
}

func init() {
	proto.RegisterFile("site/api_group_update.proto", fileDescriptor_api_group_update_e92a0520fb613eec)
}

var fileDescriptor_api_group_update_e92a0520fb613eec = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xed, 0x6a, 0x13, 0x41,
	0x14, 0xcd, 0x26, 0x69, 0x3e, 0x6e, 0x24, 0x2e, 0x93, 0x36, 0x4e, 0xab, 0x85, 0x35, 0xa0, 0x04,
	0xd1, 0x0d, 0x5d, 0xf1, 0x8f, 0xff, 0x5a, 0x0b, 0x4d, 0x85, 0x4a, 0x98, 0x5a, 0x95, 0x52, 0x28,
	0x93, 0xdd, 0xdb, 0x74, 0xe8, 0x7e, 0xb9, 0x33, 0xdb, 0x12, 0x5f, 0xc5, 0xe7, 0xf0, 0x25, 0x7c,
	0x27, 0x41, 0x76, 0xb6, 0x49, 0xb3, 0x9b, 0xfe, 0xcb, 0x3d, 0xe7, 0xdc, 0x7b, 0xb3, 0xf7, 0x9c,
	0x81, 0xe7, 0x52, 0x28, 0x1c, 0xf1, 0x58, 0x5c, 0xce, 0x92, 0x28, 0x8d, 0x2f, 0xd3, 0xd8, 0xe3,
	0x0a, 0xed, 0x38, 0x89, 0x54, 0x44, 0xea, 0x19, 0xb9, 0x63, 0xba, 0x51, 0x82, 0x23, 0x4d, 0xe7,
	0xf8, 0x4e, 0x57, 0x23, 0x21, 0xaa, 0xbc, 0x1e, 0xfc, 0xa9, 0x41, 0x6f, 0x3f, 0x16, 0x47, 0x99,
	0xe4, 0x4c, 0x0f, 0xf8, 0xc6, 0xfd, 0x14, 0xc9, 0x5b, 0xa8, 0xab, 0x79, 0x8c, 0xd4, 0xb0, 0x8c,
	0x61, 0xd7, 0xa1, 0x76, 0x36, 0xce, 0x2e, 0x0a, 0xbf, 0xce, 0x63, 0x64, 0x5a, 0x45, 0xf6, 0xa0,
	0x7d, 0x97, 0x88, 0x1c, 0xa2, 0x55, 0xdd, 0xd2, 0xb3, 0xb3, 0x4d, 0xf6, 0x21, 0x57, 0xfc, 0xfb,
	0x82, 0x62, 0x0f, 0x2a, 0xb2, 0x09, 0xf5, 0x90, 0x07, 0x48, 0x6b, 0x96, 0x31, 0x6c, 0x8f, 0x2b,
	0x4c, 0x57, 0xe4, 0x08, 0xba, 0x31, 0x26, 0x81, 0x90, 0x52, 0x44, 0xe1, 0xe7, 0x48, 0x84, 0xb4,
	0xae, 0xa7, 0xed, 0xe6, 0xd3, 0xf4, 0xf6, 0x0c, 0x9e, 0x2c, 0x45, 0xd9, 0xb0, 0x71, 0x85, 0x95,
	0xda, 0x88, 0x03, 0x3d, 0x97, 0x87, 0x47, 0x29, 0x4a, 0xc5, 0x90, 0x7b, 0x27, 0x28, 0x25, 0x9f,
	0x21, 0xdd, 0xb0, 0x8c, 0x61, 0x6b, 0x5c, 0x61, 0x8f, 0x91, 0xe4, 0x23, 0x74, 0x3c, 0x94, 0x6e,
	0x22, 0x62, 0x25, 0xa2, 0x90, 0x36, 0x2c, 0x63, 0xd8, 0x71, 0xfa, 0x2b, 0x9b, 0x0f, 0x1f, 0xd8,
	0x71, 0x85, 0xad, 0x8a, 0x09, 0x85, 0x86, 0x90, 0x27, 0xa9, 0x42, 0xda, 0xbe, 0x5f, 0x71, 0x5f,
	0x93, 0x01, 0x3c, 0xe1, 0x5e, 0x20, 0xc2, 0x33, 0x89, 0xc9, 0xb1, 0x27, 0x69, 0xd3, 0xaa, 0x0d,
	0xdb, 0xac, 0x80, 0x91, 0xd7, 0xd0, 0x95, 0x31, 0xf2, 0x1b, 0x4c, 0x16, 0xaa, 0x96, 0x56, 0x95,
	0xd0, 0x83, 0x16, 0x34, 0xae, 0x04, 0xfa, 0x9e, 0x1c, 0x78, 0xb0, 0x55, 0x74, 0x83, 0xe1, 0xcf,
	0xec, 0x7b, 0x08, 0x85, 0xa6, 0xf6, 0xfb, 0xd8, 0xd3, 0xde, 0xb5, 0xd9, 0xa2, 0x24, 0x7b, 0xd0,
	0xb8, 0xcd, 0xbc, 0x95, 0xb4, 0x6a, 0xd5, 0x86, 0x1d, 0x67, 0xfb, 0x31, 0x53, 0xb5, 0xfb, 0xec,
	0x5e, 0x38, 0xf8, 0x6d, 0x40, 0xbf, 0xbc, 0x46, 0xc6, 0x51, 0x28, 0x91, 0x38, 0xd0, 0x8c, 0x93,
	0xe8, 0x4a, 0xf8, 0x79, 0x46, 0x3a, 0x0e, 0xcd, 0x0f, 0x35, 0x49, 0xa7, 0xbe, 0x70, 0x75, 0xc7,
	0x24, 0xe7, 0xd9, 0x42, 0x48, 0xfa, 0xcb, 0x23, 0x65, 0x19, 0x69, 0x2d, 0x4f, 0xf4, 0x01, 0x20,
	0xc0, 0x60, 0x8a, 0x89, 0xce, 0x4f, 0x4d, 0x3b, 0xbe, 0xb5, 0x72, 0xf7, 0x93, 0x25, 0xc9, 0x56,
	0x84, 0x6f, 0xfe, 0x19, 0x40, 0xd6, 0x23, 0x49, 0xb6, 0xcb, 0xa7, 0x39, 0x0e, 0x6f, 0xb9, 0x2f,
	0x3c, 0xb3, 0x42, 0xfa, 0xe5, 0x86, 0x2f, 0x3c, 0x40, 0xd3, 0x20, 0x16, 0xbc, 0x28, 0xe2, 0x93,
	0x42, 0x9a, 0xcc, 0x2a, 0x79, 0x05, 0x2f, 0x8b, 0x8a, 0x4f, 0xeb, 0x01, 0x32, 0x6b, 0x64, 0x17,
	0xb6, 0x8b, 0xb2, 0x95, 0xc8, 0x98, 0x75, 0xf2, 0xac, 0xfc, 0xd8, 0xf6, 0xb3, 0x14, 0x98, 0x1b,
	0xeb, 0xff, 0xf9, 0x34, 0x37, 0xde, 0x6c, 0x10, 0x0a, 0x9b, 0xa5, 0xcf, 0xd1, 0x47, 0x33, 0x9b,
	0x07, 0x3f, 0xa0, 0xe7, 0x46, 0x81, 0xfd, 0x8b, 0xfb, 0xf3, 0xfc, 0x35, 0x6b, 0x43, 0xcf, 0x47,
	0x33, 0xa1, 0xae, 0xd3, 0xa9, 0xed, 0x46, 0xc1, 0xc8, 0x4b, 0xdd, 0x1b, 0xf7, 0x9a, 0xab, 0xe5,
	0x8f, 0x77, 0x33, 0xae, 0xf0, 0x8e, 0xcf, 0x47, 0xba, 0x61, 0x94, 0x35, 0xfc, 0xad, 0x3e, 0x3d,
	0xe7, 0xfe, 0xfc, 0x62, 0x92, 0x21, 0x17, 0xa7, 0x42, 0xe1, 0xb4, 0xa1, 0xd9, 0xf7, 0xff, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x55, 0xac, 0xe3, 0x77, 0x63, 0x04, 0x00, 0x00,
}
