// Code generated by protoc-gen-go.
// source: sample.proto
// DO NOT EDIT!

/*
Package pingpong is a generated protocol buffer package.

It is generated from these files:
	sample.proto

It has these top-level messages:
	PingRequest
	PingResponse
*/
package pingpong

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type PingRequest struct {
	Ping string `protobuf:"bytes,1,opt,name=ping" json:"ping,omitempty"`
	Pong string `protobuf:"bytes,2,opt,name=pong" json:"pong,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingRequest) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

func (m *PingRequest) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

type PingResponse struct {
	Ping string `protobuf:"bytes,1,opt,name=ping" json:"ping,omitempty"`
	Pong string `protobuf:"bytes,2,opt,name=pong" json:"pong,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingResponse) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

func (m *PingResponse) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "pingpong.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "pingpong.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pingpong service

type PingpongClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type pingpongClient struct {
	cc *grpc.ClientConn
}

func NewPingpongClient(cc *grpc.ClientConn) PingpongClient {
	return &pingpongClient{cc}
}

func (c *pingpongClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/pingpong.pingpong/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pingpong service

type PingpongServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterPingpongServer(s *grpc.Server, srv PingpongServer) {
	s.RegisterService(&_Pingpong_serviceDesc, srv)
}

func _Pingpong_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingpongServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pingpong.pingpong/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingpongServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pingpong_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pingpong.pingpong",
	HandlerType: (*PingpongServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Pingpong_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sample.proto",
}

func init() { proto.RegisterFile("sample.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0xc8, 0xcc, 0x4b, 0x2f, 0xc8,
	0xcf, 0x4b, 0x57, 0x32, 0xe5, 0xe2, 0x0e, 0xc8, 0xcc, 0x4b, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x12, 0xe2, 0x62, 0x01, 0x49, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9,
	0x60, 0xb1, 0xfc, 0xbc, 0x74, 0x09, 0x26, 0xa8, 0x18, 0x48, 0x9b, 0x19, 0x17, 0x0f, 0x44, 0x5b,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0xb1, 0xfa, 0x8c, 0x9c, 0xb9, 0xe0, 0x56, 0x0b, 0x99, 0x73,
	0xb1, 0x80, 0xcc, 0x10, 0x12, 0xd5, 0x83, 0x09, 0xe9, 0x21, 0x39, 0x45, 0x4a, 0x0c, 0x5d, 0x18,
	0x62, 0x95, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x13, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d,
	0x9f, 0x78, 0x4b, 0xd4, 0x00, 0x00, 0x00,
}