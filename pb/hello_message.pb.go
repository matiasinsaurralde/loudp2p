// Code generated by protoc-gen-go.
// source: hello_message.proto
// DO NOT EDIT!

package loud

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HelloMessage struct {
	Origin  string `protobuf:"bytes,1,opt,name=origin" json:"origin,omitempty"`
	RpcPort int64  `protobuf:"varint,2,opt,name=rpc_port,json=rpcPort" json:"rpc_port,omitempty"`
}

func (m *HelloMessage) Reset()                    { *m = HelloMessage{} }
func (m *HelloMessage) String() string            { return proto.CompactTextString(m) }
func (*HelloMessage) ProtoMessage()               {}
func (*HelloMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*HelloMessage)(nil), "loud.HelloMessage")
}

func init() { proto.RegisterFile("hello_message.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0x8f, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0xc9, 0xc9, 0x2f, 0x4d, 0x51, 0x72, 0xe4, 0xe2, 0xf1, 0x00, 0x49, 0xfa, 0x42, 0xe4, 0x84,
	0xc4, 0xb8, 0xd8, 0xf2, 0x8b, 0x32, 0xd3, 0x33, 0xf3, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0xa0, 0x3c, 0x21, 0x49, 0x2e, 0x8e, 0xa2, 0x82, 0xe4, 0xf8, 0x82, 0xfc, 0xa2, 0x12, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0xe6, 0x20, 0xf6, 0xa2, 0x82, 0xe4, 0x80, 0xfc, 0xa2, 0x92, 0x24, 0x36, 0xb0,
	0x79, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0x4f, 0xa0, 0x76, 0x66, 0x00, 0x00, 0x00,
}
