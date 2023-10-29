/*
 * @Author: huangsihao7
 * @Date: 2023-10-29 17:09:12
 * @LastEditors: huangsihao7 1057434651@qq.com
 * @LastEditTime: 2023-10-29 19:40:00
 * @FilePath: /scooter-WSVA/frontend/src/apis/follow.ts
 * @Description: 关注
 */
import { service } from "@/axios";

export function getFollowList(uid: number) {
  return service({
    url: `/relation/favoriteList?uid=${uid}`,
    method: "get",
  });
}
export function followOne(to_user_id: number) {
  return service({
    url: `/relation/action`,
    method: "post",
    data: {
      to_user_id: to_user_id,
      action: 1,
    },
  });
}
export function canclefollowOne(to_user_id: number) {
  return service({
    url: `/relation/action`,
    method: "post",
    data: {
      to_user_id: to_user_id,
      action: 2,
    },
  });
}
