// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google.golang.org/appengine/internal/base/api_base.proto

package base

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StringProto struct {
	Value                *string  `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringProto) Reset()         { *m = StringProto{} }
func (m *StringProto) String() string { return proto.CompactTextString(m) }
func (*StringProto) ProtoMessage()    {}
func (*StringProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_base_9d49f8792e0c1140, []int{0}
}
func (m *StringProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringProto.Unmarshal(m, b)
}
func (m *StringProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringProto.Marshal(b, m, deterministic)
}
func (dst *StringProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringProto.Merge(dst, src)
}
func (m *StringProto) XXX_Size() int {
	return xxx_messageInfo_StringProto.Size(m)
}
func (m *StringProto) XXX_Discardmyuser() {
	xxx_messageInfo_StringProto.Discardmyuser(m)
}

var xxx_messageInfo_StringProto proto.InternalMessageInfo

func (m *StringProto) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Integer32Proto struct {
	Value                *int32   `protobuf:"varint,1,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Integer32Proto) Reset()         { *m = Integer32Proto{} }
func (m *Integer32Proto) String() string { return proto.CompactTextString(m) }
func (*Integer32Proto) ProtoMessage()    {}
func (*Integer32Proto) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_base_9d49f8792e0c1140, []int{1}
}
func (m *Integer32Proto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Integer32Proto.Unmarshal(m, b)
}
func (m *Integer32Proto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Integer32Proto.Marshal(b, m, deterministic)
}
func (dst *Integer32Proto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Integer32Proto.Merge(dst, src)
}
func (m *Integer32Proto) XXX_Size() int {
	return xxx_messageInfo_Integer32Proto.Size(m)
}
func (m *Integer32Proto) XXX_Discardmyuser() {
	xxx_messageInfo_Integer32Proto.Discardmyuser(m)
}

var xxx_messageInfo_Integer32Proto proto.InternalMessageInfo

func (m *Integer32Proto) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Integer64Proto struct {
	Value                *int64   `protobuf:"varint,1,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Integer64Proto) Reset()         { *m = Integer64Proto{} }
func (m *Integer64Proto) String() string { return proto.CompactTextString(m) }
func (*Integer64Proto) ProtoMessage()    {}
func (*Integer64Proto) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_base_9d49f8792e0c1140, []int{2}
}
func (m *Integer64Proto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Integer64Proto.Unmarshal(m, b)
}
func (m *Integer64Proto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Integer64Proto.Marshal(b, m, deterministic)
}
func (dst *Integer64Proto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Integer64Proto.Merge(dst, src)
}
func (m *Integer64Proto) XXX_Size() int {
	return xxx_messageInfo_Integer64Proto.Size(m)
}
func (m *Integer64Proto) XXX_Discardmyuser() {
	xxx_messageInfo_Integer64Proto.Discardmyuser(m)
}

var xxx_messageInfo_Integer64Proto proto.InternalMessageInfo

func (m *Integer64Proto) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type BoolProto struct {
	Value                *bool    `protobuf:"varint,1,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolProto) Reset()         { *m = BoolProto{} }
func (m *BoolProto) String() string { return proto.CompactTextString(m) }
func (*BoolProto) ProtoMessage()    {}
func (*BoolProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_base_9d49f8792e0c1140, []int{3}
}
func (m *BoolProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolProto.Unmarshal(m, b)
}
func (m *BoolProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolProto.Marshal(b, m, deterministic)
}
func (dst *BoolProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolProto.Merge(dst, src)
}
func (m *BoolProto) XXX_Size() int {
	return xxx_messageInfo_BoolProto.Size(m)
}
func (m *BoolProto) XXX_Discardmyuser() {
	xxx_messageInfo_BoolProto.Discardmyuser(m)
}

var xxx_messageInfo_BoolProto proto.InternalMessageInfo

func (m *BoolProto) GetValue() bool {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return false
}

type DoubleProto struct {
	Value                *float64 `protobuf:"fixed64,1,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleProto) Reset()         { *m = DoubleProto{} }
func (m *DoubleProto) String() string { return proto.CompactTextString(m) }
func (*DoubleProto) ProtoMessage()    {}
func (*DoubleProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_base_9d49f8792e0c1140, []int{4}
}
func (m *DoubleProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleProto.Unmarshal(m, b)
}
func (m *DoubleProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleProto.Marshal(b, m, deterministic)
}
func (dst *DoubleProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleProto.Merge(dst, src)
}
func (m *DoubleProto) XXX_Size() int {
	return xxx_messageInfo_DoubleProto.Size(m)
}
func (m *DoubleProto) XXX_Discardmyuser() {
	xxx_messageInfo_DoubleProto.Discardmyuser(m)
}

var xxx_messageInfo_DoubleProto proto.InternalMessageInfo

func (m *DoubleProto) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type BytesProto struct {
	Value                []byte   `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesProto) Reset()         { *m = BytesProto{} }
func (m *BytesProto) String() string { return proto.CompactTextString(m) }
func (*BytesProto) ProtoMessage()    {}
func (*BytesProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_base_9d49f8792e0c1140, []int{5}
}
func (m *BytesProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesProto.Unmarshal(m, b)
}
func (m *BytesProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesProto.Marshal(b, m, deterministic)
}
func (dst *BytesProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesProto.Merge(dst, src)
}
func (m *BytesProto) XXX_Size() int {
	return xxx_messageInfo_BytesProto.Size(m)
}
func (m *BytesProto) XXX_Discardmyuser() {
	xxx_messageInfo_BytesProto.Discardmyuser(m)
}

var xxx_messageInfo_BytesProto proto.InternalMessageInfo

func (m *BytesProto) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type VoidProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoidProto) Reset()         { *m = VoidProto{} }
func (m *VoidProto) String() string { return proto.CompactTextString(m) }
func (*VoidProto) ProtoMessage()    {}
func (*VoidProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_base_9d49f8792e0c1140, []int{6}
}
func (m *VoidProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoidProto.Unmarshal(m, b)
}
func (m *VoidProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoidProto.Marshal(b, m, deterministic)
}
func (dst *VoidProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoidProto.Merge(dst, src)
}
func (m *VoidProto) XXX_Size() int {
	return xxx_messageInfo_VoidProto.Size(m)
}
func (m *VoidProto) XXX_Discardmyuser() {
	xxx_messageInfo_VoidProto.Discardmyuser(m)
}

