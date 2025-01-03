// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: knapsack/knapsacks.proto

// KnapsackService is a service for knapsack
// 背包服务 用于背包物品管理

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
	KnapsackService_GetKnapsack_FullMethodName       = "/knapsack.v1.KnapsackService/GetKnapsack"
	KnapsackService_AddItem_FullMethodName           = "/knapsack.v1.KnapsackService/AddItem"
	KnapsackService_RemoveItem_FullMethodName        = "/knapsack.v1.KnapsackService/RemoveItem"
	KnapsackService_RemoveThenAddItem_FullMethodName = "/knapsack.v1.KnapsackService/RemoveThenAddItem"
)

// KnapsackServiceClient is the client API for KnapsackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KnapsackServiceClient interface {
	// GetKnapsack get knapsack items
	GetKnapsack(ctx context.Context, in *GetKnapsackRequest, opts ...grpc.CallOption) (*GetKnapsackResponse, error)
	// AddItem add items to knapsack
	AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error)
	// RemoveItem remove items
	RemoveItem(ctx context.Context, in *RemoveItemRequest, opts ...grpc.CallOption) (*RemoveItemResponse, error)
	// RemoveThenAddItem remove items then add items
	RemoveThenAddItem(ctx context.Context, in *RemoveThenAddItemRequest, opts ...grpc.CallOption) (*RemoveThenAddItemResponse, error)
}

type knapsackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKnapsackServiceClient(cc grpc.ClientConnInterface) KnapsackServiceClient {
	return &knapsackServiceClient{cc}
}

