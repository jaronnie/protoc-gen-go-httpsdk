// Code generated by protoc-gen-go. DO NOT EDIT.
// source: machine.proto

package corev1

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Machine struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Machine) Reset()         { *m = Machine{} }
func (m *Machine) String() string { return proto.CompactTextString(m) }
func (*Machine) ProtoMessage()    {}
func (*Machine) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b4e4a03b74bd47d, []int{0}
}

func (m *Machine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Machine.Unmarshal(m, b)
}
func (m *Machine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Machine.Marshal(b, m, deterministic)
}
func (m *Machine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Machine.Merge(m, src)
}
func (m *Machine) XXX_Size() int {
	return xxx_messageInfo_Machine.Size(m)
}
func (m *Machine) XXX_DiscardUnknown() {
	xxx_messageInfo_Machine.DiscardUnknown(m)
}

var xxx_messageInfo_Machine proto.InternalMessageInfo

func (m *Machine) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Machine) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*Machine)(nil), "proto.Machine")
}

func init() { proto.RegisterFile("machine.proto", fileDescriptor_4b4e4a03b74bd47d) }

var fileDescriptor_4b4e4a03b74bd47d = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x4c, 0xce,
	0xc8, 0xcc, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xba, 0x5c,
	0xec, 0xbe, 0x10, 0x71, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6,
	0x20, 0xa6, 0xcc, 0x14, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0x09, 0x66, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x30, 0xdb, 0x89, 0x2b, 0x8a, 0x43, 0x4f, 0x3f, 0x39, 0xbf, 0x28, 0xb5, 0xcc,
	0x30, 0x89, 0x0d, 0x6c, 0x82, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x20, 0x89, 0x1c, 0x59,
	0x00, 0x00, 0x00,
}
