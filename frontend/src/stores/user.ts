/*
 * @Author: li.yunhao
 * @Date: 2023-04-21 23:52:58
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-25 14:26:45
 * @FilePath: \scooter-wsva\frontend\src\stores\user.ts
 * @Description:
 */
import { defineStore } from 'pinia'

interface userState {
  isLoggedIn: boolean
  token: string
  avatar: string
  username: string
}

export const userStore = defineStore({
  id: 'user',
  state: (): userState => ({
    isLoggedIn: false,
    token: '',
    avatar: '',
    username: ''
  }),
  actions: {
    async logout() {
      this.isLoggedIn = false
      this.token = ''
      this.avatar = '1111'
      this.username = ''
    },
    async login(token: string, avatar: string, username: string) {
      this.isLoggedIn = true
      this.token = token
      this.avatar = avatar
      this.username = username
    }
  },
  persist: true
})
