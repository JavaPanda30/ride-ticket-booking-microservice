// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: proto/passenger.proto

package passenger

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
	PassengerService_RequestRide_FullMethodName           = "/passenger.PassengerService/RequestRide"
	PassengerService_UpdatePassengerStatus_FullMethodName = "/passenger.PassengerService/UpdatePassengerStatus"
)

// PassengerServiceClient is the client API for PassengerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PassengerServiceClient interface {
	RequestRide(ctx context.Context, in *PassengerRequest, opts ...grpc.CallOption) (*PassengerResponse, error)
	UpdatePassengerStatus(ctx context.Context, in *PassengerStatusUpdate, opts ...grpc.CallOption) (*PassengerResponse, error)
}

type passengerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPassengerServiceClient(cc grpc.ClientConnInterface) PassengerServiceClient {
	return &passengerServiceClient{cc}
}

func (c *passengerServiceClient) RequestRide(ctx context.Context, in *PassengerRequest, opts ...grpc.CallOption) (*PassengerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PassengerResponse)
	err := c.cc.Invoke(ctx, PassengerService_RequestRide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerServiceClient) UpdatePassengerStatus(ctx context.Context, in *PassengerStatusUpdate, opts ...grpc.CallOption) (*PassengerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PassengerResponse)
	err := c.cc.Invoke(ctx, PassengerService_UpdatePassengerStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PassengerServiceServer is the server API for PassengerService service.
// All implementations must embed UnimplementedPassengerServiceServer
// for forward compatibility.
type PassengerServiceServer interface {
	RequestRide(context.Context, *PassengerRequest) (*PassengerResponse, error)
	UpdatePassengerStatus(context.Context, *PassengerStatusUpdate) (*PassengerResponse, error)
	mustEmbedUnimplementedPassengerServiceServer()
}

// UnimplementedPassengerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPassengerServiceServer struct{}

func (UnimplementedPassengerServiceServer) RequestRide(context.Context, *PassengerRequest) (*PassengerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestRide not implemented")
}
func (UnimplementedPassengerServiceServer) UpdatePassengerStatus(context.Context, *PassengerStatusUpdate) (*PassengerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassengerStatus not implemented")
}
func (UnimplementedPassengerServiceServer) mustEmbedUnimplementedPassengerServiceServer() {}
func (UnimplementedPassengerServiceServer) testEmbeddedByValue()                          {}

// UnsafePassengerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PassengerServiceServer will
// result in compilation errors.
type UnsafePassengerServiceServer interface {
	mustEmbedUnimplementedPassengerServiceServer()
}

func RegisterPassengerServiceServer(s grpc.ServiceRegistrar, srv PassengerServiceServer) {
	// If the following call pancis, it indicates UnimplementedPassengerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PassengerService_ServiceDesc, srv)
}

func _PassengerService_RequestRide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PassengerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerServiceServer).RequestRide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PassengerService_RequestRide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerServiceServer).RequestRide(ctx, req.(*PassengerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerService_UpdatePassengerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PassengerStatusUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerServiceServer).UpdatePassengerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PassengerService_UpdatePassengerStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerServiceServer).UpdatePassengerStatus(ctx, req.(*PassengerStatusUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

// PassengerService_ServiceDesc is the grpc.ServiceDesc for PassengerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PassengerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "passenger.PassengerService",
	HandlerType: (*PassengerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestRide",
			Handler:    _PassengerService_RequestRide_Handler,
		},
		{
			MethodName: "UpdatePassengerStatus",
			Handler:    _PassengerService_UpdatePassengerStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/passenger.proto",
}