var xxx_messageInfo_VoidProto proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StringProto)(nil), "appengine.base.StringProto")
	proto.RegisterType((*Integer32Proto)(nil), "appengine.base.Integer32Proto")
	proto.RegisterType((*Integer64Proto)(nil), "appengine.base.Integer64Proto")
	proto.RegisterType((*BoolProto)(nil), "appengine.base.BoolProto")
	proto.RegisterType((*DoubleProto)(nil), "appengine.base.DoubleProto")
	proto.RegisterType((*BytesProto)(nil), "appengine.base.BytesProto")
	proto.RegisterType((*VoidProto)(nil), "appengine.base.VoidProto")
}

func init() {
	proto.RegisterFile("google.golang.org/appengine/internal/base/api_base.proto", fileDescriptor_api_base_9d49f8792e0c1140)
}

var fileDescriptor_api_base_9d49f8792e0c1140 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xcf, 0x3f, 0x4b, 0xc6, 0x30,
	0x10, 0x06, 0x70, 0x5a, 0xad, 0xb4, 0x57, 0xe9, 0x20, 0x0e, 0x1d, 0xb5, 0x05, 0x71, 0x4a, 0x40,
	0x45, 0x9c, 0x83, 0x8b, 0x9b, 0x28, 0x38, 0xb8, 0x48, 0x8a, 0xc7, 0x11, 0x08, 0xb9, 0x90, 0xa6,
	0x82, 0xdf, 0x5e, 0xda, 0xd2, 0xfa, 0xc2, 0x9b, 0xed, 0xfe, 0xfc, 0xe0, 0xe1, 0x81, 0x27, 0x62,
	0x26, 0x8b, 0x82, 0xd8, 0x6a, 0x47, 0x82, 0x03, 0x49, 0xed, 0x3d, 0x3a, 0x32, 0x0e, 0xa5, 0x71,
	0x11, 0x83, 0xd3, 0x56, 0x0e, 0x7a, 0x44, 0xa9, 0xbd, 0xf9, 0x9a, 0x07, 0xe1, 0x03, 0x47, 0xbe,
	0x68, 0x76, 0x27, 0xe6, 0x6b, 0xd7, 0x43, 0xfd, 0x1e, 0x83, 0x71, 0xf4, 0xba, 0xbc, 0x2f, 0xa1,
	0xf8, 0xd1, 0x76, 0xc2, 0x36, 0xbb, 0xca, 0x6f, 0xab, 0xb7, 0x75, 0xe9, 0x6e, 0xa0, 0x79, 0x71,
	0x11, 0x09, 0xc3, 0xfd, 0x5d, 0xc2, 0x15, 0xc7, 0xee, 0xf1, 0x21, 0xe1, 0x4e, 0x36, 0x77, 0x0d,
	0x95, 0x62, 0xb6, 0x09, 0x52, 0x6e, 0xa4, 0x87, 0xfa, 0x99, 0xa7, 0xc1, 0x62, 0x02, 0x65, 0xff,
	0x79, 0xa0, 0x7e, 0x23, 0x8e, 0xab, 0x69, 0x0f, 0xcd, 0xb9, 0xca, 0xcb, 0xdd, 0xd5, 0x50, 0x7d,
	0xb0, 0xf9, 0x5e, 0x98, 0x3a, 0xfb, 0x3c, 0x9d, 0x9b, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xba,
	0x37, 0x25, 0xea, 0x44, 0x01, 0x00, 0x00,
}
