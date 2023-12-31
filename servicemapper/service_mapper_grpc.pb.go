// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: service_mapper/service_mapper.proto

package servicemapper

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

// ServiceMapperClient is the client API for ServiceMapper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceMapperClient interface {
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error)
}

type serviceMapperClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceMapperClient(cc grpc.ClientConnInterface) ServiceMapperClient {
	return &serviceMapperClient{cc}
}

func (c *serviceMapperClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error) {
	out := new(CreateServiceResponse)
	err := c.cc.Invoke(ctx, "/serviceMapper.ServiceMapper/CreateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceMapperServer is the server API for ServiceMapper service.
// All implementations must embed UnimplementedServiceMapperServer
// for forward compatibility
type ServiceMapperServer interface {
	CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error)
	mustEmbedUnimplementedServiceMapperServer()
}

// UnimplementedServiceMapperServer must be embedded to have forward compatible implementations.
type UnimplementedServiceMapperServer struct {
}

func (UnimplementedServiceMapperServer) CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedServiceMapperServer) mustEmbedUnimplementedServiceMapperServer() {}

// UnsafeServiceMapperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceMapperServer will
// result in compilation errors.
type UnsafeServiceMapperServer interface {
	mustEmbedUnimplementedServiceMapperServer()
}

func RegisterServiceMapperServer(s grpc.ServiceRegistrar, srv ServiceMapperServer) {
	s.RegisterService(&ServiceMapper_ServiceDesc, srv)
}

func _ServiceMapper_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceMapperServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceMapper.ServiceMapper/CreateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceMapperServer).CreateService(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceMapper_ServiceDesc is the grpc.ServiceDesc for ServiceMapper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceMapper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serviceMapper.ServiceMapper",
	HandlerType: (*ServiceMapperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateService",
			Handler:    _ServiceMapper_CreateService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_mapper/service_mapper.proto",
}
