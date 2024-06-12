// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.6
// source: fastrunner.proto

package grpcgen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	FastRunner_RunGoLang_FullMethodName = "/fastrunner.FastRunner/RunGoLang"
)

// FastRunnerClient is the client API for FastRunner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FastRunnerClient interface {
	RunGoLang(ctx context.Context, in *GoRequest, opts ...grpc.CallOption) (*RunResponse, error)
}

type fastRunnerClient struct {
	cc grpc.ClientConnInterface
}

func NewFastRunnerClient(cc grpc.ClientConnInterface) FastRunnerClient {
	return &fastRunnerClient{cc}
}

func (c *fastRunnerClient) RunGoLang(ctx context.Context, in *GoRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RunResponse)
	err := c.cc.Invoke(ctx, FastRunner_RunGoLang_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FastRunnerServer is the server API for FastRunner service.
// All implementations must embed UnimplementedFastRunnerServer
// for forward compatibility
type FastRunnerServer interface {
	RunGoLang(context.Context, *GoRequest) (*RunResponse, error)
	mustEmbedUnimplementedFastRunnerServer()
}

// UnimplementedFastRunnerServer must be embedded to have forward compatible implementations.
type UnimplementedFastRunnerServer struct {
}

func (UnimplementedFastRunnerServer) RunGoLang(context.Context, *GoRequest) (*RunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunGoLang not implemented")
}
func (UnimplementedFastRunnerServer) mustEmbedUnimplementedFastRunnerServer() {}

// UnsafeFastRunnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FastRunnerServer will
// result in compilation errors.
type UnsafeFastRunnerServer interface {
	mustEmbedUnimplementedFastRunnerServer()
}

func RegisterFastRunnerServer(s grpc.ServiceRegistrar, srv FastRunnerServer) {
	s.RegisterService(&FastRunner_ServiceDesc, srv)
}

func _FastRunner_RunGoLang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FastRunnerServer).RunGoLang(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FastRunner_RunGoLang_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FastRunnerServer).RunGoLang(ctx, req.(*GoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FastRunner_ServiceDesc is the grpc.ServiceDesc for FastRunner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FastRunner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fastrunner.FastRunner",
	HandlerType: (*FastRunnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunGoLang",
			Handler:    _FastRunner_RunGoLang_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fastrunner.proto",
}