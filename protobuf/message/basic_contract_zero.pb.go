// Code generated by protoc-gen-go. DO NOT EDIT.
// source: basic_contract_zero.proto

package zero

import (
	_ "aelf"
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

func init() { proto.RegisterFile("basic_contract_zero.proto", fileDescriptor_feca40e9e32ac99b) }

var fileDescriptor_feca40e9e32ac99b = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x4a, 0x2c, 0xce,
	0x4c, 0x8e, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0x89, 0xaf, 0x4a, 0x2d, 0xca, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0x84, 0x12, 0x53, 0x73, 0xd2, 0xf4,
	0xf3, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0x32, 0x52, 0x5c, 0x89, 0xc9, 0xc5, 0x06, 0x10,
	0xb6, 0x51, 0x18, 0x97, 0xa0, 0x13, 0xc8, 0x08, 0x67, 0xa8, 0x09, 0x51, 0x20, 0x4d, 0x8e, 0xa7,
	0x4e, 0x7d, 0x63, 0x44, 0x52, 0xb5, 0xe9, 0xcc, 0x37, 0x46, 0x5d, 0x47, 0xd7, 0x9c, 0x34, 0x3d,
	0x98, 0xb2, 0x62, 0x3d, 0xf7, 0xd4, 0xbc, 0xd4, 0xe2, 0xcc, 0x62, 0x3d, 0x0c, 0xdd, 0xc1, 0x25,
	0x89, 0x25, 0xa9, 0x4e, 0x92, 0xab, 0x98, 0xc4, 0xb0, 0xeb, 0x48, 0x62, 0x03, 0x9b, 0x69, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xa2, 0x12, 0xc8, 0xbc, 0x00, 0x00, 0x00,
}