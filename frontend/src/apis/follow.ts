/*
 * @Author: huangsihao7
 * @Date: 2023-10-29 17:09:12
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 18:08:37
 * @FilePath: \scooter-WSVA\frontend\src\apis\follow.ts
 * @Description: 关注操作接口
 */
import { service } from "@/axios";

export function getFollowsList(user_id: number) {
  return service({
    url: `/relation/favoriteList?uid=${user_id}`,
    method: "get",
  });
}

export function getFollowersList(user_id: number) {
  return service({
    url: `/relation/followerList?uid=${user_id}`,
    method: "get",
  });
}

export function getFriendsList(user_id: number) {
  return service({
    url: `/relation/friendList?uid=${user_id}`,
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

export function doFollow(to_user_id: number, action_type: number) {
  return service({
    url: "/relation/action",
    method: "post",
    data: {
      to_user_id: to_user_id,
      action_type: action_type,
    },
  });
}
