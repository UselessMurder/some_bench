// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package tutorial

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

type Msg struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Content              []byte   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_514f7deeb382e468, []int{0}
}
func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (dst *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(dst, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Msg) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Msg) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*Msg)(nil), "tutorial.msg")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_msg_514f7deeb382e468) }

var fileDescriptor_msg_514f7deeb382e468 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x29, 0x2d, 0xc9, 0x2f, 0xca, 0x4c, 0xcc, 0x51,
	0x72, 0xe6, 0x62, 0xce, 0x2d, 0x4e, 0x17, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x98, 0x14,
	0x18, 0x35, 0x78, 0x83, 0x98, 0x32, 0x53, 0x84, 0x24, 0xb8, 0xd8, 0x93, 0xf3, 0xf3, 0x4a, 0x52,
	0xf3, 0x4a, 0x24, 0x98, 0x15, 0x18, 0x35, 0x78, 0x82, 0x60, 0xdc, 0x24, 0x36, 0xb0, 0xa9, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xc2, 0xd9, 0x4d, 0x62, 0x00, 0x00, 0x00,
}
