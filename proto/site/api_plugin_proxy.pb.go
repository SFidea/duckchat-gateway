// Code generated by protoc-gen-go. DO NOT EDIT.
// source: site/api_plugin_proxy.proto

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
// action: api.plugin.proxy
//
type ApiPluginProxyRequest struct {
	PluginId             int32              `protobuf:"varint,1,opt,name=pluginId,proto3" json:"pluginId,omitempty"`
	Url                  string             `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Method               core.HttpQueryType `protobuf:"varint,3,opt,name=method,proto3,enum=core.HttpQueryType" json:"method,omitempty"`
	Body                 string             `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Cookie               string             `protobuf:"bytes,5,opt,name=cookie,proto3" json:"cookie,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ApiPluginProxyRequest) Reset()         { *m = ApiPluginProxyRequest{} }
func (m *ApiPluginProxyRequest) String() string { return proto.CompactTextString(m) }
func (*ApiPluginProxyRequest) ProtoMessage()    {}
func (*ApiPluginProxyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_plugin_proxy_2771ec85cab78e45, []int{0}
}
func (m *ApiPluginProxyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiPluginProxyRequest.Unmarshal(m, b)
}
func (m *ApiPluginProxyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiPluginProxyRequest.Marshal(b, m, deterministic)
}
func (dst *ApiPluginProxyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiPluginProxyRequest.Merge(dst, src)
}
func (m *ApiPluginProxyRequest) XXX_Size() int {
	return xxx_messageInfo_ApiPluginProxyRequest.Size(m)
}
func (m *ApiPluginProxyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiPluginProxyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApiPluginProxyRequest proto.InternalMessageInfo

func (m *ApiPluginProxyRequest) GetPluginId() int32 {
	if m != nil {
		return m.PluginId
	}
	return 0
}

func (m *ApiPluginProxyRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ApiPluginProxyRequest) GetMethod() core.HttpQueryType {
	if m != nil {
		return m.Method
	}
	return core.HttpQueryType_HttpQueryInvalid
}

func (m *ApiPluginProxyRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *ApiPluginProxyRequest) GetCookie() string {
	if m != nil {
		return m.Cookie
	}
	return ""
}

type ApiPluginProxyResponse struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Cookie               string   `protobuf:"bytes,2,opt,name=cookie,proto3" json:"cookie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiPluginProxyResponse) Reset()         { *m = ApiPluginProxyResponse{} }
func (m *ApiPluginProxyResponse) String() string { return proto.CompactTextString(m) }
func (*ApiPluginProxyResponse) ProtoMessage()    {}
func (*ApiPluginProxyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_plugin_proxy_2771ec85cab78e45, []int{1}
}
func (m *ApiPluginProxyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiPluginProxyResponse.Unmarshal(m, b)
}
func (m *ApiPluginProxyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiPluginProxyResponse.Marshal(b, m, deterministic)
}
func (dst *ApiPluginProxyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiPluginProxyResponse.Merge(dst, src)
}
func (m *ApiPluginProxyResponse) XXX_Size() int {
	return xxx_messageInfo_ApiPluginProxyResponse.Size(m)
}
func (m *ApiPluginProxyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiPluginProxyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApiPluginProxyResponse proto.InternalMessageInfo

func (m *ApiPluginProxyResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *ApiPluginProxyResponse) GetCookie() string {
	if m != nil {
		return m.Cookie
	}
	return ""
}

func init() {
	proto.RegisterType((*ApiPluginProxyRequest)(nil), "site.ApiPluginProxyRequest")
	proto.RegisterType((*ApiPluginProxyResponse)(nil), "site.ApiPluginProxyResponse")
}

func init() {
	proto.RegisterFile("site/api_plugin_proxy.proto", fileDescriptor_api_plugin_proxy_2771ec85cab78e45)
}

var fileDescriptor_api_plugin_proxy_2771ec85cab78e45 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0xf4, 0x30,
	0x14, 0xc5, 0xc9, 0xfc, 0xe3, 0xfb, 0xb2, 0x18, 0x25, 0x83, 0x43, 0x19, 0x37, 0xc3, 0xac, 0x0a,
	0x62, 0x02, 0xfa, 0x04, 0x8a, 0x0b, 0xdd, 0xd5, 0xea, 0x42, 0x86, 0x81, 0x21, 0x4d, 0x2f, 0x6d,
	0x98, 0xb6, 0x89, 0x6d, 0x82, 0xc6, 0x87, 0xf1, 0x81, 0x7c, 0x2a, 0x49, 0xa6, 0x14, 0x11, 0x77,
	0xe7, 0x9e, 0x7b, 0xce, 0x2f, 0xe1, 0xe2, 0xf3, 0x4e, 0x1a, 0x60, 0x5c, 0xcb, 0xbd, 0xae, 0x6c,
	0x21, 0x9b, 0xbd, 0x6e, 0xd5, 0xbb, 0xa3, 0xba, 0x55, 0x46, 0x91, 0x89, 0x5f, 0xae, 0xe6, 0x42,
	0xb5, 0xc0, 0x1a, 0x30, 0x47, 0x77, 0xf3, 0x89, 0xf0, 0xd9, 0x8d, 0x96, 0x49, 0xc8, 0x27, 0x3e,
	0x9e, 0xc2, 0xab, 0x85, 0xce, 0x90, 0x15, 0xfe, 0x77, 0xa4, 0x3c, 0xe4, 0x11, 0x5a, 0xa3, 0x78,
	0x9a, 0x0e, 0x33, 0x39, 0xc5, 0x63, 0xdb, 0x56, 0xd1, 0x68, 0x8d, 0xe2, 0xff, 0xa9, 0x97, 0xe4,
	0x02, 0xcf, 0x6a, 0x30, 0xa5, 0xca, 0xa3, 0xf1, 0x1a, 0xc5, 0xf3, 0xab, 0x05, 0xf5, 0x0f, 0xd1,
	0x7b, 0x63, 0xf4, 0xa3, 0x85, 0xd6, 0x3d, 0x3b, 0x0d, 0x69, 0x1f, 0x21, 0x04, 0x4f, 0x32, 0x95,
	0xbb, 0x68, 0x12, 0xfa, 0x41, 0x93, 0x25, 0x9e, 0x09, 0xa5, 0x0e, 0x12, 0xa2, 0x69, 0x70, 0xfb,
	0x69, 0x73, 0x87, 0x97, 0xbf, 0xff, 0xd7, 0x69, 0xd5, 0x74, 0x30, 0x50, 0xd0, 0x9f, 0x94, 0xd1,
	0x4f, 0xca, 0xed, 0x0b, 0x5e, 0x08, 0x55, 0xd3, 0x0f, 0x5e, 0xf5, 0xe7, 0xa0, 0xfe, 0x1a, 0x5b,
	0x56, 0x48, 0x53, 0xda, 0x8c, 0x0a, 0x55, 0xb3, 0xdc, 0x8a, 0x83, 0x28, 0xb9, 0x19, 0xc4, 0x65,
	0xc1, 0x0d, 0xbc, 0x71, 0xc7, 0x42, 0x81, 0xf9, 0xc2, 0xd7, 0xe8, 0x64, 0xcb, 0x2b, 0xb7, 0x4b,
	0xbc, 0xb3, 0x7b, 0x92, 0x06, 0xb2, 0x59, 0xd8, 0x5e, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xed,
	0x94, 0x86, 0xb9, 0x7c, 0x01, 0x00, 0x00,
}
