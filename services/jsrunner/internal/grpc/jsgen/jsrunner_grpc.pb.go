// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.6
// source: jsrunner.proto

package jsgen

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
	JsRunner_RunJs_FullMethodName = "/jsrunner.JsRunner/RunJs"
)

// JsRunnerClient is the client API for JsRunner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JsRunnerClient interface {
	RunJs(ctx context.Context, in *JsRunRequest, opts ...grpc.CallOption) (*JsRunResponse, error)
}

type jsRunnerClient struct {
	cc grpc.ClientConnInterface
}

func NewJsRunnerClient(cc grpc.ClientConnInterface) JsRunnerClient {
	return &jsRunnerClient{cc}
}

func (c *jsRunnerClient) RunJs(ctx context.Context, in *JsRunRequest, opts ...grpc.CallOption) (*JsRunResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JsRunResponse)
	err := c.cc.Invoke(ctx, JsRunner_RunJs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JsRunnerServer is the server API for JsRunner service.
// All implementations must embed UnimplementedJsRunnerServer
// for forward compatibility
type JsRunnerServer interface {
	RunJs(context.Context, *JsRunRequest) (*JsRunResponse, error)
	mustEmbedUnimplementedJsRunnerServer()
}

// UnimplementedJsRunnerServer must be embedded to have forward compatible implementations.
type UnimplementedJsRunnerServer struct {
}

func (UnimplementedJsRunnerServer) RunJs(context.Context, *JsRunRequest) (*JsRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunJs not implemented")
}
func (UnimplementedJsRunnerServer) mustEmbedUnimplementedJsRunnerServer() {}

// UnsafeJsRunnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JsRunnerServer will
// result in compilation errors.
type UnsafeJsRunnerServer interface {
	mustEmbedUnimplementedJsRunnerServer()
}

func RegisterJsRunnerServer(s grpc.ServiceRegistrar, srv JsRunnerServer) {
	s.RegisterService(&JsRunner_ServiceDesc, srv)
}

func _JsRunner_RunJs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JsRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JsRunnerServer).RunJs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JsRunner_RunJs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JsRunnerServer).RunJs(ctx, req.(*JsRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JsRunner_ServiceDesc is the grpc.ServiceDesc for JsRunner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JsRunner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jsrunner.JsRunner",
	HandlerType: (*JsRunnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunJs",
			Handler:    _JsRunner_RunJs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jsrunner.proto",
}