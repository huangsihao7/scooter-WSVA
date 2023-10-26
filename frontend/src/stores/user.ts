/*
 * @Author: li.yunhao
 * @Date: 2023-04-21 23:52:58
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-26 15:15:57
 * @FilePath: \scooter-wsva\frontend\src\stores\user.ts
 * @Description:
 */
import { defineStore } from 'pinia'
export const userStore = defineStore({
  id: 'user',
  state: () => ({
    isLoggedIn: false,
    token: '',
    avatar: 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png',
    username: ''
  })
})