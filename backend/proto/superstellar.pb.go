// Code generated by protoc-gen-go.
// source: superstellar.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	superstellar.proto

It has these top-level messages:
	IntVector
	Vector
	Spaceship
	Space
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type IntVector struct {
	X int32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
}

func (m *IntVector) Reset()                    { *m = IntVector{} }
func (m *IntVector) String() string            { return proto1.CompactTextString(m) }
func (*IntVector) ProtoMessage()               {}
func (*IntVector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Vector struct {
	X float32 `protobuf:"fixed32,1,opt,name=x" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y" json:"y,omitempty"`
}

func (m *Vector) Reset()                    { *m = Vector{} }
func (m *Vector) String() string            { return proto1.CompactTextString(m) }
func (*Vector) ProtoMessage()               {}
func (*Vector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Spaceship struct {
	Id          uint32     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Position    *IntVector `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
	Velocity    *Vector    `protobuf:"bytes,3,opt,name=velocity" json:"velocity,omitempty"`
	Facing      float32    `protobuf:"fixed32,4,opt,name=facing" json:"facing,omitempty"`
	InputThrust bool       `protobuf:"varint,5,opt,name=input_thrust,json=inputThrust" json:"input_thrust,omitempty"`
}

func (m *Spaceship) Reset()                    { *m = Spaceship{} }
func (m *Spaceship) String() string            { return proto1.CompactTextString(m) }
func (*Spaceship) ProtoMessage()               {}
func (*Spaceship) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Spaceship) GetPosition() *IntVector {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Spaceship) GetVelocity() *Vector {
	if m != nil {
		return m.Velocity
	}
	return nil
}

type Space struct {
	Spaceships []*Spaceship `protobuf:"bytes,1,rep,name=spaceships" json:"spaceships,omitempty"`
}

func (m *Space) Reset()                    { *m = Space{} }
func (m *Space) String() string            { return proto1.CompactTextString(m) }
func (*Space) ProtoMessage()               {}
func (*Space) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Space) GetSpaceships() []*Spaceship {
	if m != nil {
		return m.Spaceships
	}
	return nil
}

func init() {
	proto1.RegisterType((*IntVector)(nil), "superstellar.IntVector")
	proto1.RegisterType((*Vector)(nil), "superstellar.Vector")
	proto1.RegisterType((*Spaceship)(nil), "superstellar.Spaceship")
	proto1.RegisterType((*Space)(nil), "superstellar.Space")
}

func init() { proto1.RegisterFile("superstellar.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x14, 0x84, 0x49, 0xd7, 0xd6, 0xee, 0xdb, 0xea, 0xe1, 0x21, 0x9a, 0xa3, 0x16, 0x41, 0x4f, 0x8b,
	0xb8, 0x07, 0xaf, 0xe2, 0xcd, 0x6b, 0x14, 0x0f, 0x5e, 0xa4, 0xd6, 0xe8, 0x06, 0x4a, 0x12, 0x92,
	0x57, 0xd9, 0xfe, 0x38, 0xff, 0x9b, 0xdd, 0xa7, 0x2d, 0x5b, 0x3d, 0x25, 0x33, 0xf9, 0x06, 0x66,
	0x02, 0x18, 0x5b, 0xaf, 0x43, 0x24, 0xdd, 0x34, 0x55, 0x58, 0xfa, 0xe0, 0xc8, 0x61, 0xb1, 0xeb,
	0x95, 0x17, 0x30, 0xbf, 0xb7, 0xf4, 0xa4, 0x6b, 0x72, 0x01, 0x0b, 0x10, 0x1b, 0x29, 0x4e, 0xc5,
	0x65, 0xaa, 0xc4, 0x66, 0xab, 0x3a, 0x99, 0xfc, 0xa8, 0xae, 0x3c, 0x87, 0xec, 0x2f, 0x95, 0x4c,
	0xa8, 0x64, 0x4b, 0x7d, 0x09, 0x98, 0x3f, 0xf8, 0xaa, 0xd6, 0x71, 0x6d, 0x3c, 0x1e, 0x42, 0x62,
	0xde, 0x18, 0x3d, 0x50, 0xfd, 0x0d, 0x57, 0x90, 0x7b, 0x17, 0x0d, 0x19, 0x67, 0x39, 0xb2, 0xb8,
	0x3e, 0x59, 0x4e, 0x1a, 0x8e, 0x55, 0xd4, 0x08, 0xe2, 0x15, 0xe4, 0x9f, 0xba, 0x71, 0xb5, 0xa1,
	0x4e, 0xce, 0x38, 0x74, 0x34, 0x0d, 0x0d, 0x89, 0x81, 0xc2, 0x63, 0xc8, 0xde, 0xab, 0xda, 0xd8,
	0x0f, 0xb9, 0xc7, 0xbd, 0x7e, 0x15, 0x9e, 0x41, 0x61, 0xac, 0x6f, 0xe9, 0x85, 0xd6, 0xa1, 0x8d,
	0x24, 0xd3, 0xfe, 0x35, 0x57, 0x0b, 0xf6, 0x1e, 0xd9, 0x2a, 0x6f, 0x21, 0xe5, 0xfa, 0x78, 0x03,
	0x10, 0x87, 0x1d, 0xb1, 0x9f, 0x30, 0xfb, 0x5f, 0x76, 0xdc, 0xa9, 0x76, 0xd0, 0xbb, 0xfd, 0xe7,
	0x94, 0xff, 0xf9, 0x35, 0xe3, 0x63, 0xf5, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x75, 0x68, 0xbb, 0x04,
	0x84, 0x01, 0x00, 0x00,
}
