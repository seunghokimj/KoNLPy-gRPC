// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mecab.proto

package konlpy_v0alpha

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("mecab.proto", fileDescriptor_82861a7a195e55ea) }

var fileDescriptor_82861a7a195e55ea = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x4d, 0x4d, 0x4e,
	0x4c, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0xce, 0xcf, 0xcb, 0x29, 0xa8, 0xd4,
	0x2b, 0x33, 0x48, 0xcc, 0x29, 0xc8, 0x48, 0x94, 0xe2, 0x49, 0xcf, 0xc9, 0x4f, 0x4a, 0xcc, 0x81,
	0xc8, 0x1a, 0xbd, 0x61, 0xe4, 0x62, 0xf5, 0x05, 0xa9, 0x16, 0xf2, 0xe0, 0x62, 0x0e, 0xc8, 0x2f,
	0x16, 0x92, 0xd5, 0x43, 0x55, 0xaf, 0x17, 0x5c, 0x52, 0x94, 0x99, 0x97, 0x1e, 0x94, 0x5a, 0x58,
	0x9a, 0x5a, 0x5c, 0x22, 0xa5, 0x84, 0x2e, 0x1d, 0x52, 0x5a, 0x90, 0x93, 0xea, 0x58, 0x54, 0x94,
	0x58, 0x19, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xe4, 0xcd, 0xc5, 0xea, 0x97, 0x5f,
	0x9a, 0x47, 0xd0, 0x2c, 0x65, 0xec, 0xd2, 0xa8, 0x86, 0xf9, 0x70, 0xb1, 0xf9, 0xe6, 0x17, 0x15,
	0x64, 0x50, 0xc5, 0xb4, 0x24, 0x36, 0xb0, 0xaf, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7b,
	0x6f, 0x05, 0x6e, 0x22, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MecabClient is the client API for Mecab service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MecabClient interface {
	Pos(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*TupleArrayResponse, error)
	Nouns(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringArrayResponse, error)
	Morphs(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringArrayResponse, error)
}

type mecabClient struct {
	cc *grpc.ClientConn
}

func NewMecabClient(cc *grpc.ClientConn) MecabClient {
	return &mecabClient{cc}
}

func (c *mecabClient) Pos(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*TupleArrayResponse, error) {
	out := new(TupleArrayResponse)
	err := c.cc.Invoke(ctx, "/konlpy.v0alpha.Mecab/Pos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mecabClient) Nouns(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringArrayResponse, error) {
	out := new(StringArrayResponse)
	err := c.cc.Invoke(ctx, "/konlpy.v0alpha.Mecab/Nouns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mecabClient) Morphs(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringArrayResponse, error) {
	out := new(StringArrayResponse)
	err := c.cc.Invoke(ctx, "/konlpy.v0alpha.Mecab/Morphs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MecabServer is the server API for Mecab service.
type MecabServer interface {
	Pos(context.Context, *StringRequest) (*TupleArrayResponse, error)
	Nouns(context.Context, *StringRequest) (*StringArrayResponse, error)
	Morphs(context.Context, *StringRequest) (*StringArrayResponse, error)
}

// UnimplementedMecabServer can be embedded to have forward compatible implementations.
type UnimplementedMecabServer struct {
}

func (*UnimplementedMecabServer) Pos(ctx context.Context, req *StringRequest) (*TupleArrayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pos not implemented")
}
func (*UnimplementedMecabServer) Nouns(ctx context.Context, req *StringRequest) (*StringArrayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nouns not implemented")
}
func (*UnimplementedMecabServer) Morphs(ctx context.Context, req *StringRequest) (*StringArrayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Morphs not implemented")
}

func RegisterMecabServer(s *grpc.Server, srv MecabServer) {
	s.RegisterService(&_Mecab_serviceDesc, srv)
}

func _Mecab_Pos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MecabServer).Pos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/konlpy.v0alpha.Mecab/Pos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MecabServer).Pos(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mecab_Nouns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MecabServer).Nouns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/konlpy.v0alpha.Mecab/Nouns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MecabServer).Nouns(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mecab_Morphs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MecabServer).Morphs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/konlpy.v0alpha.Mecab/Morphs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MecabServer).Morphs(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mecab_serviceDesc = grpc.ServiceDesc{
	ServiceName: "konlpy.v0alpha.Mecab",
	HandlerType: (*MecabServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pos",
			Handler:    _Mecab_Pos_Handler,
		},
		{
			MethodName: "Nouns",
			Handler:    _Mecab_Nouns_Handler,
		},
		{
			MethodName: "Morphs",
			Handler:    _Mecab_Morphs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mecab.proto",
}