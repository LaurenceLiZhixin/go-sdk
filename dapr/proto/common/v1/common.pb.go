// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dapr/proto/common/v1/common.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Type of HTTP 1.1 Methods
// RFC 7231: https://tools.ietf.org/html/rfc7231#page-24
type HTTPExtension_Verb int32

const (
	HTTPExtension_NONE    HTTPExtension_Verb = 0
	HTTPExtension_GET     HTTPExtension_Verb = 1
	HTTPExtension_HEAD    HTTPExtension_Verb = 2
	HTTPExtension_POST    HTTPExtension_Verb = 3
	HTTPExtension_PUT     HTTPExtension_Verb = 4
	HTTPExtension_DELETE  HTTPExtension_Verb = 5
	HTTPExtension_CONNECT HTTPExtension_Verb = 6
	HTTPExtension_OPTIONS HTTPExtension_Verb = 7
	HTTPExtension_TRACE   HTTPExtension_Verb = 8
)

var HTTPExtension_Verb_name = map[int32]string{
	0: "NONE",
	1: "GET",
	2: "HEAD",
	3: "POST",
	4: "PUT",
	5: "DELETE",
	6: "CONNECT",
	7: "OPTIONS",
	8: "TRACE",
}

var HTTPExtension_Verb_value = map[string]int32{
	"NONE":    0,
	"GET":     1,
	"HEAD":    2,
	"POST":    3,
	"PUT":     4,
	"DELETE":  5,
	"CONNECT": 6,
	"OPTIONS": 7,
	"TRACE":   8,
}

func (x HTTPExtension_Verb) String() string {
	return proto.EnumName(HTTPExtension_Verb_name, int32(x))
}

func (HTTPExtension_Verb) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c448f683ad359d5, []int{0, 0}
}

// HTTPExtension includes HTTP verb and querystring
// when Dapr runtime delivers HTTP content.
//
// For example, when callers calls http invoke api
// POST http://localhost:3500/v1.0/invoke/<app_id>/method/<method>?query1=value1&query2=value2
//
// Dapr runtime will parse POST as a verb and extract querystring to quersytring map.
type HTTPExtension struct {
	// verb is HTTP verb.
	//
	// This is required.
	Verb HTTPExtension_Verb `protobuf:"varint,1,opt,name=verb,proto3,enum=dapr.proto.common.v1.HTTPExtension_Verb" json:"verb,omitempty"`
	// querystring includes HTTP querystring.
	Querystring          map[string]string `protobuf:"bytes,2,rep,name=querystring,proto3" json:"querystring,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HTTPExtension) Reset()         { *m = HTTPExtension{} }
func (m *HTTPExtension) String() string { return proto.CompactTextString(m) }
func (*HTTPExtension) ProtoMessage()    {}
func (*HTTPExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c448f683ad359d5, []int{0}
}

func (m *HTTPExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPExtension.Unmarshal(m, b)
}
func (m *HTTPExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPExtension.Marshal(b, m, deterministic)
}
func (m *HTTPExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPExtension.Merge(m, src)
}
func (m *HTTPExtension) XXX_Size() int {
	return xxx_messageInfo_HTTPExtension.Size(m)
}
func (m *HTTPExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPExtension.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPExtension proto.InternalMessageInfo

func (m *HTTPExtension) GetVerb() HTTPExtension_Verb {
	if m != nil {
		return m.Verb
	}
	return HTTPExtension_NONE
}

func (m *HTTPExtension) GetQuerystring() map[string]string {
	if m != nil {
		return m.Querystring
	}
	return nil
}

type InvokeRequest struct {
	// method is a method name which will be invoked by caller.
	//
	// This field is required.
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// data conveys bytes value or Protobuf message which caller sent.
	// Dapr treats Any.value as bytes type if Any.type_url is unset.
	//
	// This field is required.
	Data *any.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// content_type is the type of data content.
	//
	// This field is required if data delivers http request body
	// Otherwise, this is optional.
	ContentType string `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// http_extension includes http specific fields if request conveys
	// http-compatible request.
	//
	// This field is optional.
	HttpExtension        *HTTPExtension `protobuf:"bytes,4,opt,name=http_extension,json=httpExtension,proto3" json:"http_extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *InvokeRequest) Reset()         { *m = InvokeRequest{} }
func (m *InvokeRequest) String() string { return proto.CompactTextString(m) }
func (*InvokeRequest) ProtoMessage()    {}
func (*InvokeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c448f683ad359d5, []int{1}
}

func (m *InvokeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeRequest.Unmarshal(m, b)
}
func (m *InvokeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeRequest.Marshal(b, m, deterministic)
}
func (m *InvokeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeRequest.Merge(m, src)
}
func (m *InvokeRequest) XXX_Size() int {
	return xxx_messageInfo_InvokeRequest.Size(m)
}
func (m *InvokeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeRequest proto.InternalMessageInfo

func (m *InvokeRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *InvokeRequest) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InvokeRequest) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *InvokeRequest) GetHttpExtension() *HTTPExtension {
	if m != nil {
		return m.HttpExtension
	}
	return nil
}

type InvokeResponse struct {
	// data conveys the content body of InvokeService response.
	//
	// This field is required.
	Data *any.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// content_type is the type of data content.
	//
	// This field is required.
	ContentType          string   `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvokeResponse) Reset()         { *m = InvokeResponse{} }
