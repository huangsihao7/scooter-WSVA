// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: feed.proto

package feed

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
	Feed_CreateVideo_FullMethodName           = "/feed.Feed/CreateVideo"
	Feed_ListVideo_FullMethodName             = "/feed.Feed/ListVideo"
	Feed_CountVideo_FullMethodName            = "/feed.Feed/CountVideo"
	Feed_ListVideosByRecommend_FullMethodName = "/feed.Feed/ListVideosByRecommend"
	Feed_ListVideos_FullMethodName            = "/feed.Feed/ListVideos"
)

// FeedClient is the client API for Feed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedClient interface {
	CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*CreateVideoResponse, error)
	ListVideo(ctx context.Context, in *ListVideoRequest, opts ...grpc.CallOption) (*ListVideoResponse, error)
	CountVideo(ctx context.Context, in *CountVideoRequest, opts ...grpc.CallOption) (*CountVideoResponse, error)
	ListVideosByRecommend(ctx context.Context, in *ListFeedRequest, opts ...grpc.CallOption) (*ListFeedResponse, error)
	ListVideos(ctx context.Context, in *ListFeedRequest, opts ...grpc.CallOption) (*ListFeedResponse, error)
}

type feedClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedClient(cc grpc.ClientConnInterface) FeedClient {
	return &feedClient{cc}
}

func (c *feedClient) CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*CreateVideoResponse, error) {
	out := new(CreateVideoResponse)
	err := c.cc.Invoke(ctx, Feed_CreateVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) ListVideo(ctx context.Context, in *ListVideoRequest, opts ...grpc.CallOption) (*ListVideoResponse, error) {
	out := new(ListVideoResponse)
	err := c.cc.Invoke(ctx, Feed_ListVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) CountVideo(ctx context.Context, in *CountVideoRequest, opts ...grpc.CallOption) (*CountVideoResponse, error) {
	out := new(CountVideoResponse)
	err := c.cc.Invoke(ctx, Feed_CountVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) ListVideosByRecommend(ctx context.Context, in *ListFeedRequest, opts ...grpc.CallOption) (*ListFeedResponse, error) {
	out := new(ListFeedResponse)
	err := c.cc.Invoke(ctx, Feed_ListVideosByRecommend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) ListVideos(ctx context.Context, in *ListFeedRequest, opts ...grpc.CallOption) (*ListFeedResponse, error) {
	out := new(ListFeedResponse)
	err := c.cc.Invoke(ctx, Feed_ListVideos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServer is the server API for Feed service.
// All implementations must embed UnimplementedFeedServer
// for forward compatibility
type FeedServer interface {
	CreateVideo(context.Context, *CreateVideoRequest) (*CreateVideoResponse, error)
	ListVideo(context.Context, *ListVideoRequest) (*ListVideoResponse, error)
	CountVideo(context.Context, *CountVideoRequest) (*CountVideoResponse, error)
	ListVideosByRecommend(context.Context, *ListFeedRequest) (*ListFeedResponse, error)
	ListVideos(context.Context, *ListFeedRequest) (*ListFeedResponse, error)
	mustEmbedUnimplementedFeedServer()
}

// UnimplementedFeedServer must be embedded to have forward compatible implementations.
type UnimplementedFeedServer struct {
}

func (UnimplementedFeedServer) CreateVideo(context.Context, *CreateVideoRequest) (*CreateVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVideo not implemented")
}
func (UnimplementedFeedServer) ListVideo(context.Context, *ListVideoRequest) (*ListVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVideo not implemented")
}
func (UnimplementedFeedServer) CountVideo(context.Context, *CountVideoRequest) (*CountVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountVideo not implemented")
}
func (UnimplementedFeedServer) ListVideosByRecommend(context.Context, *ListFeedRequest) (*ListFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVideosByRecommend not implemented")
}
func (UnimplementedFeedServer) ListVideos(context.Context, *ListFeedRequest) (*ListFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVideos not implemented")
}
func (UnimplementedFeedServer) mustEmbedUnimplementedFeedServer() {}

// UnsafeFeedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedServer will
// result in compilation errors.
type UnsafeFeedServer interface {
	mustEmbedUnimplementedFeedServer()
}

func RegisterFeedServer(s grpc.ServiceRegistrar, srv FeedServer) {
	s.RegisterService(&Feed_ServiceDesc, srv)
}

func _Feed_CreateVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).CreateVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_CreateVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).CreateVideo(ctx, req.(*CreateVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_ListVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).ListVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_ListVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).ListVideo(ctx, req.(*ListVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_CountVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).CountVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_CountVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).CountVideo(ctx, req.(*CountVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_ListVideosByRecommend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).ListVideosByRecommend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_ListVideosByRecommend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).ListVideosByRecommend(ctx, req.(*ListFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_ListVideos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).ListVideos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_ListVideos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).ListVideos(ctx, req.(*ListFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Feed_ServiceDesc is the grpc.ServiceDesc for Feed service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Feed_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "feed.Feed",
	HandlerType: (*FeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVideo",
			Handler:    _Feed_CreateVideo_Handler,
		},
		{
			MethodName: "ListVideo",
			Handler:    _Feed_ListVideo_Handler,
		},
		{
			MethodName: "CountVideo",
			Handler:    _Feed_CountVideo_Handler,
		},
		{
			MethodName: "ListVideosByRecommend",
			Handler:    _Feed_ListVideosByRecommend_Handler,
		},
		{
			MethodName: "ListVideos",
			Handler:    _Feed_ListVideos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feed.proto",
}
