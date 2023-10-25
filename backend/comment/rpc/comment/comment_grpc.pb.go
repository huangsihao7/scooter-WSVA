// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: comment.proto

package comment

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
	CommentSrv_CommentAction_FullMethodName = "/comment.CommentSrv/CommentAction"
	CommentSrv_CommentList_FullMethodName   = "/comment.CommentSrv/CommentList"
)

// CommentSrvClient is the client API for CommentSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentSrvClient interface {
	CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error)
	CommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error)
}

type commentSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentSrvClient(cc grpc.ClientConnInterface) CommentSrvClient {
	return &commentSrvClient{cc}
}

func (c *commentSrvClient) CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error) {
	out := new(CommentActionResponse)
	err := c.cc.Invoke(ctx, CommentSrv_CommentAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentSrvClient) CommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error) {
	out := new(CommentListResponse)
	err := c.cc.Invoke(ctx, CommentSrv_CommentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentSrvServer is the server API for CommentSrv service.
// All implementations must embed UnimplementedCommentSrvServer
// for forward compatibility
type CommentSrvServer interface {
	CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error)
	CommentList(context.Context, *CommentListRequest) (*CommentListResponse, error)
	mustEmbedUnimplementedCommentSrvServer()
}

// UnimplementedCommentSrvServer must be embedded to have forward compatible implementations.
type UnimplementedCommentSrvServer struct {
}

func (UnimplementedCommentSrvServer) CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedCommentSrvServer) CommentList(context.Context, *CommentListRequest) (*CommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentList not implemented")
}
func (UnimplementedCommentSrvServer) mustEmbedUnimplementedCommentSrvServer() {}

// UnsafeCommentSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentSrvServer will
// result in compilation errors.
type UnsafeCommentSrvServer interface {
	mustEmbedUnimplementedCommentSrvServer()
}

func RegisterCommentSrvServer(s grpc.ServiceRegistrar, srv CommentSrvServer) {
	s.RegisterService(&CommentSrv_ServiceDesc, srv)
}

func _CommentSrv_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentSrvServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentSrv_CommentAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentSrvServer).CommentAction(ctx, req.(*CommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentSrv_CommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentSrvServer).CommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentSrv_CommentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentSrvServer).CommentList(ctx, req.(*CommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentSrv_ServiceDesc is the grpc.ServiceDesc for CommentSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.CommentSrv",
	HandlerType: (*CommentSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommentAction",
			Handler:    _CommentSrv_CommentAction_Handler,
		},
		{
			MethodName: "CommentList",
			Handler:    _CommentSrv_CommentList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
