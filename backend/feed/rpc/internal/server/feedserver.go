// Code generated by goctl. DO NOT EDIT.
// Source: feed.proto

package server

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/logic"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"
)

type FeedServer struct {
	svcCtx *svc.ServiceContext
	feed.UnimplementedFeedServer
}

func NewFeedServer(svcCtx *svc.ServiceContext) *FeedServer {
	return &FeedServer{
		svcCtx: svcCtx,
	}
}

func (s *FeedServer) CreateVideo(ctx context.Context, in *feed.CreateVideoRequest) (*feed.CreateVideoResponse, error) {
	l := logic.NewCreateVideoLogic(ctx, s.svcCtx)
	return l.CreateVideo(in)
}

func (s *FeedServer) ListVideo(ctx context.Context, in *feed.ListVideoRequest) (*feed.ListVideoResponse, error) {
	l := logic.NewListVideoLogic(ctx, s.svcCtx)
	return l.ListVideo(in)
}

// rpc CountVideo(CountVideoRequest) returns (CountVideoResponse) {}
func (s *FeedServer) ListVideosByRecommend(ctx context.Context, in *feed.ListFeedRequest) (*feed.ListFeedResponse, error) {
	l := logic.NewListVideosByRecommendLogic(ctx, s.svcCtx)
	return l.ListVideosByRecommend(in)
}

func (s *FeedServer) ListVideos(ctx context.Context, in *feed.ListFeedRequest) (*feed.ListFeedResponse, error) {
	l := logic.NewListVideosLogic(ctx, s.svcCtx)
	return l.ListVideos(in)
}

func (s *FeedServer) ListCategoryVideos(ctx context.Context, in *feed.CategoryFeedRequest) (*feed.CategoryFeedResponse, error) {
	l := logic.NewListCategoryVideosLogic(ctx, s.svcCtx)
	return l.ListCategoryVideos(in)
}

func (s *FeedServer) ListPopularVideos(ctx context.Context, in *feed.ListFeedRequest) (*feed.ListFeedResponse, error) {
	l := logic.NewListPopularVideosLogic(ctx, s.svcCtx)
	return l.ListPopularVideos(in)
}

func (s *FeedServer) CreateVideoTest(ctx context.Context, in *feed.CreateVideoRequest) (*feed.CreateVideoResponse, error) {
	l := logic.NewCreateVideoTestLogic(ctx, s.svcCtx)
	return l.CreateVideoTest(in)
}

func (s *FeedServer) VideoDuration(ctx context.Context, in *feed.VideoDurationReq) (*feed.VideoDurationResp, error) {
	l := logic.NewVideoDurationLogic(ctx, s.svcCtx)
	return l.VideoDuration(in)
}

func (s *FeedServer) ListHistoryVideos(ctx context.Context, in *feed.HistoryReq) (*feed.HistoryResp, error) {
	l := logic.NewListHistoryVideosLogic(ctx, s.svcCtx)
	return l.ListHistoryVideos(in)
}

func (s *FeedServer) ListNeighborVideos(ctx context.Context, in *feed.NeighborsReq) (*feed.NeighborsResp, error) {
	l := logic.NewListNeighborVideosLogic(ctx, s.svcCtx)
	return l.ListNeighborVideos(in)
}

func (s *FeedServer) DeleteVideo(ctx context.Context, in *feed.DeleteVideoReq) (*feed.DeleteVideoResp, error) {
	l := logic.NewDeleteVideoLogic(ctx, s.svcCtx)
	return l.DeleteVideo(in)
}

func (s *FeedServer) SearchES(ctx context.Context, in *feed.EsSearchReq) (*feed.EsSearchResp, error) {
	l := logic.NewSearchESLogic(ctx, s.svcCtx)
	return l.SearchES(in)
}
