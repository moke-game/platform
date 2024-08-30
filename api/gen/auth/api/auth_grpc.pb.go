// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: auth/auth.proto

// Auth service is used for token refresh, token validation, custom token packing, etc.
// 认证服务,用于token刷新，token验证，自定义封装token 等功能

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthService_Authenticate_FullMethodName  = "/auth.v1.AuthService/Authenticate"
	AuthService_RefreshToken_FullMethodName  = "/auth.v1.AuthService/RefreshToken"
	AuthService_ValidateToken_FullMethodName = "/auth.v1.AuthService/ValidateToken"
	AuthService_PackToken_FullMethodName     = "/auth.v1.AuthService/PackToken"
	AuthService_ClearToken_FullMethodName    = "/auth.v1.AuthService/ClearToken"
	AuthService_Delete_FullMethodName        = "/auth.v1.AuthService/Delete"
	AuthService_AddBlocked_FullMethodName    = "/auth.v1.AuthService/AddBlocked"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	// Authenticate request a jwt token by id
	// return a jwt token and a refresh token
	// 请求一个 jwt token,用于身份认证
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	// RefreshToken request a new jwt token by refresh token
	// return a jwt token and a refresh token
	// 刷新jwt token 过期时间
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	// ValidateToken validate a jwt token
	// return a uid and data
	// 验证jwt token是否合法
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
	// PacketToken pack some custom data to a jwt token, suitable for game room auth
	// 将一些自定义数据打包到jwt token中,适用于游戏房间验证
	PackToken(ctx context.Context, in *PackTokenRequest, opts ...grpc.CallOption) (*PackTokenResponse, error)
	// ClearToken clear the uid's token
	// 清除uid对应的token
	ClearToken(ctx context.Context, in *ClearTokenRequest, opts ...grpc.CallOption) (*ClearTokenResponse, error)
	// Delete delete the id and uid info from db
	// 删除id对应的uid信息
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// AddBlocked add a uid to block list
	// if is_block is true, will block the uid for duration seconds
	// 添加到黑名单
	AddBlocked(ctx context.Context, in *AddBlockedRequest, opts ...grpc.CallOption) (*AddBlockedResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, AuthService_Authenticate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_RefreshToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_ValidateToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) PackToken(ctx context.Context, in *PackTokenRequest, opts ...grpc.CallOption) (*PackTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PackTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_PackToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ClearToken(ctx context.Context, in *ClearTokenRequest, opts ...grpc.CallOption) (*ClearTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClearTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_ClearToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, AuthService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AddBlocked(ctx context.Context, in *AddBlockedRequest, opts ...grpc.CallOption) (*AddBlockedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddBlockedResponse)
	err := c.cc.Invoke(ctx, AuthService_AddBlocked_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations should embed UnimplementedAuthServiceServer
// for forward compatibility.
type AuthServiceServer interface {
	// Authenticate request a jwt token by id
	// return a jwt token and a refresh token
	// 请求一个 jwt token,用于身份认证
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	// RefreshToken request a new jwt token by refresh token
	// return a jwt token and a refresh token
	// 刷新jwt token 过期时间
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	// ValidateToken validate a jwt token
	// return a uid and data
	// 验证jwt token是否合法
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
	// PacketToken pack some custom data to a jwt token, suitable for game room auth
	// 将一些自定义数据打包到jwt token中,适用于游戏房间验证
	PackToken(context.Context, *PackTokenRequest) (*PackTokenResponse, error)
	// ClearToken clear the uid's token
	// 清除uid对应的token
	ClearToken(context.Context, *ClearTokenRequest) (*ClearTokenResponse, error)
	// Delete delete the id and uid info from db
	// 删除id对应的uid信息
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// AddBlocked add a uid to block list
	// if is_block is true, will block the uid for duration seconds
	// 添加到黑名单
	AddBlocked(context.Context, *AddBlockedRequest) (*AddBlockedResponse, error)
}

// UnimplementedAuthServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedAuthServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthServiceServer) ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedAuthServiceServer) PackToken(context.Context, *PackTokenRequest) (*PackTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PackToken not implemented")
}
func (UnimplementedAuthServiceServer) ClearToken(context.Context, *ClearTokenRequest) (*ClearTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearToken not implemented")
}
func (UnimplementedAuthServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAuthServiceServer) AddBlocked(context.Context, *AddBlockedRequest) (*AddBlockedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlocked not implemented")
}
func (UnimplementedAuthServiceServer) testEmbeddedByValue() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Authenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ValidateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_PackToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).PackToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_PackToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).PackToken(ctx, req.(*PackTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ClearToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ClearToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ClearToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ClearToken(ctx, req.(*ClearTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AddBlocked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlockedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddBlocked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_AddBlocked_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddBlocked(ctx, req.(*AddBlockedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AuthService_Authenticate_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AuthService_RefreshToken_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _AuthService_ValidateToken_Handler,
		},
		{
			MethodName: "PackToken",
			Handler:    _AuthService_PackToken_Handler,
		},
		{
			MethodName: "ClearToken",
			Handler:    _AuthService_ClearToken_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AuthService_Delete_Handler,
		},
		{
			MethodName: "AddBlocked",
			Handler:    _AuthService_AddBlocked_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}