/*
 * @Author: Xu Ning
 * @Date: 2023-10-26 13:13:00
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-30 21:20:36
 * @Description: route
 * @FilePath: \scooter-WSVA\frontend\src\stores\video.ts
 */
import { defineStore } from "pinia";
export const videoStore = defineStore({
  id: "video",
  state: () => ({
    video_id: -1,
    search_value : ''
  }),
});
