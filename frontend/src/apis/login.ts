/*
 * @Author: Xu Ning
 * @Date: 2023-10-25 14:19:40
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-25 14:19:49
 * @Description: 
 * @FilePath: \scooter-wsva\frontend\src\apis\login.ts
 */
import { service } from '@/axios'
export function login(username: string, password: string) {
  return service({
    url: '/user/login',
    method: 'post',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    data: {
      username,
      password
    }
  })
}
