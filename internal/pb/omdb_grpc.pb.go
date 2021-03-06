// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// OmdbServiceClient is the client API for OmdbService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OmdbServiceClient interface {
	FindAll(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchPaginatedResponse, error)
	FindById(ctx context.Context, in *GetDetailRequest, opts ...grpc.CallOption) (*SearchDetailResponse, error)
}

type omdbServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOmdbServiceClient(cc grpc.ClientConnInterface) OmdbServiceClient {
	return &omdbServiceClient{cc}
}

func (c *omdbServiceClient) FindAll(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchPaginatedResponse, error) {
	out := new(SearchPaginatedResponse)
	err := c.cc.Invoke(ctx, "/proto.OmdbService/FindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omdbServiceClient) FindById(ctx context.Context, in *GetDetailRequest, opts ...grpc.CallOption) (*SearchDetailResponse, error) {
	out := new(SearchDetailResponse)
	err := c.cc.Invoke(ctx, "/proto.OmdbService/FindById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OmdbServiceServer is the server API for OmdbService service.
// All implementations must embed UnimplementedOmdbServiceServer
// for forward compatibility
type OmdbServiceServer interface {
	FindAll(context.Context, *SearchRequest) (*SearchPaginatedResponse, error)
	FindById(context.Context, *GetDetailRequest) (*SearchDetailResponse, error)
	mustEmbedUnimplementedOmdbServiceServer()
}

// UnimplementedOmdbServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOmdbServiceServer struct {
}

func (UnimplementedOmdbServiceServer) FindAll(context.Context, *SearchRequest) (*SearchPaginatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedOmdbServiceServer) FindById(context.Context, *GetDetailRequest) (*SearchDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedOmdbServiceServer) mustEmbedUnimplementedOmdbServiceServer() {}

// UnsafeOmdbServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OmdbServiceServer will
// result in compilation errors.
type UnsafeOmdbServiceServer interface {
	mustEmbedUnimplementedOmdbServiceServer()
}

func RegisterOmdbServiceServer(s grpc.ServiceRegistrar, srv OmdbServiceServer) {
	s.RegisterService(&OmdbService_ServiceDesc, srv)
}

func _OmdbService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmdbServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OmdbService/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmdbServiceServer).FindAll(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmdbService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmdbServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OmdbService/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmdbServiceServer).FindById(ctx, req.(*GetDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OmdbService_ServiceDesc is the grpc.ServiceDesc for OmdbService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OmdbService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.OmdbService",
	HandlerType: (*OmdbServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAll",
			Handler:    _OmdbService_FindAll_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _OmdbService_FindById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "omdb.proto",
}
