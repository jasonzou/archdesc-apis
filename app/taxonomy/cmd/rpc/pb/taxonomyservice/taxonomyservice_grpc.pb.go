// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: rpc/pb/taxonomyservice.proto

package taxonomyservice

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

// TaxonomyServiceClient is the client API for TaxonomyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaxonomyServiceClient interface {
	Get(ctx context.Context, in *ReqTaxonomyId, opts ...grpc.CallOption) (*Taxonomy, error)
	GetAll(ctx context.Context, in *ReqGetAll, opts ...grpc.CallOption) (*RespGetAll, error)
}

type taxonomyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaxonomyServiceClient(cc grpc.ClientConnInterface) TaxonomyServiceClient {
	return &taxonomyServiceClient{cc}
}

func (c *taxonomyServiceClient) Get(ctx context.Context, in *ReqTaxonomyId, opts ...grpc.CallOption) (*Taxonomy, error) {
	out := new(Taxonomy)
	err := c.cc.Invoke(ctx, "/taxonomyservice.TaxonomyService/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxonomyServiceClient) GetAll(ctx context.Context, in *ReqGetAll, opts ...grpc.CallOption) (*RespGetAll, error) {
	out := new(RespGetAll)
	err := c.cc.Invoke(ctx, "/taxonomyservice.TaxonomyService/getAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaxonomyServiceServer is the server API for TaxonomyService service.
// All implementations must embed UnimplementedTaxonomyServiceServer
// for forward compatibility
type TaxonomyServiceServer interface {
	Get(context.Context, *ReqTaxonomyId) (*Taxonomy, error)
	GetAll(context.Context, *ReqGetAll) (*RespGetAll, error)
	mustEmbedUnimplementedTaxonomyServiceServer()
}

// UnimplementedTaxonomyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaxonomyServiceServer struct {
}

func (UnimplementedTaxonomyServiceServer) Get(context.Context, *ReqTaxonomyId) (*Taxonomy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTaxonomyServiceServer) GetAll(context.Context, *ReqGetAll) (*RespGetAll, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedTaxonomyServiceServer) mustEmbedUnimplementedTaxonomyServiceServer() {}

// UnsafeTaxonomyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaxonomyServiceServer will
// result in compilation errors.
type UnsafeTaxonomyServiceServer interface {
	mustEmbedUnimplementedTaxonomyServiceServer()
}

func RegisterTaxonomyServiceServer(s grpc.ServiceRegistrar, srv TaxonomyServiceServer) {
	s.RegisterService(&TaxonomyService_ServiceDesc, srv)
}

func _TaxonomyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqTaxonomyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxonomyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taxonomyservice.TaxonomyService/get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxonomyServiceServer).Get(ctx, req.(*ReqTaxonomyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaxonomyService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetAll)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxonomyServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taxonomyservice.TaxonomyService/getAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxonomyServiceServer).GetAll(ctx, req.(*ReqGetAll))
	}
	return interceptor(ctx, in, info, handler)
}

// TaxonomyService_ServiceDesc is the grpc.ServiceDesc for TaxonomyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaxonomyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "taxonomyservice.TaxonomyService",
	HandlerType: (*TaxonomyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _TaxonomyService_Get_Handler,
		},
		{
			MethodName: "getAll",
			Handler:    _TaxonomyService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/pb/taxonomyservice.proto",
}
