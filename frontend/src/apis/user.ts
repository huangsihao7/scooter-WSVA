/*
 * @Author: Xu Ning
 * @Date: 2023-10-28 13:48:11
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 18:09:03
 * @Description: 用户接口
 * @FilePath: \scooter-WSVA\frontend\src\apis\user.ts
 */
import { service } from "@/axios";

export function getUserInfo(uid: number) {
  return service({
    url: "/user/userinfo",
    method: "post",
    data: { uid },
  });
}

export function updateUserInfo(
  name: string,
  gender: number,
  avatar: string,
  dec: string,
  background_image: string,
) {
  return service({
    url: "/user/update",
    method: "post",
    data: { name, gender, avatar, dec, background_image },
  });
}
