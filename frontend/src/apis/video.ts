/*
 * @Author: huangsihao7
 * @Date: 2023-10-29 13:04:21
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-31 21:15:44
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

export function userVideoListReq(to_user_id: number) {
  return service({
    url: "/feed/UserVideosList",
    method: "post",
    data: { to_user_id },
  });
}

export function getRecommendVideos(offset: number, readed_videoId: number) {
  return service({
    url: "/feed/recommends",
    method: "post",
    data: { offset, readed_videoId },
  });
}

export function getPopularVideos(offset: number, readed_videoId: number) {
  return service({
    url: "/feed/populars",
    method: "post",
    data: { offset, readed_videoId },
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

export function getRecommendVideosList(video_id: number) {
  return service({
    url: `/feed/neighbors?video_id=${video_id}`, // 使用字符串模板来拼接runId
    method: "get",
  });
}

export function getHistoryVideosListReq() {
  return service({
    url: `/feed/history`,
    method: "get",
  });
}
