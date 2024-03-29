// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: protos/devices.proto

package messages

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

const (
	Devices_GetDevices_FullMethodName       = "/Devices/GetDevices"
	Devices_UpdatePrivateKey_FullMethodName = "/Devices/UpdatePrivateKey"
)

// DevicesClient is the client API for Devices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DevicesClient interface {
	GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (*GetDevicesResponse, error)
	UpdatePrivateKey(ctx context.Context, in *UpdatePrivateKeyRequest, opts ...grpc.CallOption) (*UpdatePrivateKeyResponse, error)
}

type devicesClient struct {
	cc grpc.ClientConnInterface
}

func NewDevicesClient(cc grpc.ClientConnInterface) DevicesClient {
	return &devicesClient{cc}
}

func (c *devicesClient) GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (*GetDevicesResponse, error) {
	out := new(GetDevicesResponse)
	err := c.cc.Invoke(ctx, Devices_GetDevices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) UpdatePrivateKey(ctx context.Context, in *UpdatePrivateKeyRequest, opts ...grpc.CallOption) (*UpdatePrivateKeyResponse, error) {
	out := new(UpdatePrivateKeyResponse)
	err := c.cc.Invoke(ctx, Devices_UpdatePrivateKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevicesServer is the server API for Devices service.
// All implementations must embed UnimplementedDevicesServer
// for forward compatibility
type DevicesServer interface {
	GetDevices(context.Context, *GetDevicesRequest) (*GetDevicesResponse, error)
	UpdatePrivateKey(context.Context, *UpdatePrivateKeyRequest) (*UpdatePrivateKeyResponse, error)
	mustEmbedUnimplementedDevicesServer()
}

// UnimplementedDevicesServer must be embedded to have forward compatible implementations.
type UnimplementedDevicesServer struct {
}

func (UnimplementedDevicesServer) GetDevices(context.Context, *GetDevicesRequest) (*GetDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevices not implemented")
}
func (UnimplementedDevicesServer) UpdatePrivateKey(context.Context, *UpdatePrivateKeyRequest) (*UpdatePrivateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrivateKey not implemented")
}
func (UnimplementedDevicesServer) mustEmbedUnimplementedDevicesServer() {}

// UnsafeDevicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevicesServer will
// result in compilation errors.
type UnsafeDevicesServer interface {
	mustEmbedUnimplementedDevicesServer()
}

func RegisterDevicesServer(s grpc.ServiceRegistrar, srv DevicesServer) {
	s.RegisterService(&Devices_ServiceDesc, srv)
}

func _Devices_GetDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Devices_GetDevices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetDevices(ctx, req.(*GetDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_UpdatePrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePrivateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).UpdatePrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Devices_UpdatePrivateKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).UpdatePrivateKey(ctx, req.(*UpdatePrivateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Devices_ServiceDesc is the grpc.ServiceDesc for Devices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Devices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Devices",
	HandlerType: (*DevicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDevices",
			Handler:    _Devices_GetDevices_Handler,
		},
		{
			MethodName: "UpdatePrivateKey",
			Handler:    _Devices_UpdatePrivateKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/devices.proto",
}
