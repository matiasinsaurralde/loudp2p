// Code generated by protoc-gen-go.
// source: dummy_message.proto
// DO NOT EDIT!

/*
Package loud is a generated protocol buffer package.

It is generated from these files:
	dummy_message.proto
	hello_message.proto
	service.proto

It has these top-level messages:
	DummyMessage
	HelloMessage
*/
package loud

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

type DummyMessage struct {
}

func (m *DummyMessage) Reset()                    { *m = DummyMessage{} }
func (m *DummyMessage) String() string            { return proto.CompactTextString(m) }
func (*DummyMessage) ProtoMessage()               {}
func (*DummyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*DummyMessage)(nil), "loud.DummyMessage")
}

func init() { proto.RegisterFile("dummy_message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 66 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x29, 0xcd, 0xcd,
	0xad, 0x8c, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0xc9, 0xc9, 0x2f, 0x4d, 0x51, 0xe2, 0xe3, 0xe2, 0x71, 0x01, 0x49, 0xfa, 0x42, 0xe4, 0x92,
	0xd8, 0xc0, 0x92, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0xf7, 0xb5, 0x80, 0x33, 0x00,
	0x00, 0x00,
}