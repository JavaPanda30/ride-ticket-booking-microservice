// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: proto/ride.proto

package ride

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
	RideService_CreateRide_FullMethodName       = "/ride.RideService/CreateRide"
	RideService_UpdateRideStatus_FullMethodName = "/ride.RideService/UpdateRideStatus"
)

// RideServiceClient is the client API for RideService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RideServiceClient interface {
	CreateRide(ctx context.Context, in *RideRequest, opts ...grpc.CallOption) (*RideResponse, error)
	UpdateRideStatus(ctx context.Context, in *RideStatusUpdate, opts ...grpc.CallOption) (*RideResponse, error)
}

type rideServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRideServiceClient(cc grpc.ClientConnInterface) RideServiceClient {
	return &rideServiceClient{cc}
}

func (c *rideServiceClient) CreateRide(ctx context.Context, in *RideRequest, opts ...grpc.CallOption) (*RideResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RideResponse)
	err := c.cc.Invoke(ctx, RideService_CreateRide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rideServiceClient) UpdateRideStatus(ctx context.Context, in *RideStatusUpdate, opts ...grpc.CallOption) (*RideResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RideResponse)
	err := c.cc.Invoke(ctx, RideService_UpdateRideStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RideServiceServer is the server API for RideService service.
// All implementations must embed UnimplementedRideServiceServer
// for forward compatibility.
type RideServiceServer interface {
	CreateRide(context.Context, *RideRequest) (*RideResponse, error)
	UpdateRideStatus(context.Context, *RideStatusUpdate) (*RideResponse, error)
	mustEmbedUnimplementedRideServiceServer()
}

// UnimplementedRideServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRideServiceServer struct{}

func (UnimplementedRideServiceServer) CreateRide(context.Context, *RideRequest) (*RideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRide not implemented")
}
func (UnimplementedRideServiceServer) UpdateRideStatus(context.Context, *RideStatusUpdate) (*RideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRideStatus not implemented")
}
func (UnimplementedRideServiceServer) mustEmbedUnimplementedRideServiceServer() {}
func (UnimplementedRideServiceServer) testEmbeddedByValue()                     {}

// UnsafeRideServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RideServiceServer will
// result in compilation errors.
type UnsafeRideServiceServer interface {
	mustEmbedUnimplementedRideServiceServer()
}

func RegisterRideServiceServer(s grpc.ServiceRegistrar, srv RideServiceServer) {
	// If the following call pancis, it indicates UnimplementedRideServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RideService_ServiceDesc, srv)
}

func _RideService_CreateRide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RideServiceServer).CreateRide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RideService_CreateRide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RideServiceServer).CreateRide(ctx, req.(*RideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RideService_UpdateRideStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RideStatusUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RideServiceServer).UpdateRideStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RideService_UpdateRideStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RideServiceServer).UpdateRideStatus(ctx, req.(*RideStatusUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

// RideService_ServiceDesc is the grpc.ServiceDesc for RideService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RideService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ride.RideService",
	HandlerType: (*RideServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRide",
			Handler:    _RideService_CreateRide_Handler,
		},
		{
			MethodName: "UpdateRideStatus",
			Handler:    _RideService_UpdateRideStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ride.proto",
}
