// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ekko_service.proto

package protobuf

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type GetUniqueIDRequest struct {
	Product              int32    `protobuf:"varint,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUniqueIDRequest) Reset()         { *m = GetUniqueIDRequest{} }
func (m *GetUniqueIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetUniqueIDRequest) ProtoMessage()    {}
func (*GetUniqueIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_95c81757a4c2c6af, []int{0}
}

func (m *GetUniqueIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUniqueIDRequest.Unmarshal(m, b)
}
func (m *GetUniqueIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUniqueIDRequest.Marshal(b, m, deterministic)
}
func (m *GetUniqueIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUniqueIDRequest.Merge(m, src)
}
func (m *GetUniqueIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetUniqueIDRequest.Size(m)
}
func (m *GetUniqueIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUniqueIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUniqueIDRequest proto.InternalMessageInfo

func (m *GetUniqueIDRequest) GetProduct() int32 {
	if m != nil {
		return m.Product
	}
	return 0
}

type GetUniqueIDResponse struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUniqueIDResponse) Reset()         { *m = GetUniqueIDResponse{} }
func (m *GetUniqueIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetUniqueIDResponse) ProtoMessage()    {}
func (*GetUniqueIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_95c81757a4c2c6af, []int{1}
}

func (m *GetUniqueIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUniqueIDResponse.Unmarshal(m, b)
}
func (m *GetUniqueIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUniqueIDResponse.Marshal(b, m, deterministic)
}
func (m *GetUniqueIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUniqueIDResponse.Merge(m, src)
}
func (m *GetUniqueIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetUniqueIDResponse.Size(m)
}
func (m *GetUniqueIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUniqueIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUniqueIDResponse proto.InternalMessageInfo

func (m *GetUniqueIDResponse) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type MGetUniqueIDRequest struct {
	Product              int32    `protobuf:"varint,1,opt,name=product,proto3" json:"product,omitempty"`
	Count                uint32   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MGetUniqueIDRequest) Reset()         { *m = MGetUniqueIDRequest{} }
func (m *MGetUniqueIDRequest) String() string { return proto.CompactTextString(m) }
func (*MGetUniqueIDRequest) ProtoMessage()    {}
func (*MGetUniqueIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_95c81757a4c2c6af, []int{2}
}

func (m *MGetUniqueIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MGetUniqueIDRequest.Unmarshal(m, b)
}
func (m *MGetUniqueIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MGetUniqueIDRequest.Marshal(b, m, deterministic)
}
func (m *MGetUniqueIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MGetUniqueIDRequest.Merge(m, src)
}
func (m *MGetUniqueIDRequest) XXX_Size() int {
	return xxx_messageInfo_MGetUniqueIDRequest.Size(m)
}
func (m *MGetUniqueIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MGetUniqueIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MGetUniqueIDRequest proto.InternalMessageInfo

func (m *MGetUniqueIDRequest) GetProduct() int32 {
	if m != nil {
		return m.Product
	}
	return 0
}

func (m *MGetUniqueIDRequest) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type MGetUniqueIDResponse struct {
	LowerUid             uint64   `protobuf:"varint,1,opt,name=lower_uid,json=lowerUid,proto3" json:"lower_uid,omitempty"`
	UpperUid             uint64   `protobuf:"varint,2,opt,name=upper_uid,json=upperUid,proto3" json:"upper_uid,omitempty"`
	Count                uint32   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MGetUniqueIDResponse) Reset()         { *m = MGetUniqueIDResponse{} }
func (m *MGetUniqueIDResponse) String() string { return proto.CompactTextString(m) }
func (*MGetUniqueIDResponse) ProtoMessage()    {}
func (*MGetUniqueIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_95c81757a4c2c6af, []int{3}
}

func (m *MGetUniqueIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MGetUniqueIDResponse.Unmarshal(m, b)
}
func (m *MGetUniqueIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MGetUniqueIDResponse.Marshal(b, m, deterministic)
}
func (m *MGetUniqueIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MGetUniqueIDResponse.Merge(m, src)
}
func (m *MGetUniqueIDResponse) XXX_Size() int {
	return xxx_messageInfo_MGetUniqueIDResponse.Size(m)
}
func (m *MGetUniqueIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MGetUniqueIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MGetUniqueIDResponse proto.InternalMessageInfo

func (m *MGetUniqueIDResponse) GetLowerUid() uint64 {
	if m != nil {
		return m.LowerUid
	}
	return 0
}

func (m *MGetUniqueIDResponse) GetUpperUid() uint64 {
	if m != nil {
		return m.UpperUid
	}
	return 0
}

func (m *MGetUniqueIDResponse) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*GetUniqueIDRequest)(nil), "protobuf.GetUniqueIDRequest")
	proto.RegisterType((*GetUniqueIDResponse)(nil), "protobuf.GetUniqueIDResponse")
	proto.RegisterType((*MGetUniqueIDRequest)(nil), "protobuf.MGetUniqueIDRequest")
	proto.RegisterType((*MGetUniqueIDResponse)(nil), "protobuf.MGetUniqueIDResponse")
}

func init() { proto.RegisterFile("ekko_service.proto", fileDescriptor_95c81757a4c2c6af) }

var fileDescriptor_95c81757a4c2c6af = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xcd, 0xce, 0xce,
	0x8f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x00, 0x53, 0x49, 0xa5, 0x69, 0x4a, 0x7a, 0x5c, 0x42, 0xee, 0xa9, 0x25, 0xa1, 0x79, 0x99, 0x85,
	0xa5, 0xa9, 0x9e, 0x2e, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x12, 0x5c, 0xec, 0x05,
	0x45, 0xf9, 0x29, 0xa5, 0xc9, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x30, 0xae, 0x92,
	0x3a, 0x97, 0x30, 0x8a, 0xfa, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x01, 0x2e, 0xe6, 0xd2,
	0xcc, 0x14, 0xb0, 0x62, 0x96, 0x20, 0x10, 0x53, 0xc9, 0x95, 0x4b, 0xd8, 0x97, 0x14, 0x93, 0x85,
	0x44, 0xb8, 0x58, 0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0x24, 0x98, 0x14, 0x18, 0x35, 0x78, 0x83, 0x20,
	0x1c, 0xa5, 0x34, 0x2e, 0x11, 0x5f, 0x6c, 0x16, 0x4a, 0x73, 0x71, 0xe6, 0xe4, 0x97, 0xa7, 0x16,
	0xc5, 0x23, 0xac, 0xe5, 0x00, 0x0b, 0x84, 0x66, 0xa6, 0x80, 0x24, 0x4b, 0x0b, 0x0a, 0xa0, 0x92,
	0x4c, 0x10, 0x49, 0xb0, 0x00, 0x48, 0x12, 0x6e, 0x0f, 0x33, 0x92, 0x3d, 0x49, 0x6c, 0xe0, 0x10,
	0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x0c, 0x84, 0x04, 0x2e, 0x01, 0x00, 0x00,
}
