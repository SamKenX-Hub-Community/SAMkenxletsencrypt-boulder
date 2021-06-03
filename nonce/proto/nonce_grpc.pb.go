// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NonceServiceClient is the client API for NonceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NonceServiceClient interface {
	Nonce(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NonceMessage, error)
	Redeem(ctx context.Context, in *NonceMessage, opts ...grpc.CallOption) (*ValidMessage, error)
}

type nonceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNonceServiceClient(cc grpc.ClientConnInterface) NonceServiceClient {
	return &nonceServiceClient{cc}
}

func (c *nonceServiceClient) Nonce(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NonceMessage, error) {
	out := new(NonceMessage)
	err := c.cc.Invoke(ctx, "/nonce.NonceService/Nonce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nonceServiceClient) Redeem(ctx context.Context, in *NonceMessage, opts ...grpc.CallOption) (*ValidMessage, error) {
	out := new(ValidMessage)
	err := c.cc.Invoke(ctx, "/nonce.NonceService/Redeem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NonceServiceServer is the server API for NonceService service.
// All implementations must embed UnimplementedNonceServiceServer
// for forward compatibility
type NonceServiceServer interface {
	Nonce(context.Context, *emptypb.Empty) (*NonceMessage, error)
	Redeem(context.Context, *NonceMessage) (*ValidMessage, error)
	mustEmbedUnimplementedNonceServiceServer()
}

// UnimplementedNonceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNonceServiceServer struct {
}

func (UnimplementedNonceServiceServer) Nonce(context.Context, *emptypb.Empty) (*NonceMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nonce not implemented")
}
func (UnimplementedNonceServiceServer) Redeem(context.Context, *NonceMessage) (*ValidMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Redeem not implemented")
}
func (UnimplementedNonceServiceServer) mustEmbedUnimplementedNonceServiceServer() {}

// UnsafeNonceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NonceServiceServer will
// result in compilation errors.
type UnsafeNonceServiceServer interface {
	mustEmbedUnimplementedNonceServiceServer()
}

func RegisterNonceServiceServer(s grpc.ServiceRegistrar, srv NonceServiceServer) {
	s.RegisterService(&NonceService_ServiceDesc, srv)
}

func _NonceService_Nonce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NonceServiceServer).Nonce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nonce.NonceService/Nonce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NonceServiceServer).Nonce(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NonceService_Redeem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonceMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NonceServiceServer).Redeem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nonce.NonceService/Redeem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NonceServiceServer).Redeem(ctx, req.(*NonceMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// NonceService_ServiceDesc is the grpc.ServiceDesc for NonceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NonceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nonce.NonceService",
	HandlerType: (*NonceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Nonce",
			Handler:    _NonceService_Nonce_Handler,
		},
		{
			MethodName: "Redeem",
			Handler:    _NonceService_Redeem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nonce.proto",
}
