// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bookInfo

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

// BookInfoClient is the client API for BookInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookInfoClient interface {
	// GetRate returns the exchange rate for the two provided currency codes
	GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
}

type bookInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewBookInfoClient(cc grpc.ClientConnInterface) BookInfoClient {
	return &bookInfoClient{cc}
}

func (c *bookInfoClient) GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/bookInfo/GetRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookInfoServer is the server API for BookInfo service.
// All implementations must embed UnimplementedBookInfoServer
// for forward compatibility
type BookInfoServer interface {
	// GetRate returns the exchange rate for the two provided currency codes
	GetRate(context.Context, *RateRequest) (*RateResponse, error)
	mustEmbedUnimplementedBookInfoServer()
}

// UnimplementedBookInfoServer must be embedded to have forward compatible implementations.
type UnimplementedBookInfoServer struct {
}

func (UnimplementedBookInfoServer) GetRate(context.Context, *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}
func (UnimplementedBookInfoServer) mustEmbedUnimplementedBookInfoServer() {}

// UnsafeBookInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookInfoServer will
// result in compilation errors.
type UnsafeBookInfoServer interface {
	mustEmbedUnimplementedBookInfoServer()
}

func RegisterBookInfoServer(s grpc.ServiceRegistrar, srv BookInfoServer) {
	s.RegisterService(&BookInfo_ServiceDesc, srv)
}

func _BookInfo_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookInfoServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookInfo/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookInfoServer).GetRate(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookInfo_ServiceDesc is the grpc.ServiceDesc for BookInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookInfo",
	HandlerType: (*BookInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _BookInfo_GetRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/hack.proto",
}
