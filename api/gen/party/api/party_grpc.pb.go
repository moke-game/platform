// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: party/party.proto

// PartyService is a service for party, used to manage a party member, sync party info etc.
// 组队服务，用于管理队伍成员，同步队伍信息等

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
	PartyService_GetParty_FullMethodName         = "/party.v1.PartyService/GetParty"
	PartyService_JoinParty_FullMethodName        = "/party.v1.PartyService/JoinParty"
	PartyService_JoinPartyReply_FullMethodName   = "/party.v1.PartyService/JoinPartyReply"
	PartyService_LeaveParty_FullMethodName       = "/party.v1.PartyService/LeaveParty"
	PartyService_KickOut_FullMethodName          = "/party.v1.PartyService/KickOut"
	PartyService_ManageParty_FullMethodName      = "/party.v1.PartyService/ManageParty"
	PartyService_UpdateMember_FullMethodName     = "/party.v1.PartyService/UpdateMember"
	PartyService_InviteJoinParty_FullMethodName  = "/party.v1.PartyService/InviteJoinParty"
	PartyService_InviteJoinReplay_FullMethodName = "/party.v1.PartyService/InviteJoinReplay"
)

// PartyServiceClient is the client API for PartyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// PartyService is a service for party
// 组队服务
type PartyServiceClient interface {
	// GetParty get party info
	// 获取当前队伍信息
	GetParty(ctx context.Context, in *GetPartyRequest, opts ...grpc.CallOption) (*GetPartyResponse, error)
	// JoinParty join party
	// 加入队伍
	JoinParty(ctx context.Context, in *JoinPartyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[JoinPartyResponse], error)
	// JoinPartyReply join party reply
	// 加入队伍回复
	JoinPartyReply(ctx context.Context, in *JoinPartyReplyRequest, opts ...grpc.CallOption) (*JoinPartyReplyResponse, error)
	// LeaveParty leave party
	// 离开队伍
	LeaveParty(ctx context.Context, in *LeavePartyRequest, opts ...grpc.CallOption) (*LeavePartyResponse, error)
	// KickOut kick out member
	// 踢出队伍
	KickOut(ctx context.Context, in *KickOutRequest, opts ...grpc.CallOption) (*KickOutResponse, error)
	// ManageParty manage party
	// 管理队伍
	ManageParty(ctx context.Context, in *ManagePartyRequest, opts ...grpc.CallOption) (*ManagePartyResponse, error)
	// UpdateMember update member
	// 队伍成员更新信息
	UpdateMember(ctx context.Context, in *UpdateMemberRequest, opts ...grpc.CallOption) (*UpdateMemberResponse, error)
	// InviteJoin invite join party
	// 邀请加入队伍
	InviteJoinParty(ctx context.Context, in *InviteJoinRequest, opts ...grpc.CallOption) (*InviteJoinResponse, error)
	// InviteJoinReplay invite join party replay
	// 邀请加入队伍回复
	InviteJoinReplay(ctx context.Context, in *InviteReplayRequest, opts ...grpc.CallOption) (*InviteReplayResponse, error)
}

type partyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartyServiceClient(cc grpc.ClientConnInterface) PartyServiceClient {
	return &partyServiceClient{cc}
}

