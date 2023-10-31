/*
 * @Author: li.yunhao
 * @Date: 2023-04-21 23:52:58
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-27 22:36:59
 * @FilePath: \scooter-WSVA\frontend\src\stores\user.ts
 * @Description:
 */
import { defineStore } from "pinia";

export const userStore = defineStore({
  id: "user",
  state: () => ({
    isLoggedIn: false,
    token: "",
    user_id: -1,
    name: "",
    avatar: "",
    gender: 1,
    signature: "",
    background_image: "",
    phoneNum: "",
  }),
  persist: true,
});
