/*
 * @Author: huangsihao7
 * @Date: 2023-10-29 13:04:21
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-30 16:25:04
 * @FilePath: \scooter-WSVA\frontend\src\apis\video.ts
 * @Description: 视频接口
 */
import { service } from "@/axios";

export function postVideo(
  url: string,
  coverUrl: string,
  title: string,
  category: number,
) {
  return service({
    url: "/feed/create",
    method: "post",
    data: { url, coverUrl, title, category },
  });
}

export function userVideoListReq(toUid: number) {
  return service({
    url: "/feed/UserVideosList",
    method: "post",
    data: { toUid },
  });
}

export function getRecommendVideos(offset: number) {
  return service({
    url: "/feed/recommends",
    method: "post",
    data: { offset },
  });
}

export function getPopularVideos(offset: number) {
  return service({
    url: "/feed/populars",
    method: "post",
    data: { offset },
  });
}

export function getVideosList() {
  return service({
    url: `/feed/VideosList`, // 使用字符串模板来拼接runId
    method: "get",
  });
}

export function getCategoryVideosList(category: number) {
  return service({
    url: `/feed/CategoryVideosList?category=${category}`, // 使用字符串模板来拼接runId
    method: "get",
  });
}
