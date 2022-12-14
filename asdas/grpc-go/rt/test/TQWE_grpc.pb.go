// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package test

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TQWEClient is the client API for TQWE service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TQWEClient interface {
	Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error)
	Helloworld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type tQWEClient struct {
	cc grpc.ClientConnInterface
}

func NewTQWEClient(cc grpc.ClientConnInterface) TQWEClient {
	return &tQWEClient{cc}
}

func (c *tQWEClient) Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error) {
	out := new(HealthcheckResponse)
	err := c.cc.Invoke(ctx, "/rt.test.TQWE/Heathcheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tQWEClient) Helloworld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/rt.test.TQWE/Helloworld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TQWEServer is the server API for TQWE service.
// All implementations must embed UnimplementedTQWEServer
// for forward compatibility
type TQWEServer interface {
	Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error)
	Helloworld(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedTQWEServer()
}

// UnimplementedTQWEServer must be embedded to have forward compatible implementations.
type UnimplementedTQWEServer struct {
}

func (UnimplementedTQWEServer) Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heathcheck not implemented")
}
func (UnimplementedTQWEServer) Helloworld(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Helloworld not implemented")
}
func (UnimplementedTQWEServer) mustEmbedUnimplementedTQWEServer() {}

// UnsafeTQWEServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TQWEServer will
// result in compilation errors.
type UnsafeTQWEServer interface {
	mustEmbedUnimplementedTQWEServer()
}

func RegisterTQWEServer(s grpc.ServiceRegistrar, srv TQWEServer) {
	s.RegisterService(&TQWE_ServiceDesc, srv)
}

func _TQWE_Heathcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthcheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TQWEServer).Heathcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rt.test.TQWE/Heathcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TQWEServer).Heathcheck(ctx, req.(*HealthcheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TQWE_Helloworld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TQWEServer).Helloworld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rt.test.TQWE/Helloworld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TQWEServer).Helloworld(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TQWE_ServiceDesc is the grpc.ServiceDesc for TQWE service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TQWE_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rt.test.TQWE",
	HandlerType: (*TQWEServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heathcheck",
			Handler:    _TQWE_Heathcheck_Handler,
		},
		{
			MethodName: "Helloworld",
			Handler:    _TQWE_Helloworld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "TQWE.proto",
}
