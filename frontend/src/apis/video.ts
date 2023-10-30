/*
 * @Author: huangsihao7
 * @Date: 2023-10-29 13:04:21
 * @LastEditors: huangsihao7 1057434651@qq.com
 * @LastEditTime: 2023-10-30 09:48:38
 * @FilePath: /scooter-WSVA/frontend/src/apis/video.ts
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

export function getVideosList() {
  return service({
    url: `/feed/VideosList`, // 使用字符串模板来拼接runId
    method: "get",
  });
}

export function getCategoryVideosList(category: string) {
  return service({
    url: `/feed/CategoryVideosList/${category}`, // 使用字符串模板来拼接runId
    method: "get",
  });
}
