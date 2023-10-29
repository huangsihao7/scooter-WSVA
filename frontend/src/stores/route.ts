/*
 * @Author: Xu Ning
 * @Date: 2023-10-26 13:13:00
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-26 14:55:32
 * @Description: route
 * @FilePath: \scooter-wsva\frontend\src\stores\route.ts
 */
import { defineStore } from "pinia";
export const routeStore = defineStore({
  id: "route",
  state: () => ({
    name: "",
  }),
});