func (c *knapsackServiceClient) GetKnapsack(ctx context.Context, in *GetKnapsackRequest, opts ...grpc.CallOption) (*GetKnapsackResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetKnapsackResponse)
	err := c.cc.Invoke(ctx, KnapsackService_GetKnapsack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knapsackServiceClient) AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddItemResponse)
	err := c.cc.Invoke(ctx, KnapsackService_AddItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knapsackServiceClient) RemoveItem(ctx context.Context, in *RemoveItemRequest, opts ...grpc.CallOption) (*RemoveItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveItemResponse)
	err := c.cc.Invoke(ctx, KnapsackService_RemoveItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knapsackServiceClient) RemoveThenAddItem(ctx context.Context, in *RemoveThenAddItemRequest, opts ...grpc.CallOption) (*RemoveThenAddItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveThenAddItemResponse)
	err := c.cc.Invoke(ctx, KnapsackService_RemoveThenAddItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KnapsackServiceServer is the server API for KnapsackService service.
// All implementations should embed UnimplementedKnapsackServiceServer
// for forward compatibility.
type KnapsackServiceServer interface {
	// GetKnapsack get knapsack items
	GetKnapsack(context.Context, *GetKnapsackRequest) (*GetKnapsackResponse, error)
	// AddItem add items to knapsack
	AddItem(context.Context, *AddItemRequest) (*AddItemResponse, error)
	// RemoveItem remove items
	RemoveItem(context.Context, *RemoveItemRequest) (*RemoveItemResponse, error)
	// RemoveThenAddItem remove items then add items
	RemoveThenAddItem(context.Context, *RemoveThenAddItemRequest) (*RemoveThenAddItemResponse, error)
}

// UnimplementedKnapsackServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKnapsackServiceServer struct{}

func (UnimplementedKnapsackServiceServer) GetKnapsack(context.Context, *GetKnapsackRequest) (*GetKnapsackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKnapsack not implemented")
}
func (UnimplementedKnapsackServiceServer) AddItem(context.Context, *AddItemRequest) (*AddItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedKnapsackServiceServer) RemoveItem(context.Context, *RemoveItemRequest) (*RemoveItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveItem not implemented")
}
func (UnimplementedKnapsackServiceServer) RemoveThenAddItem(context.Context, *RemoveThenAddItemRequest) (*RemoveThenAddItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveThenAddItem not implemented")
}
func (UnimplementedKnapsackServiceServer) testEmbeddedByValue() {}

// UnsafeKnapsackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KnapsackServiceServer will
// result in compilation errors.
type UnsafeKnapsackServiceServer interface {
	mustEmbedUnimplementedKnapsackServiceServer()
}

func RegisterKnapsackServiceServer(s grpc.ServiceRegistrar, srv KnapsackServiceServer) {
	// If the following call pancis, it indicates UnimplementedKnapsackServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KnapsackService_ServiceDesc, srv)
}

func _KnapsackService_GetKnapsack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKnapsackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnapsackServiceServer).GetKnapsack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnapsackService_GetKnapsack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnapsackServiceServer).GetKnapsack(ctx, req.(*GetKnapsackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnapsackService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnapsackServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnapsackService_AddItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnapsackServiceServer).AddItem(ctx, req.(*AddItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnapsackService_RemoveItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnapsackServiceServer).RemoveItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnapsackService_RemoveItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnapsackServiceServer).RemoveItem(ctx, req.(*RemoveItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnapsackService_RemoveThenAddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveThenAddItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnapsackServiceServer).RemoveThenAddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnapsackService_RemoveThenAddItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnapsackServiceServer).RemoveThenAddItem(ctx, req.(*RemoveThenAddItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KnapsackService_ServiceDesc is the grpc.ServiceDesc for KnapsackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KnapsackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "knapsack.v1.KnapsackService",
	HandlerType: (*KnapsackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKnapsack",
			Handler:    _KnapsackService_GetKnapsack_Handler,
		},
		{
			MethodName: "AddItem",
			Handler:    _KnapsackService_AddItem_Handler,
		},
		{
			MethodName: "RemoveItem",
			Handler:    _KnapsackService_RemoveItem_Handler,
		},
		{
			MethodName: "RemoveThenAddItem",
			Handler:    _KnapsackService_RemoveThenAddItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "knapsack/knapsacks.proto",
}

const (
	KnapsackPrivateService_AddItem_FullMethodName     = "/knapsack.v1.KnapsackPrivateService/AddItem"
	KnapsackPrivateService_RemoveItem_FullMethodName  = "/knapsack.v1.KnapsackPrivateService/RemoveItem"
	KnapsackPrivateService_GetItemById_FullMethodName = "/knapsack.v1.KnapsackPrivateService/GetItemById"
	KnapsackPrivateService_GetKnapsack_FullMethodName = "/knapsack.v1.KnapsackPrivateService/GetKnapsack"
)

// KnapsackPrivateServiceClient is the client API for KnapsackPrivateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// KnapsackPrivateService provide private service for game/gm service
type KnapsackPrivateServiceClient interface {
	// AddItem add items
	AddItem(ctx context.Context, in *AddItemPrivateRequest, opts ...grpc.CallOption) (*AddItemPrivateResponse, error)
	// RemoveItem remove items
	RemoveItem(ctx context.Context, in *RemoveItemPrivateRequest, opts ...grpc.CallOption) (*RemoveItemPrivateResponse, error)
	// GetItemById get item by id
	GetItemById(ctx context.Context, in *GetItemByIdPrivateRequest, opts ...grpc.CallOption) (*GetItemByIdPrivateResponse, error)
	// GetKnapsack get knapsack info
	GetKnapsack(ctx context.Context, in *GetKnapsackRequest, opts ...grpc.CallOption) (*GetKnapsackResponse, error)
}

type knapsackPrivateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKnapsackPrivateServiceClient(cc grpc.ClientConnInterface) KnapsackPrivateServiceClient {
	return &knapsackPrivateServiceClient{cc}
}

func (c *knapsackPrivateServiceClient) AddItem(ctx context.Context, in *AddItemPrivateRequest, opts ...grpc.CallOption) (*AddItemPrivateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddItemPrivateResponse)
	err := c.cc.Invoke(ctx, KnapsackPrivateService_AddItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knapsackPrivateServiceClient) RemoveItem(ctx context.Context, in *RemoveItemPrivateRequest, opts ...grpc.CallOption) (*RemoveItemPrivateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveItemPrivateResponse)
	err := c.cc.Invoke(ctx, KnapsackPrivateService_RemoveItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knapsackPrivateServiceClient) GetItemById(ctx context.Context, in *GetItemByIdPrivateRequest, opts ...grpc.CallOption) (*GetItemByIdPrivateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetItemByIdPrivateResponse)
	err := c.cc.Invoke(ctx, KnapsackPrivateService_GetItemById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knapsackPrivateServiceClient) GetKnapsack(ctx context.Context, in *GetKnapsackRequest, opts ...grpc.CallOption) (*GetKnapsackResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetKnapsackResponse)
	err := c.cc.Invoke(ctx, KnapsackPrivateService_GetKnapsack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KnapsackPrivateServiceServer is the server API for KnapsackPrivateService service.
// All implementations should embed UnimplementedKnapsackPrivateServiceServer
// for forward compatibility.
//
// KnapsackPrivateService provide private service for game/gm service
type KnapsackPrivateServiceServer interface {
	// AddItem add items
	AddItem(context.Context, *AddItemPrivateRequest) (*AddItemPrivateResponse, error)
	// RemoveItem remove items
	RemoveItem(context.Context, *RemoveItemPrivateRequest) (*RemoveItemPrivateResponse, error)
	// GetItemById get item by id
	GetItemById(context.Context, *GetItemByIdPrivateRequest) (*GetItemByIdPrivateResponse, error)
	// GetKnapsack get knapsack info
	GetKnapsack(context.Context, *GetKnapsackRequest) (*GetKnapsackResponse, error)
}

// UnimplementedKnapsackPrivateServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKnapsackPrivateServiceServer struct{}

func (UnimplementedKnapsackPrivateServiceServer) AddItem(context.Context, *AddItemPrivateRequest) (*AddItemPrivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedKnapsackPrivateServiceServer) RemoveItem(context.Context, *RemoveItemPrivateRequest) (*RemoveItemPrivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveItem not implemented")
}
func (UnimplementedKnapsackPrivateServiceServer) GetItemById(context.Context, *GetItemByIdPrivateRequest) (*GetItemByIdPrivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemById not implemented")
}
func (UnimplementedKnapsackPrivateServiceServer) GetKnapsack(context.Context, *GetKnapsackRequest) (*GetKnapsackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKnapsack not implemented")
}
func (UnimplementedKnapsackPrivateServiceServer) testEmbeddedByValue() {}

// UnsafeKnapsackPrivateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KnapsackPrivateServiceServer will
// result in compilation errors.
type UnsafeKnapsackPrivateServiceServer interface {
	mustEmbedUnimplementedKnapsackPrivateServiceServer()
}

func RegisterKnapsackPrivateServiceServer(s grpc.ServiceRegistrar, srv KnapsackPrivateServiceServer) {
	// If the following call pancis, it indicates UnimplementedKnapsackPrivateServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KnapsackPrivateService_ServiceDesc, srv)
}

func _KnapsackPrivateService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemPrivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnapsackPrivateServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnapsackPrivateService_AddItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnapsackPrivateServiceServer).AddItem(ctx, req.(*AddItemPrivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnapsackPrivateService_RemoveItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveItemPrivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnapsackPrivateServiceServer).RemoveItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnapsackPrivateService_RemoveItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnapsackPrivateServiceServer).RemoveItem(ctx, req.(*RemoveItemPrivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnapsackPrivateService_GetItemById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemByIdPrivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnapsackPrivateServiceServer).GetItemById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnapsackPrivateService_GetItemById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnapsackPrivateServiceServer).GetItemById(ctx, req.(*GetItemByIdPrivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnapsackPrivateService_GetKnapsack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKnapsackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnapsackPrivateServiceServer).GetKnapsack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnapsackPrivateService_GetKnapsack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnapsackPrivateServiceServer).GetKnapsack(ctx, req.(*GetKnapsackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KnapsackPrivateService_ServiceDesc is the grpc.ServiceDesc for KnapsackPrivateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KnapsackPrivateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "knapsack.v1.KnapsackPrivateService",
	HandlerType: (*KnapsackPrivateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItem",
			Handler:    _KnapsackPrivateService_AddItem_Handler,
		},
		{
			MethodName: "RemoveItem",
			Handler:    _KnapsackPrivateService_RemoveItem_Handler,
		},
		{
			MethodName: "GetItemById",
			Handler:    _KnapsackPrivateService_GetItemById_Handler,
		},
		{
			MethodName: "GetKnapsack",
			Handler:    _KnapsackPrivateService_GetKnapsack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "knapsack/knapsacks.proto",
}