func (m *InvokeResponse) String() string { return proto.CompactTextString(m) }
func (*InvokeResponse) ProtoMessage()    {}
func (*InvokeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c448f683ad359d5, []int{2}
}

func (m *InvokeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeResponse.Unmarshal(m, b)
}
func (m *InvokeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeResponse.Marshal(b, m, deterministic)
}
func (m *InvokeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeResponse.Merge(m, src)
}
func (m *InvokeResponse) XXX_Size() int {
	return xxx_messageInfo_InvokeResponse.Size(m)
}
func (m *InvokeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeResponse proto.InternalMessageInfo

func (m *InvokeResponse) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InvokeResponse) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func init() {
	proto.RegisterEnum("dapr.proto.common.v1.HTTPExtension_Verb", HTTPExtension_Verb_name, HTTPExtension_Verb_value)
	proto.RegisterType((*HTTPExtension)(nil), "dapr.proto.common.v1.HTTPExtension")
	proto.RegisterMapType((map[string]string)(nil), "dapr.proto.common.v1.HTTPExtension.QuerystringEntry")
	proto.RegisterType((*InvokeRequest)(nil), "dapr.proto.common.v1.InvokeRequest")
	proto.RegisterType((*InvokeResponse)(nil), "dapr.proto.common.v1.InvokeResponse")
}

func init() { proto.RegisterFile("dapr/proto/common/v1/common.proto", fileDescriptor_0c448f683ad359d5) }

