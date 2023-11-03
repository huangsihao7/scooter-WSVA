package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/mq/format"
	"github.com/huangsihao7/scooter-WSVA/mq/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"strconv"
	"time"
)

type ParseJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewParseJob(ctx context.Context, svcCtx *svc.ServiceContext) *ParseJob {
	return &ParseJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ParseJob) Consume(key, val string) error {
	logx.Infof("job key :%s , val :%s", key, val)
	var job format.JobBody
	err := json.Unmarshal([]byte(val), &job)
	if err != nil {
		fmt.Println("解析 JSON 失败:", err.Error())
		return err
	}
	status, vid, suggestion, err := format.GetJobBack(job.Job, l.svcCtx.Config.SecretKey, l.svcCtx.Config.AccessKey)
	if err != nil {
		println(err)
		return err
	}
	if status != "FINISHED" {
		//未完成或者失败，放进队列
		jobKq := format.JobBody{
			Job: job.Job,
			Id:  job.Id,
			Url: job.Url,
			Uid: job.Uid,
		}
		time.Sleep(5 * time.Second)
		data, err := json.Marshal(jobKq)
		if err != nil {
			logx.Error("[Video] marshal msg: %v error: %v", jobKq, err)
			return err
		}
		err = l.svcCtx.KqPusherJobClient.Push(string(data))
		if err != nil {
			logx.Error("[Video] kq push data: %s error: %v", data, err)
			return err
		}

	} else {
		if suggestion != "pass" {
			//审核不通过，删除该视频
			num, _ := strconv.ParseInt(vid, 10, 32)
			video, err := l.svcCtx.Feeder.DeleteVideo(l.ctx, &feed.DeleteVideoReq{Vid: int32(num)})
			if err != nil {
				println(video.StatusMsg)
				return err
			}
			println("审核不通过，删除视频")
			return nil

		} else {
			messagekq := format.UploadFile{
				Id:  job.Id,
				Url: job.Url,
				Uid: job.Uid,
			}
			// 发送kafka消息，异步
			threading.GoSafe(func() {
				data, err := json.Marshal(messagekq)
				if err != nil {
					return
				}
				err = l.svcCtx.KqPusherClient.Push(string(data))
			})
			println("审核通过")
			return nil
		}
	}
	return nil

}
