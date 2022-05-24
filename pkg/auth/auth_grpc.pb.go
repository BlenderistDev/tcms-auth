// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auth

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

// TcmsAuthClient is the client API for TcmsAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TcmsAuthClient interface {
	Register(ctx context.Context, in *AuthData, opts ...grpc.CallOption) (*Result, error)
	Login(ctx context.Context, in *AuthData, opts ...grpc.CallOption) (*LoginResult, error)
	CheckAuth(ctx context.Context, in *LoginResult, opts ...grpc.CallOption) (*CheckAuthResult, error)
	TelegramAuth(ctx context.Context, in *TelegramAuthRequest, opts ...grpc.CallOption) (*TelegramAuthResponse, error)
	TelegramSign(ctx context.Context, in *TelegramSignRequest, opts ...grpc.CallOption) (*TelegramAuthResponse, error)
}

type tcmsAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewTcmsAuthClient(cc grpc.ClientConnInterface) TcmsAuthClient {
	return &tcmsAuthClient{cc}
}

func (c *tcmsAuthClient) Register(ctx context.Context, in *AuthData, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/auth.TcmsAuth/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tcmsAuthClient) Login(ctx context.Context, in *AuthData, opts ...grpc.CallOption) (*LoginResult, error) {
	out := new(LoginResult)
	err := c.cc.Invoke(ctx, "/auth.TcmsAuth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tcmsAuthClient) CheckAuth(ctx context.Context, in *LoginResult, opts ...grpc.CallOption) (*CheckAuthResult, error) {
	out := new(CheckAuthResult)
	err := c.cc.Invoke(ctx, "/auth.TcmsAuth/CheckAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tcmsAuthClient) TelegramAuth(ctx context.Context, in *TelegramAuthRequest, opts ...grpc.CallOption) (*TelegramAuthResponse, error) {
	out := new(TelegramAuthResponse)
	err := c.cc.Invoke(ctx, "/auth.TcmsAuth/TelegramAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tcmsAuthClient) TelegramSign(ctx context.Context, in *TelegramSignRequest, opts ...grpc.CallOption) (*TelegramAuthResponse, error) {
	out := new(TelegramAuthResponse)
	err := c.cc.Invoke(ctx, "/auth.TcmsAuth/TelegramSign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TcmsAuthServer is the server API for TcmsAuth service.
// All implementations must embed UnimplementedTcmsAuthServer
// for forward compatibility
type TcmsAuthServer interface {
	Register(context.Context, *AuthData) (*Result, error)
	Login(context.Context, *AuthData) (*LoginResult, error)
	CheckAuth(context.Context, *LoginResult) (*CheckAuthResult, error)
	TelegramAuth(context.Context, *TelegramAuthRequest) (*TelegramAuthResponse, error)
	TelegramSign(context.Context, *TelegramSignRequest) (*TelegramAuthResponse, error)
	mustEmbedUnimplementedTcmsAuthServer()
}

// UnimplementedTcmsAuthServer must be embedded to have forward compatible implementations.
type UnimplementedTcmsAuthServer struct {
}

func (UnimplementedTcmsAuthServer) Register(context.Context, *AuthData) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedTcmsAuthServer) Login(context.Context, *AuthData) (*LoginResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedTcmsAuthServer) CheckAuth(context.Context, *LoginResult) (*CheckAuthResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuth not implemented")
}
func (UnimplementedTcmsAuthServer) TelegramAuth(context.Context, *TelegramAuthRequest) (*TelegramAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TelegramAuth not implemented")
}
func (UnimplementedTcmsAuthServer) TelegramSign(context.Context, *TelegramSignRequest) (*TelegramAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TelegramSign not implemented")
}
func (UnimplementedTcmsAuthServer) mustEmbedUnimplementedTcmsAuthServer() {}

// UnsafeTcmsAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TcmsAuthServer will
// result in compilation errors.
type UnsafeTcmsAuthServer interface {
	mustEmbedUnimplementedTcmsAuthServer()
}

func RegisterTcmsAuthServer(s grpc.ServiceRegistrar, srv TcmsAuthServer) {
	s.RegisterService(&TcmsAuth_ServiceDesc, srv)
}

func _TcmsAuth_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TcmsAuthServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.TcmsAuth/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TcmsAuthServer).Register(ctx, req.(*AuthData))
	}
	return interceptor(ctx, in, info, handler)
}

func _TcmsAuth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TcmsAuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.TcmsAuth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TcmsAuthServer).Login(ctx, req.(*AuthData))
	}
	return interceptor(ctx, in, info, handler)
}

func _TcmsAuth_CheckAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TcmsAuthServer).CheckAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.TcmsAuth/CheckAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TcmsAuthServer).CheckAuth(ctx, req.(*LoginResult))
	}
	return interceptor(ctx, in, info, handler)
}

func _TcmsAuth_TelegramAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelegramAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TcmsAuthServer).TelegramAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.TcmsAuth/TelegramAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TcmsAuthServer).TelegramAuth(ctx, req.(*TelegramAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TcmsAuth_TelegramSign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelegramSignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TcmsAuthServer).TelegramSign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.TcmsAuth/TelegramSign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TcmsAuthServer).TelegramSign(ctx, req.(*TelegramSignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TcmsAuth_ServiceDesc is the grpc.ServiceDesc for TcmsAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TcmsAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.TcmsAuth",
	HandlerType: (*TcmsAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _TcmsAuth_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _TcmsAuth_Login_Handler,
		},
		{
			MethodName: "CheckAuth",
			Handler:    _TcmsAuth_CheckAuth_Handler,
		},
		{
			MethodName: "TelegramAuth",
			Handler:    _TcmsAuth_TelegramAuth_Handler,
		},
		{
			MethodName: "TelegramSign",
			Handler:    _TcmsAuth_TelegramSign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth.proto",
}
