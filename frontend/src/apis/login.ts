/*
 * @Author: Xu Ning
 * @Date: 2023-10-25 14:19:40
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 13:28:26
 * @Description:
 * @FilePath: \scooter-WSVA\frontend\src\apis\login.ts
 */
import { service } from "@/axios";
export function login(mobile: string, password: string) {
  return service({
    url: "/user/login",
    method: "post",
    data: {
      mobile,
      password,
    },
  });
}

export function register(
  name: string,
  gender: number,
  mobile: string,
  password: string,
  dec: string,
  avatar: string,
  background_image: string,
) {
  return service({
    url: "/user/register",
    method: "post",
    data: {
      name,
      gender,
      mobile,
      password,
      dec,
      avatar,
      background_image,
    },
  });
}
