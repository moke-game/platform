// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: analytics/analytics.proto

// Analytics service for sending analytics events,
// support multi delivery type: local,thinkingdata,clickhouse,mixpanel etc.
// 分析服务用于发送分析事件
// 支持多种投递方式: local,thinkingdata,clickhouse,mixpanel等

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
	AnalyticsService_Analytics_FullMethodName = "/analytics.v1.AnalyticsService/Analytics"
)

// AnalyticsServiceClient is the client API for AnalyticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyticsServiceClient interface {
	// Analytics send a batch of analytics events to the analytics service，return Nothing
	// Recommend to use async/multi events at once
	// 发送一批分析事件到分析服务，返回Nothing
	// 建议使用异步+批量事件一次发送
	Analytics(ctx context.Context, in *AnalyticsEvents, opts ...grpc.CallOption) (*Nothing, error)
}

type analyticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticsServiceClient(cc grpc.ClientConnInterface) AnalyticsServiceClient {
	return &analyticsServiceClient{cc}
}

func (c *analyticsServiceClient) Analytics(ctx context.Context, in *AnalyticsEvents, opts ...grpc.CallOption) (*Nothing, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Nothing)
	err := c.cc.Invoke(ctx, AnalyticsService_Analytics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyticsServiceServer is the server API for AnalyticsService service.
// All implementations should embed UnimplementedAnalyticsServiceServer
// for forward compatibility.
type AnalyticsServiceServer interface {
	// Analytics send a batch of analytics events to the analytics service，return Nothing
	// Recommend to use async/multi events at once
	// 发送一批分析事件到分析服务，返回Nothing
	// 建议使用异步+批量事件一次发送
	Analytics(context.Context, *AnalyticsEvents) (*Nothing, error)
}

// UnimplementedAnalyticsServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAnalyticsServiceServer struct{}

func (UnimplementedAnalyticsServiceServer) Analytics(context.Context, *AnalyticsEvents) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Analytics not implemented")
}
func (UnimplementedAnalyticsServiceServer) testEmbeddedByValue() {}

// UnsafeAnalyticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyticsServiceServer will
// result in compilation errors.
type UnsafeAnalyticsServiceServer interface {
	mustEmbedUnimplementedAnalyticsServiceServer()
}

func RegisterAnalyticsServiceServer(s grpc.ServiceRegistrar, srv AnalyticsServiceServer) {
	// If the following call pancis, it indicates UnimplementedAnalyticsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AnalyticsService_ServiceDesc, srv)
}

func _AnalyticsService_Analytics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyticsEvents)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServiceServer).Analytics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnalyticsService_Analytics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServiceServer).Analytics(ctx, req.(*AnalyticsEvents))
	}
	return interceptor(ctx, in, info, handler)
}

// AnalyticsService_ServiceDesc is the grpc.ServiceDesc for AnalyticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnalyticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "analytics.v1.AnalyticsService",
	HandlerType: (*AnalyticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Analytics",
			Handler:    _AnalyticsService_Analytics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analytics/analytics.proto",
}