func (c *partyServiceClient) GetParty(ctx context.Context, in *GetPartyRequest, opts ...grpc.CallOption) (*GetPartyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPartyResponse)
	err := c.cc.Invoke(ctx, PartyService_GetParty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) JoinParty(ctx context.Context, in *JoinPartyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[JoinPartyResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PartyService_ServiceDesc.Streams[0], PartyService_JoinParty_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[JoinPartyRequest, JoinPartyResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PartyService_JoinPartyClient = grpc.ServerStreamingClient[JoinPartyResponse]

func (c *partyServiceClient) JoinPartyReply(ctx context.Context, in *JoinPartyReplyRequest, opts ...grpc.CallOption) (*JoinPartyReplyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinPartyReplyResponse)
	err := c.cc.Invoke(ctx, PartyService_JoinPartyReply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) LeaveParty(ctx context.Context, in *LeavePartyRequest, opts ...grpc.CallOption) (*LeavePartyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeavePartyResponse)
	err := c.cc.Invoke(ctx, PartyService_LeaveParty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) KickOut(ctx context.Context, in *KickOutRequest, opts ...grpc.CallOption) (*KickOutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KickOutResponse)
	err := c.cc.Invoke(ctx, PartyService_KickOut_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) ManageParty(ctx context.Context, in *ManagePartyRequest, opts ...grpc.CallOption) (*ManagePartyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ManagePartyResponse)
	err := c.cc.Invoke(ctx, PartyService_ManageParty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) UpdateMember(ctx context.Context, in *UpdateMemberRequest, opts ...grpc.CallOption) (*UpdateMemberResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMemberResponse)
	err := c.cc.Invoke(ctx, PartyService_UpdateMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) InviteJoinParty(ctx context.Context, in *InviteJoinRequest, opts ...grpc.CallOption) (*InviteJoinResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InviteJoinResponse)
	err := c.cc.Invoke(ctx, PartyService_InviteJoinParty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) InviteJoinReplay(ctx context.Context, in *InviteReplayRequest, opts ...grpc.CallOption) (*InviteReplayResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InviteReplayResponse)
	err := c.cc.Invoke(ctx, PartyService_InviteJoinReplay_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartyServiceServer is the server API for PartyService service.
// All implementations should embed UnimplementedPartyServiceServer
// for forward compatibility.
//
// PartyService is a service for party
// 组队服务
type PartyServiceServer interface {
	// GetParty get party info
	// 获取当前队伍信息
	GetParty(context.Context, *GetPartyRequest) (*GetPartyResponse, error)
	// JoinParty join party
	// 加入队伍
	JoinParty(*JoinPartyRequest, grpc.ServerStreamingServer[JoinPartyResponse]) error
	// JoinPartyReply join party reply
	// 加入队伍回复
	JoinPartyReply(context.Context, *JoinPartyReplyRequest) (*JoinPartyReplyResponse, error)
	// LeaveParty leave party
	// 离开队伍
	LeaveParty(context.Context, *LeavePartyRequest) (*LeavePartyResponse, error)
	// KickOut kick out member
	// 踢出队伍
	KickOut(context.Context, *KickOutRequest) (*KickOutResponse, error)
	// ManageParty manage party
	// 管理队伍
	ManageParty(context.Context, *ManagePartyRequest) (*ManagePartyResponse, error)
	// UpdateMember update member
	// 队伍成员更新信息
	UpdateMember(context.Context, *UpdateMemberRequest) (*UpdateMemberResponse, error)
	// InviteJoin invite join party
	// 邀请加入队伍
	InviteJoinParty(context.Context, *InviteJoinRequest) (*InviteJoinResponse, error)
	// InviteJoinReplay invite join party replay
	// 邀请加入队伍回复
	InviteJoinReplay(context.Context, *InviteReplayRequest) (*InviteReplayResponse, error)
}

// UnimplementedPartyServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPartyServiceServer struct{}

func (UnimplementedPartyServiceServer) GetParty(context.Context, *GetPartyRequest) (*GetPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParty not implemented")
}
func (UnimplementedPartyServiceServer) JoinParty(*JoinPartyRequest, grpc.ServerStreamingServer[JoinPartyResponse]) error {
	return status.Errorf(codes.Unimplemented, "method JoinParty not implemented")
}
func (UnimplementedPartyServiceServer) JoinPartyReply(context.Context, *JoinPartyReplyRequest) (*JoinPartyReplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPartyReply not implemented")
}
func (UnimplementedPartyServiceServer) LeaveParty(context.Context, *LeavePartyRequest) (*LeavePartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveParty not implemented")
}
func (UnimplementedPartyServiceServer) KickOut(context.Context, *KickOutRequest) (*KickOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickOut not implemented")
}
func (UnimplementedPartyServiceServer) ManageParty(context.Context, *ManagePartyRequest) (*ManagePartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManageParty not implemented")
}
func (UnimplementedPartyServiceServer) UpdateMember(context.Context, *UpdateMemberRequest) (*UpdateMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMember not implemented")
}
func (UnimplementedPartyServiceServer) InviteJoinParty(context.Context, *InviteJoinRequest) (*InviteJoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteJoinParty not implemented")
}
func (UnimplementedPartyServiceServer) InviteJoinReplay(context.Context, *InviteReplayRequest) (*InviteReplayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteJoinReplay not implemented")
}
func (UnimplementedPartyServiceServer) testEmbeddedByValue() {}

// UnsafePartyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartyServiceServer will
// result in compilation errors.
type UnsafePartyServiceServer interface {
	mustEmbedUnimplementedPartyServiceServer()
}

func RegisterPartyServiceServer(s grpc.ServiceRegistrar, srv PartyServiceServer) {
	// If the following call pancis, it indicates UnimplementedPartyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PartyService_ServiceDesc, srv)
}

func _PartyService_GetParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).GetParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_GetParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).GetParty(ctx, req.(*GetPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_JoinParty_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinPartyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PartyServiceServer).JoinParty(m, &grpc.GenericServerStream[JoinPartyRequest, JoinPartyResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PartyService_JoinPartyServer = grpc.ServerStreamingServer[JoinPartyResponse]

func _PartyService_JoinPartyReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinPartyReplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).JoinPartyReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_JoinPartyReply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).JoinPartyReply(ctx, req.(*JoinPartyReplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_LeaveParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeavePartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).LeaveParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_LeaveParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).LeaveParty(ctx, req.(*LeavePartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_KickOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).KickOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_KickOut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).KickOut(ctx, req.(*KickOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_ManageParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManagePartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).ManageParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_ManageParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).ManageParty(ctx, req.(*ManagePartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_UpdateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).UpdateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_UpdateMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).UpdateMember(ctx, req.(*UpdateMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_InviteJoinParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteJoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).InviteJoinParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_InviteJoinParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).InviteJoinParty(ctx, req.(*InviteJoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_InviteJoinReplay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteReplayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).InviteJoinReplay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_InviteJoinReplay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).InviteJoinReplay(ctx, req.(*InviteReplayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PartyService_ServiceDesc is the grpc.ServiceDesc for PartyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "party.v1.PartyService",
	HandlerType: (*PartyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetParty",
			Handler:    _PartyService_GetParty_Handler,
		},
		{
			MethodName: "JoinPartyReply",
			Handler:    _PartyService_JoinPartyReply_Handler,
		},
		{
			MethodName: "LeaveParty",
			Handler:    _PartyService_LeaveParty_Handler,
		},
		{
			MethodName: "KickOut",
			Handler:    _PartyService_KickOut_Handler,
		},
		{
			MethodName: "ManageParty",
			Handler:    _PartyService_ManageParty_Handler,
		},
		{
			MethodName: "UpdateMember",
			Handler:    _PartyService_UpdateMember_Handler,
		},
		{
			MethodName: "InviteJoinParty",
			Handler:    _PartyService_InviteJoinParty_Handler,
		},
		{
			MethodName: "InviteJoinReplay",
			Handler:    _PartyService_InviteJoinReplay_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinParty",
			Handler:       _PartyService_JoinParty_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "party/party.proto",
}