/*
 * @Author: Xu Ning
 * @Date: 2023-10-29 22:43:26
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-01 12:37:49
 * @Description:
 * @FilePath: \scooter-WSVA\frontend\src\apis\comment.ts
 */
import { service } from "@/axios";

export function getCommentList(video_id: number) {
  return service({
    url: `/comment/list?video_id=${video_id}`,
    method: "get",
  });
}

export function getDanmuList(video_id: number) {
  return service({
    url: `/danmu/list?video_id=${video_id}`,
    method: "get",
  });
}

export function doComment(
  video_id: number,
  action_type: number,
  comment_text: string,
  comment_id: number,
) {
  return service({
    url: "/comment/action",
    method: "post",
    data: {
      video_id: video_id,
      action_type: action_type,
      comment_text: comment_text,
      comment_id: comment_id,
    },
  });
}
