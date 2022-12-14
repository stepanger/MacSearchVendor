// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: config/conf_server.proto

package gen

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

// SearchVendorClient is the client API for SearchVendor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchVendorClient interface {
	GetSearchVendor(ctx context.Context, in *Mac, opts ...grpc.CallOption) (*Vendor, error)
}

type searchVendorClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchVendorClient(cc grpc.ClientConnInterface) SearchVendorClient {
	return &searchVendorClient{cc}
}

func (c *searchVendorClient) GetSearchVendor(ctx context.Context, in *Mac, opts ...grpc.CallOption) (*Vendor, error) {
	out := new(Vendor)
	err := c.cc.Invoke(ctx, "/conf_server.SearchVendor/GetSearchVendor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchVendorServer is the server API for SearchVendor service.
// All implementations must embed UnimplementedSearchVendorServer
// for forward compatibility
type SearchVendorServer interface {
	GetSearchVendor(context.Context, *Mac) (*Vendor, error)
	mustEmbedUnimplementedSearchVendorServer()
}

// UnimplementedSearchVendorServer must be embedded to have forward compatible implementations.
type UnimplementedSearchVendorServer struct {
}

func (UnimplementedSearchVendorServer) GetSearchVendor(context.Context, *Mac) (*Vendor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearchVendor not implemented")
}
func (UnimplementedSearchVendorServer) mustEmbedUnimplementedSearchVendorServer() {}

// UnsafeSearchVendorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchVendorServer will
// result in compilation errors.
type UnsafeSearchVendorServer interface {
	mustEmbedUnimplementedSearchVendorServer()
}

func RegisterSearchVendorServer(s grpc.ServiceRegistrar, srv SearchVendorServer) {
	s.RegisterService(&SearchVendor_ServiceDesc, srv)
}

func _SearchVendor_GetSearchVendor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mac)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchVendorServer).GetSearchVendor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conf_server.SearchVendor/GetSearchVendor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchVendorServer).GetSearchVendor(ctx, req.(*Mac))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchVendor_ServiceDesc is the grpc.ServiceDesc for SearchVendor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchVendor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "conf_server.SearchVendor",
	HandlerType: (*SearchVendorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSearchVendor",
			Handler:    _SearchVendor_GetSearchVendor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config/conf_server.proto",
}
