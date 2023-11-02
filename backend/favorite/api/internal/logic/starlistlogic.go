package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/types"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favorite"

	"github.com/zeromicro/go-zero/core/logx"
)

type StarListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStarListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StarListLogic {
	return &StarListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StarListLogic) StarList(req *types.ListReq) (resp *types.ListResp, err error) {

	usrId, _ := l.ctx.Value("uid").(json.Number).Int64()

	res, err := l.svcCtx.Favor.StarList(l.ctx, &favorite.StarListRequest{UserId: usrId, ActorId: req.UserId})
	if err != nil {
		return nil, err
	}

	resLists := make([]types.VideoInfo, 0)

	for i := 0; i < len(res.VideoList); i++ {
		newUser := types.User{
			Id:             res.VideoList[i].Author.Id,
			Name:           res.VideoList[i].Author.Name,
			Gender:         res.VideoList[i].Author.Gender,
			FollowCount:    *res.VideoList[i].Author.FollowCount,
			FollowerCount:  *res.VideoList[i].Author.FollowerCount,
			IsFollow:       res.VideoList[i].Author.IsFollow,
			Avatar:         *res.VideoList[i].Author.Avatar,
			Signature:      *res.VideoList[i].Author.Signature,
			TotalFavorited: *res.VideoList[i].Author.TotalFavorited,
			WorkCount:      *res.VideoList[i].Author.WorkCount,
			FavoriteCount:  *res.VideoList[i].Author.FavoriteCount,
			FriendCount:    res.VideoList[i].Author.FriendCount,
		}
		videoDetail := types.VideoInfo{
			VideoId:       res.VideoList[i].Id,
			User:          newUser,
			PlayUrl:       res.VideoList[i].PlayUrl,
			CommentCount:  res.VideoList[i].CommentCount,
			CoverUrl:      res.VideoList[i].CoverUrl,
			FavoriteCount: res.VideoList[i].FavoriteCount,
			StarCount:     res.VideoList[i].StarCount,
			IsFavorite:    res.VideoList[i].IsFavorite,
			IsStar:        res.VideoList[i].IsStar,
			Title:         res.VideoList[i].Title,
			CreateTime:    res.VideoList[i].CreateTime,
			Duration:      res.VideoList[i].Duration,
		}
		resLists = append(resLists, videoDetail)

	}
	return &types.ListResp{
		StatusCode: int(res.StatusCode),
		StatusMsg:  res.StatusMsg,
		VideoList:  resLists,
	}, nil
}
