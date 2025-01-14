// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: internal/infra/proto/order_item.proto

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
	OrderService_CreateOrder_FullMethodName                    = "/pb.OrderService/CreateOrder"
	OrderService_CreateOrderStream_FullMethodName              = "/pb.OrderService/CreateOrderStream"
	OrderService_CreateOrderStreamBidirectional_FullMethodName = "/pb.OrderService/CreateOrderStreamBidirectional"
	OrderService_ListOrders_FullMethodName                     = "/pb.OrderService/ListOrders"
	OrderService_GetOrder_FullMethodName                       = "/pb.OrderService/GetOrder"
	OrderService_AddItem_FullMethodName                        = "/pb.OrderService/AddItem"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error)
	CreateOrderStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[CreateOrderRequest, OrderList], error)
	CreateOrderStreamBidirectional(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[CreateOrderRequest, Order], error)
	ListOrders(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*OrderList, error)
	GetOrder(ctx context.Context, in *OrderGetRequest, opts ...grpc.CallOption) (*Order, error)
	AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*Item, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderService_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOrderStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[CreateOrderRequest, OrderList], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[0], OrderService_CreateOrderStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CreateOrderRequest, OrderList]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OrderService_CreateOrderStreamClient = grpc.ClientStreamingClient[CreateOrderRequest, OrderList]

func (c *orderServiceClient) CreateOrderStreamBidirectional(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[CreateOrderRequest, Order], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[1], OrderService_CreateOrderStreamBidirectional_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CreateOrderRequest, Order]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OrderService_CreateOrderStreamBidirectionalClient = grpc.BidiStreamingClient[CreateOrderRequest, Order]

func (c *orderServiceClient) ListOrders(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*OrderList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderList)
	err := c.cc.Invoke(ctx, OrderService_ListOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *OrderGetRequest, opts ...grpc.CallOption) (*Order, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderService_GetOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*Item, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Item)
	err := c.cc.Invoke(ctx, OrderService_AddItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility.
type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*Order, error)
	CreateOrderStream(grpc.ClientStreamingServer[CreateOrderRequest, OrderList]) error
	CreateOrderStreamBidirectional(grpc.BidiStreamingServer[CreateOrderRequest, Order]) error
	ListOrders(context.Context, *Blank) (*OrderList, error)
	GetOrder(context.Context, *OrderGetRequest) (*Order, error)
	AddItem(context.Context, *AddItemRequest) (*Item, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderServiceServer struct{}

func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) CreateOrderStream(grpc.ClientStreamingServer[CreateOrderRequest, OrderList]) error {
	return status.Errorf(codes.Unimplemented, "method CreateOrderStream not implemented")
}
func (UnimplementedOrderServiceServer) CreateOrderStreamBidirectional(grpc.BidiStreamingServer[CreateOrderRequest, Order]) error {
	return status.Errorf(codes.Unimplemented, "method CreateOrderStreamBidirectional not implemented")
}
func (UnimplementedOrderServiceServer) ListOrders(context.Context, *Blank) (*OrderList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedOrderServiceServer) GetOrder(context.Context, *OrderGetRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServiceServer) AddItem(context.Context, *AddItemRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}
func (UnimplementedOrderServiceServer) testEmbeddedByValue()                      {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateOrderStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderServiceServer).CreateOrderStream(&grpc.GenericServerStream[CreateOrderRequest, OrderList]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OrderService_CreateOrderStreamServer = grpc.ClientStreamingServer[CreateOrderRequest, OrderList]

func _OrderService_CreateOrderStreamBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderServiceServer).CreateOrderStreamBidirectional(&grpc.GenericServerStream[CreateOrderRequest, Order]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OrderService_CreateOrderStreamBidirectionalServer = grpc.BidiStreamingServer[CreateOrderRequest, Order]

func _OrderService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ListOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOrders(ctx, req.(*Blank))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*OrderGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_AddItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).AddItem(ctx, req.(*AddItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _OrderService_ListOrders_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
		{
			MethodName: "AddItem",
			Handler:    _OrderService_AddItem_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateOrderStream",
			Handler:       _OrderService_CreateOrderStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateOrderStreamBidirectional",
			Handler:       _OrderService_CreateOrderStreamBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/infra/proto/order_item.proto",
}
