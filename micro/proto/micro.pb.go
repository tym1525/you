// Code generated by protoc-gen-go. DO NOT EDIT.
// source: micro.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	micro.proto

It has these top-level messages:
	Testrequest
	Resposen
*/
package pb

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

type Testrequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
}

func (m *Testrequest) Reset()                    { *m = Testrequest{} }
func (m *Testrequest) String() string            { return proto.CompactTextString(m) }
func (*Testrequest) ProtoMessage()               {}
func (*Testrequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Testrequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Testrequest) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type Resposen struct {
	Ret int32 `protobuf:"varint,1,opt,name=ret" json:"ret,omitempty"`
}

func (m *Resposen) Reset()                    { *m = Resposen{} }
func (m *Resposen) String() string            { return proto.CompactTextString(m) }
func (*Resposen) ProtoMessage()               {}
func (*Resposen) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Resposen) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func init() {
	proto.RegisterType((*Testrequest)(nil), "pb.Testrequest")
	proto.RegisterType((*Resposen)(nil), "pb.Resposen")
}

func init() { proto.RegisterFile("micro.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcd, 0x4c, 0x2e,
	0xca, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x32, 0xe6, 0xe2, 0x0e,
	0x49, 0x2d, 0x2e, 0x29, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b,
	0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x13,
	0xd3, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0x40, 0x4c, 0x25, 0x19, 0x2e, 0x8e, 0xa0,
	0xd4, 0xe2, 0x82, 0xfc, 0xe2, 0xd4, 0x3c, 0x90, 0x6c, 0x51, 0x6a, 0x09, 0x58, 0x03, 0x6b, 0x10,
	0x88, 0x69, 0x64, 0xc3, 0xc5, 0x1d, 0x5a, 0x9c, 0x5a, 0x14, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c,
	0x2a, 0xa4, 0xcb, 0xc5, 0x05, 0xb2, 0xc1, 0xb7, 0xd2, 0xad, 0x34, 0x2f, 0x59, 0x88, 0x5f, 0xaf,
	0x20, 0x49, 0x0f, 0xc9, 0x46, 0x29, 0x1e, 0x90, 0x00, 0xcc, 0x34, 0x25, 0x86, 0x24, 0x36, 0xb0,
	0xdb, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xa9, 0x62, 0x96, 0xaa, 0x00, 0x00, 0x00,
}