var fileDescriptor_0c448f683ad359d5 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x5f, 0x6f, 0xda, 0x3c,
	0x14, 0xc6, 0xdf, 0x84, 0x00, 0xed, 0xa1, 0x20, 0xcb, 0x42, 0xaf, 0x58, 0x77, 0xd3, 0xb2, 0x1b,
	0xa4, 0x69, 0x8e, 0x60, 0xbb, 0x98, 0xa6, 0x69, 0x12, 0x05, 0xab, 0xed, 0x34, 0x01, 0x4b, 0xbd,
	0x5e, 0x4c, 0x9a, 0xaa, 0x04, 0xbc, 0x10, 0x01, 0x76, 0x9a, 0x38, 0xd1, 0xf2, 0x95, 0xf6, 0x29,
	0xf6, 0x09, 0xf6, 0x99, 0x26, 0x3b, 0xa1, 0xfb, 0xd7, 0x8b, 0xee, 0xee, 0x9c, 0xe3, 0xe7, 0x79,
	0xf4, 0xb3, 0x7d, 0xe0, 0x74, 0xe5, 0xc7, 0x89, 0x1b, 0x27, 0x52, 0x49, 0x77, 0x29, 0x77, 0x3b,
	0x29, 0xdc, 0x7c, 0x58, 0x55, 0xc4, 0x8c, 0x71, 0x57, 0x4b, 0xca, 0x9a, 0x54, 0x07, 0xf9, 0xf0,
	0xf8, 0x51, 0x28, 0x65, 0xb8, 0xe5, 0xa5, 0x35, 0xc8, 0x3e, 0xbb, 0xbe, 0x28, 0x4a, 0x51, 0xff,
	0xbb, 0x0d, 0xed, 0x0b, 0xc6, 0x16, 0xf4, 0x8b, 0xe2, 0x22, 0x8d, 0xa4, 0xc0, 0xaf, 0xc1, 0xc9,
	0x79, 0x12, 0xf4, 0xac, 0x13, 0x6b, 0xd0, 0x19, 0x0d, 0xc8, 0x7d, 0x89, 0xe4, 0x37, 0x0b, 0xb9,
	0xe6, 0x49, 0xe0, 0x19, 0x17, 0xbe, 0x86, 0xd6, 0x6d, 0xc6, 0x93, 0x22, 0x55, 0x49, 0x24, 0xc2,
	0x9e, 0x7d, 0x52, 0x1b, 0xb4, 0x46, 0x2f, 0x1e, 0x12, 0xf2, 0xfe, 0xa7, 0x8d, 0x0a, 0x95, 0x14,
	0xde, 0xaf, 0x41, 0xc7, 0x6f, 0x00, 0xfd, 0x29, 0xc0, 0x08, 0x6a, 0x1b, 0x5e, 0x18, 0xd0, 0x43,
	0x4f, 0x97, 0xb8, 0x0b, 0xf5, 0xdc, 0xdf, 0x66, 0xbc, 0x67, 0x9b, 0x59, 0xd9, 0xbc, 0xb2, 0x5f,
	0x5a, 0xfd, 0x10, 0x1c, 0x4d, 0x89, 0x0f, 0xc0, 0x99, 0xcd, 0x67, 0x14, 0xfd, 0x87, 0x9b, 0x50,
	0x3b, 0xa7, 0x0c, 0x59, 0x7a, 0x74, 0x41, 0xc7, 0x53, 0x64, 0xeb, 0x6a, 0x31, 0xbf, 0x62, 0xa8,
	0xa6, 0x0f, 0x17, 0x1f, 0x18, 0x72, 0x30, 0x40, 0x63, 0x4a, 0xdf, 0x51, 0x46, 0x51, 0x1d, 0xb7,
	0xa0, 0x39, 0x99, 0xcf, 0x66, 0x74, 0xc2, 0x50, 0x43, 0x37, 0xf3, 0x05, 0xbb, 0x9c, 0xcf, 0xae,
	0x50, 0x13, 0x1f, 0x42, 0x9d, 0x79, 0xe3, 0x09, 0x45, 0x07, 0xfd, 0x6f, 0x16, 0xb4, 0x2f, 0x45,
	0x2e, 0x37, 0xdc, 0xe3, 0xb7, 0x19, 0x4f, 0x15, 0xfe, 0x1f, 0x1a, 0x3b, 0xae, 0xd6, 0x72, 0x55,
	0x91, 0x56, 0x1d, 0x1e, 0x80, 0xb3, 0xf2, 0x95, 0x6f, 0x58, 0x5b, 0xa3, 0x2e, 0x29, 0x3f, 0x89,
	0xec, 0x3f, 0x89, 0x8c, 0x45, 0xe1, 0x19, 0x05, 0x3e, 0x85, 0xa3, 0xa5, 0x14, 0x8a, 0x0b, 0x75,
	0xa3, 0x8a, 0x98, 0xf7, 0x6a, 0x26, 0xa7, 0x55, 0xcd, 0x58, 0x11, 0x73, 0xfc, 0x16, 0x3a, 0x6b,
	0xa5, 0xe2, 0x1b, 0xbe, 0x7f, 0xcf, 0x9e, 0x63, 0x62, 0x9f, 0x3c, 0xe0, 0xe9, 0xbd, 0xb6, 0xb6,
	0xde, 0xb5, 0xfd, 0x4f, 0xd0, 0xd9, 0xdf, 0x20, 0x8d, 0xa5, 0x48, 0xf9, 0x1d, 0xaa, 0xf5, 0xcf,
	0xa8, 0xf6, 0x5f, 0xa8, 0x67, 0x1c, 0x20, 0x92, 0x25, 0x56, 0x3e, 0x3c, 0x3b, 0x9a, 0x18, 0xa8,
	0x85, 0x8e, 0x4a, 0x3f, 0x3e, 0x0d, 0x23, 0xb5, 0xce, 0x02, 0x4d, 0xea, 0x9a, 0x6d, 0x0f, 0xe5,
	0xb3, 0x74, 0xb5, 0x71, 0xef, 0xdb, 0xfc, 0xaf, 0xf6, 0xe3, 0xa9, 0x0e, 0x99, 0x6c, 0x23, 0x2e,
	0x14, 0x19, 0x67, 0x4a, 0x86, 0x5c, 0x90, 0xf3, 0x24, 0x5e, 0x92, 0x7c, 0x18, 0x34, 0x8c, 0xfc,
	0xf9, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x8c, 0xec, 0xbf, 0x36, 0x03, 0x00, 0x00,
}
