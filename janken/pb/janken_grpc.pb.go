// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// JankenServiceClient is the client API for JankenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JankenServiceClient interface {
	PlayJanken(ctx context.Context, in *PlayJankenRequest, opts ...grpc.CallOption) (*PlayJankenResponse, error)
	PlayJankenResults(ctx context.Context, in *PlayResultRequest, opts ...grpc.CallOption) (*PlayResultResponse, error)
}

type jankenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJankenServiceClient(cc grpc.ClientConnInterface) JankenServiceClient {
	return &jankenServiceClient{cc}
}

func (c *jankenServiceClient) PlayJanken(ctx context.Context, in *PlayJankenRequest, opts ...grpc.CallOption) (*PlayJankenResponse, error) {
	out := new(PlayJankenResponse)
	err := c.cc.Invoke(ctx, "/janken.JankenService/PlayJanken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jankenServiceClient) PlayJankenResults(ctx context.Context, in *PlayResultRequest, opts ...grpc.CallOption) (*PlayResultResponse, error) {
	out := new(PlayResultResponse)
	err := c.cc.Invoke(ctx, "/janken.JankenService/PlayJankenResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JankenServiceServer is the server API for JankenService service.
// All implementations should embed UnimplementedJankenServiceServer
// for forward compatibility
type JankenServiceServer interface {
	PlayJanken(context.Context, *PlayJankenRequest) (*PlayJankenResponse, error)
	PlayJankenResults(context.Context, *PlayResultRequest) (*PlayResultResponse, error)
}

// UnimplementedJankenServiceServer should be embedded to have forward compatible implementations.
type UnimplementedJankenServiceServer struct {
}

func (UnimplementedJankenServiceServer) PlayJanken(context.Context, *PlayJankenRequest) (*PlayJankenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayJanken not implemented")
}
func (UnimplementedJankenServiceServer) PlayJankenResults(context.Context, *PlayResultRequest) (*PlayResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayJankenResults not implemented")
}

// UnsafeJankenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JankenServiceServer will
// result in compilation errors.
type UnsafeJankenServiceServer interface {
	mustEmbedUnimplementedJankenServiceServer()
}

func RegisterJankenServiceServer(s grpc.ServiceRegistrar, srv JankenServiceServer) {
	s.RegisterService(&JankenService_ServiceDesc, srv)
}

func _JankenService_PlayJanken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayJankenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JankenServiceServer).PlayJanken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/janken.JankenService/PlayJanken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JankenServiceServer).PlayJanken(ctx, req.(*PlayJankenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JankenService_PlayJankenResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JankenServiceServer).PlayJankenResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/janken.JankenService/PlayJankenResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JankenServiceServer).PlayJankenResults(ctx, req.(*PlayResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JankenService_ServiceDesc is the grpc.ServiceDesc for JankenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JankenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "janken.JankenService",
	HandlerType: (*JankenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlayJanken",
			Handler:    _JankenService_PlayJanken_Handler,
		},
		{
			MethodName: "PlayJankenResults",
			Handler:    _JankenService_PlayJankenResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/janken.proto",
}
