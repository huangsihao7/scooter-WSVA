package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/types"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentListLogic) GetCommentList(req *types.ListReq) (resp *types.ListResp, err error) {

	usrId, _ := l.ctx.Value("uid").(json.Number).Int64()

	res, err := l.svcCtx.Commenter.CommentList(l.ctx, &comment.CommentListRequest{UserId: usrId, VideoId: req.VideoId})

	if err != nil {
		return nil, err
	}
	resLists := make([]types.CommentInfo, 0)

	for i := 0; i < len(res.CommentList); i++ {
		newUser := types.User{
			Id:             res.CommentList[i].User.Id,
			Name:           res.CommentList[i].User.Name,
			FollowCount:    *res.CommentList[i].User.FollowCount,
			FollowerCount:  *res.CommentList[i].User.FollowerCount,
			Gender:         res.CommentList[i].User.Gender,
			IsFollow:       res.CommentList[i].User.IsFollow,
			Avatar:         *res.CommentList[i].User.Avatar,
			Signature:      *res.CommentList[i].User.Signature,
			TotalFavorited: *res.CommentList[i].User.TotalFavorited,
			WorkCount:      *res.CommentList[i].User.WorkCount,
			FavoriteCount:  *res.CommentList[i].User.FavoriteCount,
			FriendCount:    res.CommentList[i].User.FriendCount,
		}
		commentDetail := types.CommentInfo{
			CommentId:  res.CommentList[i].Id,
			User:       newUser,
			Content:    res.CommentList[i].Content,
			CreateDate: res.CommentList[i].CreateDate,
		}
		resLists = append(resLists, commentDetail)

	}
	return &types.ListResp{
		StatusCode:  int(res.StatusCode),
		StatusMsg:   res.StatusMsg,
		CommentList: resLists,
	}, nil
}
