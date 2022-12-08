// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.17.3
// source: server.proto

package pb

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

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerClient interface {
	Hello(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HelloResponse, error)
}

type serverClient struct {
	cc grpc.ClientConnInterface
}

func NewServerClient(cc grpc.ClientConnInterface) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Hello(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/Server/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
// All implementations should embed UnimplementedServerServer
// for forward compatibility
type ServerServer interface {
	Hello(context.Context, *Empty) (*HelloResponse, error)
}

// UnimplementedServerServer should be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (UnimplementedServerServer) Hello(context.Context, *Empty) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}

// UnsafeServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerServer will
// result in compilation errors.
type UnsafeServerServer interface {
	mustEmbedUnimplementedServerServer()
}

func RegisterServerServer(s grpc.ServiceRegistrar, srv ServerServer) {
	s.RegisterService(&Server_ServiceDesc, srv)
}

func _Server_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Hello(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Server_ServiceDesc is the grpc.ServiceDesc for Server service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Server_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Server_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}