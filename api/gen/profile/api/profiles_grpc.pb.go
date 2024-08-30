// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: profile/profiles.proto

// ProfileService is the service for profile
// 玩家基本信息服务, 用于读取/更新玩家基本信息

package pb

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
	ProfileService_IsProfileExist_FullMethodName   = "/profile.v1.ProfileService/IsProfileExist"
	ProfileService_GetProfile_FullMethodName       = "/profile.v1.ProfileService/GetProfile"
	ProfileService_CreateProfile_FullMethodName    = "/profile.v1.ProfileService/CreateProfile"
	ProfileService_UpdateProfile_FullMethodName    = "/profile.v1.ProfileService/UpdateProfile"
	ProfileService_GetProfileStatus_FullMethodName = "/profile.v1.ProfileService/GetProfileStatus"
	ProfileService_WatchProfile_FullMethodName     = "/profile.v1.ProfileService/WatchProfile"
)

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ProfileService is the service for profile
// 玩家基本信息服务
type ProfileServiceClient interface {
	//IsProfileExist check if profile exist
	// 当前玩家是否存在
	IsProfileExist(ctx context.Context, in *IsProfileExistRequest, opts ...grpc.CallOption) (*IsProfileExistResponse, error)
	// GetProfile get profile by uid
	// 获取玩家基本信息
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	// CreateProfile create profile
	// 创建玩家基本信息
	CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*CreateProfileResponse, error)
	// UpdateProfile update profile
	// 更新玩家基本信息
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error)
	// GetProfileStatus get profile status
	// 获取玩家在线状态
	GetProfileStatus(ctx context.Context, in *GetProfileStatusRequest, opts ...grpc.CallOption) (*GetProfileStatusResponse, error)
	// WatchProfile watch profile
	// 监听玩家基本信息变化
	WatchProfile(ctx context.Context, in *WatchProfileRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchProfileResponse], error)
}

type profileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileServiceClient(cc grpc.ClientConnInterface) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) IsProfileExist(ctx context.Context, in *IsProfileExistRequest, opts ...grpc.CallOption) (*IsProfileExistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsProfileExistResponse)
	err := c.cc.Invoke(ctx, ProfileService_IsProfileExist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, ProfileService_GetProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*CreateProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProfileResponse)
	err := c.cc.Invoke(ctx, ProfileService_CreateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProfileResponse)
	err := c.cc.Invoke(ctx, ProfileService_UpdateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetProfileStatus(ctx context.Context, in *GetProfileStatusRequest, opts ...grpc.CallOption) (*GetProfileStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileStatusResponse)
	err := c.cc.Invoke(ctx, ProfileService_GetProfileStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) WatchProfile(ctx context.Context, in *WatchProfileRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchProfileResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ProfileService_ServiceDesc.Streams[0], ProfileService_WatchProfile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[WatchProfileRequest, WatchProfileResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProfileService_WatchProfileClient = grpc.ServerStreamingClient[WatchProfileResponse]

// ProfileServiceServer is the server API for ProfileService service.
// All implementations should embed UnimplementedProfileServiceServer
// for forward compatibility.
//
// ProfileService is the service for profile
// 玩家基本信息服务
type ProfileServiceServer interface {
	//IsProfileExist check if profile exist
	// 当前玩家是否存在
	IsProfileExist(context.Context, *IsProfileExistRequest) (*IsProfileExistResponse, error)
	// GetProfile get profile by uid
	// 获取玩家基本信息
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	// CreateProfile create profile
	// 创建玩家基本信息
	CreateProfile(context.Context, *CreateProfileRequest) (*CreateProfileResponse, error)
	// UpdateProfile update profile
	// 更新玩家基本信息
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error)
	// GetProfileStatus get profile status
	// 获取玩家在线状态
	GetProfileStatus(context.Context, *GetProfileStatusRequest) (*GetProfileStatusResponse, error)
	// WatchProfile watch profile
	// 监听玩家基本信息变化
	WatchProfile(*WatchProfileRequest, grpc.ServerStreamingServer[WatchProfileResponse]) error
}

// UnimplementedProfileServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProfileServiceServer struct{}

func (UnimplementedProfileServiceServer) IsProfileExist(context.Context, *IsProfileExistRequest) (*IsProfileExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsProfileExist not implemented")
}
func (UnimplementedProfileServiceServer) GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedProfileServiceServer) CreateProfile(context.Context, *CreateProfileRequest) (*CreateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (UnimplementedProfileServiceServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedProfileServiceServer) GetProfileStatus(context.Context, *GetProfileStatusRequest) (*GetProfileStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileStatus not implemented")
}
func (UnimplementedProfileServiceServer) WatchProfile(*WatchProfileRequest, grpc.ServerStreamingServer[WatchProfileResponse]) error {
	return status.Errorf(codes.Unimplemented, "method WatchProfile not implemented")
}
func (UnimplementedProfileServiceServer) testEmbeddedByValue() {}

// UnsafeProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServiceServer will
// result in compilation errors.
type UnsafeProfileServiceServer interface {
	mustEmbedUnimplementedProfileServiceServer()
}

func RegisterProfileServiceServer(s grpc.ServiceRegistrar, srv ProfileServiceServer) {
	// If the following call pancis, it indicates UnimplementedProfileServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProfileService_ServiceDesc, srv)
}

func _ProfileService_IsProfileExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsProfileExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).IsProfileExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileService_IsProfileExist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).IsProfileExist(ctx, req.(*IsProfileExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileService_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileService_CreateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).CreateProfile(ctx, req.(*CreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileService_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetProfileStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfileStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileService_GetProfileStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfileStatus(ctx, req.(*GetProfileStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_WatchProfile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchProfileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProfileServiceServer).WatchProfile(m, &grpc.GenericServerStream[WatchProfileRequest, WatchProfileResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProfileService_WatchProfileServer = grpc.ServerStreamingServer[WatchProfileResponse]

// ProfileService_ServiceDesc is the grpc.ServiceDesc for ProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profile.v1.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsProfileExist",
			Handler:    _ProfileService_IsProfileExist_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _ProfileService_GetProfile_Handler,
		},
		{
			MethodName: "CreateProfile",
			Handler:    _ProfileService_CreateProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _ProfileService_UpdateProfile_Handler,
		},
		{
			MethodName: "GetProfileStatus",
			Handler:    _ProfileService_GetProfileStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchProfile",
			Handler:       _ProfileService_WatchProfile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "profile/profiles.proto",
}

const (
	ProfilePrivateService_GetProfilePrivate_FullMethodName = "/profile.v1.ProfilePrivateService/GetProfilePrivate"
	ProfilePrivateService_SetProfileStatus_FullMethodName  = "/profile.v1.ProfilePrivateService/SetProfileStatus"
	ProfilePrivateService_GetProfileBasics_FullMethodName  = "/profile.v1.ProfilePrivateService/GetProfileBasics"
)

// ProfilePrivateServiceClient is the client API for ProfilePrivateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfilePrivateServiceClient interface {
	// GetProfilePrivate get profile info for private use,like gm,admin etc
	// 获取玩家信息，如gm，admin等
	GetProfilePrivate(ctx context.Context, in *GetProfilePrivateRequest, opts ...grpc.CallOption) (*GetProfilePrivateResponse, error)
	// SetProfileStatus set profile status
	// 设置玩家在线状态
	SetProfileStatus(ctx context.Context, in *SetProfileStatusRequest, opts ...grpc.CallOption) (*SetProfileStatusResponse, error)
	// GetProfileBasics multiple get profile basics, for friends, leaderboard etc
	// 批量获取玩家基本信息,适用于好友，排行榜等等
	GetProfileBasics(ctx context.Context, in *GetProfileBasicsRequest, opts ...grpc.CallOption) (*GetProfileBasicsResponse, error)
}

type profilePrivateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilePrivateServiceClient(cc grpc.ClientConnInterface) ProfilePrivateServiceClient {
	return &profilePrivateServiceClient{cc}
}

func (c *profilePrivateServiceClient) GetProfilePrivate(ctx context.Context, in *GetProfilePrivateRequest, opts ...grpc.CallOption) (*GetProfilePrivateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfilePrivateResponse)
	err := c.cc.Invoke(ctx, ProfilePrivateService_GetProfilePrivate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilePrivateServiceClient) SetProfileStatus(ctx context.Context, in *SetProfileStatusRequest, opts ...grpc.CallOption) (*SetProfileStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetProfileStatusResponse)
	err := c.cc.Invoke(ctx, ProfilePrivateService_SetProfileStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilePrivateServiceClient) GetProfileBasics(ctx context.Context, in *GetProfileBasicsRequest, opts ...grpc.CallOption) (*GetProfileBasicsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileBasicsResponse)
	err := c.cc.Invoke(ctx, ProfilePrivateService_GetProfileBasics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilePrivateServiceServer is the server API for ProfilePrivateService service.
// All implementations should embed UnimplementedProfilePrivateServiceServer
// for forward compatibility.
type ProfilePrivateServiceServer interface {
	// GetProfilePrivate get profile info for private use,like gm,admin etc
	// 获取玩家信息，如gm，admin等
	GetProfilePrivate(context.Context, *GetProfilePrivateRequest) (*GetProfilePrivateResponse, error)
	// SetProfileStatus set profile status
	// 设置玩家在线状态
	SetProfileStatus(context.Context, *SetProfileStatusRequest) (*SetProfileStatusResponse, error)
	// GetProfileBasics multiple get profile basics, for friends, leaderboard etc
	// 批量获取玩家基本信息,适用于好友，排行榜等等
	GetProfileBasics(context.Context, *GetProfileBasicsRequest) (*GetProfileBasicsResponse, error)
}

// UnimplementedProfilePrivateServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProfilePrivateServiceServer struct{}

func (UnimplementedProfilePrivateServiceServer) GetProfilePrivate(context.Context, *GetProfilePrivateRequest) (*GetProfilePrivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfilePrivate not implemented")
}
func (UnimplementedProfilePrivateServiceServer) SetProfileStatus(context.Context, *SetProfileStatusRequest) (*SetProfileStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProfileStatus not implemented")
}
func (UnimplementedProfilePrivateServiceServer) GetProfileBasics(context.Context, *GetProfileBasicsRequest) (*GetProfileBasicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileBasics not implemented")
}
func (UnimplementedProfilePrivateServiceServer) testEmbeddedByValue() {}

// UnsafeProfilePrivateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfilePrivateServiceServer will
// result in compilation errors.
type UnsafeProfilePrivateServiceServer interface {
	mustEmbedUnimplementedProfilePrivateServiceServer()
}

func RegisterProfilePrivateServiceServer(s grpc.ServiceRegistrar, srv ProfilePrivateServiceServer) {
	// If the following call pancis, it indicates UnimplementedProfilePrivateServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProfilePrivateService_ServiceDesc, srv)
}

func _ProfilePrivateService_GetProfilePrivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfilePrivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilePrivateServiceServer).GetProfilePrivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfilePrivateService_GetProfilePrivate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilePrivateServiceServer).GetProfilePrivate(ctx, req.(*GetProfilePrivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilePrivateService_SetProfileStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetProfileStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilePrivateServiceServer).SetProfileStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfilePrivateService_SetProfileStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilePrivateServiceServer).SetProfileStatus(ctx, req.(*SetProfileStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilePrivateService_GetProfileBasics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileBasicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilePrivateServiceServer).GetProfileBasics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfilePrivateService_GetProfileBasics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilePrivateServiceServer).GetProfileBasics(ctx, req.(*GetProfileBasicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfilePrivateService_ServiceDesc is the grpc.ServiceDesc for ProfilePrivateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfilePrivateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profile.v1.ProfilePrivateService",
	HandlerType: (*ProfilePrivateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfilePrivate",
			Handler:    _ProfilePrivateService_GetProfilePrivate_Handler,
		},
		{
			MethodName: "SetProfileStatus",
			Handler:    _ProfilePrivateService_SetProfileStatus_Handler,
		},
		{
			MethodName: "GetProfileBasics",
			Handler:    _ProfilePrivateService_GetProfileBasics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile/profiles.proto",
}