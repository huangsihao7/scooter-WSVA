/*
 * @Author: Xu Ning
 * @Date: 2023-10-25 14:19:40
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-27 11:59:04
 * @Description:
 * @FilePath: \scooter-WSVA\frontend\src\apis\login.ts
 */
import { service } from "@/axios";
export function login(mobile: string, password: string) {
  return service({
    url: "/user/login",
    method: "post",
    // headers: {
    //   'Content-Type': 'application/x-www-form-urlencoded'
    // },
    data: {
      mobile,
      password,
    },
  });
}
