// Code generated by goctl. DO NOT EDIT.
// Source: recommend.proto

package recommendsrv

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/recommend/rpc/recommend"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CategoryFeedRequest  = recommend.CategoryFeedRequest
	CategoryFeedResponse = recommend.CategoryFeedResponse
	CountVideoRequest    = recommend.CountVideoRequest
	CountVideoResponse   = recommend.CountVideoResponse
	CreateVideoRequest   = recommend.CreateVideoRequest
	CreateVideoResponse  = recommend.CreateVideoResponse
	ListFeedRequest      = recommend.ListFeedRequest
	ListFeedResponse     = recommend.ListFeedResponse
	ListVideoRequest     = recommend.ListVideoRequest
	ListVideoResponse    = recommend.ListVideoResponse
	User                 = recommend.User
	VideoInfo            = recommend.VideoInfo

	RecommendSrv interface {
		CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*CreateVideoResponse, error)
	}

	defaultRecommendSrv struct {
		cli zrpc.Client
	}
)

func NewRecommendSrv(cli zrpc.Client) RecommendSrv {
	return &defaultRecommendSrv{
		cli: cli,
	}
}

func (m *defaultRecommendSrv) CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*CreateVideoResponse, error) {
	client := recommend.NewRecommendSrvClient(m.cli.Conn())
	return client.CreateVideo(ctx, in, opts...)
}
