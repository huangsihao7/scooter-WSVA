// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: recommend.proto

package recommend

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
	RecommendSrv_CreateVideo_FullMethodName = "/feed.RecommendSrv/CreateVideo"
)

// RecommendSrvClient is the client API for RecommendSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecommendSrvClient interface {
	CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*CreateVideoResponse, error)
}

type recommendSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommendSrvClient(cc grpc.ClientConnInterface) RecommendSrvClient {
	return &recommendSrvClient{cc}
}

func (c *recommendSrvClient) CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*CreateVideoResponse, error) {
	out := new(CreateVideoResponse)
	err := c.cc.Invoke(ctx, RecommendSrv_CreateVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendSrvServer is the server API for RecommendSrv service.
// All implementations must embed UnimplementedRecommendSrvServer
// for forward compatibility
type RecommendSrvServer interface {
	CreateVideo(context.Context, *CreateVideoRequest) (*CreateVideoResponse, error)
	mustEmbedUnimplementedRecommendSrvServer()
}

// UnimplementedRecommendSrvServer must be embedded to have forward compatible implementations.
type UnimplementedRecommendSrvServer struct {
}

func (UnimplementedRecommendSrvServer) CreateVideo(context.Context, *CreateVideoRequest) (*CreateVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVideo not implemented")
}
func (UnimplementedRecommendSrvServer) mustEmbedUnimplementedRecommendSrvServer() {}

// UnsafeRecommendSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecommendSrvServer will
// result in compilation errors.
type UnsafeRecommendSrvServer interface {
	mustEmbedUnimplementedRecommendSrvServer()
}

func RegisterRecommendSrvServer(s grpc.ServiceRegistrar, srv RecommendSrvServer) {
	s.RegisterService(&RecommendSrv_ServiceDesc, srv)
}

func _RecommendSrv_CreateVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendSrvServer).CreateVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecommendSrv_CreateVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendSrvServer).CreateVideo(ctx, req.(*CreateVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecommendSrv_ServiceDesc is the grpc.ServiceDesc for RecommendSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecommendSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "feed.RecommendSrv",
	HandlerType: (*RecommendSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVideo",
			Handler:    _RecommendSrv_CreateVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recommend.proto",
}
