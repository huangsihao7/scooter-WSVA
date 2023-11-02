/*
 * @Author: Xu Ning
 * @Date: 2023-10-26 13:13:00
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-02 16:21:14
 * @Description: route
 * @FilePath: \scooter-WSVA\frontend\src\stores\relation.ts
 */
import { defineStore } from "pinia";
export const relationStore = defineStore({
  id: "relation",
  state: () => ({
    to_user_id: -1,
    type:-1
  }),
  persist: true
});
