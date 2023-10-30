/*
 * @Author: huangsihao7
 * @Date: 2023-10-30 14:23:59
 * @LastEditors: huangsihao7 1057434651@qq.com
 * @LastEditTime: 2023-10-30 15:03:06
 * @FilePath: /scooter-WSVA/frontend/src/apis/favourite.ts
 * @Description:
 */
import { service } from "@/axios";

export function doStar(video_id: number, action_type: number) {
  return service({
    url: "/star/action",
    method: "post",
    data: {
      video_id: video_id,
      action_type: action_type,
    },
  });
}

export function doFavourite(video_id: number, action_type: number) {
  return service({
    url: "/favorite/action",
    method: "post",
    data: {
      video_id: video_id,
      action_type: action_type,
    },
  });
}

export function userFavouriteListReq(uid: number) {
  return service({
    url: `/favorite/list?uid=${uid}`,
    method: "get",
  });
}
export function userStarListReq(uid: number) {
  return service({
    url: `/star/list?uid=${uid}`,
    method: "get",
  });
}
