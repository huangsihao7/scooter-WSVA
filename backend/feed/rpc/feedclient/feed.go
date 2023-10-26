// Code generated by goctl. DO NOT EDIT.
// Source: feed.proto

package feedclient

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CountVideoRequest   = feed.CountVideoRequest
	CountVideoResponse  = feed.CountVideoResponse
	CreateVideoRequest  = feed.CreateVideoRequest
	CreateVideoResponse = feed.CreateVideoResponse
	ListFeedRequest     = feed.ListFeedRequest
	ListFeedResponse    = feed.ListFeedResponse
	ListVideoRequest    = feed.ListVideoRequest
	ListVideoResponse   = feed.ListVideoResponse
	User                = feed.User
	VideoInfo           = feed.VideoInfo

	Feed interface {
		CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*CreateVideoResponse, error)
		ListVideo(ctx context.Context, in *ListVideoRequest, opts ...grpc.CallOption) (*ListVideoResponse, error)
		// rpc CountVideo(CountVideoRequest) returns (CountVideoResponse) {}
		ListVideos(ctx context.Context, in *ListFeedRequest, opts ...grpc.CallOption) (*ListFeedResponse, error)
	}

	defaultFeed struct {
		cli zrpc.Client
	}
)

func NewFeed(cli zrpc.Client) Feed {
	return &defaultFeed{
		cli: cli,
	}
}

func (m *defaultFeed) CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*CreateVideoResponse, error) {
	client := feed.NewFeedClient(m.cli.Conn())
	return client.CreateVideo(ctx, in, opts...)
}

func (m *defaultFeed) ListVideo(ctx context.Context, in *ListVideoRequest, opts ...grpc.CallOption) (*ListVideoResponse, error) {
	client := feed.NewFeedClient(m.cli.Conn())
	return client.ListVideo(ctx, in, opts...)
}

// rpc CountVideo(CountVideoRequest) returns (CountVideoResponse) {}
func (m *defaultFeed) ListVideos(ctx context.Context, in *ListFeedRequest, opts ...grpc.CallOption) (*ListFeedResponse, error) {
	client := feed.NewFeedClient(m.cli.Conn())
	return client.ListVideos(ctx, in, opts...)
}
