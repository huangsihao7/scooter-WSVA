/*
 * @Author: huangsihao7
 * @Date: 2023-10-29 17:09:12
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-02 14:44:41
 * @FilePath: \scooter-WSVA\frontend\src\apis\follow.ts
 * @Description: 关注
 */
import { service } from "@/axios";

export function getFollowList(user_id: number) {
  return service({
    url: `/relation/favoriteList?uid=${user_id}`,
    method: "get",
  });
}
export function followOne(to_user_id: number) {
  return service({
    url: "/relation/action",
    method: "post",
    data: {
      to_user_id: to_user_id,
      action_type: 1,
    },
  });
}
export function canclefollowOne(to_user_id: number) {
  return service({
    url: "/relation/action",
    method: "post",
    data: {
      to_user_id: to_user_id,
      action_type: 2,
    },
  });
}
